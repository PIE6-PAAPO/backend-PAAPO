# PAAPO Backend API

Bem-vindo ao backend do PAAPO! Este projeto fornece uma API RESTful para gerenciamento de usuários, sessões de treinamento, informações de saúde e dados médicos. Este guia foi feito para ser fácil de entender, especialmente para estagiários e novos desenvolvedores.

## 🚀 Visão Geral
- **Linguagem:** Go (Golang)
- **Framework:** Gin
- **Banco de Dados:** PostgreSQL
- **Gerenciamento:** Docker e Docker Compose

## 🛠️ Como Rodar o Projeto

1. **Pré-requisitos:**
   - Docker e Docker Compose instalados

2. **Subir o ambiente:**
   ```sh
   docker-compose up -d --build
   ```
   Isso irá iniciar o banco de dados e a API na porta `8080`.

3. **Testar se está rodando:**
   ```sh
   curl http://localhost:8080/api/v1/health-info
   ```

## 🔑 Autenticação
A maioria dos endpoints requer um token JWT. Registre-se e faça login para obter seu token.

## 📚 Endpoints Principais

### Autenticação
- **POST /api/v1/auth/register**
  - Registra um novo usuário.
  - Exemplo:
    ```json
    {
      "first_name": "João",
      "last_name": "Silva",
      "email": "joao@exemplo.com",
      "password": "senha123"
    }
    ```
- **POST /api/v1/auth/login**
  - Faz login e retorna o token JWT.
  - Exemplo:
    ```json
    {
      "email": "joao@exemplo.com",
      "password": "senha123"
    }
    ```

### Usuários
- **Recuperação de senha:**
  - **POST /api/v1/auth/forgot-password**
  - **POST /api/v1/auth/reset-password**

### Informações de Saúde
- **POST /api/v1/health-info** — Cria informações de saúde
- **GET /api/v1/health-info** — Consulta informações de saúde
- **PUT /api/v1/health-info** — Atualiza informações de saúde
- **DELETE /api/v1/health-info** — Remove informações de saúde

### Dados Médicos
- **POST /api/v1/medical-data** — Cria dados médicos
- **GET /api/v1/medical-data** — Consulta dados médicos
- **PUT /api/v1/medical-data** — Atualiza dados médicos
- **DELETE /api/v1/medical-data** — Remove dados médicos

### Sessões de Treinamento (Check-in)
- **POST /api/v1/training-sessions** — Inicia uma nova sessão
  - Campos: `start_date`, `start_time`, `category` (opcional)
- **PUT /api/v1/training-sessions/end** — Encerra a sessão ativa
  - Campos: `comments` (opcional)
- **GET /api/v1/training-sessions/active** — Consulta a sessão ativa
- **GET /api/v1/session/active** — Alias para a sessão ativa
- **GET /api/v1/training-sessions** — Lista sessões do usuário (filtros: `from`, `to`, `category`, `page`, `limit`)

#### Relatórios de Sessões
- **GET /api/v1/sessions** — Lista sessões com filtros e paginação
- **GET /api/v1/sessions/total-hours?from=...&to=...** — Total de horas treinadas
- **GET /api/v1/sessions/average-duration?from=...&to=...** — Duração média das sessões
- **GET /api/v1/sessions/category-count** — Contagem por categoria

### Grupos de Usuários (Balanceamento)
- Novos usuários são automaticamente distribuídos entre Grupo 1 e Grupo 2 para manter o equilíbrio.
- A diferença entre os grupos nunca será maior que 1.

## 🧪 Testes e Scripts Úteis
- **test_group_assignment.sh** — Testa o balanceamento de grupos
- **check_group_distribution.sh** — Mostra a distribuição atual dos grupos no banco

## 💡 Dicas para Estagiários
- Sempre leia os exemplos de requisição e resposta.
- Use ferramentas como Postman ou Insomnia para testar os endpoints.
- Tokens JWT devem ser enviados no header: `Authorization: Bearer <token>`
- Se tiver dúvidas, pergunte! O código está documentado e a equipe está pronta para ajudar.

## 👨‍💻 Contribuindo
- Faça um fork, crie uma branch e envie seu PR.
- Siga o padrão de código e escreva comentários claros.

---

**Dúvidas?**
Abra uma issue ou fale com o time de desenvolvimento!
