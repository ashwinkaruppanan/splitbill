package main

import (
	"log"
	"os"
	"splitbill/api/handler"
	"splitbill/api/router"
	"splitbill/db"
	"splitbill/internal/repository"
	"splitbill/internal/service"
	"splitbill/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	smtpUsername := os.Getenv("SMTP_USERNAME")
	smtpPassword := os.Getenv("SMTP_PASSWORD")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		log.Fatal(err)
	}

	mailObj := utils.NewMailService(smtpHost, smtpPort, smtpUsername, smtpPassword)

	newDB := db.NewDB()
	db := newDB.GetDB()

	rep := repository.NewUserRepository(db)
	ser := service.NewUserService(rep, mailObj)
	handler := handler.NewUserHandler(ser)

	router.NewRouter(r, handler)

	err = r.Run(":" + os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}

	// query := "INSERT INTO users (name, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)"
	// _, err = db.Exec(query, "bob", "bob@gmail.com", "password", time.Now(), time.Now())
	// if err != nil {
	// 	log.Fatal(err)
	// }

}
