package com.hyosakura.signin

import com.hyosakura.signin.forum.Huahuo
import com.hyosakura.signin.forum.Sign
import com.hyosakura.signin.forum.Zdfx
import com.hyosakura.signin.util.Formatter

fun main() {
    val forumList = listOf<Class<out Sign>>(Zdfx::class.java, Huahuo::class.java)
    for (forum in forumList) {
        val instance = forum.getConstructor().newInstance()
        val upLine = Formatter.outlineFormat(instance.name, "=")
        println(upLine)
        val currentCookie = Environment.getCookie(forum.simpleName)
        if (currentCookie == null) {
            println("未设置${instance.name}的COOKIE,跳过此论坛的操作!")
        } else {
            instance.sign(currentCookie)
        }
        val downLine = Formatter.outlineFormat("", "=")
        println(downLine)
        println()
    }
}