package com.hyosakura.signin.sign.forum.discuz

import com.hyosakura.signin.sign.Response
import com.hyosakura.signin.sign.Result
import com.hyosakura.signin.util.Request
import kotlinx.coroutines.delay
import kotlinx.coroutines.runBlocking
import org.openqa.selenium.By
import org.openqa.selenium.Cookie
import org.openqa.selenium.chrome.ChromeDriver
import org.openqa.selenium.chrome.ChromeOptions
import org.openqa.selenium.support.ui.ExpectedConditions
import org.openqa.selenium.support.ui.WebDriverWait
import java.time.Clock
import java.time.Duration
import java.time.LocalDateTime
import java.time.ZoneOffset
import java.util.*

/**
 * @author LovesAsuna
 **/
open class Zdfx(cookie: String) : Discuz(cookie) {
    override val name: String = "终点论坛"
    override val baseUrl = "https://bbs.zdfx.net/"

    override suspend fun sign(): Result {
        return listOf(lottery(cookie), forumSign(cookie))
    }

    private suspend fun lottery(cookie: String): Response {
        System.setProperty("webdriver.chrome.driver", "chromedriver")
        val option = ChromeOptions()
        option.setHeadless(true)
        val driver = ChromeDriver(option)
        driver.manage().window().maximize()
        driver.get("${baseUrl}k_misign-sign.html")
        cookie.split(";").forEach {
            val entry = it.split("=")
            driver.manage().addCookie(
                Cookie(
                    entry[0].trim(), entry[1].trim(), "bbs.zdfx.net", "/", Date.from(
                        LocalDateTime.now().plusDays(1).toInstant(
                            ZoneOffset.UTC
                        )
                    )
                )
            )
        }
         fun getWait(timeout: Duration): WebDriverWait {
            return WebDriverWait(
                driver,
                timeout,
                Duration.ofMillis(500L),
                Clock.systemDefaultZone()
            ) { duration ->
                runBlocking {
                    delay(duration.toMillis())
                }
            }
        }
        driver.navigate().refresh()
        driver.get("${baseUrl}plugin.php?id=yinxingfei_zzza:yaoyao")
        val button = driver.findElement(By.cssSelector(".num_box > .btn"))
        getWait(Duration.ofSeconds(5)).until(ExpectedConditions.elementToBeClickable(button))
        button.click()
        val res = driver.findElement(By.cssSelector("#res"))
        getWait(Duration.ofSeconds(5)).until(ExpectedConditions.textToBePresentInElement(res, "已经"))
        driver.quit()
        return true to res.text
    }

    private suspend fun forumSign(cookie: String): Response {
        val signUrl =
            "${baseUrl}k_misign-sign.html?operation=qiandao&format=global_usernav_extra&formhash=${formHash}&inajax=1&ajaxtarget=k_misign_topb"
        val response = Request.get(signUrl, headers = mapOf("Cookie" to cookie))
        return getText(response, "#fx_checkin_b", "root", true)
    }
}