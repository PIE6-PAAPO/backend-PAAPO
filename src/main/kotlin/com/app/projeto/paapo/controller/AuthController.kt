package com.app.projeto.paapo.controller

import com.app.projeto.paapo.dto.AuthResponse
import com.app.projeto.paapo.dto.LoginRequest
import com.app.projeto.paapo.dto.RegisterRequest
import com.app.projeto.paapo.model.UserDetailsImpl
import com.app.projeto.paapo.security.JwtTokenProvider
import com.app.projeto.paapo.service.UserService
import jakarta.validation.Valid
import org.springframework.http.ResponseEntity
import org.springframework.security.authentication.AuthenticationManager
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken
import org.springframework.security.core.context.SecurityContextHolder
import org.springframework.web.bind.annotation.PostMapping
import org.springframework.web.bind.annotation.RequestBody
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RestController

@RestController
@RequestMapping("/api/auth")
class AuthController(
    private val authenticationManager: AuthenticationManager,
    private val userService: UserService,
    private val jwtTokenProvider: JwtTokenProvider
) {

    @PostMapping("/signin")
    fun authenticateUser(@Valid @RequestBody loginRequest: LoginRequest): ResponseEntity<AuthResponse> {
        val authentication = authenticationManager.authenticate(
            UsernamePasswordAuthenticationToken(loginRequest.username, loginRequest.password)
        )

        SecurityContextHolder.getContext().authentication = authentication
        val jwt = jwtTokenProvider.generateJwtToken(authentication)
        val userDetails = authentication.principal as UserDetailsImpl

        val roles = userDetails.authorities.map { it.authority }

        return ResponseEntity.ok(
            AuthResponse(
                token = jwt,
                id = userDetails.id,
                username = userDetails.username,
                email = userDetails.email,
                roles = roles
            )
        )
    }

    @PostMapping("/signup")
    fun registerUser(@Valid @RequestBody signUpRequest: RegisterRequest): ResponseEntity<*> {
        val user = userService.createUser(signUpRequest)

        return ResponseEntity.ok("Usuário registrado com sucesso!")
    }
}