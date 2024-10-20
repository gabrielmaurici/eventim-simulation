package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gabrielmaurici/eventim-simulation/ticket-purchase/internal/database"
	"github.com/gabrielmaurici/eventim-simulation/ticket-purchase/internal/usecase/reserve_ticket"
	"github.com/gabrielmaurici/eventim-simulation/ticket-purchase/internal/web"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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
	reserveTicketUseCase := reserve_ticket.NewReserveTicket(ticketDb, ticketReservationDb)
	webTicketsReservationHandler := web.NewWebTicketsReservationHandler(*reserveTicketUseCase)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Post("/api/tickets/reserve", webTicketsReservationHandler.Reserve)
	fmt.Println("Server is running!")

	http.ListenAndServe(":3001", router)
}
