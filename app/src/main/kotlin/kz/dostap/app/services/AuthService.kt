package kz.dostap.app.services

import arrow.core.Either
import kz.dostap.app.models.Token
import kz.dostap.app.models.UserSignUpRequest

interface AuthService {
    suspend fun login(username: String, password: String): Either<LoginError, Token>
    suspend fun signup(request: UserSignUpRequest): Either<SignUpError, Token>
    suspend fun refreshToken(token: String): Token?
    suspend fun confirmEmail(code: String): Boolean
}

sealed class LoginError {
    data object InvalidCredentials : LoginError()
}

sealed class SignUpError {
    data object EmailOrUsernameAlreadyTaken : SignUpError()
    data object UsernameAlreadyTaken : SignUpError()
    data object CityNotFound : SignUpError()
}