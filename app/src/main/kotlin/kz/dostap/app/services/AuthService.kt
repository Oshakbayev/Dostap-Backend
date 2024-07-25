package kz.dostap.app.services

import kz.dostap.app.models.Token
import kz.dostap.app.models.UserSignUpRequest

interface AuthService {
    suspend fun login(username: String, password: String): Token?
    suspend fun signup(request: UserSignUpRequest): Token?
    suspend fun refreshToken(token: String): Token?
    suspend fun confirmEmail(code: String): Boolean
}