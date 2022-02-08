package routers

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/swkkd/crud/database"
	"github.com/swkkd/crud/handlers"
	"log"
	"net/http"
)

func Setup() {
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

	handler := cors.Default().Handler(r)

	if err := http.ListenAndServe(":8080", handler); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
	log.Printf("Finished")
}
