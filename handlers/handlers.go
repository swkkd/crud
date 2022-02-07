package handlers

import (
	_ "database/sql"
	"encoding/json"
	_ "encoding/json" // package to encode and decode the json into struct and vice versa
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/swkkd/crud/database"
	"github.com/swkkd/crud/models"
	"gorm.io/gorm"
	"html/template"
	"log"
	_ "log"
	"net/http" // used to access the request and response object of the api
	"strconv"
)

type DbConn struct {
	DB *gorm.DB
}

func (db *DbConn) GetCustomers(w http.ResponseWriter, r *http.Request) {

	customers, err := database.GetCustomers(db.DB)
	if err != nil {
		log.Fatal(err)
	}
	tmpl, err := template.ParseFiles("static/customers.html")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	if err := tmpl.Execute(w, customers); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func (db *DbConn) GetCustomer(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	customer, exists, err := database.GetCustomerByID(id, db.DB)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if !exists {
		http.Error(w, err.Error(), 400)
	}

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	tmpl, err := template.ParseFiles("static/customers.html")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	if err := tmpl.Execute(w, customer); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

}

func (db *DbConn) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	log.Printf("ID %v", id)
	_, exists, err := database.GetCustomerByID(id, db.DB)
	if !exists {
		log.Println(err)
		http.NotFound(w, r)
		return
	}
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 400)
		return
	}
	err = database.DeleteCustomer(id, db.DB)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 400)
		return
	}

	http.Redirect(w, r, "/", 301)

}

func (db *DbConn) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var customer models.Customer
	if r.Method == "POST" {
		customer.FirstName = r.FormValue("FirstName")
		customer.LastName = r.FormValue("LastName")
		customer.Email = r.FormValue("Email")
		customer.Gender = r.FormValue("Gender")
		customer.Birthday = r.FormValue("Birthday")
		customer.Address = r.FormValue("Address")

	}
	json.NewDecoder(r.Body).Decode(&customer)
	if err := database.CreateCustomer(customer, db.DB); err != nil {
		log.Println(err)
		return
	}
	http.Redirect(w, r, "/", 301)
}
func (db *DbConn) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	customer, exists, err := database.GetCustomerByID(id, db.DB)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 400)
		return
	}
	if !exists {
		http.NotFound(w, r)
		return
	}
	var updatedCustomer models.Customer
	if r.Method == "POST" {
		log.Printf("FormValue: %v", r.FormValue("Index"))
		customerID, err := strconv.ParseUint(r.FormValue("Index"), 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		updatedCustomer.ID = uint(customerID)
		updatedCustomer.FirstName = r.FormValue("FirstName")
		updatedCustomer.LastName = r.FormValue("LastName")
		updatedCustomer.Email = r.FormValue("Email")
		updatedCustomer.Gender = r.FormValue("Gender")
		updatedCustomer.Birthday = r.FormValue("Birthday")
		updatedCustomer.Address = r.FormValue("Address")
		customer = updatedCustomer
		err = database.UpdateCustomer(updatedCustomer, db.DB)
		if err != nil {
			log.Fatal(err)
		}
		http.Redirect(w, r, "/", 301)
	}
	tmpl, err := template.ParseFiles("static/edit.html")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if err := tmpl.Execute(w, customer); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

}
