package com.app.projeto.paapo.security

import com.fasterxml.jackson.annotation.JsonIgnore
import io.jsonwebtoken.*
import io.jsonwebtoken.io.Decoders
import io.jsonwebtoken.security.Keys
import org.slf4j.LoggerFactory
import org.springframework.beans.factory.annotation.Value
import org.springframework.security.core.Authentication
import org.springframework.security.core.userdetails.UserDetails
import org.springframework.stereotype.Component
import java.util.*
import javax.crypto.SecretKey

@Component
class JwtTokenProvider {

    private val logger = LoggerFactory.getLogger(JwtTokenProvider::class.java)

    @Value("\${app.jwt.secret}")
    private lateinit var jwtSecret: String

    @Value("\${app.jwt.expiration:3600000}")
    private var jwtExpirationMs: Int = 0

    fun generateJwtToken(authentication: Authentication): String {
        val userPrincipal = authentication.principal as UserDetails

        return Jwts.builder()
            .setSubject(userPrincipal.username)
            .setIssuedAt(Date())
            .setExpiration(Date(Date().time + jwtExpirationMs))
            .signWith(key(), SignatureAlgorithm.HS256)
            .compact()
    }

    private fun key(): SecretKey {
        val keyBytes = try {
            Decoders.BASE64.decode(jwtSecret)
        } catch (e: IllegalArgumentException) {
            logger.error("Erro na decodificação da chave secreta: Certifique-se de usar uma chave Base64 válida")
            throw JwtConfigurationException("Chave JWT inválida", e)
        }
        return Keys.hmacShaKeyFor(keyBytes)
    }

    fun getUserNameFromJwtToken(token: String): String {
        return Jwts.parserBuilder()
            .setSigningKey(key())
            .build()
            .parseClaimsJws(token)
            .body
            .subject
    }

    fun validateJwtToken(authToken: String): Boolean {
        try {
            Jwts.parserBuilder().setSigningKey(key()).build().parseClaimsJws(authToken)
            return true
        } catch (e: SecurityException) {
            logger.error("Assinatura JWT inválida: {}", e.message, e)
        } catch (e: MalformedJwtException) {
            logger.error("Token JWT inválido: {}", e.message, e)
        } catch (e: ExpiredJwtException) {
            logger.error("Token JWT expirado: {}", e.message, e)
        } catch (e: UnsupportedJwtException) {
            logger.error("Token JWT não suportado: {}", e.message, e)
        } catch (e: IllegalArgumentException) {
            logger.error("JWT claims string está vazia: {}", e.message, e)
        }
        return false
    }
}

class JwtConfigurationException(message: String, cause: Throwable) : RuntimeException(message, cause)