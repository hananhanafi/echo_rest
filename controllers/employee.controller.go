package controllers

import (
	"echo-rest/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func FetchAllEmployee(c echo.Context) error {
	result, err := models.GetAllEmployee()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreEmployee(c echo.Context) error {
	name := c.FormValue("name")
	address := c.FormValue("address")
	phone := c.FormValue("phone")

	result, err := models.StoreEmployee(name, address, phone)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateEmployee(c echo.Context) error {
	id := c.FormValue("id")
	name := c.FormValue("name")
	address := c.FormValue("address")
	phone := c.FormValue("phone")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateEmployee(conv_id, name, address, phone)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteEmployee(c echo.Context) error {
	id := c.FormValue("id")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteEmployee(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
