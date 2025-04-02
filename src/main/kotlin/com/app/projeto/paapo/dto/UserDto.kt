package com.app.projeto.paapo.dto

data class UserDto(
    val id: String,
    val username: String,
    val email: String,
    val roles: List<String>
)
