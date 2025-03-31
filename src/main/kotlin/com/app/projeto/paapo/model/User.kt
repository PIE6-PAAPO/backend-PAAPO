package com.app.projeto.paapo.model

import jakarta.persistence.*
import jakarta.validation.constraints.Email
import jakarta.validation.constraints.NotBlank

@Entity
@Table(name = "users")
data class User(
    @Id
    @GeneratedValue(strategy = GenerationType.UUID)
    val id: String = "",
    @NotBlank
    val userName: String = "",
    @NotBlank
    @Email
    val email: String = "",
    @NotBlank
    val passWord: String = "",
    @ManyToMany(fetch = FetchType.EAGER)
    @JoinTable(name = "user_roles",
        joinColumns = [JoinColumn(name = "user_id")],
        inverseJoinColumns = [JoinColumn(name = "role_id")])
    val roles: Set<Role> = HashSet()
)
