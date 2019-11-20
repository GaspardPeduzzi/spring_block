package main

import (
	"github.com/gaspardpeduzzi/spring_block/display"
	"log"
	"sync"

	"github.com/gaspardpeduzzi/spring_block/data"
	"github.com/gaspardpeduzzi/spring_block/graph"
)


var oldestIndex = 0


type Optimizer struct {
	Endpoint string
	Graph    graph.Graph
	Channel  chan int
}

func NewOptimizer(endpoint string, c chan int) *Optimizer {
	graph := graph.Graph{
		NGraph: make(map[string]map[string]*graph.OrderBook),
		AccountRoots: make(map[string]map[int]*graph.Offer),
		Lock:  sync.RWMutex{},
	}
	return &Optimizer{endpoint, graph, c}
}


func (lo *Optimizer) NConstructTxGraph() {
	lastIndex := data.GetLastLedgerSeq(&lo.Endpoint)
	if lastIndex > oldestIndex {
		display.DisplayVerbose("New block index: ", lastIndex)
		oldestIndex = lastIndex
		txs := data.GetLedgerData(&lo.Endpoint, lastIndex)
		var tmpCreate []data.Transaction
		for _, tx := range txs {
			if tx.TransactionType == "OfferCreate" {
				tmpCreate = append(tmpCreate, tx)
			} else if tx.TransactionType == "OfferCancel" {
				tmpCreate = append(tmpCreate, tx)
			}
		}
		lo.ParseOfferCreateTransactions(tmpCreate)


		//lo.Channel <- 1
		lo.NConstructTxGraph()
	}
	lo.NConstructTxGraph()
}



func (lo *Optimizer) ParseOfferCreateTransactions(transactions []data.Transaction) {
	log.Println("ADDED", len(transactions), " OfferCreate transaction(s)")
	for _, tx := range transactions {
		lo.Graph.ParseOfferCreate(tx)
	}
	//lo.Graph.SortGraphWithTxs()
}


func (lo *Optimizer) parseTransactionsCancel(transactions []data.Transaction) {
	display.DisplayVerbose("ADDED", len(transactions), " OfferCancel transaction(s)")

	deleted := make([]string,1)


	for _, tx := range transactions {
		for _, v := range tx.MetaData.AffectedNodes {
			//log.Println("====================================================================================")
			//display.DisplayVerbose("MODIFIED", v.ModifiedNode.LedgerIndex, "CREATE", v.CreatedNode.LedgerIndex, "DELETED", v.DeletedNode.LedgerIndex)
			//display.DisplayVerbose(v.DeletedNode.FinalFields.PreviousTxnID)
			if v.DeletedNode.FinalFields.PreviousTxnID != "" {
				deleted = append(deleted, v.DeletedNode.FinalFields.PreviousTxnID)
			}
			//log.Println("====================================================================================")
		}
	}
	if deleted[0] != "" {
		display.DisplayVerbose("DELETING OFFERS")
		//lo.Graph.DeleteOffers(deleted)
	}

}





func (lo *Optimizer) parsePay(transactions []data.Transaction){
	log.Println("ADDED", len(transactions), " Payment transaction(s)")

}
