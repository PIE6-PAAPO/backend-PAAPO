package com.app.projeto.paapo.dto

import jakarta.validation.constraints.Email
import jakarta.validation.constraints.NotBlank
import jakarta.validation.constraints.Size

data class RegisterRequest(

    @NotBlank
    @Size(min = 3, max = 20)
    val firstName: String,

    @NotBlank
    @Size(min = 3, max = 20)
    val lastName: String,

    @Size(min = 3, max = 50)
    @Email
    @NotBlank
    val email: String,

    @NotBlank
    @Size(min = 6, max = 40)
    val password: String,

    val role: Set<String> = setOf("user")

)