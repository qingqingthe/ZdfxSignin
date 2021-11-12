package com.hyosakura.signin

object Environment {
    lateinit var cookieMap : Map<String, String>

    private fun loadCookie() {
        cookieMap = System.getenv().mapKeys {
            it.key.uppercase()
        }.filter {
            it.key.endsWith("_COOKIE")
        }
    }

    fun getCookie(key: String) : String? {
        return cookieMap["${key.uppercase()}_COOKIE"]
    }

    init {
        loadCookie()
    }
}