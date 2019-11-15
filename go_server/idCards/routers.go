package idCards

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func UserRegister(router *gin.RouterGroup) {
	router.POST("/", createIdCard)
	router.GET("/:id", getIdCard)
	router.DELETE("/:id", deleteIdCard)
	router.PUT("/:id", updateIdCard)
}

func UsersList(router *gin.RouterGroup) {
	router.GET("/", getAllUsers)
}

func createIdCard(c *gin.Context) {
	fmt.Print("Creating")
	user := UserModel{}
	c.BindJSON(&user)
	id, err := Create(user)
	if err != nil {
		fmt.Print(err)
		c.JSON(409, gin.H{"error": err, "user": user})
		return
	}
	user.ID = id

	c.JSON(201, gin.H{"user": user})
	return
}
func getIdCard(c *gin.Context) {
	fmt.Print("details")
	fmt.Println(c.Param("id"))
	user := UserModel{}
	user.Name = c.Param("id")
	err := Details(user)
	if err != nil {
		fmt.Print(err)
		c.JSON(409, gin.H{"error": err, "user": user})
		return
	}

	c.JSON(200, gin.H{"msg": user})

}
func deleteIdCard(c *gin.Context) {
	fmt.Print("Delete")
	fmt.Println(c.Param("id"))
	user := UserModel{}
	user.Name = c.Param("id")
	err := Delete(user)
	if err != nil {
		fmt.Print(err)
		c.JSON(409, gin.H{"error": err, "user": user})
		return
	}

	c.JSON(200, gin.H{"msg": user})
}
func updateIdCard(c *gin.Context) {
	fmt.Print("Update")
	fmt.Println(c.Param("id"))
	user := UserModel{}
	c.BindJSON(&user)
	user.Name = c.Param("id")
	err := Update(user)
	if err != nil {
		fmt.Print(err)
		c.JSON(409, gin.H{"error": err, "user": user})
		return
	}

	c.JSON(200, gin.H{"msg": user})
}

func getAllUsers(c *gin.Context) {
	res, err := ListAll()
	fmt.Println("List", res)
	c.JSON(200, gin.H{	"msg": res,
						"error": err})
}
