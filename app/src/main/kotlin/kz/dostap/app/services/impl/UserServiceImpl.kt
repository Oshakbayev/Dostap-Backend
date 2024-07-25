package kz.dostap.app.services.impl

import arrow.core.Either
import arrow.core.raise.either
import kz.dostap.app.entities.CityEntity
import kz.dostap.app.entities.UserEntity
import kz.dostap.app.models.User
import kz.dostap.app.models.UserUpdateRequest
import kz.dostap.app.models.toUser
import kz.dostap.app.services.UpdateError
import kz.dostap.app.services.UserService
import org.jetbrains.exposed.sql.transactions.experimental.newSuspendedTransaction

class UserServiceImpl : UserService {
    override suspend fun getUser(id: Long): User? = newSuspendedTransaction {
        UserEntity.findById(id)?.toUser()
    }

    override suspend fun updateUser(id: Long, request: UserUpdateRequest): Either<UpdateError, User> = either {
        newSuspendedTransaction {
            val city = CityEntity.findById(id) ?: raise(UpdateError.CityNotFound)

            val user = UserEntity.findById(id) ?: raise(UpdateError.UserNotFound)
            user.firstName = request.firstName
            user.lastName = request.lastName
            user.residenceCity = city
            user.avatarLink = request.avatarLink
            user.gender = request.gender
            user.age = request.age
            user.phoneNumber = request.phoneNumber
            user.description = request.description
            user.toUser()
        }
    }
}