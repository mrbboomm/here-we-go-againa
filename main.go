package main

import (
	"fmt"
	"go-nf/config"
	"go-nf/deliveries"
	"go-nf/entities"
	"go-nf/kafka/producer"
	"go-nf/mongodb"
	repositories "go-nf/repositories/user"
	"go-nf/tier"
	usecases "go-nf/usecases/user"
	"go-nf/user"
	"go-nf/utils"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	//  Load env
	if err := godotenv.Load("local.env"); err != nil {
		fmt.Println("NOT HAVE LOCAL ENV")
	}

	KAFKA_HOST := os.Getenv("KAFKA_HOST")
	// Connection part
	cfg := config.KafkaConnCfg{
		Url:    KAFKA_HOST,
		Topics: config.KafkaTopics,
	}
	kafkaHandler := utils.KafkaConn(&cfg)

	// Check topics
	if topics := utils.ListTopic(kafkaHandler.Conn); len(topics) == 0 {
		utils.CreateTopic(kafkaHandler.Conn)
	}

	tier := &tier.Tier{Id: 1, Name: tier.Lang{En: "t", Th: "a"}}
	user := &user.User{Username: "hello", Password: "world", Tier: tier}
	fmt.Println("hello world")
	fmt.Println(user)
	fmt.Println(user.Tier)

	// Initialize Fiber
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/kafka/list-topic", kafkaHandler.GetListTopic)
	app.Post("/kafka/topic", kafkaHandler.CreateTopics)
	app.Delete("/kafka/topic", kafkaHandler.DeleteTopic)
	app.Post("/kafka/producer", producer.SendMassage)

	// connect mongo
	mongodb.ConnectToMongo()
	app.Post("/create-user", mongodb.CreateUserLogin)
	app.Get("/user", mongodb.GetAllUserLogin)
	app.Get("/user/:username", mongodb.GetUserLoginByUsername)
	app.Get("/user-id/:id", mongodb.GetUserLoginById)
	app.Put("/update-user/:id", mongodb.UpdateUserLoginById)
	app.Delete("/delete-user/:id", mongodb.DeleteUserLoginById)

	// mock users data
	users := []entities.UserEntity{{Id: "1", Name: "name1"}, {Id: "2", Name: "name2"}, {Id: "3", Name: "name3"}}
	usersRepo := repositories.NewUserRepo(users)
	userUseCase := usecases.NewUserUseCase((usersRepo))
	userHandlers := deliveries.NewUserHandler((userUseCase))

	app.Get("/users", userHandlers.GetAllUsers)

	app.Listen(":3000")

}
