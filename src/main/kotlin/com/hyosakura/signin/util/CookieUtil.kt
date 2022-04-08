package com.hyosakura.signin.util

import org.openqa.selenium.Cookie
import org.openqa.selenium.WebDriver.Options
import java.time.LocalDateTime
import java.time.ZoneOffset
import java.util.*

object CookieUtil {
    fun String.addCookie(options : Options) {
        this.split(";").forEach {
            val entry = it.split("=")
            options.addCookie(
                Cookie(
                    entry[0].trim(), entry[1].trim(), "bbs.zdfx.net", "/", Date.from(
                        LocalDateTime.now().plusDays(1).toInstant(
                            ZoneOffset.UTC
                        )
                    )
                )
            )
        }
    }
}