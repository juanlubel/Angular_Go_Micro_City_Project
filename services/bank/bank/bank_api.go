package bank

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//BankAPI :  Is in charge to provide all the posible transactions to the router
type BankAPI struct {
	BankService BankService
}

//ProvideBankAPI :  Provides to the router the entry api
func ProvideBankAPI(p BankService) BankAPI {
	return BankAPI{BankService: p}
}

//FindAll : The method to obtain all the data of the table
func (p *BankAPI) FindAll(c *gin.Context) {
	banks := p.BankService.FindAll()

	c.JSON(http.StatusOK, gin.H{"banks": ToBankDTOs(banks)})
}

//FindByID :  The method to obtain specific data
func (p *BankAPI) FindByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	bank := p.BankService.FindByID(uint(id))

	c.JSON(http.StatusOK, gin.H{"bank": ToBankDTO(bank)})
}

// Create : Create a new entry in the table
func (p *BankAPI) Create(c *gin.Context) {
	var bankDTO BankDTO
	err := c.BindJSON(&bankDTO)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	createdBank := p.BankService.Save(ToBank(bankDTO))

	c.JSON(http.StatusOK, gin.H{"bank": ToBankDTO(createdBank)})
}

// Update :  Updates an specific content of the table
func (p *BankAPI) Update(c *gin.Context) {
	var bankDTO BankDTO
	err := c.BindJSON(&bankDTO)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	bank := p.BankService.FindByID(uint(id))
	if bank == (Bank{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	bank.BankName = bankDTO.BankName
	p.BankService.Save(bank)

	c.Status(http.StatusOK)
}

// Delete :  Deletes an specific content of the table
func (p *BankAPI) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	bank := p.BankService.FindByID(uint(id))
	if bank == (Bank{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	p.BankService.Delete(bank)

	c.Status(http.StatusOK)
}
