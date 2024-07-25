package kz.dostap.app.models

import kotlinx.serialization.Serializable
import kz.dostap.app.entities.Gender
import kz.dostap.app.entities.UserEntity

@Serializable
data class User(
    val id: Long,
    val firstName: String,
    val lastName: String,
    val email: String,
    val username: String,
    val cityId: Long,
    val avatarLink: String?,
    val gender: Gender?,
    val age: Int?,
    val phoneNumber: String?,
    val description: String?,
    val interests: List<Interest>
)

fun UserEntity.toUser() = User(
    id.value,
    firstName,
    lastName,
    email,
    username,
    residenceCityId.value,
    avatarLink,
    gender,
    age,
    phoneNumber,
    description,
    interests.map { it.toInterest() }
)
