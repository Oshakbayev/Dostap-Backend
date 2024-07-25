package kz.dostap.app.entities

import kotlinx.datetime.Clock
import org.jetbrains.exposed.dao.UUIDEntity
import org.jetbrains.exposed.dao.UUIDEntityClass
import org.jetbrains.exposed.dao.id.EntityID
import org.jetbrains.exposed.dao.id.UUIDTable
import org.jetbrains.exposed.sql.and
import org.jetbrains.exposed.sql.kotlin.datetime.timestamp
import java.util.*

object EmailVerificationTable : UUIDTable("email_verifications") {
    val user = reference("user_id", UserTable)
    val email = text("email")
    val secretCode = text("secret_code")
    val isUsed = bool("is_used").default(false)
    val expiresAt = timestamp("expires_at")
}

class EmailVerificationEntity(id: EntityID<UUID>) : UUIDEntity(id) {
    companion object : UUIDEntityClass<EmailVerificationEntity>(EmailVerificationTable) {
        fun findValidBySecretCode(secretCode: String): EmailVerificationEntity? = find {
            (EmailVerificationTable.secretCode eq secretCode) and
                    (EmailVerificationTable.isUsed eq false) and
                    (EmailVerificationTable.expiresAt greaterEq Clock.System.now())
        }.firstOrNull()
    }
    var user by UserEntity referencedOn EmailVerificationTable.user
    var email by EmailVerificationTable.email
    var secretCode by EmailVerificationTable.secretCode
    var isUsed by EmailVerificationTable.isUsed
    var expiresAt by EmailVerificationTable.expiresAt
}