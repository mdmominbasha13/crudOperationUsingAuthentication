package controller

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/CrudOperationUsingAuthentication/pkg/authentication"
	"github.com/CrudOperationUsingAuthentication/pkg/models"
	"github.com/dgriJalva/jwt-go"

	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	var newBook models.Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		log.Fatal(err.Error() + "during create request")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	b := newBook.CreateBook()
	c.JSON(http.StatusCreated, b)
}

func GetBook(c *gin.Context) {
	newBooks := models.GetBook()

	c.JSON(http.StatusOK, newBooks)
}

func GetBookById(c *gin.Context) {
	bookid := c.Param("bookid")

	id, err := strconv.ParseInt(bookid, 0, 0)
	if err != nil {
		log.Fatal("error while parsing")
	}
	bookDetails, _ := models.GetBookById(id)
	c.JSON(http.StatusOK, bookDetails)

}

func UpdateBookById(c *gin.Context) {
	var updateBook = &models.Book{}
	bookid := c.Param("bookid")
	if err := c.ShouldBindJSON(&updateBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := strconv.ParseInt(bookid, 0, 0)
	if err != nil {

		log.Fatal("error while parsing")
	}
	bookDetails, db := models.GetBookById(id)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name

	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author

	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication

	}
	db.Save(bookDetails)
	c.JSON(http.StatusOK, bookDetails)

}

func DeleteBookById(c *gin.Context) {
	bookid := c.Param("bookid")
	id, err := strconv.ParseInt(bookid, 0, 0)
	if err != nil {
		log.Fatal("error while parsing")
	}
	book := models.DeleteBookById(id)
	c.JSON(http.StatusOK, book)

}

var users = map[string]string{
	"user1": "password1",
}

func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	expectedPassword, exists := users[user.UserName]
	if !exists || expectedPassword != user.Password {

		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
		return

	}
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &models.Claims{
		UserName: user.UserName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(authentication.Jwtkey)
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
