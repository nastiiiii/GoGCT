package Controller

import (
	"GCT/Structure/Services"
	"GCT/Structure/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type AccountController struct {
	Services *Services.AccountService
}

func NewAccountController(service Services.AccountService) *AccountController {
	return &AccountController{Services: &service}
}

func (ac *AccountController) Register(c *gin.Context) {
	var request struct {
		ContactInfo  string `json:"contactInfo"`
		IsSocialClub bool   `json:"isSocialClub"`
		UserDOB      string `json:"userDOB"`
		Username     string `json:"username"`
		Password     string `json:"password"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	parsedDOB, err := time.Parse("2006-01-02", request.UserDOB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD."})
		return
	}

	account := models.Account{
		ContactInfo:    request.ContactInfo,
		IsSocialClub:   request.IsSocialClub,
		UserDOB:        parsedDOB,
		Username:       request.Username,
		AccountBalance: 0, // Default balance
	}

	createdAccount, err := ac.Services.Register(account, request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"accountID":    createdAccount.AccountId,
		"contactInfo":  createdAccount.ContactInfo,
		"isSocialClub": createdAccount.IsSocialClub,
		"userDOB":      createdAccount.UserDOB,
		"username":     createdAccount.Username,
	})
}

func (ac *AccountController) Login(c *gin.Context) {
	var cred struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBind(&cred); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := ac.Services.Login(cred.Username, cred.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (ac *AccountController) GetUserByToken(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token required"})
		return
	}

	user, err := ac.Services.GetUserByToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (ac *AccountController) GetUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	account, err := ac.Services.GetAccountById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, account)
}

func (ac *AccountController) UpdateAccount(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var account models.Account
	if err := c.ShouldBind(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedAccount, err := ac.Services.UpdateAccount(id, account)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedAccount)
}

func (ac *AccountController) DeleteAccount(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ac.Services.DeleteAccount(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Account deleted"})
}

func SetupAccountRouter(router *gin.Engine, service Services.AccountService) {
	controller := NewAccountController(service)

	accountRoutes := router.Group("/account")
	{
		accountRoutes.POST("/register", controller.Register)
		accountRoutes.POST("/login", controller.Login)
		accountRoutes.GET("/me", controller.GetUserByToken)
		accountRoutes.GET("/:id", controller.GetUserById)
		accountRoutes.PUT("/:id", controller.UpdateAccount)
		accountRoutes.DELETE("/:id", controller.DeleteAccount)
	}

} //TODO THE GETTICKETS AND ANOTHER FUNCTION
