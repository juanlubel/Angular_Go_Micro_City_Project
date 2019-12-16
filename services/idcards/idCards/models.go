package idCards

import (
	"database/sql"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"

	"idcards/db"

	"log"
)

type UserModel struct {
	ID      int64   `json:"id"`
	Name    string  `json:"name"`
	Surname string  `json:"surname"`
	Slug    string  `json:"slug"`
	Email   string  `json:"email"`
	NCard   float64 `json:"nCard"`
	Pass    string  `json:"pass"`
	// created_at time.Time `json:"created_at"`
	// updated_at time.Time `json:"updated_at"`
}
type UserDTO struct {
	Name    string  `json:"name"`
	Surname string  `json:"surname"`
	Slug    string  `json:"slug"`
	Email   string  `json:"email"`
	NCard   float64 `json:"nCard"`
}

type UserLoggedDTO struct {
	Slug    string  `json:"slug"`
	NCard   float64 `json:"nCard"`
	Token   string	`json:"token"`
}

type UserLogInDTO struct {
	Name 	string `json:"name"`
	Pass	string `json:"pass"`
}



func LogIn(userLogInDTO UserLogInDTO) (u UserModel, err error) {
	conn, err := db.ConnectSQL()

	if err != nil {
		return
	}
	defer conn.Close()

	sentence := fmt.Sprintf("SELECT id, name, surname, slug, nCard, email, pass FROM Users WHERE name='%v'", userLogInDTO.Name)

	err = conn.QueryRow(sentence).Scan(&u.ID, &u.Name, &u.Surname, &u.Slug, &u.NCard, &u.Email, &u.Pass)
	if err != nil {
		return
	}

	err = checkPassword(userLogInDTO.Pass, u.Pass)
	if err != nil {
		u = UserModel{}
	}

	return
}

func Details(id int64) (u UserModel, err error) {
	conn, err := db.ConnectSQL()

	if err != nil {
		return
	}
	defer conn.Close()

	sentence := fmt.Sprintf("SELECT id, name, surname, slug, nCard, email FROM Users WHERE id='%v'", id)

	fmt.Println(sentence)
	err = conn.QueryRow(sentence).Scan(&u.ID, &u.Name, &u.Surname, &u.Slug, &u.NCard, &u.Email)
	if err != nil {
		return
	}

	return
}

func Delete(u UserModel) (err error) {
	conn, err := db.ConnectSQL()

	if err != nil {
		return
	}
	defer conn.Close()

	//sentence, err := db.Prepare("INSERT INTO db_test (name) VALUES(?)")
	sentence := fmt.Sprintf("DELETE FROM Users WHERE name='%v'", u.Name)

	fmt.Println(sentence)
	_, err = conn.Query(sentence)
	if err != nil {
		return
	}
	return
}

func Update(u UserDTO, newSlug string) (err error) {
	conn, err := db.ConnectSQL()

	if err != nil {
		return
	}
	defer conn.Close()

	//sentence, err := db.Prepare("INSERT INTO db_test (name) VALUES(?)")
	sentence := fmt.Sprintf("UPDATE Users SET name='%v', slug='%v', email='%v' WHERE slug='%v'",u.Name, newSlug, u.Email, u.Slug)

	fmt.Println(sentence)
	_ , err = conn.Query(sentence)
	if err != nil {
		return
	}
	return
}

func Create(u UserModel) (v UserModel, err error) {
	conn, err := db.ConnectSQL()

	if err != nil {
		return
	}
	defer conn.Close()

	u.Pass, err = setPassword(u.Pass)
	if err != nil {
		log.Fatalln(err)
		return
	}

	num, err := setNCard(conn)
	if err != nil {
		log.Fatalln(err)
		return
	}

	u.NCard = num + 1
	//sentence, err := db.Prepare("INSERT INTO db_test (name) VALUES(?)")
	fmt.Print(u.Slug)
	sentence := fmt.Sprintf(
		"INSERT INTO Users VALUES(null, '%v', '%v', '%v','%v', '%v', '%v')",
		u.Name, u.Surname, u.Slug, u.Email, u.NCard, u.Pass)

	_, err = conn.Exec(sentence)
	if err != nil {
		return
	}
	//defer insert.close()

	//id, _ = insert.LastInsertId()
	v = u
	return
}

func ListAll() ([]UserDTO, error) {
	var users []UserDTO
	conn, err := db.ConnectSQL()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.Query("SELECT name, surname, slug, email, nCard FROM Users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var u UserModel
	for rows.Next() {
		err = rows.Scan(&u.Name, &u.Surname, &u.Slug, &u.Email, &u.NCard)
		if err != nil {
			return nil, err
		}
		users = append(users, ToUsersDTO(u))
	}

	return users, err
}

func setPassword(password string) (string, error) {

	if len(password) == 0 {
		return "", errors.New("password should not be empty!")
	}
	bytePassword := []byte(password)
	// Make sure the second param `bcrypt generator cost` between [4, 32)
	pass, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	PassStringified := string(pass)

	return PassStringified, nil
}

func checkPassword(password string, dbPassword string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(dbPassword)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}

func setNCard(conn *sql.DB) (num float64, err error) {
	sentence := fmt.Sprintf("SELECT MAX(nCard) as max_nCadr FROM Users")

	fmt.Println(sentence)
	rows, err := conn.Query(sentence)

	if err != nil {
		fmt.Print(err)
		return
	}
	for rows.Next() {
		_ = rows.Scan(&num)
	}
	if num == 0 {
		num = 1000
	}
	fmt.Printf("El numero es %v \n", num)
	return
}
