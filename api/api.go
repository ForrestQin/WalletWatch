package api

import (
	"WalletWatch/db"
	"WalletWatch/pkg/user"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"strconv"
)

func SetUpUserAPI(router *gin.Engine, dbClient *mongo.Client) {
	db.InitializeDatabaseSetting(dbClient)
	userGroup := router.Group("/users")
	userGroup.GET("/", getUserList)
	userGroup.GET("/:id", getUserByID)
	userGroup.POST("/", createUser(dbClient))
	userGroup.PUT("/:id", updateUser)
	userGroup.DELETE("/:id", deleteUser)
}

func getUserList(c *gin.Context) {
	limit := getIniQueryParam(c, "limit", 5)
	offset := getIniQueryParam(c, "offset", 0)
	users, err := db.GetUserList(limit, offset)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func getUserByID(c *gin.Context) {

}

func createUser(dbClient *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		body, err := c.GetRawData()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		var newUser user.User

		if err := newUser.UnmarshalJSON(body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//newUser.CreatedDate = time.Now()
		if err := db.CreateUser(newUser); err != nil {
			log.Printf("Erroe creat user: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
	}
}

func updateUser(c *gin.Context) {

}

func deleteUser(c *gin.Context) {

}

func getIniQueryParam(c *gin.Context, key string, defaultValue int64) int64 {
	value, err := strconv.ParseInt(c.DefaultQuery(key, strconv.FormatInt(defaultValue, 10)), 10, 64)
	if err != nil {
		return defaultValue
	}
	return value
}
