package idCards

import (
	"fmt"
	"idcards/db"
)

type UserModel struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	// created_at time.Time `json:"created_at"`
	// updated_at time.Time `json:"updated_at"`
}


func Details(u UserModel) (err error) {
	conn, err := db.ConnectSQL()

	if err != nil {
		return
	}
	defer conn.Close()

	sentence := fmt.Sprintf("SELECT name FROM Users WHERE name='%v'", u.Name)

	fmt.Println(sentence)
	res, err := conn.Query(sentence)
	if err != nil {
		return
	}

	fmt.Println(res)
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

func Update(u UserModel) (err error) {
	conn, err := db.ConnectSQL()

	if err != nil {
		return
	}
	defer conn.Close()

	//sentence, err := db.Prepare("INSERT INTO db_test (name) VALUES(?)")
	sentence := fmt.Sprintf("UPDATE Users SET name='Joselito' WHERE name='%v'", u.Name)

	fmt.Println(sentence)
	_, err = conn.Query(sentence)
	if err != nil {
		return
	}
	return
}

func Create(u UserModel) (id int64, err error) {
	conn, err := db.ConnectSQL()

	if err != nil {
		return
	}
	defer conn.Close()

	//sentence, err := db.Prepare("INSERT INTO db_test (name) VALUES(?)")
	sentence := fmt.Sprintf("INSERT INTO Users VALUES(null, '%v', '%v')", u.Name, u.Email)

	fmt.Println(sentence)
	insert, err := conn.Exec(sentence)
	if err != nil {
		return
	}
	//defer insert.close()

	id, _ = insert.LastInsertId()
	fmt.Printf("%T", id)
	fmt.Println("id")
	return
}

func ListAll() ([]UserModel, error) {
	var users []UserModel
	conn, err := db.ConnectSQL()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.Query("SELECT * FROM Users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var u UserModel
	for rows.Next() {
		err = rows.Scan(&u.ID, &u.Name, &u.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, err
}
