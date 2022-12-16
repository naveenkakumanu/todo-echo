package controllers

import (
	"github.com/labstack/echo"
	"github.com/naveenkakumanu/todo-echo/db"
)

func Create(c echo.Context) error {
	task := new(db.Task)
	conn := db.DB()
	conn.Create(&task)
	return nil
}

func Read(c echo.Context) error {
	task := new(db.Task)
	id := c.Param("id")
	conn := db.DB()
	conn.Find(&task, "id = ?", id)
	return nil
}

func Update(c echo.Context) error {
	task := new(db.Task)
	id := c.Param("id")
	c.Bind(&task)
	conn := db.DB()
	conn.Model(&task).Where("id", id).Update("name", task.Name)
	c.JSON(200, task)
	return nil
}

func Delete(c echo.Context) error {
	task := new(db.Task)
	id := c.Param("id")
	conn := db.DB()
	conn.Find(&task, "id = ?", id)
	c.JSON(200, task)
	return nil
}

func CreateUser(c echo.Context) error {
	user := new(db.User)
	c.Bind(&user)
	conn := db.DB()
	conn.Create(&user)

	return nil
}

func ReadUser(c echo.Context) error {
	var user db.User
	id := c.Param("id")
	conn := db.DB()
	conn.Find(&user, "id = ?", id)
	c.JSON(200, user)
	return nil
}

func UpdateUser(c echo.Context) error {
	user := new(db.User)
	id := c.Param("id")
	conn := db.DB()
	conn.Model(&user).Where("id", id).Update("name", user.Name)
	return nil
}

// DeleteUser
func DeleteUser(c echo.Context) error {
	var user db.User
	id := c.Param("id")
	conn := db.DB()
	conn.Delete(&user, "id = ?", id)
	return nil
}
