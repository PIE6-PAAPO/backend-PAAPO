package com.app.projeto.paapo.repository

import com.app.projeto.paapo.model.User
import org.springframework.data.jpa.repository.JpaRepository
import org.springframework.stereotype.Repository
import java.util.*

@Repository
interface UserRepository : JpaRepository<User, Long> {
    fun findByUserName(username: String): Optional<User>
    fun existsByUserName(username: String): Boolean
    fun existsByEmail(email: String): Boolean
}