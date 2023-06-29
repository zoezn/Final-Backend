package main

import (
	"database/sql"
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

	db, err := sql.Open("mysql", "user1:secret_password@/my_db")

	if err != nil {
		panic(err.Error())
	}

	errPing := db.Ping()
	if errPing != nil {
		panic(errPing.Error())
	}

	// storage := store.SqlStore{db}

	repo := dentista.NewRepository(storage)
	service := dentista.NewService(repo)
	dentistaHandler := handler.NewDentistaHandler(service)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	dentistas := r.Group("/dentistas")
	{
		dentistas.GET(":id", dentistaHandler.GetByID())
		dentistas.POST("", dentistaHandler.Post())
		dentistas.DELETE(":id", dentistaHandler.Delete())
		dentistas.PATCH(":id", dentistaHandler.Patch())
		dentistas.PUT(":id", dentistaHandler.Put())
	}

	r.Run(":8080")
}
