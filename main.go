package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port, err := strconv.Atoi(os.Getenv("CASSANDRA_PORT"))
	if err != nil {
		log.Fatal(err)
	}

	cluster := gocql.NewCluster(os.Getenv("CASSANDRA_HOST"))
	cluster.Port = port
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: os.Getenv("CASSANDRA_USERNAME"),
		Password: os.Getenv("CASSANDRA_PASSWORD"),
	}
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()
	controller := controller.NewHotelController(repo)

	router := gin.Default()
	router.POST("/hotels", controller.AddHotel)
	router.PUT("/hotels/:name", controller.UpdateHotel)
	router.GET("/hotels/:name", controller.GetHotel)

	router.Run()
	fmt.Println("Connected to Cassandra!")
}
