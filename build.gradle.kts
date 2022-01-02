import org.jetbrains.kotlin.gradle.tasks.KotlinCompile

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
    implementation("io.ktor:ktor-client-core:1.6.7")
    implementation("io.ktor:ktor-client-okhttp:1.6.7")
    implementation("org.jsoup:jsoup:1.14.3")
    implementation("com.fasterxml.jackson.module:jackson-module-kotlin:2.13.1")
    implementation("org.seleniumhq.selenium:selenium-java:4.1.1")
    implementation("org.seleniumhq.selenium:selenium-chrome-driver:4.1.1")
    testImplementation(kotlin("test"))
}

application {
    mainClass.set("com.hyosakura.signin.MainKt")
}

tasks.test {
    useJUnitPlatform()
}

tasks.withType<KotlinCompile>() {
    kotlinOptions.jvmTarget = "1.8"
}