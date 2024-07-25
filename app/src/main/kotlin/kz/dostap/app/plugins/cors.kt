package kz.dostap.app.plugins

import io.ktor.http.*
import io.ktor.server.application.*
import io.ktor.server.plugins.cors.routing.*

fun Application.cors() {
    val app = this
    install(CORS) {
        if (app.environment.developmentMode) {
            anyHost()
            allowHeaders { true }
        }
        HttpMethod.DefaultMethods.forEach {
            allowMethod(it)
        }
    }
}