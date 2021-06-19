package main

import (
	"fmt"
	"log"
	"net/http"
	_ "os/exec"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

func createUser(c echo.Context) (err error) {

	us := &user{}

	if err = c.Bind(us); err != nil {
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(us.Password), bcrypt.DefaultCost)
	// log.Println("hashedPassword= ", string(hashedPassword))
	if err != nil {
		log.Println("Error generating hashedPassword :", err)
		return
	}

	us.Created_at = time.Now()
	us.Updated_at = time.Now()

	q := `INSERT INTO users(name,email,password,created_at,updated_at) VALUES($1,$2,$3,$4,$5) RETURNING id`
	err = db.QueryRow(q, us.Name, us.Email, hashedPassword, us.Created_at, us.Updated_at).Scan(&us.ID)
	if err != nil {
		fmt.Println(err)
		return
	}

	return c.JSON(http.StatusOK, us)
}

func getUsers(c echo.Context) (err error) {

	var usr []user

	q := `SELECT id,name,email,password,created_at,updated_at FROM users`

	rows, err := db.Query(q)
	if err != nil {
		fmt.Println(err)
		return
	}

	for rows.Next() {
		var us user

		err = rows.Scan(&us.ID, &us.Name, &us.Email, &us.Password, &us.Created_at, &us.Updated_at)
		if err != nil {
			fmt.Println(err)
		}

		usr = append(usr, us)
	}

	return c.JSON(http.StatusOK, usr)
}

func deleteUser(c echo.Context) (err error) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = db.Exec(`Delete from users where id = $1 `, id)
	if err != nil {
		fmt.Println(err)
		return
	}

	return c.NoContent(http.StatusOK)
}

func updateUserName(c echo.Context) (err error) {
	us := &user{}

	us.Updated_at = time.Now()

	if err = c.Bind(us); err != nil {
		return
	}
	q := `UPDATE users SET name = $1, updated_at=$2 where id =$3`
	_, err = db.Exec(q, us.Name, us.Updated_at, us.ID)
	if err != nil {
		fmt.Println(err)
		return
	}
	return c.String(http.StatusOK, "Name updated successfully")
}

func updateUserEmail(c echo.Context) (err error) {
	us := &user{}

	us.Updated_at = time.Now()

	if err = c.Bind(us); err != nil {
		return
	}
	q := `UPDATE users SET email = $1, updated_at=$2 where id =$3`
	_, err = db.Exec(q, us.Email, us.Updated_at, us.ID)
	if err != nil {
		fmt.Println(err)
		return
	}
	return c.String(http.StatusOK, "Email updated successfully")
}

func updateUserPassword(c echo.Context) (err error) {
	us := &user{}

	us.Updated_at = time.Now()

	if err = c.Bind(us); err != nil {
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(us.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error generating hashedPassword :", err)
		return
	}
	q := `UPDATE users SET password = $1, updated_at=$2 where id =$3`
	_, err = db.Exec(q, hashedPassword, us.Updated_at, us.ID)
	if err != nil {
		fmt.Println(err)
		return
	}
	return c.String(http.StatusOK, "Password updated successfully")
}
