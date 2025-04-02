package com.app.projeto.paapo.service
import com.app.projeto.paapo.model.RoleEnum
import com.app.projeto.paapo.model.Roles
import com.app.projeto.paapo.repository.RoleRepository
import org.springframework.stereotype.Service
import java.util.Optional

@Service
class RoleService(private val roleRepository: RoleRepository) {

    fun findByName(name: RoleEnum): Optional<Roles> {
        return roleRepository.findByRole(name)
    }
}