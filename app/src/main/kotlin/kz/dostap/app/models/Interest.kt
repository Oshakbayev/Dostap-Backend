package kz.dostap.app.models

import kotlinx.serialization.Serializable
import kz.dostap.app.entities.InterestEntity

@Serializable
data class Interest(
    val id: Long,
    val name: String,
    val category: String
)

fun InterestEntity.toInterest() = Interest(
    id.value,
    name,
    category.name
)