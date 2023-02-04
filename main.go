package main

import (
	"Je-Devlop/sunday-api/sunday"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(postgres.Open("postgres://sunday:secret@localhost:5432/sunday"))
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&sunday.Scoop{})

	scoopsHandler := sunday.NewSundayHandler(db)

	r := gin.Default()
	r.POST("/create-scoops", scoopsHandler.CreateScoops)

	server := newServer(r)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	gracefullyShuttingDown(server)
}

func gracefullyShuttingDown(server *http.Server) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()
	stop()
	fmt.Println("shutting down gracefully, press Ctrl+C again to force")

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(timeoutCtx); err != nil {
		fmt.Println(err)
	}
}

func newServer(r http.Handler) *http.Server {
	return &http.Server{
		Addr:              ":" + os.Getenv("PORT"),
		Handler:           r,
		ReadTimeout:       60 * time.Second,
		WriteTimeout:      300 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}
}
