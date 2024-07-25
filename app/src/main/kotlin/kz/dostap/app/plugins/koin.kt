package kz.dostap.app.plugins

import io.ktor.server.application.*
import kz.dostap.app.appModule
import org.koin.ktor.plugin.Koin

fun Application.koin() {
    install(Koin) {
        modules(appModule)
    }
}