package users

import (
	"backend-PAAPO/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup, service *Service) {
	r.POST("/login", func(c *gin.Context) {
		var loginDTO dto.LoginDTO
		if err := c.ShouldBindJSON(&loginDTO); err != nil {
			c.JSON(400, gin.H{"error": "Invalid input"})
			return
		}
		user, err := service.Login(loginDTO)
		if err != nil {
			c.JSON(401, gin.H{"error": "Invalid credentials"})
			return
		}

		accessToken, err := GenerateJWT(user.ID, user.Role)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to generate token"})
			return
		}

		c.JSON(200, gin.H{
			"access_token": accessToken,
			"expires_in":   900,
		})
	})

	r.POST("/register", func(c *gin.Context) {
		var registerDTO dto.RegisterDTO
		if err := c.ShouldBindJSON(&registerDTO); err != nil {
			c.JSON(400, gin.H{"error": "Invalid input"})
			return
		}
		newUser, err := service.Register(registerDTO)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		accessToken, err := GenerateJWT(newUser.ID, newUser.Role)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to generate token"})
			return
		}

		c.JSON(200, gin.H{
			"access_token": accessToken,
			"expires_in":   900,
		})
	})

	// Rota para solicitar recuperação de senha
	r.POST("/forgot-password", func(c *gin.Context) {
		// Pega o email do corpo da requisição
		var req dto.ForgotPasswordDTO
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
			return
		}

		// Chama o serviço para processar a solicitação
		err := service.SolicitarRecuperacaoSenha(req.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao processar solicitação"})
			return
		}

		// Retorna sucesso mesmo se o email não existir (por segurança)
		c.JSON(http.StatusOK, gin.H{
			"message": "Se o email existir, você receberá um código de recuperação",
		})
	})

	// Rota para resetar a senha com o código de recuperação
	r.POST("/reset-password", func(c *gin.Context) {
		// Pega os dados do corpo da requisição
		var req dto.ResetPasswordDTO
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
			return
		}

		// Valida se as senhas são iguais (já validado pelo binding, mas é bom ter)
		if req.NewPassword != req.ConfirmPassword {
			c.JSON(http.StatusBadRequest, gin.H{"error": "As senhas não conferem"})
			return
		}

		// Chama o serviço para confirmar a troca de senha
		err := service.ConfirmarTrocaSenha(req.Email, req.Code, req.NewPassword)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Retorna sucesso
		c.JSON(http.StatusOK, gin.H{
			"message": "Senha alterada com sucesso",
		})
	})
}
