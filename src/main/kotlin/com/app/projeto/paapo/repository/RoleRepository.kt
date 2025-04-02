package com.app.projeto.paapo.repository

import com.app.projeto.paapo.model.RoleEnum
import com.app.projeto.paapo.model.Roles
import org.springframework.data.jpa.repository.JpaRepository
import org.springframework.stereotype.Repository
import java.util.Optional

@Repository
interface RoleRepository : JpaRepository<Roles, Long> {
    fun findByRole(name: RoleEnum): Optional<Roles>;
    fun existsByRole(role: RoleEnum): Boolean
}