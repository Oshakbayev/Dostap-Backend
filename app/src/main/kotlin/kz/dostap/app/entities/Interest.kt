package kz.dostap.app.entities

import org.jetbrains.exposed.dao.LongEntity
import org.jetbrains.exposed.dao.LongEntityClass
import org.jetbrains.exposed.dao.id.EntityID
import org.jetbrains.exposed.dao.id.LongIdTable

object InterestTable : LongIdTable("interests") {
    val name = varchar("name", 255)
    val category = reference("category_id", InterestCategoryTable)
}

object InterestCategoryTable : LongIdTable("category") {
    val name = varchar("name", 255)
}

class InterestEntity(id: EntityID<Long>) : LongEntity(id) {
    companion object : LongEntityClass<InterestEntity>(InterestTable)
    var name by InterestTable.name
    var category by InterestCategoryEntity referencedOn InterestTable.category
}

class InterestCategoryEntity(id: EntityID<Long>) : LongEntity(id) {
    companion object : LongEntityClass<InterestCategoryEntity>(InterestTable)
    var name by InterestTable.name
}