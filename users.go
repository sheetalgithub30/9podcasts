package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type user struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

func createUsers(c echo.Context) (err error) {

	us := &user{}

	if err = c.Bind(us); err != nil {
		return
	}

	q := `INSERT INTO users(u_name,email,u_password,created_at,updated_at) VALUES($1,$2,$3,$4,$5) RETURNING user_id`
	err = db.QueryRow(q, us.Name, us.Email, us.Password, us.Created_at, us.Updated_at).Scan(&us.Id)
	if err != nil {
		fmt.Println(err)
		return
	}

	return c.JSON(http.StatusOK, us)
}

func getUsers(c echo.Context) (err error) {

	var usr []user

	q := `SELECT user_id,u_name,email,u_password,created_at,updated_at FROM users`

	rows, err := db.Query(q)
	if err != nil {
		fmt.Println(err)
		return
	}

	for rows.Next() {
		var us user

		err = rows.Scan(&us.Id, &us.Name, &us.Email, &us.Password, &us.Created_at, &us.Updated_at)
		if err != nil {
			fmt.Println(err)
		}

		usr = append(usr, us)
	}

	return c.JSON(http.StatusOK, usr)
}
