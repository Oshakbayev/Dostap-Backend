package kz.dostap.app.entities

import org.jetbrains.exposed.dao.LongEntity
import org.jetbrains.exposed.dao.LongEntityClass
import org.jetbrains.exposed.dao.id.EntityID
import org.jetbrains.exposed.dao.id.LongIdTable

object FriendRequestTable : LongIdTable("friend_requests") {
    val userId = reference("user_id", UserTable)
    val friendId = reference("friend_id", UserTable)
    val status = enumerationByName<Status>("status", 10).index()
}

enum class Status {
    PENDING, ACCEPTED, REJECTED
}

class FriendRequestEntity(id: EntityID<Long>) : LongEntity(id) {
    companion object : LongEntityClass<FriendRequestEntity>(FriendRequestTable)
    var user by UserEntity referencedOn FriendRequestTable.userId
    var friend by UserEntity referencedOn FriendRequestTable.friendId
    var status by FriendRequestTable.status
}