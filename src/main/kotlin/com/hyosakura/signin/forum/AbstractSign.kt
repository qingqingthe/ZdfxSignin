package com.hyosakura.signin.forum

import okhttp3.Response
import org.jsoup.Jsoup
import org.jsoup.select.Elements

/**
 * @author LovesAsuna
 **/
abstract class AbstractSign : Sign {
    fun getText(
        response: Response,
        successCssSelector: String,
        failCssSelector: String,
        hasCDATA: Boolean = false
    ): Boolean {
        val html = response.body?.string()
        if (html == null) {
            println("请求失败")
            return false
        }
        return getText(html, successCssSelector, failCssSelector, hasCDATA)
    }

    private fun getText(
        html: String,
        successCssSelector: String,
        failCssSelector: String,
        hasCDATA: Boolean,
    ): Boolean {
        val formatHtml = if (hasCDATA) {
            Jsoup.parse(html).select("root").text()
        } else {
            html
        }
        val successElement: Elements = Jsoup.parse(formatHtml).select(successCssSelector)
        return if (successElement.isEmpty()) {
            val failElement: Elements = Jsoup.parse(formatHtml).select(failCssSelector)
            if (failElement.isEmpty()) {
                println("无法解析HTML!")
                false
            } else {
                failElement.print()
                true
            }
        } else {
            successElement.print()
            true
        }
    }

    private fun Elements.print() {
        forEach {
            if (it.`is`("img")) {
                println(it.attr("alt"))
            } else {
                println(it.text())
            }
        }
    }
}