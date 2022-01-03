package com.hyosakura.signin.sign.forum.discuz

import com.hyosakura.signin.sign.Response
import com.hyosakura.signin.sign.Result
import com.hyosakura.signin.util.Request

/**
 * @author LovesAsuna
 **/
class Huahuo(cookie: String) : Discuz(cookie) {
    override val name: String = "花火学园"
    override val baseUrl = "https://www.sayhuahuo.com/"

    override suspend fun sign(): Result {
        return listOf(forumSign())
    }

    private suspend fun forumSign(): Response {
        val signUrl =
            "${baseUrl}plugin.php?id=dsu_paulsign:sign&operation=qiandao&infloat=1&inajax=1"
        val response = Request.submitForm(
            signUrl,
            mapOf("formhash" to formHash, "qdxq" to "kx", "qdmode" to "1", "todaysay" to "签到", "fastreply" to "0"),
            headers = mapOf("Cookie" to cookie)
        )
        return getText(response, "div.c", "div.c", true)
    }
}