package main

import (
	"database/sql"
	"log"

	"github.com/Gin-CURD-API/internal/controllers"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	router := gin.Default()
	port := ":3000"

	var db *sql.DB
	var e error

	if db, e = sql.Open("sqlite3", "./data.db"); e != nil {
		log.Fatalf("Error: %v", e)
	}

	defer db.Close()

	if e := db.Ping(); e != nil {
		log.Fatalf("Error: %v", e)

	}

	router.GET("/products", controllers.GetProducts(db))
	router.GET("/products/:id", controllers.GetProduct(db))
	router.POST("/products", controllers.AddProducts(db))
	router.PUT("/products/:id", controllers.UpdateProducts(db))
	router.DELETE("/products/:id", controllers.DelProducts(db))

	log.Fatal(router.Run(port))
}
