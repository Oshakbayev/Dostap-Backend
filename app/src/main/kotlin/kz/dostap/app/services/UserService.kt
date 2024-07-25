package kz.dostap.app.services

import arrow.core.Either
import kz.dostap.app.models.User
import kz.dostap.app.models.UserUpdateRequest

interface UserService {
    suspend fun getUser(id: Long): User?
    suspend fun updateUser(id: Long, request: UserUpdateRequest): Either<UpdateError, User>
}

sealed class UpdateError {
    data object UserNotFound : UpdateError()
    data object CityNotFound : UpdateError()
}