package kz.dostap.app

import io.ktor.server.application.*
import kz.dostap.app.plugins.*

fun main(args: Array<String>) = io.ktor.server.netty.EngineMain.main(args)


@Suppress("unused")
fun Application.module() {
    json()
    call()
    cors()
    security()
    koin()
}

