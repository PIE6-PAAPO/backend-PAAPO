package com.app.projeto.paapo.service

import com.app.projeto.paapo.model.UserDetailsImpl
import com.app.projeto.paapo.repository.UserRepository
import jakarta.transaction.Transactional
import org.springframework.security.core.userdetails.UserDetailsService
import org.springframework.security.core.userdetails.UsernameNotFoundException
import org.springframework.stereotype.Service

@Service
class UserDetailsServiceImpl(private val userRepository: UserRepository) : UserDetailsService {

    @Transactional
    override fun loadUserByUsername(username: String): UserDetailsImpl {
        val user = userRepository.findByUserName(username)
            .orElseThrow { UsernameNotFoundException("Usuário não encontrado: $username") }

        return UserDetailsImpl.build(user)
    }
}