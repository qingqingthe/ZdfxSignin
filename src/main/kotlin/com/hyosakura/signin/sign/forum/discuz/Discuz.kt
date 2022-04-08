package com.hyosakura.signin.sign.forum.discuz

import com.hyosakura.signin.sign.AbstractSign
import com.hyosakura.signin.util.CookieUtil.addCookie
import com.hyosakura.signin.util.Request
import io.ktor.client.statement.*
import kotlinx.coroutines.delay
import kotlinx.coroutines.runBlocking
import org.jsoup.Jsoup
import org.jsoup.select.Elements
import org.openqa.selenium.WebDriver
import org.openqa.selenium.chrome.ChromeDriver
import org.openqa.selenium.chrome.ChromeOptions
import org.openqa.selenium.support.ui.WebDriverWait
import java.time.Clock
import java.time.Duration

/**
 * @author LovesAsuna
 **/
abstract class Discuz(cookie: String) : AbstractSign(cookie) {
    init {
        System.setProperty("webdriver.chrome.driver", "C:\\SeleniumWebDrivers\\ChromeDriver\\chromedriver.exe")
    }

    protected val logger = Logger()

    protected val formHash: String by lazy {
        runBlocking {
            val base = Jsoup.parse(Request.get(baseUrl, headers = mapOf("Cookie" to cookie)).readText())
            val form = base.select("#scbar_form")
            form.select("input:nth-child(2)").attr("value")
        }
    }

    suspend fun getText(
        response: HttpResponse,
        successCssSelector: String,
        failCssSelector: String,
        hasCDATA: Boolean = false
    ): Pair<Boolean, String> {
        val html = response.readText()
        return getText(html, successCssSelector, failCssSelector, hasCDATA)
    }

    fun startDriver(): WebDriver {
        val option = ChromeOptions()
        val driver = ChromeDriver(option)
        driver.manage().window().maximize()
        driver.get(baseUrl)
        cookie.addCookie(driver.manage())
        driver.navigate().refresh()
        return driver
    }

    fun WebDriver.wait(timeout: Duration): WebDriverWait {
        return WebDriverWait(
            this,
            timeout,
            Duration.ofMillis(500L),
            Clock.systemDefaultZone()
        ) { duration ->
            runBlocking {
                delay(duration.toMillis())
            }
        }
    }

    private fun getText(
        html: String,
        successCssSelector: String,
        failCssSelector: String,
        hasCDATA: Boolean,
    ): Pair<Boolean, String> {
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
                    true to formatHtml
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

typealias Logger = StringBuilder
