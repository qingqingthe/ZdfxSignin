package com.hyosakura.signin.sign.forum.discuz

import com.hyosakura.signin.util.OkHttpUtil

/**
 * @author LovesAsuna
 **/
open class Zdfx(cookie: String) : Discuz(cookie) {
    override val name: String = "终点论坛"
    override val baseUrl = "https://bbs.zdfx.net/"

    override fun sign(): Boolean {
        return lottery(cookie) && forumSign(cookie)
    }

    private fun lottery(cookie: String): Boolean {
        val lotteryUrl = "${baseUrl}plugin.php?id=yinxingfei_zzza:yinxingfei_zzza_post"
        val response = OkHttpUtil.post(lotteryUrl, mapOf("formhash" to formHash), OkHttpUtil.addHeaders(cookie))
        return getText(response, ".zzza_hall_top_left_infor", "#messagetext > p:first-child")
    }

    private fun forumSign(cookie: String): Boolean {
        val signUrl =
            "${baseUrl}k_misign-sign.html?operation=qiandao&format=global_usernav_extra&formhash=${formHash}&inajax=1&ajaxtarget=k_misign_topb"
        val response = OkHttpUtil[signUrl, OkHttpUtil.addHeaders(cookie)]
        return getText(response, "#fx_checkin_b", "root", true)
    }
}