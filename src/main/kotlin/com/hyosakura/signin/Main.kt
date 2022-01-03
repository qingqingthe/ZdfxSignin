package com.hyosakura.signin

import com.hyosakura.signin.sign.Sign
import com.hyosakura.signin.sign.forum.discuz.Huahuo
import com.hyosakura.signin.sign.forum.discuz.Zdfx
import com.hyosakura.signin.util.Formatter
import com.hyosakura.signin.util.Request
import kotlinx.coroutines.launch
import kotlinx.coroutines.runBlocking

fun main() = runBlocking {
    val forumList = listOf<Class<out Sign>>(Zdfx::class.java, Huahuo::class.java)
    for (forum in forumList) {
        val currentCookie = Environment.getCookie(forum.simpleName)
        val instance = forum.getConstructor(String::class.java).newInstance(currentCookie ?: "")
        if (currentCookie == null) {
            println("未设置${instance.name}的COOKIE,跳过此论坛的操作!")
        } else {
            launch {
                println("开始执行${instance.name}的签到操作!")
                val builder = StringBuilder()
                val upLine = Formatter.outlineFormat(instance.name, "=")
                builder.append(upLine).append("\n")
                val result = instance.sign()
                result.forEach {response->
                    if (response.first) {
                        builder.append(response.second).append("\n")
                    }
                }
                val downLine = Formatter.outlineFormat("", "=")
                builder.append(downLine)
                builder.append("\n")
                println(builder.toString())
            }
        }
    }
}