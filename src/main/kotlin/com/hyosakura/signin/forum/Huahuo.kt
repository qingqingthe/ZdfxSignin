package com.hyosakura.signin.forum

import com.hyosakura.signin.util.OkHttpUtil

/**
 * @author LovesAsuna
 **/
class Huahuo : AbstractSign() {
    override val baseUrl = "https://www.sayhuahuo.com/"

    override fun sign(cookie: String): Boolean {
        return forumSign(cookie)
    }

    private fun forumSign(cookie: String): Boolean {
        val signUrl =
            "${baseUrl}plugin.php?id=dsu_paulsign:sign&operation=qiandao&infloat=1&inajax=1"
        val response = OkHttpUtil.post(
            signUrl,
            mapOf("formhash" to "eddae219", "qdxq" to "kx", "qdmode" to "1", "todaysay" to "签到", "fastreply" to "0"),
            OkHttpUtil.addHeaders(cookie)
        )
        return getText(response, "l", "div.c", 1)
    }
}