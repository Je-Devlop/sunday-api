package main

import (
	"Je-Devlop/sunday-api/router"
	"Je-Devlop/sunday-api/store/db"
	"Je-Devlop/sunday-api/sunday"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
)

func main() {
	liveFile := "/tmp/live"
	_, err := os.Create(liveFile)
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove("/tmp/live")

	err = godotenv.Load("local.env")
	if err != nil {
		log.Printf("please consider environment variable: %s\n", err)
	}

	store, err := db.NewPostgres(os.Getenv("DB_CONN"))
	if err != nil {
		panic(err.Error())
	}

	r := router.NewMyRouter()
	r.Static("/images", "public/images")

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"http://localhost:3000",
	}
	config.AllowHeaders = []string{
		"Origin",
		"Authorization",
		"content-type",
	}
	r.Use(cors.New(config))

	scoopsHandler := sunday.NewSundayHandler(store)

	r.POST("/create-scoops", scoopsHandler.CreateScoops)
	r.GET("/scoops", scoopsHandler.GetSundayScoops)
	r.POST("/create-toppings", scoopsHandler.CreateTopping)
	r.GET("/toppings", scoopsHandler.GetSundayTopping)
	r.POST("/order", scoopsHandler.OrderIceCream)

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
