package com.hyosakura.signin.sign.forum.discuz

import com.hyosakura.signin.util.OkHttpUtil

/**
 * @author LovesAsuna
 **/
class Huahuo(cookie: String) : Discuz(cookie) {
    override val name: String = "花火学园"
    override val baseUrl = "https://www.sayhuahuo.com/"

    override fun sign(): Boolean {
        return forumSign()
    }

    private fun forumSign(): Boolean {
        val signUrl =
            "${baseUrl}plugin.php?id=dsu_paulsign:sign&operation=qiandao&infloat=1&inajax=1"
        val response = OkHttpUtil.post(
            signUrl,
            mapOf("formhash" to formHash, "qdxq" to "kx", "qdmode" to "1", "todaysay" to "签到", "fastreply" to "0"),
            OkHttpUtil.addHeaders(cookie)
        )
        return getText(response, "div.c", "div.c", true)
    }
}