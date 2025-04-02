package com.app.projeto.paapo.dto

import jakarta.validation.constraints.NotBlank

data class LoginRequest(
    @NotBlank
    val username: String,

    @NotBlank
    val password: String
)
