package com.hyosakura.signin.sign.forum.discuz

import com.hyosakura.signin.util.Formatter
import org.openqa.selenium.By
import org.openqa.selenium.WebDriver
import org.openqa.selenium.WebElement
import org.openqa.selenium.support.ui.ExpectedConditions
import java.time.Duration

/**
 * @author LovesAsuna
 **/
open class Zdfx(cookie: String) : Discuz(cookie) {
    override val name: String = "Zdfx"
    override val baseUrl = "https://bbs.zdfx.net/"

    override suspend fun sign(): Boolean {
        var result = true
        val driver = startDriver()
        logger.append(Formatter.outlineFormat(name, "=")).append("\n")
        phase("start signing action on $name") {
            result = result && forumSign(driver)
        }

        phase("start lottery action on $name") {
            result = result && lottery(driver)
        }
        driver.quit()
        logger.append(Formatter.outlineFormat("", "="))
        println(logger.toString())
        return result
    }

    private fun lottery(driver: WebDriver): Boolean {
        driver.get("${baseUrl}plugin.php?id=yinxingfei_zzza:yaoyao")
        val button = driver.findElement(By.cssSelector(".num_box > .btn"))
        val res: WebElement?
        val resText: String?
        try {
            res = driver.findElement(By.cssSelector("#res"))
            val originText = res.text
            driver.wait(Duration.ofSeconds(20)).until(ExpectedConditions.elementToBeClickable(button))
            button.click()
            driver.wait(Duration.ofSeconds(10)).until(ExpectedConditions.not(ExpectedConditions.textToBePresentInElement(res, originText)))
            resText = res.text
        } catch (e: Exception) {
            e.printStackTrace()
            logger.append("lottery failed")
            return false
        }
        logger.append(resText)
        return true
    }

    private fun forumSign(driver: WebDriver): Boolean {
        driver.get("${baseUrl}k_misign-sign.html")
        try {
            val list = driver.findElements(By.cssSelector("#JD_sign"))
            if (list.isEmpty()) {
                logger.append("already signed in!")
                return true
            }
            val button = list[0]
            // TODO get success text
            driver.wait(Duration.ofSeconds(20)).until(ExpectedConditions.elementToBeClickable(button))
            button.click()
        } catch (e: Exception) {
            e.printStackTrace()
            logger.append("sign in failed!")
            return false
        }
        return true
    }

    override fun prePhase(phase: String) {
        logger.append(phase).append("\n")
    }

    override fun postPhase(phase: String) {
        logger.append("\n")
    }
}