package com.hyosakura.signin.sign

/**
 * @author LovesAsuna
 **/
abstract class AbstractSign(override val cookie: String) : Sign {
    inline fun phase(phase: String, block: (String) -> Unit) {
        prePhase(phase)
        block(phase)
        postPhase(phase)
    }

    abstract fun prePhase(phase: String)

    abstract fun postPhase(phase: String)
}