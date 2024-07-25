package kz.dostap.app.routes

import arrow.core.Either
import io.ktor.http.*
import io.ktor.server.application.*
import io.ktor.server.request.*
import io.ktor.server.response.*
import io.ktor.server.routing.*
import kz.dostap.app.models.UserLoginRequest
import kz.dostap.app.models.UserSignUpRequest
import kz.dostap.app.services.AuthService
import org.koin.ktor.ext.inject

fun Route.auth() {
    val authService: AuthService by inject()

    post("login") {
        val request = call.receive<UserLoginRequest>()
        when (val token = authService.login(request.username, request.password)) {
            is Either.Left -> call.respond(HttpStatusCode.Unauthorized, token.value)
            is Either.Right -> call.respond(token.value)
        }
    }

    post("signup") {
        val request = call.receive<UserSignUpRequest>()
        when (val token = authService.signup(request)) {
            is Either.Left -> call.respond(HttpStatusCode.BadRequest, token.value)
            is Either.Right -> call.respond(token.value)
        }
    }

    get("verify") {
        val code = call.request.queryParameters["code"] ?: throw IllegalArgumentException("code is required")
        if (authService.confirmEmail(code)) {
            call.respond(HttpStatusCode.OK, "Your verification is successful")
        } else {
            call.respond(HttpStatusCode.NotFound)
        }
    }
}