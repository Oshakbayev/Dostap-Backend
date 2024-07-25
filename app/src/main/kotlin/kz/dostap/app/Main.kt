package kz.dostap.app

import io.ktor.server.application.*
import io.ktor.server.routing.*
import kz.dostap.app.plugins.*
import kz.dostap.app.routes.auth
import kz.dostap.app.routes.users

fun main(args: Array<String>) = io.ktor.server.netty.EngineMain.main(args)


@Suppress("unused")
fun Application.module() {
    json()
    call()
    cors()
    security()
    koin()

    routing {
        route("api/v1") {
            route("auth") { auth() }
            route("users") { users() }
        }
    }
}

