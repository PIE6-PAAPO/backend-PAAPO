package com.app.projeto.paapo.model

import jakarta.persistence.*

@Entity
@Table(name="roles")
data class Roles (
    @Id
    @GeneratedValue(strategy = GenerationType.UUID)
    val id: String = "",
    @Enumerated
    val role: Role
)

enum class Role{
    ADMIN, USER
}