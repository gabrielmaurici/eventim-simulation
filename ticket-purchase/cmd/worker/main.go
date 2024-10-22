package main

import (
	"database/sql"
	"fmt"

	"github.com/gabrielmaurici/eventim-simulation/ticket-purchase/internal/database"
	"github.com/gabrielmaurici/eventim-simulation/ticket-purchase/internal/usecase/processing_expired_tickets"
	"github.com/gabrielmaurici/eventim-simulation/ticket-purchase/internal/worker"
	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
)

func main() {
	redisDb := redis.NewClient(&redis.Options{
		Addr:     "redis-ticket-purchase:6379",
		Password: "",
		DB:       0,
	})

	mysqlDb, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "mysql-ticket-purchase", "3306", "eventim"))
	if err != nil {
		panic(err)
	}
	defer mysqlDb.Close()

	ticketDb := database.NewTicketDb(mysqlDb)
	ticketReservationDb := database.NewTicketReservationDb(redisDb)
	processingExpiredTicketsUseCase := processing_expired_tickets.NewProcessingExpiredTicketsUseCase(ticketDb, ticketReservationDb)
	processingExpiredTicketsWorker := worker.NewProcessingExpiredTicketsWorker(*processingExpiredTicketsUseCase)
	fmt.Println("Worker is running!")
	processingExpiredTicketsWorker.Start()
}
