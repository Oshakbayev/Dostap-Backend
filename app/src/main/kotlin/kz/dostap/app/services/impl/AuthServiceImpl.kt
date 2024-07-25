package kz.dostap.app.services.impl

import arrow.core.Either
import arrow.core.raise.either
import com.auth0.jwt.JWT
import com.auth0.jwt.algorithms.Algorithm
import kotlinx.datetime.Clock
import kz.dostap.app.entities.*
import kz.dostap.app.models.Token
import kz.dostap.app.models.UserSignUpRequest
import kz.dostap.app.plugins.EmailVerificationConfig
import kz.dostap.app.plugins.JWTConfig
import kz.dostap.app.services.AuthService
import kz.dostap.app.services.LoginError
import kz.dostap.app.services.SignUpError
import org.jetbrains.exposed.sql.or
import org.jetbrains.exposed.sql.transactions.experimental.newSuspendedTransaction
import org.springframework.security.crypto.bcrypt.BCrypt
import java.util.*
import kotlin.time.Duration.Companion.minutes

class AuthServiceImpl(
    private val config: JWTConfig,
    emailVerificationConfig: EmailVerificationConfig
) : AuthService {
    private val emailVerificationExpiresIn = emailVerificationConfig.expiresIn.minutes

    override suspend fun login(username: String, password: String): Either<LoginError.InvalidCredentials, Token> = either {
        val user = newSuspendedTransaction user@{
             UserEntity.find { UserTable.username eq username }
                .firstOrNull()
        } ?: raise(LoginError.InvalidCredentials)
        if (user.encryptedPassword != password) {
            raise(LoginError.InvalidCredentials)
        }
        Token(generateToken(user), config.expiresIn)
    }

    override suspend fun signup(request: UserSignUpRequest) = either {
        val user = newSuspendedTransaction user@{
            UserEntity.find { (UserTable.username eq request.username) or (UserTable.email eq request.email) }
                .firstOrNull()
                ?.let { raise(
                    if (it.username == request.username) SignUpError.UsernameAlreadyTaken
                    else SignUpError.EmailOrUsernameAlreadyTaken
                ) }

            val city = CityEntity.findById(request.cityId) ?: raise(SignUpError.CityNotFound)
            UserEntity.new {
                firstName = request.firstName
                lastName = request.lastName
                email = request.email
                username = request.username
                encryptedPassword = request.password
                residenceCity = city
                avatarLink = request.avatarLink
                gender = request.gender
                age = request.age
                phoneNumber = request.phoneNumber
                description = request.description
            }.also { createEmailVerification(it) }
        }

        Token(generateToken(user), config.expiresIn)
    }

    override suspend fun refreshToken(token: String): Token? {
        TODO("Not yet implemented")
    }

    override suspend fun confirmEmail(code: String): Boolean = newSuspendedTransaction confirm@{
        val verification = EmailVerificationEntity.findValidBySecretCode(code)
            ?: return@confirm false
        verification.isUsed = true
        verification.user.isEmailVerified = true
        true
    }

    private fun createEmailVerification(user: UserEntity): EmailVerificationEntity = EmailVerificationEntity.new {
        this.user = user
        this.email = user.email
        this.secretCode = generateVerificationCode()
        this.expiresAt = Clock.System.now() + emailVerificationExpiresIn
    }

    private fun generateVerificationCode(): String {
        val result = UUID.randomUUID().toString()
        val code = BCrypt.hashpw(result, BCrypt.gensalt())
        val alreadyExistingCode = EmailVerificationEntity.findValidBySecretCode(code)
        if (alreadyExistingCode != null) {
            return generateVerificationCode()
        }
        return code
    }

    private fun generateToken(user: UserEntity): String = JWT.create()
        .withAudience(config.audience)
        .withIssuer(config.issuer)
        .withSubject(user.username)
        .withClaim("id", user.id.value)
        .withExpiresAt(Date(System.currentTimeMillis() + config.expiresIn))
        .sign(Algorithm.HMAC256(config.secret))
}