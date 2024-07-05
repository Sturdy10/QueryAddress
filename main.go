package main

import (
	"queryAddress/database"
	"queryAddress/handlers"
	"queryAddress/repositories"
	"queryAddress/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db := database.Postgresql()
	defer db.Close()

	r := repositories.NewRepositoryAdapter(db)
	s := services.NewServiceAdapter(r)
	h := handlers.NewHanerhandlerAdapter(s)

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "X-Auth-Token", "Authorization"}
	router.Use(cors.New(config))

	router.GET("/api/getCountries", h.GetCountriesHandlers)
	router.GET("/api/getProvinceAmphoeTambonZipcode", h.GetProvinceAmphoeTambonZipcodeHandlers)

	err := router.Run(":8888")
	if err != nil {
		panic(err.Error())
	}
}
