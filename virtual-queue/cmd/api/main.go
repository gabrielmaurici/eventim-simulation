package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gabrielmaurici/eventim-simulation/internal/database"
	"github.com/gabrielmaurici/eventim-simulation/internal/usecase/entry_virtual_queue"
	"github.com/gabrielmaurici/eventim-simulation/internal/web"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/redis/go-redis/v9"
)

func main() {
	redisDb := redis.NewClient(&redis.Options{
		Addr:     "redis-virtual-queue:6379",
		Password: "",
		DB:       0,
	})

	ctx := context.Background()
	buyersActivesDb := database.NewBuyersActivesDb(redisDb, ctx)
	virtualQueueDb := database.NewVirtualQueueDb(redisDb, ctx)
	entryQueueUsecase := entry_virtual_queue.NewEntryQueueUseCase(buyersActivesDb, virtualQueueDb)
	webVirtualQueueHandler := web.NewWebVirtualQueueHandler(*entryQueueUsecase)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Post("/api/virtual-queue", webVirtualQueueHandler.EntryQueue)
	fmt.Println("Server is running!")

	http.ListenAndServe(":3000", router)
}
