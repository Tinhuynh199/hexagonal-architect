package app

import (
	"database/sql"
	"github.com/gorilla/mux"
	"context"
	"log"
	_ "github.com/go-sql-driver/mysql"

	userHandler "hexrestapi1/internal/infrastructure/adapter/user/handler"
	userRepo "hexrestapi1/internal/infrastructure/adapter/user/repository"
	. "hexrestapi1/internal/infrastructure/port/user"
	userService "hexrestapi1/internal/infrastructure/service/user"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	studentHandler "hexrestapi1/internal/infrastructure/adapter/student/handler"
	studentRepo "hexrestapi1/internal/infrastructure/adapter/student/repository"
	. "hexrestapi1/internal/infrastructure/port/student"
	studentService "hexrestapi1/internal/infrastructure/service/student"
)

type App struct {
	Router *mux.Router
	User UserTransport
	Student StudentTransport
}

func (a *App) Initialize(config *Config) {

	// User Application
	sqlDB, err := sql.Open(config.Sql.Driver, config.Sql.Data_Source)
	if err != nil {
		log.Fatal("Could not connect to MySQL Database !!!")
	}

	userRepository := userRepo.NewUserAdapter(sqlDB)
	userService := userService.NewUserService(sqlDB, userRepository)
	a.User = userHandler.NewUserHandler(userService)
	
	// Student Application
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(config.Mongo.URI))
	if err != nil {
		log.Fatal("Could not connect to Mongo Database !!!")
	}
	mongoDB := client.Database(config.Mongo.Database)

	studentRepository := studentRepo.NewStudentAdapter(mongoDB)
	studentService := studentService.NewStudentService(mongoDB, studentRepository)
	a.Student = studentHandler.NewStudentHandler(studentService)

	a.Router = mux.NewRouter()
	a.setRouters()
}


