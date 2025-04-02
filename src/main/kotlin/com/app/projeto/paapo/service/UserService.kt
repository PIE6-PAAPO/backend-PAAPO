package com.app.projeto.paapo.service


import com.app.projeto.paapo.dto.RegisterRequest
import com.app.projeto.paapo.dto.UserDto
import com.app.projeto.paapo.model.RoleEnum
import com.app.projeto.paapo.model.Roles
import com.app.projeto.paapo.model.User
import com.app.projeto.paapo.repository.UserRepository
import org.springframework.security.crypto.password.PasswordEncoder
import org.springframework.stereotype.Service
import org.springframework.transaction.annotation.Transactional
import java.util.*

@Service
class UserService(
    private val userRepository: UserRepository,
    private val passwordEncoder: PasswordEncoder,
    private val roleService: RoleService
) {

    @Transactional
    fun createUser(signUpRequest: RegisterRequest): User {
        if (userRepository.existsByUserName(signUpRequest.email)) {
            throw RuntimeException("Erro: Nome de usuário já está em uso!")
        }

        if (userRepository.existsByEmail(signUpRequest.email)) {
            throw RuntimeException("Erro: Email já está em uso!")
        }

        val user = User(
            userName = signUpRequest.email,
            email = signUpRequest.email,
            firstName = signUpRequest.firstName,
            lastName = signUpRequest.lastName,
            passWord = passwordEncoder.encode(signUpRequest.password)
        )

        val roles = mutableSetOf<Roles>()

        if (signUpRequest.role.isEmpty()) {
            val userRole = roleService.findByName(RoleEnum.USER)
                .orElseThrow { RuntimeException("Erro: Role não encontrada.") }
            roles.add(userRole)
        } else {
            signUpRequest.role.forEach { role ->
                when (role.lowercase()) {
                    "admin" -> {
                        val adminRole = roleService.findByName(RoleEnum.ADMIN)
                            .orElseThrow { RuntimeException("Erro: Role não encontrada.") }
                        roles.add(adminRole)
                    }
                    else -> {
                        val userRole = roleService.findByName(RoleEnum.USER)
                            .orElseThrow { RuntimeException("Erro: Role não encontrada.") }
                        roles.add(userRole)
                    }
                }
            }
        }

        val userWithRoles = user.copy(roles = roles)
        return userRepository.save(userWithRoles)
    }

    fun getUserById(id: Long): UserDto {
        val user = userRepository.findById(id)
            .orElseThrow { RuntimeException("Usuário não encontrado com id: $id") }

        return UserDto(
            id = user.id,
            username = user.email,
            email = user.email,
            roles = user.roles.map { it.role.name }
        )
    }
}