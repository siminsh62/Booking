package main

import (
	"Booking/models"
	"Booking/repository"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"net/http"
)
var db *gorm.DB
func main() {
	fmt.Println("hello")

	db, _ =repository.Connect()


	app := fiber.New()

	//	app.Get("/", hello)

	userGroup:=app.Group("/user")

	userGroup.Get("/",getUsers)
	userGroup.Post("/",addUser)
	//userGroup.Get("/:id",getUserById)
	//userGroup.Put("/:id",editUserById)
	//userGroup.Delete("/:id",deleteUserById)
	//
	//hotelGroup:=app.Group("/hotel")
	//
	//hotelGroup.Get("/",getHotels)
	//hotelGroup.Post("/",addHotel)
	//hotelGroup.Get("/:id", getHotelById)
	//hotelGroup.Put("/:id",editHotelById)
	//hotelGroup.Delete("/:id",deleteHotelById)
	//
	//bookGroup:=app.Group("/book")
	//
	//bookGroup.Get("/",getBooks)
	//bookGroup.Post("/",addBook)
	//bookGroup.Get("/:id", getBookById)
	//bookGroup.Put("/:id",editBookById)
	//bookGroup.Delete("/:id",deleteBookById)

	app.Listen(":3000")
}

//
//func deleteBookById(ctx *fiber.Ctx) error {
//
//}
//
//func editBookById(ctx *fiber.Ctx) error {
//
//}
//
//func getBookById(ctx *fiber.Ctx) error {
//
//}
//
//func addBook(ctx *fiber.Ctx) error {
//
//}
//
//func getBooks(ctx *fiber.Ctx) error {
//
//}
//
//func getHotels(ctx *fiber.Ctx) error {
//
//}
//
//func deleteHotelById(ctx *fiber.Ctx) error {
//
//}
//
//func editHotelById(ctx *fiber.Ctx) error {
//
//}
//
//func getHotelById(ctx *fiber.Ctx) error {
//
//}
//
//func addHotel(ctx *fiber.Ctx) error {
//
//}
//
//
//func deleteUserById(ctx *fiber.Ctx) error {
//
//}
//
//func editUserById(ctx *fiber.Ctx) error {
//
//}
//
//func getUserById(ctx *fiber.Ctx) error {
//
//}
//
func addUser(ctx *fiber.Ctx) error {
	user := &models.User{}

	err := ctx.BodyParser(user)

	if err != nil {
		message :=models.Response{Message: err.Error()}

		return ctx.Status(http.StatusNotAcceptable).JSON(message)
	}


	result :=db.Create(user)
	if result.Error !=nil{
		message :=models.Response{Message: err.Error()}
		return ctx.Status(http.StatusNotAcceptable).JSON(message)
	}


	message :=models.Response{Message: "User created"}
	return ctx.JSON(message)
}

func getUsers(ctx *fiber.Ctx) error {
	var users []models.User
	result := db.Find(&users)
	if result.Error != nil {
		fmt.Println(result.Error)
	}

	return ctx.JSON(users)
}
//
//func hello(ctx *fiber.Ctx) error {
//	user1 := models.User{
//		Id:          0,
//		FirstName:   "omid",
//		LastName:    "hojabri",
//		Email:       "",
//		PhoneNumber: "",
//	}
//	return ctx.JSON(user1)
//}
//
