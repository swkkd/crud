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
	r.HandleFunc("/create/customer", handleFunc.CreateCustomer).Methods("POST")
	r.HandleFunc("/update", handleFunc.UpdateCustomer).Queries("id", "{id}")
	handler := cors.Default().Handler(r)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
