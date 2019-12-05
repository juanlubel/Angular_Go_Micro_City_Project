package idCards

import (
	"fmt"
	"github.com/bariseser/Go-Seeder"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"strconv"
)

func UserRegister(router *gin.RouterGroup) {
	router.POST("/", createIdCard)
	router.GET("/:id", getIdCard)
	router.DELETE("/:id", deleteIdCard)
	router.PUT("/:id", updateIdCard)
}

func UserLogin(router *gin.RouterGroup) {
	router.POST("", loginIdCard)
}

func UsersList(router *gin.RouterGroup) {
	router.GET("/", getAllUsers)
}

func UsersSeed(router *gin.RouterGroup) {
	router.GET("/", createRows)
}

func loginIdCard(c *gin.Context) {
	fmt.Print("Login")
	c.JSON(200, gin.H{"user": "login"})
}

func createIdCard(c *gin.Context) {
	fmt.Print("Creating")
	user := UserModel{}
	c.BindJSON(&user)
	user.Slug = slug.Make(user.Name +" "+ user.Surname)
	newUser, err := Create(user)
	if err != nil {
		fmt.Print(err)
		c.JSON(409, gin.H{"error": err, "user": user})
		return
	}
	c.JSON(201, gin.H{"user": ToUsersDTO(newUser)})
	return
}
func getIdCard(c *gin.Context) {
	fmt.Print("details")
	fmt.Println(c.Param("id"))
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := Details(int64(id))
	if err != nil {
		fmt.Print(err)
		c.JSON(409, gin.H{"error": err, "user": user})
		return
	}

	c.JSON(200, gin.H{"user": ToUsersDTO(user)})

}
func deleteIdCard(c *gin.Context) {
	fmt.Print("Delete")
	fmt.Println(c.Param("id"))
	user := UserModel{}
	user.Name = c.Param("id")
	err := Delete(user)
	if err != nil {
		fmt.Print(err)
		c.JSON(409, gin.H{"error": err, "user": ToUsersDTO(user)})
		return
	}

	c.JSON(200, gin.H{"user": user})
}
func updateIdCard(c *gin.Context) {
	fmt.Print("Update")
	fmt.Println(c.Param("id"))
	user := UserModel{}
	c.BindJSON(&user)
	//no esta terminado, no updatea parametros según peticion del usuario.
	user.Name = c.Param("id")
	err := Update(user)
	if err != nil {
		fmt.Print(err)
		c.JSON(409, gin.H{"error": err, "user": ToUsersDTO(user)})
		return
	}

	c.JSON(200, gin.H{"user": ToUsersDTO(user)})
}

func getAllUsers(c *gin.Context) {
	fmt.Print("List All")
	res, err := ListAll()
	c.JSON(200, gin.H{	"list": res,
						"error": err})
}


func createRows(c *gin.Context) {
	n := 0
	for n <= 100 {
		surname := seeder.Name()
		name := seeder.FirstNameMale()
		slugify := slug.Make(name +" "+ surname)
		email :=  slugify + "@gmail.com"
		u := UserModel{
			Name:   	name,
			Surname: 	surname,
			Slug:		slugify,
			Email:  	email,
			Pass:  		surname,
		}
		Create(u)
		n++
	}

}
