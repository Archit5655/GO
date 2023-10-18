package main

import (
	// "context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/archit5655/go-gorm-postgress/models"
	"github.com/archit5655/go-gorm-postgress/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"gorm.io/gorm"
)

type Book struct {
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
}

type Repo struct {
	DB *gorm.DB
}

func (r *Repo) CreateBook(context *fiber.Ctx) error {
	book := Book{}
	err := context.BodyParser(&book)
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}
	err = r.DB.Create(&book).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create book"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "book has been added"})
	return nil

}

func (r *Repo) GetBooks(context *fiber.Ctx) error {
	bookModels := &[]models.Books{}
	err := r.DB.Find(bookModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Could not get the books"})
		return err
	}
	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "Books fetched successsfully", "data": bookModels})

	return nil

}
func (r *Repo) DeleteBook(context *fiber.Ctx) error {
	bookModel := models.Books{}
	id := context.Params("id")
	if id == "" {

		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "id canno tbe  empty "})
		return nil
	}
	err := r.DB.Delete(bookModel, id)
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not delete book"})
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "books delete successfully"})
	return err.Error
}
func (r *Repo) GetBookByID(context *fiber.Ctx) error {
	id := context.Params("id")
	bookModels := &models.Books{}
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "id cannot be empty"})
		return nil

	}
	fmt.Println("THe id is", id)
	err := r.DB.Where("id = ?", id).First(bookModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not ge the book"})
		return err

	}
	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "books id fetched successfully", "data": bookModels})
	return nil

}
func (r *Repo) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_books", r.CreateBook)
	api.Delete("/delete_book/:id", r.DeleteBook)
	api.Get("/get_books/:id", r.GetBookByID)
	api.Get("/get_books", r.GetBooks)

}
func main() {
	// this allows to use enviornment varavbles from .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Passwrod: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		Dbname:   os.Getenv("DB_NAME"),
		SSLModel: os.Getenv("DB_SSLMODE"),
	}
	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal(err)
	}

	err = models.MigrateBooks(db)
	if err != nil {
		log.Fatal(err)

	}

	r := Repo{
		DB: db,
	}

	app := fiber.New()

	r.SetupRoutes(app)
	app.Listen(":4000")

}
