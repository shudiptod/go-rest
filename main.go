// main.go
package main

import (
	"pcshop/app/routes"
	"pcshop/config"

	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    db, err := config.InitDB()
    if err != nil {
		panic(err)
        panic("Failed to connect to database")
    }
    defer db.Close()

    routes.SetupRoutes(router)

    router.Run(":8080")
}
