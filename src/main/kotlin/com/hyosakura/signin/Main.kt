package com.hyosakura.signin

import com.hyosakura.signin.sign.forum.discuz.Huahuo
import com.hyosakura.signin.sign.Sign
import com.hyosakura.signin.sign.forum.discuz.Zdfx
import com.hyosakura.signin.util.Formatter

fun main() {
    val forumList = listOf<Class<out Sign>>(Zdfx::class.java, Huahuo::class.java)
    for (forum in forumList) {
        val currentCookie = Environment.getCookie(forum.simpleName)
        val instance = forum.getConstructor(String::class.java).newInstance(currentCookie)
        val upLine = Formatter.outlineFormat(instance.name, "=")
        println(upLine)
        if (currentCookie == null) {
            println("未设置${instance.name}的COOKIE,跳过此论坛的操作!")
        } else {
            instance.sign()
        }
        val downLine = Formatter.outlineFormat("", "=")
        println(downLine)
        println()
    }
}