package kz.dostap.app.routes

import arrow.core.Either
import io.ktor.http.*
import io.ktor.server.application.*
import io.ktor.server.auth.*
import io.ktor.server.request.*
import io.ktor.server.response.*
import io.ktor.server.routing.*
import kz.dostap.app.models.UserUpdateRequest
import kz.dostap.app.plugins.UserPrincipal
import kz.dostap.app.services.UserService
import org.koin.ktor.ext.inject

fun Route.users() {
    val userService: UserService by inject()

    authenticate {
        route("me") {
            get {
                val principal = call.authentication.principal<UserPrincipal>()!!
                val user = userService.getUser(principal.id)
                if (user == null) {
                    call.respond(HttpStatusCode.NotFound)
                } else {
                    call.respond(user)
                }
            }

            put {
                val principal = call.authentication.principal<UserPrincipal>()!!
                val request = call.receive<UserUpdateRequest>()
                when (val result = userService.updateUser(principal.id, request)) {
                    is Either.Left -> call.respond(HttpStatusCode.BadRequest, result.value)
                    is Either.Right -> call.respond(result.value)
                }
            }
        }
    }


}