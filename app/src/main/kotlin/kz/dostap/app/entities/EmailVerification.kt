package kz.dostap.app.entities

import org.jetbrains.exposed.dao.UUIDEntity
import org.jetbrains.exposed.dao.UUIDEntityClass
import org.jetbrains.exposed.dao.id.EntityID
import org.jetbrains.exposed.dao.id.UUIDTable
import org.jetbrains.exposed.sql.kotlin.datetime.datetime
import java.util.*

object EmailVerificationTable : UUIDTable("email_verifications") {
    val user = reference("user_id", UserTable)
    val email = text("email")
    val secretCode = text("secret_code")
    val isUsed = bool("is_used").default(false)
    val expiresAt = datetime("expires_at")
}

class EmailVerificationEntity(id: EntityID<UUID>) : UUIDEntity(id) {
    companion object : UUIDEntityClass<EmailVerificationEntity>(EmailVerificationTable)
    var user by UserEntity referencedOn EmailVerificationTable.user
    var email by EmailVerificationTable.email
    var secretCode by EmailVerificationTable.secretCode
    var isUsed by EmailVerificationTable.isUsed
    var expiresAt by EmailVerificationTable.expiresAt
}