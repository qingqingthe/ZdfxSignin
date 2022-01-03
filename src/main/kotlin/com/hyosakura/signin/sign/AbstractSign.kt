package com.hyosakura.signin.sign

import io.ktor.client.statement.*
import org.jsoup.Jsoup
import org.jsoup.select.Elements

/**
 * @author LovesAsuna
 **/
abstract class AbstractSign(override val cookie : String) : Sign {
    suspend fun getText(
        response: HttpResponse,
        successCssSelector: String,
        failCssSelector: String,
        hasCDATA: Boolean = false
    ): Response {
        val html = response.readText()
        return getText(html, successCssSelector, failCssSelector, hasCDATA)
    }

    private fun getText(
        html: String,
        successCssSelector: String,
        failCssSelector: String,
        hasCDATA: Boolean,
    ): Response {
        val formatHtml = if (hasCDATA) {
            Jsoup.parse(html).select("root").text()
        } else {
            html
        }
        val successElement: Elements = Jsoup.parse(formatHtml).select(successCssSelector)
        return if (successElement.isEmpty()) {
            val failElement: Elements = Jsoup.parse(formatHtml).select(failCssSelector)
            if (failElement.isEmpty()) {
                if (formatHtml.isEmpty()) {
                    false to "无法解析HTML!"
                } else {
                    true to  formatHtml
                }
            } else {
                true to failElement.trueText()
            }
        } else {
            true to successElement.trueText()
        }
    }

    private fun Elements.trueText(): String {
        val builder = StringBuilder()
        forEach {
            if (it.`is`("img")) {
                builder.append(it.attr("alt"))
            } else {
                builder.append(it.text())
            }
        }
        return builder.toString()
    }
}