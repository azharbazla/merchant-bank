package services

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"merchant-bank/models"
)

var loggedInCustomers = make(map[string]bool)

func LoginCustomer(customer models.Customer) error {
	data, err := ioutil.ReadFile("storage/customer_data.json")
	if err != nil {
		return err
	}

	var customers []models.Customer
	if err := json.Unmarshal(data, &customers); err != nil {
		return err
	}

	for _, c := range customers {
		if c.Email == customer.Email && c.Password == customer.Password {
			loggedInCustomers[customer.ID] = true
			return nil
		}
	}
	return errors.New("customer not found")
}

func ProcessPayment(customerID string, amount float64) error {
	if !loggedInCustomers[customerID] {
		return errors.New("customer not logged in")
	}

	data, err := ioutil.ReadFile("storage/customer_data.json")
	if err != nil {
		return err
	}

	var customers []models.Customer
	if err := json.Unmarshal(data, &customers); err != nil {
		return err
	}

	for i, c := range customers {
		if c.ID == customerID {
			if c.Balance < amount {
				return errors.New("insufficient balance")
			}
			customers[i].Balance -= amount
			break
		}
	}

	newData, err := json.Marshal(customers)
	if err != nil {
		return err
	}
	return ioutil.WriteFile("storage/customer_data.json", newData, 0644)
}

func LogoutCustomer(customerID string) error {
	if !loggedInCustomers[customerID] {
		return errors.New("customer not logged in")
	}

	delete(loggedInCustomers, customerID)
	return nil
}
