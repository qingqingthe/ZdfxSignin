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
        depth: Int = 0
    ): Boolean {
        val html = response.body?.string()
        if (html == null) {
            println("请求失败")
            return false
        }
        return getText(html, successCssSelector, failCssSelector, depth)
    }

    private fun getText(
        html: String,
        successCssSelector: String,
        failCssSelector: String,
        depth: Int = 0
    ): Boolean {
        val successElement: Elements = Jsoup.parse(html).select(successCssSelector)
        return if (successElement.isEmpty()) {
            val failElement: Elements = Jsoup.parse(html).select(failCssSelector)
            if (failElement.isEmpty()) {
                if (depth == 0) {
                    println("无法解析HTML!")
                }
                getText(Jsoup.parse(html).text(), successCssSelector, failCssSelector, depth - 1)
            } else {
                println(failElement.text())
                true
            }
        } else {
            println(successElement.text())
            true
        }
    }
}