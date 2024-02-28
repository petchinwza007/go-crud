package controller

import (
	"example/go-crud/initializers"
	"example/go-crud/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context){
	var body struct{
		Username string
		Password string
	}

	if c.Bind(&body) != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Failed to read body",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error" : err,
		})
		return
	}
	sign := models.User{
		Username : body.Username,
		Password: string(hash),
	}
	result := initializers.DB.Create(&sign)
	if result.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error":"Failed to create user",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"message" : "Create Successful",
	})
}

func Login(c *gin.Context) {
	var body struct{
		Username string
		Password string
	}

	if c.Bind(&body) != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Failed to read body",
		})
		return
	}
	var user models.User
	initializers.DB.First(&user, "username = ?", body.Username)

	if user.ID == 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Username or Password",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Username or Password",
		})
		return
	}

	//jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
        jwt.MapClaims{ 
        "username": user.ID, 
        "exp": time.Now().Add(time.Hour * 24).Unix(), 
        })

    tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
   	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid to create token",
		})
		return
   	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authirization", tokenString, 3600 * 24 * 30, "", "", false, true)
	c.IndentedJSON(http.StatusOK, gin.H{
		
	})

}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{"message": user})

}