package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/zoezn/Final-Backend/cmd/server/handler"
	"github.com/zoezn/Final-Backend/internal/dentista"
	"github.com/zoezn/Final-Backend/pkg/store"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al intentar cargar archivo .env")
	}
	storage := store.NewJsonStore("./dentistas.json")

	repo := dentista.NewRepository(storage)
	service := dentista.NewService(repo)
	productHandler := handler.NewProductHandler(service)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	r.GET("/hola", func(c *gin.Context) { c.String(200, "hola mundo :p") })
	products := r.Group("/products")
	{
		products.GET(":id", productHandler.GetByID())
		products.POST("", productHandler.Post())
		products.DELETE(":id", productHandler.Delete())
		products.PATCH(":id", productHandler.Patch())
		products.PUT(":id", productHandler.Put())
	}

	r.Run(":8080")
}
