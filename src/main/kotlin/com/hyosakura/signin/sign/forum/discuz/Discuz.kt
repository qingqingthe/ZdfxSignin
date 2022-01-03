package com.hyosakura.signin.sign.forum.discuz

import com.hyosakura.signin.sign.AbstractSign
import com.hyosakura.signin.util.Request
import io.ktor.client.statement.*
import kotlinx.coroutines.runBlocking
import org.jsoup.Jsoup

/**
 * @author LovesAsuna
 **/
abstract class Discuz(cookie: String) : AbstractSign(cookie) {
    protected val formHash: String by lazy {
        runBlocking {
            val base = Jsoup.parse(Request.get(baseUrl, headers = mapOf("Cookie" to cookie)).readText())
            val form = base.select("#scbar_form")
            form.select("input:nth-child(2)").attr("value")
        }
    }
}
