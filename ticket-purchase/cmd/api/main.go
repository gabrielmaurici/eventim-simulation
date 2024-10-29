package main

import (
	"database/sql"
	"fmt"
	"math/rand/v2"
	"net/http"

	"github.com/gabrielmaurici/eventim-simulation/ticket-purchase/internal/database"
	"github.com/gabrielmaurici/eventim-simulation/ticket-purchase/internal/usecase/buy_tickets"
	"github.com/gabrielmaurici/eventim-simulation/ticket-purchase/internal/usecase/reserve_ticket"
	"github.com/gabrielmaurici/eventim-simulation/ticket-purchase/internal/web"
	"github.com/gabrielmaurici/eventim-simulation/ticket-purchase/pkg/rabbitmq"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
	amqp "github.com/rabbitmq/amqp091-go"
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

	rabbitmqConn, err := amqp.Dial("amqp://guest:guest@rabbitmq-virtual-queue:5672/")
	if err != nil {
		panic(fmt.Errorf("erro ao conectar rabbitmq: %w", err))
	}
	defer rabbitmqConn.Close()

	producer, err := rabbitmq.NewProducer(rabbitmqConn, "buy_tickets_exchange", "direct", "buy_tickets_routing_key")
	if err != nil {
		panic(fmt.Errorf("erro ao criar produtor rabbitmq: %w", err))
	}

	ticketDb := database.NewTicketDb(mysqlDb)
	ticketReservationDb := database.NewTicketReservationDb(redisDb)
	reserveTicketUseCase := reserve_ticket.NewReserveTicket(ticketDb, ticketReservationDb)
	buyTicketsUseCase := buy_tickets.NewBuyTicketsUseCase(ticketReservationDb, producer, func() bool {
		return rand.Float64() < 0.5
	})
	webTicketsHandler := web.NewWebTicketsHandler(*reserveTicketUseCase, *buyTicketsUseCase)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Post("/api/tickets/reserve", webTicketsHandler.Reserve)
	router.Post("/api/tickets/purchase", webTicketsHandler.Purchase)
	fmt.Println("Server is running!")

	http.ListenAndServe(":3001", router)
}
