package kz.dostap.app.models

import kotlinx.serialization.Serializable

@Serializable
data class Token(
    val accessToken: String,
    val expiresIn: Int
)
