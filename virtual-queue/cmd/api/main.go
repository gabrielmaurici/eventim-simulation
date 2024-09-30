package main

import (
	"fmt"
	"net/http"

	"github.com/gabrielmaurici/eventim-simulation/internal/database"
	"github.com/gabrielmaurici/eventim-simulation/internal/usecase/entry_virtual_queue"
	"github.com/gabrielmaurici/eventim-simulation/internal/web"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-redis/redis"
)

func main() {
	redisDb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	buyersActivesDb := database.NewBuyersActivesDb(redisDb)
	virtualQueueDb := database.NewVirtualQueueDb(redisDb)
	entryQueueUsecase := entry_virtual_queue.NewEntryQueueUseCase(buyersActivesDb, virtualQueueDb)
	webVirtualQueueHandler := web.NewWebVirtualQueueHandler(*entryQueueUsecase)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Post("/api/virtual-queue", webVirtualQueueHandler.EntryQueue)
	fmt.Println("Server is running!")

	http.ListenAndServe(":3000", router)
}
