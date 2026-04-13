package main

import (
	"fmt"

	controller "github.com/ANSHSINGH050404/movie_streaming/controllers"
	"github.com/ANSHSINGH050404/movie_streaming/database"
	"github.com/ANSHSINGH050404/movie_streaming/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello World")

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World")
	})

	r.GET("/movies", controller.GetMovies())
	r.GET("/movies/:imdb_id", controller.GetMovieById())

	// Protected Routes
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/addmovie", controller.AddMovie())
	}

	r.POST("/register", controller.RegisterUser(database.Client))
	r.POST("/login", controller.LoginUser(database.Client))
	if err := r.Run(":8080"); err != nil {
		fmt.Println("Failed to start server: ", err)
	}
}
