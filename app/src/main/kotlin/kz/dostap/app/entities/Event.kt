package kz.dostap.app.entities

import org.jetbrains.exposed.dao.LongEntity
import org.jetbrains.exposed.dao.LongEntityClass
import org.jetbrains.exposed.dao.id.EntityID
import org.jetbrains.exposed.dao.id.LongIdTable
import org.jetbrains.exposed.sql.ReferenceOption
import org.jetbrains.exposed.sql.kotlin.datetime.datetime

object EventTable : LongIdTable("events") {
    val creator = reference("creator", UserTable, onDelete = ReferenceOption.CASCADE)
    val name = varchar("name", 255)
    val address = text("address").nullable()
    val latitude = double("latitude").nullable()
    val longitude = double("longitude").nullable()
    val capacity = integer("capacity").nullable()
    val description = text("description")
    val startTime = datetime("start_time")
    val endTime = datetime("end_time")
    val city = reference("city_id", CityTable)
}

object EventOrganizerTable : LongIdTable("event_organizer") {
    val organizer = reference("organizer_id", UserTable, onDelete = ReferenceOption.CASCADE)
    val event = reference("event_id", EventTable, onDelete = ReferenceOption.CASCADE)
}

object EventImageTable : LongIdTable("event_images") {
    val event = reference("event_id", EventTable, onDelete = ReferenceOption.CASCADE)
    val image = text("image")
}

class EventImageEntity(id: EntityID<Long>) : LongEntity(id) {
    companion object : LongEntityClass<EventImageEntity>(EventImageTable)
    var event by EventEntity referencedOn EventImageTable.event
    var image by EventImageTable.image
}

object EventInterestTable : LongIdTable("event_interests") {
    val event = reference("event_id", EventTable, onDelete = ReferenceOption.CASCADE)
    val interest = reference("interest_id", InterestTable, onDelete = ReferenceOption.CASCADE)
}


class EventEntity(id: EntityID<Long>) : LongEntity(id) {
    companion object : LongEntityClass<EventEntity>(EventTable)
    var creator by UserEntity referencedOn EventTable.creator
    var name by EventTable.name
    var address by EventTable.address
    var latitude by EventTable.latitude
    var longitude by EventTable.longitude
    var capacity by EventTable.capacity
    var description by EventTable.description
    var startTime by EventTable.startTime
    var endTime by EventTable.endTime
    var city by CityEntity referencedOn EventTable.city
    private val _images by EventImageEntity referrersOn EventImageTable.event
    val images get() = _images.map { it.image }
    var interests by InterestEntity via EventInterestTable
    var organizers by UserEntity via EventOrganizerTable
}