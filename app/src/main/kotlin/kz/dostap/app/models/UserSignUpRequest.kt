package kz.dostap.app.models

import kotlinx.serialization.Serializable
import kz.dostap.app.entities.Gender

@Serializable
data class UserSignUpRequest(
    val firstName: String,
    val lastName: String,
    val email: String,
    val username: String,
    val password: String,
    val cityId: Long,
    val avatarLink: String?,
    val gender: Gender?,
    val age: Int?,
    val phoneNumber: String?,
    val description: String?
)
