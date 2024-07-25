package kz.dostap.app.entities

import org.jetbrains.exposed.dao.LongEntity
import org.jetbrains.exposed.dao.LongEntityClass
import org.jetbrains.exposed.dao.id.EntityID
import org.jetbrains.exposed.dao.id.LongIdTable

object CityTable : LongIdTable("cities") {
    val name = varchar("name", 255)
}

class CityEntity(id: EntityID<Long>) : LongEntity(id) {
    companion object : LongEntityClass<CityEntity>(CityTable)
    var name by CityTable.name
}