package app

import (
	Services "MesseageMicroService/restApi/Services"
	"encoding/json"
	"log"
	"net/http"
)

// We define the handlers by wrapping them in a struct
// Customerhandlers in struct and we create the receiver functions (handlers)
// Which contains the customer service
type MessageHandlers struct {
	Service Services.MessageService
}

func (ch *MessageHandlers) GetAllMessages(w http.ResponseWriter, r *http.Request) {
	customers, err := ch.Service.GetAllMessages()
	if err != nil {
		log.Println(err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
