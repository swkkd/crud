package routers

import (
	"context"
	"flag"
	"github.com/gorilla/mux"
	"github.com/swkkd/crud/database"
	"github.com/swkkd/crud/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Setup() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	r := mux.NewRouter()

	handleFunc := &handlers.DbConn{
		DB: database.GetDB(),
	}
	r.HandleFunc("/", handleFunc.GetCustomers).Methods("GET")
	r.HandleFunc("/customer", handleFunc.GetCustomer).Methods("GET").Queries("id", "{id}")
	r.HandleFunc("/delete", handleFunc.DeleteCustomer).Queries("id", "{id}")
	r.HandleFunc("/create", handleFunc.CreateCustomer)
	r.HandleFunc("/update", handleFunc.UpdateCustomer).Queries("id", "{id}")
	r.HandleFunc("/search", handleFunc.SearchCustomers).Queries("search", "{search}")
	r.HandleFunc("/search", handleFunc.SearchCustomers)

	//handler := cors.Default().Handler(r)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("error shutting down server %s", err)
	} else {
		log.Println("Server gracefully stopped")
	}
}
