package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"merchant-bank/models"
	"os"
	"time"
)

func LogHistory(action, customerID string) {
	history := models.History{
		CustomerID: customerID,
		Action:     action,
		Timestamp:  time.Now(),
	}

	data, err := ioutil.ReadFile("storage/history_log.json")
	if err != nil && !os.IsNotExist(err) {
		log.Println("Error reading history log:", err)
		return
	}

	var historyLogs []models.History
	if len(data) != 0 {
		if err := json.Unmarshal(data, &historyLogs); err != nil {
			log.Println("Error unmarshalling history log:", err)
			return
		}
	}

	historyLogs = append(historyLogs, history)

	newData, err := json.Marshal(historyLogs)
	if err != nil {
		log.Println("Error marshalling history log:", err)
		return
	}

	if err := ioutil.WriteFile("storage/history_log.json", newData, 0644); err != nil {
		log.Println("Error writing history log:", err)
	}
}
