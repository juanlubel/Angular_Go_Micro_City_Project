package accounts

import (
	"Go_Gingonic_Server/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//AccountAPI :  Is in charge to provide all the posible transactions to the router
type AccountAPI struct {
	AccountService AccountService
	JWT            *utils.JWT
}

//ProvideAccountAPI :  Provides to the router the entry api
func ProvideAccountAPI(p AccountService, jwt *utils.JWT) AccountAPI {
	return AccountAPI{AccountService: p, JWT: jwt}
}

func (p *AccountAPI) LogIn(c *gin.Context) {
	var LogInUserDTO LogInUserDTO
	err := c.BindJSON(&LogInUserDTO)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}
	name := LogInUserDTO.Owner
	pass := LogInUserDTO.AccountNumber
	bank := LogInUserDTO.Bank
	token, err := p.JWT.GenerateJWT(pass)
	account := p.AccountService.FindByBankAndName(name, bank)
	if account == (BankAccount{}) {
		c.JSON(http.StatusNotFound, gin.H{"error": "The User doesn't exist."})
		return
	}
	if pass != account.AccountNumber {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid data"})
		return
	}
	/* err = p.checkPassword(pass, admin)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"login": "Invalid password"})
		return
	} */

	c.JSON(http.StatusOK, gin.H{"user": ToLoggedUserDTO(account, token)})
}

/* Simple Crud Opperations */

//FindAll : The method to obtain all the data of the table
func (p *AccountAPI) FindAll(c *gin.Context) {
	accounts := p.AccountService.FindAll()

	c.JSON(http.StatusOK, gin.H{"accounts": ToFullAccountDTOs(accounts)})
}

//FindByOwner :  The method to obtain specific data
func (p *AccountAPI) FindByOwner(c *gin.Context) {
	owner := c.Param("name")
	account := p.AccountService.FindByOwner(owner)
	c.JSON(http.StatusOK, gin.H{"account": ToFullAccountDTO(account)})
}

// Create : Create a new entry in the table (Used in Register Functionality)
func (p *AccountAPI) Create(c *gin.Context) {
	var fullAccountDTO FullAccountDTO
	err := c.BindJSON(&fullAccountDTO)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	createdAccount := p.AccountService.Save(ToAccount(fullAccountDTO))

	c.JSON(http.StatusOK, gin.H{"account": ToFullAccountDTO(createdAccount)})
}

/*
// Update :  Updates an specific content of the table
func (p *AccountAPI) Update(c *gin.Context) {
	var fullAccountDTO FullAccountDTO
	err := c.BindJSON(&fullAccountDTO)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	owner, _ := strconv.Atoi(c.Param("owner"))
	account := p.AccountService.FindByOwner(string(owner))
	if account == (BankAccount{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	account.Bank = fullAccountDTO.Bank
	p.AccountService.Save(account)

	c.Status(http.StatusOK)
}

// Delete :  Deletes an specific content of the table
func (p *AccountAPI) Delete(c *gin.Context) {
	owner, _ := strconv.Atoi(c.Param("owner"))
	account := p.AccountService.FindByOwner(string(owner))
	if account == (BankAccount{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	p.AccountService.Delete(account)

	c.Status(http.StatusOK)
} */
