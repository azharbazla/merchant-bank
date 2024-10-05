package main

import (
	"log"
	"merchant-bank/controllers"
	"net/http"
)

func main() {
	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/payment", controllers.Payment)
	http.HandleFunc("/logout", controllers.Logout)
	http.HandleFunc("/history", controllers.GetAllHistory)
	http.HandleFunc("/history/customer", controllers.GetCustomerHistory)

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
