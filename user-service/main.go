package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hotdog132/graphql-todo-practice/user-service/user/delivery"
	"github.com/hotdog132/graphql-todo-practice/user-service/user/usecase"

	"github.com/hotdog132/graphql-todo-practice/user-service/config"
	"github.com/hotdog132/graphql-todo-practice/user-service/user/repository"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

func main() {
	configs, err := config.NewConfigurations("development")
	if err != nil {
		log.Println("Get configs error:", err)
	}

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		configs.Database.Host,
		configs.Database.Port,
		configs.Database.User,
		configs.Database.Password,
		configs.Database.DBName)

	dbConn, err := gorm.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer dbConn.Close()

	dbConn.DB().SetMaxOpenConns(configs.Database.MaxOpenConns)
	dbConn.DB().SetMaxIdleConns(configs.Database.MaxIdleConns)

	e := echo.New()
	ur := repository.NewPsqlUserRepository(dbConn)
	uu := usecase.NewUserUsecase(ur)
	delivery.NewHttpUserHandler(e, uu)
	e.Logger.Fatal(e.Start(":7070"))
}
