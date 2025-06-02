package controllers

import (
	"go-api/internal/models"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Em GO não existe classe, então usamos STRUCTS + MÉTODOS para organizar a lógica.
// Está sendo usado aqui como um recebedor de me´todos! (Agrupando funções relacionadas)
type UserController struct {
	db *gorm.DB //db é um ponteiro para a conexão com o banco de dados, que será usado para realizar operações CRUD.
}

// NewUserController cria uma nova instância do UserController passando a conexão com o banco de dados como parâmetro.
func NewUserController(db *gorm.DB) *UserController {
	return &UserController{db: db}
}

// uc *UserController --> o método está ligado a uma instância de UserController como ponteiro.
// gin.Context --> Meu canal de resposta, é o CONTEXTO da requisição HTTP.
// -- Quando usamos c.JSON(...) estamos dizendo: "responda o client com esse conteúdo JSON e esse status code"
func (uc *UserController) GetUser(c *gin.Context) {
	id := c.Param("id") // Pega o parâmetro "id" da URL

	var user models.User
	result := uc.db.First(&user, id) // Busca o usuário pelo ID no banco de dados
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Usuário não encontrado",
		})
		return
	}
	c.JSON(http.StatusOK, user)

}
func (uc *UserController) GetAllUsers(c *gin.Context) {
	var users []models.User
	result := uc.db.Find(&users) // Busca todos os usuários no banco de dados
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao buscar usuários",
		})
		return
	}
	c.JSON(http.StatusOK, users)
}
func (uc *UserController) GetAllUsersPaginated(c *gin.Context) {
	// Step 1: Obter parâmetros de paginação (com valores padrão)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))    // Pega o parâmetro "page" da URL, com valor padrão 1
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10")) // Pega o parâmetro "limit" da URL, com valor padrão 10

	offset := (page - 1) * limit // Cálculo do offset para a consulta
	var users []models.User
	var total int64

	uc.db.Model(&models.User{}).Count(&total)                // Conta o total de usuários no banco de dados
	result := uc.db.Offset(offset).Limit(limit).Find(&users) // Busca os usuários com offset e limite
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao buscar usuários",
		})
		return
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit))) // Calcula o total de páginas
	c.JSON(http.StatusOK, models.UserPagination{
		Data:       users,
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
	})
}
func (uc *UserController) CreateUser(c *gin.Context) {
	var createUserRequest models.CreateUserRequest

	if err := c.ShouldBindJSON(&createUserRequest); err != nil {
		log.Printf("Erro no bind JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos!"})
		return
	}

	user := models.User{
		Login:       createUserRequest.Login,
		Name:        createUserRequest.Name,
		Email:       createUserRequest.Email,
		Password:    createUserRequest.Password,
		FullName:    createUserRequest.FullName,
		Active:      createUserRequest.Active,
		PhoneNumber: createUserRequest.PhoneNumber,
	}

	if err := uc.db.Create(&user).Error; err != nil {
		log.Printf("Erro ao criar usuário: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar usuário"})
		return
	}
	user.Password = ""
	c.JSON(http.StatusCreated, user)
}
func (uc *UserController) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var updateUserRequest models.UpdateUserRequest
	var user models.User

	if err := c.ShouldBindJSON(&updateUserRequest); err != nil {
		log.Printf("Erro no bind JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos!"})
		return
	}

	if err := uc.db.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado!"})
		return
	}

	user.Login = updateUserRequest.Login
	user.Password = updateUserRequest.Password

	if err := uc.db.Save(&user).Error; err != nil {
		log.Printf("Erro ao atualizar usuário: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar usuário"})
		return
	}

	user.Password = ""
	c.JSON(http.StatusOK, user)
}
