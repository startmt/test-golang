package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/startmt/test-golang/test/blockchain"
	"github.com/startmt/test-golang/test/cmd/driver"
	"log"
)

func main() {
	log.Println("START SERVER")
	err := startServer()
	if err != nil {
		log.Println(err.Error())
	}
}

func startServer() error {
	app := fiber.New()

	mongo, err := driver.ConnectMongo()
	if err != nil {
		return err
	}
	log.Printf("MONGODB CONNECTED TO DATABASE: %s", mongo.DB.Name())

	blockchain.Route(app, blockchain.CreateCollection(*mongo))

	return app.Listen(":3000")
}
