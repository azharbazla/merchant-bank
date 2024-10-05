package services

import (
	"encoding/json"
	"io/ioutil"
	"merchant-bank/models"
)

func GetHistory() ([]models.History, error) {
	data, err := ioutil.ReadFile("storage/history_log.json")
	if err != nil {
		return nil, err
	}

	var historyLogs []models.History
	if err := json.Unmarshal(data, &historyLogs); err != nil {
		return nil, err
	}

	return historyLogs, nil
}

func GetCustomerHistory(customerID string) ([]models.History, error) {
	allHistory, err := GetHistory()
	if err != nil {
		return nil, err
	}

	var customerHistory []models.History
	for _, h := range allHistory {
		if h.CustomerID == customerID {
			customerHistory = append(customerHistory, h)
		}
	}

	return customerHistory, nil
}
