package com.hyosakura.signin.forum

/**
 * @author LovesAsuna
 **/
interface Sign {
    val name: String

    val baseUrl: String

    fun sign(cookie: String): Boolean
}