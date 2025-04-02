package com.app.projeto.paapo.model

import com.fasterxml.jackson.annotation.JsonIgnore
import org.springframework.security.core.GrantedAuthority
import org.springframework.security.core.authority.SimpleGrantedAuthority
import org.springframework.security.core.userdetails.UserDetails

class UserDetailsImpl(
    val id: String,
    private val username: String,
    private val firstName: String,
    private val lastName: String,
    val email: String,
    @JsonIgnore
    private val password: String,
    private val authorities: Collection<GrantedAuthority>
) : UserDetails {

    override fun getAuthorities(): Collection<GrantedAuthority> = authorities

    override fun getPassword(): String = password

    override fun getUsername(): String = username

    override fun isAccountNonExpired(): Boolean = true

    override fun isAccountNonLocked(): Boolean = true

    override fun isCredentialsNonExpired(): Boolean = true

    override fun isEnabled(): Boolean = true

    companion object {
        fun build(user: User): UserDetailsImpl {
            val authorities = user.roles.map { role ->
                SimpleGrantedAuthority(role.role.name)
            }

            return UserDetailsImpl(
                id = user.id,
                username = user.userName,
                firstName = user.firstName,
                lastName = user.lastName,
                email = user.email,
                password = user.passWord,
                authorities = authorities
            )
        }
    }
}