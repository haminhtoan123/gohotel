package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	controllers "github.com/haminhtoan123/gohotel/controllers"
	repo "github.com/haminhtoan123/gohotel/repo"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
	"github.com/joho/godotenv"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

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
	fmt.Println("Connected to Cassandra!")

	cassandra := repo.NewCassandraHotelRepository(session)
	controller := controllers.NewHotelController(cassandra)

	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := router.Group("/api/v1")
	{
	   eg := v1.Group("/example")
	   {
		  eg.GET("/helloworld",Helloworld)
	   }
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.POST("/hotels", controller.AddHotel)
	router.PUT("/hotels/:name", controller.UpdateHotel)
	router.GET("/hotels/:name", controller.GetHotel)

	router.Run()
	defer session.Close()
}
