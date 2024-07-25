package kz.dostap.app.models

import kotlinx.serialization.Serializable
import kz.dostap.app.entities.Gender

@Serializable
data class UserUpdateRequest(
    val firstName: String,
    val lastName: String,
    val cityId: Long,
    val avatarLink: String?,
    val gender: Gender?,
    val age: Int?,
    val phoneNumber: String?,
    val description: String?
)
