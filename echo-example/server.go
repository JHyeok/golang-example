package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Cat struct {
	Name string `json:"name" validate:"required,min=3,max=128"`
	Age  int    `json:"age"`
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/dogs/:id", getDog)
	e.GET("/pets", showPet)
	e.POST("/dogs", createDog)

	e.POST("/cats", createCat)

	s := http.Server{
		Addr:    ":3000",
		Handler: e,
	}

	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

// Path Parameters (ex. http://localhost:3000/dogs/1)
func getDog(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

// Query Parameters (ex. http://localhost:3000/pets?dog=2&cat=3)
func showPet(c echo.Context) error {
	dog := c.QueryParam("dog")
	cat := c.QueryParam("cat")
	return c.String(http.StatusOK, "dog: "+dog+", cat: "+cat)
}

// Form application/x-www-form-urlencoded (ex. curl -d "name=cute dog" -d "age=5" http://localhost:3000/dogs)
// func createDog(c echo.Context) error {
// 	name := c.FormValue("name")
// 	age := c.FormValue("age")
// 	return c.String(http.StatusOK, "name: "+name+", age: "+age)
// }

func createDog(c echo.Context) error {
	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err != nil {
		return err
	}

	name := fmt.Sprintf("%v", json_map["name"])
	age := fmt.Sprintf("%v", json_map["age"])
	return c.String(http.StatusOK, "name: "+name+", age: "+age)
}

func createCat(c echo.Context) error {
	cat := new(Cat)
	if err := c.Bind(cat); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, cat)
}
