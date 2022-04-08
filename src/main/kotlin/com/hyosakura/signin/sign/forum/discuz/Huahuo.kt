package com.hyosakura.signin.sign.forum.discuz

import com.hyosakura.signin.util.Formatter
import com.hyosakura.signin.util.Request

/**
 * @author LovesAsuna
 **/
class Huahuo(cookie: String) : Discuz(cookie) {

    override val name: String = "Huahuo"
    override val baseUrl = "https://www.sayhuahuo.com/"

    override suspend fun sign(): Boolean {
        var result = true
        logger.append(Formatter.outlineFormat(name, "=")).append("\n")
        phase("start signing action on $name") {
            result = result &&  forumSign()
        }
        logger.append(Formatter.outlineFormat("", "="))
        return result
    }

    private suspend fun forumSign(): Boolean {
        val signUrl =
            "${baseUrl}plugin.php?id=dsu_paulsign:sign&operation=qiandao&infloat=1&inajax=1"
        val response = Request.submitForm(
            signUrl,
            mapOf("formhash" to formHash, "qdxq" to "kx", "qdmode" to "1", "todaysay" to "签到", "fastreply" to "0"),
            headers = mapOf("Cookie" to cookie)
        )
        val result = getText(response, "div.c", "div.c", true)
        logger.append(result.second)
        println(logger.toString())
        return result.first
    }

    override fun prePhase(phase: String) {
        logger.append(phase).append("\n")
    }

    override fun postPhase(phase: String) {
        logger.append("\n")
    }
}