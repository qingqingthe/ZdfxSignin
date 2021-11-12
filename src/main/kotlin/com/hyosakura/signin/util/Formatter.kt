package com.hyosakura.signin.util

/**
 * @author LovesAsuna
 **/
object Formatter {
    fun outlineFormat(signName: String, paddingChar: String) : String {
        val maxPadding = 20
        val length: Int = signName.length
        val padding = (maxPadding - length) / 2
        val hash = "#"
        val extra = if (length % 2 == 0) 1 else 0
        val leftPadding = "%" + padding + "s"
        val rightPadding = "%" + (padding - extra) + "s"
        val strFormat = "$leftPadding%s$rightPadding"
        val formattedString = String.format(strFormat, "", hash, "")
       return formattedString.replace(" ", paddingChar).replace(hash, signName)
    }
}