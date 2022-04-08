package com.hyosakura.signin

import com.hyosakura.signin.sign.Sign
import com.hyosakura.signin.sign.forum.discuz.Huahuo
import com.hyosakura.signin.sign.forum.discuz.Zdfx
import kotlinx.coroutines.launch
import kotlinx.coroutines.runBlocking

fun main() = runBlocking {
    val forumList = listOf<Class<out Sign>>(Zdfx::class.java, Huahuo::class.java)
    for (forum in forumList) {
        val currentCookie = Environment.getCookie(forum.simpleName)
        val instance = forum.getConstructor(String::class.java).newInstance(currentCookie ?: "")
        if (currentCookie == null) {
            println("${instance.name}‘s cookie has not set, skip sign in action !")
        } else {
            launch {
                println("start perform signing in action for ${instance.name}!")
                val result = instance.sign()
                println("sign in action for ${instance.name} ${if (result) "success" else "failed"}!")
            }
        }
    }
}