package database

import (
	"github.com/swkkd/crud/models"
	"gorm.io/gorm"
	"log"
)

func GetCustomers(db *gorm.DB) ([]models.Customer, error) {
	var customers []models.Customer
	if err := db.Order("id").Find(&customers).Error; err != nil {
		return nil, err
	}
	return customers, nil
}
func GetCustomerByID(id string, db *gorm.DB) (models.Customer, bool, error) {
	var customer models.Customer
	err := db.First(&customer, id).Error
	if err != nil && gorm.ErrRecordNotFound != nil {
		return models.Customer{}, false, err
	}
	return customer, true, nil
}
func DeleteCustomer(id string, db *gorm.DB) error {
	var customer models.Customer
	if err := db.Delete(&customer, id).Error; err != nil {
		return err
	}
	return nil
}
func CreateCustomer(customer models.Customer, db *gorm.DB) error {
	if err := db.Create(&customer); err.Error != nil {
		log.Println(err.Error)
		log.Printf("created customer: %v", customer)
		return err.Error
	}
	return nil
}
func UpdateCustomer(customer models.Customer, db *gorm.DB) error {
	if err := db.Save(&customer).Error; err != nil {
		return err
	}
	return nil
}
func SearchCustomers(query string, db *gorm.DB) ([]models.Customer, error) {
	var customers []models.Customer
	query = "%" + query + "%"
	if err := db.Where("lower(first_name) LIKE ? OR lower(last_name) LIKE ?", query, query).Find(&customers).Error; err != nil {
		return nil, err
	}
	return customers, nil
}
