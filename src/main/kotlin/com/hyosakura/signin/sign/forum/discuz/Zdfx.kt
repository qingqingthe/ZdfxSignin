package com.hyosakura.signin.sign.forum.discuz

import com.hyosakura.signin.sign.Response
import com.hyosakura.signin.sign.Result
import com.hyosakura.signin.util.Request

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
        val lotteryUrl = "${baseUrl}plugin.php?id=yinxingfei_zzza:yaoyao"
        val response =
            Request.submitForm(lotteryUrl, mapOf("formhash" to formHash), headers = mapOf("Cookie" to cookie))
        return getText(response, "#res", "#res")
    }

    private suspend fun forumSign(cookie: String): Response {
        val signUrl =
            "${baseUrl}k_misign-sign.html?operation=qiandao&format=global_usernav_extra&formhash=${formHash}&inajax=1&ajaxtarget=k_misign_topb"
        val response = Request.get(signUrl, headers = mapOf("Cookie" to cookie))
        return getText(response, "#fx_checkin_b", "root", true)
    }
}