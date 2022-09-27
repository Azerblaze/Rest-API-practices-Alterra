package controller

import (
	"net/http"
	"strconv"

	"praktikum/config"
	"praktikum/model"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
)

//get all books
func GetBooksController(c echo.Context) error {
	var books []model.Book

	if err := config.DB.Find(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all books",
		"books":   books,
	})
}

// get book by id
func GetBookController(c echo.Context) error {
	var books []model.Book

	id, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.First(&books, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all books",
		"books":   config.DB.First(&books, id),
	})
}

//create new book
func CreateBookController(c echo.Context) error {
	book := model.Book{}
	c.Bind(&book)

	if err := config.DB.Save(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new book",
		"book":    book,
	})
}

//delete book by id
func DeleteBookController(c echo.Context) error {
	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.Delete(&model.Book{}, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "book deleted",
	})
}

//update book by id
func UpdateBookController(c echo.Context) error {
	// your solution here
	var book model.Book
	newBook := model.Book{}
	c.Bind(&newBook)

	id, _ := strconv.Atoi(c.Param("id"))

	//check if data exist
	if err := config.DB.First(&book, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	book.Name = newBook.Name
	book.Author = newBook.Author
	book.Page = newBook.Page

	//update book
	if err := config.DB.Save(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update book",
		"books":   book,
	})
}
