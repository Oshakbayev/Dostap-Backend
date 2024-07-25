package kz.dostap.app.plugins

import com.auth0.jwt.JWT
import com.auth0.jwt.algorithms.Algorithm
import com.auth0.jwt.interfaces.Payload
import io.ktor.server.application.*
import io.ktor.server.auth.*
import io.ktor.server.auth.jwt.*
import kz.dostap.app.appModule

fun Application.security() {
    val secret = environment.config.property("jwt.secret").getString()
    val issuer = environment.config.property("jwt.issuer").getString()
    val audience = environment.config.property("jwt.audience").getString()
    val myRealm = environment.config.property("jwt.realm").getString()
    val expiresIn = environment.config.property("jwt.expiresIn").getString().toInt()

    appModule.single { JWTConfig(secret, issuer, audience, myRealm, expiresIn) }

    install(Authentication) {
        jwt {
            realm = myRealm
            verifier(JWT
                .require(Algorithm.HMAC256(secret))
                .withAudience(audience)
                .withIssuer(issuer)
                .build())
            validate { credential ->
                val id = credential.payload.getClaim("id").asLong()
                val username = credential.subject
                if (id != null && username != null && username != "") {
                    UserPrincipal(username, id,credential.payload)
                } else {
                    null
                }
            }
        }
    }

    val emailVerificationExpiresIn = environment.config.property("verification.expiresIn").getString().toInt()
    appModule.single { EmailVerificationConfig(emailVerificationExpiresIn) }
}

data class JWTConfig(
    val secret: String,
    val issuer: String,
    val audience: String,
    val realm: String,
    val expiresIn: Int
)

data class EmailVerificationConfig(
    val expiresIn: Int,
)

data class UserPrincipal(
    val username: String,
    val id: Long,
    val payload: Payload
) : Principal