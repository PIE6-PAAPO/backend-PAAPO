package com.app.projeto.paapo.controller

import com.app.projeto.paapo.dto.UserDto
import com.app.projeto.paapo.service.UserService
import org.springframework.http.ResponseEntity
import org.springframework.security.access.prepost.PreAuthorize
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.PathVariable
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RestController

@RestController
@RequestMapping("/api/users")
class UserController(private val userService: UserService) {

    @GetMapping("/{id}")
    @PreAuthorize("hasRole('USER') or hasRole('ADMIN')")
    fun getUserById(@PathVariable id: Long): ResponseEntity<UserDto> {
        val userDto = userService.getUserById(id)
        return ResponseEntity.ok(userDto)
    }
}