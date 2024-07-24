package kz.dostap.app

import io.ktor.http.*
import io.ktor.serialization.kotlinx.json.*
import io.ktor.server.application.*
import io.ktor.server.plugins.callid.*
import io.ktor.server.plugins.callloging.*
import io.ktor.server.plugins.contentnegotiation.*
import io.ktor.server.plugins.cors.routing.*
import kotlinx.serialization.json.Json
import org.slf4j.event.Level

fun main(args: Array<String>) = io.ktor.server.netty.EngineMain.main(args)


@Suppress("unused")
fun Application.module() {
    val app = this
    install(ContentNegotiation) {
        json(Json {
            prettyPrint = true
            ignoreUnknownKeys = true
        })
    }
    install(CallId)
    install(CallLogging) {
        level = Level.INFO
    }
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