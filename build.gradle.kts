plugins {
    kotlin("jvm") version "1.6.10"
    id("application")
}

group = "com.hyosakura"
version = "1.0-SNAPSHOT"

repositories {
    mavenCentral()
}

dependencies {
    implementation("io.ktor:ktor-client-core:1.6.8")
    implementation("io.ktor:ktor-client-okhttp:1.6.8")
    implementation("org.jsoup:jsoup:1.14.3")
    implementation("com.fasterxml.jackson.module:jackson-module-kotlin:2.13.2")
    implementation("org.seleniumhq.selenium:selenium-java:4.1.3")
    implementation("org.seleniumhq.selenium:selenium-chrome-driver:4.1.3")
    testImplementation(kotlin("test"))
}

application {
    mainClass.set("com.hyosakura.signin.MainKt")
    applicationDefaultJvmArgs = listOf("-Dfile.encoding=UTF-8")
}

tasks.test {
    useJUnitPlatform()
}

tasks.compileKotlin {
    kotlinOptions.jvmTarget = "1.8"
}