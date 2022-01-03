package com.hyosakura.signin.sign

/**
 * @author LovesAsuna
 **/
interface Sign {
    val name: String

    val baseUrl: String

    val cookie: String

    suspend fun sign(): Result
}

typealias Response = Pair<Boolean, String>
typealias Result = List<Response>