package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Contact struct {
	Id     int    `json:id`
	Name   string `json:"name"`
	Number string `json:"number"`
}

// get all contacts
func GetContacts(c *fiber.Ctx) error {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/newApi")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	res, err := db.Query("SELECT * FROM `contacts`")
	if err != nil {
		fmt.Println(err, res)
	}
	allUsers := []Contact{}

	for res.Next() {
		var user Contact
		err = res.Scan(&user.Id, &user.Name, &user.Number)
		if err != nil {
			panic(err)
		}

		allUsers = append(allUsers, user)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{

		"contacts": allUsers,
	})
}

// Create a contact
func CreateContact(c *fiber.Ctx) error {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/newApi")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	type Request struct {
		Name   string `json:"name"`
		Number string `json:"number"`
	}

	var body Request

	err2 := c.BodyParser(&body)

	// if error
	if err2 != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"result": "Cannot parse JSON",
		})
	}
	res, err := db.Query(fmt.Sprintf("SELECT * FROM `contacts` WHERE `number` = '%s'", body.Number))
	if err != nil {
		panic(err)
	}
	allUser := Contact{}
	// if error in parsing string to int
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse Id",
		})
	}
	for res.Next() {
		var user Contact
		err = res.Scan(&user.Id, &user.Name, &user.Number)
		if err != nil {
			panic(err)
		}
		allUser = user
	}
	if allUser.Number != "" {
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"contacts": "this number is already in the database",
		})
	}
	insert, err := db.Query(fmt.Sprintf("INSERT INTO `contacts` (`name`, `number`) VALUES('%s', '%s')", body.Name, body.Number))
	if err != nil {
		fmt.Println(err)
	}
	defer insert.Close()
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"contacts": "successful contact addition",
	})
}

// get a single contact
// PARAM: id
func GetContact(c *fiber.Ctx) error {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/newApi")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	id := c.Params("id")
	res, err := db.Query(fmt.Sprintf("SELECT * FROM `contacts` WHERE `id` = '%s'", id))
	if err != nil {
		panic(err)
	}
	allUser := Contact{}

	// if error in parsing string to int
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"result": "Cannot parse JSON",
		})
	}
	for res.Next() {
		var user Contact
		err = res.Scan(&user.Id, &user.Name, &user.Number)
		if err != nil {
			panic(err)
		}

		allUser = user

	}

	// find todo and return

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
      "id": allUser.Id,
			"name": allUser.Name,
      "number": allUser.Number,
	})

	// if todo not available

}

// Update a contact
// PARAM: id
func UpdateContact(c *fiber.Ctx) error {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/newApi")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// get parameter value
	id := c.Params("id")

	// request structure
	type Request struct {
		Name   *string `json:"name"`
		Number *string `json:"number"`
	}

	var body Request
	err = c.BodyParser(&body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"result": "Cannot parse JSON",
		})
	}

	upUser, err := db.Query(fmt.Sprintf("UPDATE `contacts` SET `name`='%s', `number`='%s' WHERE id='%v'", *body.Name, *body.Number, id))
	if err != nil {
		fmt.Println(err)
	}

	defer upUser.Close()

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"result": "information updated",
	})
}

// Delete a contact
// PARAM: id
func DeleteContact(c *fiber.Ctx) error {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/newApi")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// get parameter value
	id := c.Params("id")

	delete, err := db.Query(fmt.Sprintf("DELETE FROM `contacts` WHERE id = '%v'", id))
	if err != nil {
		fmt.Println(err)
	}
	defer delete.Close()

	return c.JSON(fiber.Map{
		"result": "Deleted Succesfully",
	})

}
