package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/ovalb/memory/database"
	"github.com/ovalb/memory/handler"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	c := &handler.Repository{
		Database: database.New(),
	}

	c.Database.AutoMigrate(&handler.Item{})
	// c.Database.AutoMigrate(&handler.Tag{})

	r := gin.Default()

	r.GET("/:id", c.GetItemById)
	r.POST("/items", c.AddItem)

	http.Handle("/", r)

	go func() {
		fmt.Println("server starting at port 8080")
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	fmt.Println("now waiting for sigterm or something")
	s := <-sigs
	log.Fatalf("Got signal: %v. Bye bye.", s)

}
