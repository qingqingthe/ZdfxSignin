package com.hyosakura.signin.sign.forum.discuz

import com.hyosakura.signin.sign.AbstractSign
import com.hyosakura.signin.util.OkHttpUtil
import org.jsoup.Jsoup

/**
 * @author LovesAsuna
 **/
abstract class Discuz(cookie: String) : AbstractSign(cookie) {
    protected val formHash: String by lazy {
        val base = Jsoup.parse(OkHttpUtil[baseUrl, OkHttpUtil.addHeaders(cookie)].body!!.string())
        val form = base.select("#scbar_form")
        form.select("input:nth-child(2)").attr("value")
    }
}
