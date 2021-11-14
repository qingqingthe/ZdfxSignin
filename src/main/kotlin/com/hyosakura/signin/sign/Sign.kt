package com.hyosakura.signin.sign

/**
 * @author LovesAsuna
 **/
interface Sign {
    val name: String

    val baseUrl: String

    val cookie: String

    fun sign(): Boolean
}