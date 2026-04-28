package main

import (
	"bank/internal/handlers"
	"bank/internal/service"
	"bank/internal/storage"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	store := storage.New()
	svc := service.New(store)
	h := handlers.New(svc)

	handlers.Init(r, h)

	// UI
	r.Static("/web", "./web")
	r.GET("/", func(c *gin.Context) {
		c.File("./web/index.html")
	})

	log.Println("SRBank started on :8080")

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
