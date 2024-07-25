package kz.dostap.app.entities

import org.jetbrains.exposed.crypt.Algorithms
import org.jetbrains.exposed.crypt.encryptedVarchar
import org.jetbrains.exposed.dao.LongEntity
import org.jetbrains.exposed.dao.LongEntityClass
import org.jetbrains.exposed.dao.id.EntityID
import org.jetbrains.exposed.dao.id.LongIdTable
import org.jetbrains.exposed.sql.ReferenceOption

object UserTable : LongIdTable("users") {
    val firstName = varchar("first_name", 50)
    val lastName = varchar("last_name", 50)
    val email = text("email").uniqueIndex()
    val isEmailVerified = bool("is_email_verified").default(false)
    val username = text("username").uniqueIndex()
    val encryptedPassword = encryptedVarchar(
        "encrypted_password",
        255,
        Algorithms.AES_256_PBE_CBC(
            System.getenv("PASSWORD_ENCRYPTION_KEY"),
            System.getenv("PASSWORD_ENCRYPTION_SALT")
        )
    )
    val residenceCity = reference("residence_city_id", CityTable)
    val avatarLink = text("avatar_link").nullable()
    val gender = enumerationByName<Gender>("gender", 10).nullable()
    val age = integer("age").nullable()
    val phoneNumber = varchar("phone_number", 20).nullable()
    val description = text("description").nullable()
}

object UserInterestTable : LongIdTable("user_interests") {
    val user = reference("user_id", UserTable, onDelete = ReferenceOption.CASCADE)
    val interest = reference("interest_id", InterestTable, onDelete = ReferenceOption.CASCADE)
}

enum class Gender {
    MALE, FEMALE
}

class UserEntity(id: EntityID<Long>) : LongEntity(id) {
    companion object : LongEntityClass<UserEntity>(UserTable)
    var firstName by UserTable.firstName
    var lastName by UserTable.lastName
    var email by UserTable.email
    var isEmailVerified by UserTable.isEmailVerified
    var username by UserTable.username
    var encryptedPassword by UserTable.encryptedPassword
    var residenceCity by CityEntity referencedOn UserTable.residenceCity
    var residenceCityId by UserTable.residenceCity
    var avatarLink by UserTable.avatarLink
    var gender by UserTable.gender
    var age by UserTable.age
    var phoneNumber by UserTable.phoneNumber
    var description by UserTable.description

    val interests by InterestEntity via UserInterestTable
}