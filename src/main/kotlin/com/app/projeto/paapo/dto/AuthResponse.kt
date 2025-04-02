package com.app.projeto.paapo.dto

data class AuthResponse (
    val token: String,
    val type: String = "Bearer",
    val id: String,
    val username: String,
    val email: String,
    val roles: List<String>
    )