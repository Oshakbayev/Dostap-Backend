package kz.dostap.app.plugins

import io.ktor.server.application.*
import io.ktor.server.plugins.callid.*
import io.ktor.server.plugins.callloging.*
import org.slf4j.event.Level

fun Application.call() {
    install(CallId)
    install(CallLogging) {
        level = Level.INFO
    }
}