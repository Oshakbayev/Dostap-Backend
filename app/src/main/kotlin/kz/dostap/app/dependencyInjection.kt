package kz.dostap.app

import kz.dostap.app.services.AuthService
import kz.dostap.app.services.impl.AuthServiceImpl
import org.koin.dsl.module

val appModule = module {
    single<AuthService> { AuthServiceImpl(get()) }
}