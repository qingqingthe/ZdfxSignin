package com.hyosakura.signin.forum

/**
 * @author LovesAsuna
 **/
interface Sign {
    val baseUrl: String

    fun sign(cookie: String): Boolean
}