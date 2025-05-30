package controllers

import (
	"fmt"
	"go-api/internal/models"
	"net/http"

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

// Recebe um JSON no corpo da requisição, faz o bind para a struct User e responde com o mesmo JSON.
// BindJSON lê o corpo da requisição e preenche a struct automaticamente.
func (uc *UserController) CreateUser(c *gin.Context) {
	var user models.User

	// Tenta fazer o bind do JSON da requisição para a struct User
	// ShouldBindJson só preenche os campos que existem na struct. Qualquer campo extra no JSON será ignorado silenciosamente.
	// Tipos invaliados também vai gerar erro.

	//1) Forma de tratar erro (COMPACTA)
	// if err := algumaFuncao(); err != nil { --> Faça isso, se der erro, faça isso!
	// 	// Trate o erro aqui
	// }

	//2) Forma de tratar erro (NÃO COMPACTA)
	// err := algumaFuncao()
	// if err != nil {
	// 	fmt.Println("Erro:", err)
	// 	return
	// }

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}
	//Validações
	// fmt.Println(user.Email == nil) --> Isso não funciona, pois o campo Email é uma string, e tipos primitivos não podem ser nulos.
	// Tipos básicos como string, int, etc., têm valores zero:
	// 	string → "" (string vazia)
	// 	int → 0
	// 	bool → false

	fmt.Println("Email recebido:", user.Email)
	if user.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email é obrigatório"})
		return
	}

	c.JSON(http.StatusCreated, user)
}
