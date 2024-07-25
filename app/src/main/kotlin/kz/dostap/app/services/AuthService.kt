package kz.dostap.app.services

import kz.dostap.app.models.Token
import kz.dostap.app.models.UserSignUpRequest

interface AuthService {
    fun login(username: String, password: String): Token
    fun signup(request: UserSignUpRequest): Token
    fun refreshToken(token: String): Token
    fun confirmEmail(confirmationId: String): Boolean
}