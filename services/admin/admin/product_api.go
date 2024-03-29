package admin

import (

	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
	"golang.org/x/crypto/bcrypt"

	"Go_Gingonic_Server/common"
)

//AdminApi :  Is in charge to provide all the posible transactions to the router
type Api struct {
	Service Service
}

//ProvideAdminApi :  Provides to the router the entry api
func ProvideAdminApi(p Service) Api {
	return Api{Service: p}
}

//func (p *Api) WSHandler(c *gin.Context) {
//	WS(c.Writer, c.Request)
//
//}
//
//func WS (res http.ResponseWriter, req *http.Request) {
//	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(res, req, nil)
//	if err != nil {
//		fmt.Println("Failed to set websocket upgrade: %+v", err)
//		return
//	}
//
//	for {
//		t, msg, err := conn.ReadMessage()
//		fmt.Println(t, msg, err)
//		if err != nil {
//			break
//		}
//		conn.WriteMessage(t, msg)
//	}
//
//	//conn.WriteMessage(_ ,admin)
//}

func (p *Api) LogIn(c *gin.Context) {
	c.Header("Content-Type", "application/json; charset=utf-8")
	var LogInAdminDTO LogInAdminDTO
	err := c.BindJSON(&LogInAdminDTO)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}
	name := LogInAdminDTO.Name
	pass := LogInAdminDTO.Pass
	admin := p.Service.FindByName(&Admin{Name: name})
	if admin == (Admin{}) {
		c.JSON(http.StatusNotFound, gin.H{"login": "Invalid username"})
		return
	}
	err = p.checkPassword(pass, admin)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"login": "Invalid password"})
		return
	}

	token, err := common.GenerateJWT(LogInAdminDTO.Name)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}
	res := ToLoggedAdminDTO(admin, token)

	request := gorequest.New()
	_, _, errs := request.Post("http://redis_server:3015/admin").Send(res).End()
	if errs != nil {
		fmt.Println(errs)
		os.Exit(1)
	}

	fmt.Println(res)

	c.JSON(http.StatusOK, gin.H{"logged": ToLoggedAdminDTO(admin, token)})
}

//FindByID :  The method to obtain specific data
func (p *Api) FindByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	fmt.Println(id)
	admin := p.Service.FindByID(uint(id))
	fmt.Println(admin)
	c.JSON(http.StatusOK, gin.H{"admin": ToAdminDTO(admin)})
}

// Create : Create a new entry in the table
func (p *Api) Create(c *gin.Context) {

	var adminDTO AdminDTO
	err := c.BindJSON(&adminDTO)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}
	adminDTO.Pass, err = p.setPassword(adminDTO.Pass)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	token, err := common.GenerateJWT(adminDTO.Name)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	createdProduct := p.Service.Save(ToAdmin(adminDTO))
	//fmt.Println(createdProduct)
	c.JSON(http.StatusOK, gin.H{"admin": ToLoggedAdminDTO(createdProduct, token)})
}

// Update :  Updates an specific content of the table
func (p *Api) Update(c *gin.Context) {
	var adminDTO AdminDTO
	err := c.BindJSON(&adminDTO)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))

	admin := p.Service.FindByID(uint(id))
	if admin == (Admin{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	admin.Name = adminDTO.Name
	admin.Pass = adminDTO.Pass
	p.Service.Save(admin)

	c.Status(http.StatusOK)
}

// Delete :  Deletes an specific content of the table
func (p *Api) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	admin := p.Service.FindByID(uint(id))
	if admin == (Admin{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	p.Service.Delete(admin)

	c.Status(http.StatusOK)
}

func (p *Api) setPassword(password string) (string, error) {

	if len(password) == 0 {
		return "", errors.New("password should not be empty!")
	}
	bytePassword := []byte(password)
	// Make sure the second param `bcrypt generator cost` between [4, 32)
	pass, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	PassStringified := string(pass)

	return PassStringified, nil
}

func (p *Api) checkPassword(password string, admin Admin) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(admin.Pass)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}
