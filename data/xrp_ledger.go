package data

import (
	"encoding/json"
	"flag"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"os"
	"os/signal"
)

var lastLedger = 51366888

func GetLedgerData(addr *string, indexLedger int) []Transaction {
	flag.Parse()
	log.SetFlags(0)
	// check for interrupts and cleanly close the connection
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/"}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	var lr LedgerRequest
	lr.Id = indexLedger
	lr.Command = "ledger"
	lr.LedgerIndex = "validated"
	lr.Full = false
	lr.Accounts = false
	lr.Transactions = true
	lr.Expand = true
	lr.OwnerFunds = false

	msg, _ := json.Marshal(lr)
	err = c.WriteMessage(websocket.TextMessage, []byte(string(msg)))
	if err != nil {
		log.Println("write:", err)
	}

	_, message, err := c.ReadMessage()
	if err != nil {
		log.Println("read:", err)
	}
	response := &LedgerResponseExpanded{}
	err = json.Unmarshal(message, response)
	if err != nil {
		panic("Error unmarshalling" + err.Error())
	}

	transactionsStruct := response.Result.Ledger.Transactions
	return transactionsStruct

}
