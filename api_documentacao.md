# Documentação da API PAAPO

## Sumário
- [Autenticação e Usuários](#autenticacao-e-usuarios)
- [Sessões de Treino](#sessoes-de-treino)
- [Relatórios de Sessões](#relatorios-de-sessoes)
- [Informações de Saúde](#informacoes-de-saude)
- [Dados Médicos](#dados-medicos)

---

## Autenticação e Usuários

### POST /api/v1/auth/register
Registra um novo usuário.

#### Requisição:
```json
{
  "first_name": "João",
  "last_name": "Silva",
  "email": "joao@exemplo.com",
  "password": "senha123"
}
```
#### Resposta (200):
```json
{
  "access_token": "<jwt>",
  "expires_in": 900
}
```
**Caso de uso:** Cadastro de novos usuários pelo app ou painel.

---

### POST /api/v1/auth/login
Realiza login e retorna um token JWT.

#### Requisição:
```json
{
  "email": "joao@exemplo.com",
  "password": "senha123"
}
```
#### Resposta (200):
```json
{
  "access_token": "<jwt>",
  "expires_in": 900
}
```
**Caso de uso:** Login de usuários para acessar áreas protegidas.

---

### POST /api/v1/auth/forgot-password
Solicita recuperação de senha.

#### Requisição:
```json
{
  "email": "joao@exemplo.com"
}
```
#### Resposta (200):
```json
{
  "message": "Se o email existir, você receberá um código de recuperação"
}
```
**Caso de uso:** Quando o usuário esqueceu a senha e deseja redefinir.

---

### POST /api/v1/auth/reset-password
Redefine a senha usando código enviado por email.

#### Requisição:
```json
{
  "email": "joao@exemplo.com",
  "code": "123456",
  "newPassword": "novaSenha123",
  "confirmPassword": "novaSenha123"
}
```
#### Resposta (200):
```json
{
  "message": "Senha alterada com sucesso"
}
```
**Caso de uso:** Finalizar o fluxo de recuperação de senha.

---

## Sessões de Treino

### POST /api/v1/training-sessions
Inicia uma nova sessão de treino para o usuário autenticado.

#### Requisição:
```json
{
  "start_date": "2025-07-02T14:00:00Z",
  "start_time": "2025-07-02T14:00:00Z",
  "category": "leitura" // opcional
}
```
#### Resposta (201):
```json
{
  "id": "abc123",
  "user_id": "...",
  "start_date": "2025-07-02T14:00:00Z",
  "start_time": "2025-07-02T14:00:00Z",
  "status": "active",
  "category": "leitura",
  ...
}
```
**Caso de uso:** Quando o usuário clica em "Iniciar Sessão" no painel principal.

---

### PUT /api/v1/training-sessions/end
Encerra a sessão de treino ativa do usuário.

#### Requisição:
```json
{
  "comments": "Sessão finalizada com sucesso!" // opcional
}
```
#### Resposta (200):
```json
{
  "id": "abc123",
  "status": "finished",
  "end_date": "2025-07-02T15:00:00Z",
  "duration": 3600,
  ...
}
```
**Caso de uso:** Quando o usuário encerra manualmente uma sessão.

---

### GET /api/v1/training-sessions
Lista sessões do usuário autenticado, com filtros e paginação.

#### Query Params:
- `from` (RFC3339, opcional)
- `to` (RFC3339, opcional)
- `category` (string, opcional)
- `page` (int, opcional)
- `limit` (int, opcional)

#### Resposta (200):
```json
{
  "sessions": [ { ... } ],
  "count": 2,
  "page": 1,
  "limit": 10
}
```
**Caso de uso:** Listar sessões para histórico, relatórios ou exibição no app.

---

### GET /api/v1/training-sessions/active
Obtém a sessão de treino ativa do usuário.

#### Resposta (200):
```json
{
  "id": "abc123",
  "status": "active",
  ...
}
```
#### Resposta (404):
```json
{
  "message": "no active training session found"
}
```
**Caso de uso:** Exibir sessão em andamento no painel.

---

### GET /api/v1/session/active
Alias para o endpoint acima (útil para frontend).

---

## Relatórios de Sessões

### GET /api/v1/sessions
Lista sessões com filtros e paginação (igual ao GET /training-sessions).

### GET /api/v1/sessions/total-hours?from=...&to=...
Retorna o total de horas treinadas no período.

#### Resposta (200):
```json
{
  "total_hours": 12.5
}
```
**Caso de uso:** Relatórios de desempenho.

---

### GET /api/v1/sessions/average-duration?from=...&to=...
Retorna a duração média das sessões no período (em minutos).

#### Resposta (200):
```json
{
  "average_duration_minutes": 45.2
}
```
**Caso de uso:** Relatórios de desempenho.

---

### GET /api/v1/sessions/category-count
Retorna a contagem de sessões por categoria.

#### Resposta (200):
```json
{
  "category_counts": {
    "leitura": 5,
    "escrita": 3
  }
}
```
**Caso de uso:** Análise de engajamento por tipo de atividade.

---

## Informações de Saúde

### POST /api/v1/health-info
Cria informações de saúde do usuário.

#### Requisição:
```json
{
  "weight": 70.5,
  "height": 175.0,
  "smoker": false,
  "alcohol_consumption": "occasional",
  "physical_activity_frequency": "3x semana",
  "physical_activity_type": "Caminhada"
}
```
#### Resposta (201):
```json
{
  "id": 1,
  "weight": 70.5,
  ...
}
```
**Caso de uso:** Cadastro inicial ou atualização de dados de saúde.

---

### GET /api/v1/health-info
Consulta informações de saúde do usuário.

### PUT /api/v1/health-info
Atualiza informações de saúde.

### DELETE /api/v1/health-info
Remove informações de saúde do usuário.

---

## Dados Médicos

### POST /api/v1/medical-data
Cria dados médicos do usuário.

### GET /api/v1/medical-data
Consulta dados médicos do usuário.

### PUT /api/v1/medical-data
Atualiza dados médicos do usuário.

### DELETE /api/v1/medical-data
Remove dados médicos do usuário.

---

## Observações Gerais
- Todos os endpoints (exceto login, cadastro e recuperação de senha) exigem autenticação via JWT no header:
  ```
  Authorization: Bearer <token>
  ```
- Datas devem estar no formato RFC3339 (ex: `2025-07-02T14:00:00Z`).
- Em caso de erro, a resposta segue o padrão:
  ```json
  { "error": "mensagem do erro" }
  ```

---

## Health Check

Se implementado, normalmente disponível em `/` ou `/health` (não encontrado explicitamente no código, mas pode ser adicionado conforme necessidade). 