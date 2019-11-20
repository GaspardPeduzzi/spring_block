package graph

import (
	"math"
	"sort"
	"sync"

	"github.com/gaspardpeduzzi/spring_block/data"
)

var capacityList = 0

// Graph : Data structure for graph of offers
type Graph struct {
	Graph map[string]map[string]*TxList

	NGraph map[string]map[string]*OrderBook
	AccountRoots map[string]map[int]*Offer
	Lock  sync.RWMutex
}

// SimplerGraph : Data structure for graph of best offers (nxn grid, -log(edge))
type SimplerGraph struct {
	Graph      map[string]map[string]float64
	Currencies []string
	Lock       sync.Mutex
}

type OrderBook struct {
	List []*Offer
	Pay string
	WillGet string
	//numberOfTransactions int
}

func (ng *Graph) insertNewOffer(offer *Offer){
	//graph BTC ETH donne offerCreate pour obtenir ETH en payant BTC
	//TODO: check if correct here for the A to B "policy"
	ng.initNGraph(offer.CreatorWillPay, offer.CreatorWillGet)
	ng.Lock.Lock()
	ng.NGraph[offer.CreatorWillPay][offer.CreatorWillGet].List = append(ng.NGraph[offer.CreatorWillPay][offer.CreatorWillGet].List, offer)
	ng.Lock.Unlock()
}



/*
func (ng *Graph) deleteOffer(offer *Offer){

	if ng.AccountRoots[offer.Account] == nil {
		display.DisplayVerbose("We are done since we haven't seen this guy yet")
	} else if ng.AccountRoots[offer.Account][offer.SequenceNumber] == nil {
		display.DisplayVerbose("DELETING ====================================================================================")
		for k, _ := range ng.AccountRoots[offer.Account] {
			display.DisplayVerbose(k)
		}
		display.DisplayVerbose("====================================================================================")

		//display.DisplayVerbose("We are done since we haven't seen this guy with this number sequence yet")
	} else if ng.AccountRoots[offer.Account][offer.SequenceNumber] != nil  {
		log.Println("ACTUALLY DELETED?")

		//display.DisplayVerbose(ng.Graph[offer.CreatorWillPay][offer.CreatorWillGet])
		ng.AccountRoots[offer.Account][offer.SequenceNumber] = nil
	}

	//log.Println("DELETED previous offer from", offer.Account, "with seq #", offer.SequenceNumber)

}*/


func (ng *Graph) insertNewOfferToAccount(offer *Offer) {
	if ng.AccountRoots[offer.Account] == nil {
		ng.AccountRoots[offer.Account] = make(map[int]*Offer)
	} else if ng.AccountRoots[offer.Account][offer.SequenceNumber] == nil {
		ng.AccountRoots[offer.Account][offer.SequenceNumber] = offer
	}
}




// TxList : Data structure for list of offers
type TxList struct {
	List []*Offer
}


//graph BTC ETH gives offerCreate pour obtenir ETH en payant BTC
// Offer : Data structure for an offer
type Offer struct {
	//Indexing
	XrpTx    		  data.Transaction
	TxHash   		  string
	Account  		  string
	SequenceNumber 	  int

	//Actual transaction data
	Rate   			  float64
	Quantity 		  float64
	CreatorWillPay    string
	CreatorWillGet    string
	Issuer			  string
}



func (graph *Graph) initNGraph(pay string, get string) {
	graph.Lock.Lock()

	if graph.NGraph[pay] == nil {
		graph.NGraph[pay] = make(map[string]*OrderBook)
		if graph.NGraph[pay][get] == nil {
			txlist := make([]*Offer, capacityList)
			init := OrderBook{List: txlist, Pay: pay, WillGet: get}
			graph.NGraph[pay][get] = &init
		}
	} else if graph.NGraph[pay][get] == nil {
		txlist := make([]*Offer, capacityList)
		init := OrderBook{List: txlist, Pay: pay, WillGet: get}
		graph.NGraph[pay][get] = &init

	}
	graph.Lock.Unlock()
}




// CreateSimpleGraph : function for creating a SimpleGraph
func (graph *Graph) CreateSimpleGraph() SimplerGraph {

	currencies := graph.getCurrenciesList()

	var simpleGraph = map[string]map[string]float64{}

	for _, i := range currencies {
		simpleGraph[i] = map[string]float64{}
		for _, j := range currencies {
			simpleGraph[i][j] = math.MaxFloat64
		}
	}

	for k1, v1 := range graph.Graph {
		for k2, v2 := range v1 {
			// if len(v2.List) > 2 {
			// 	log.Println("check", v2.List[0].Rate >= v2.List[1].Rate)
			// }

			simpleGraph[k1][k2] = -math.Log(v2.List[0].Rate)
		}
	}
	return SimplerGraph{Graph: simpleGraph, Currencies: currencies, Lock: sync.Mutex{}}
}





func (graph *Graph) getCurrenciesList() []string {
	currencies := make([]string, len(graph.Graph))
	i := 0
	for k := range graph.Graph {
		currencies[i] = k
		i++
	}

	return currencies
}

// SortGraphWithTxs : function for creating a new graph with offers sorted by rates
func (graph *Graph) SortGraphWithTxs() {
	sortedGraph := make(map[string]map[string]*TxList)
	for k1, v1 := range graph.Graph {
		sortedGraph[k1] = map[string]*TxList{}
		for k2, v2 := range v1 {
			list := v2.List
			sort.Slice(list, func(i, j int) bool {
				return v2.List[i].Rate > v2.List[j].Rate
			})
			//sortedGraph[k1][k2] = &TxList{List: list}
			copy(graph.Graph[k1][k2].List, list)
		}
	}
}




func (graph *Graph) addNewOffer(pay string, get string, offer *Offer) {
	graph.initGraph(pay, get)
	graph.Lock.Lock()
	graph.Graph[pay][get].List = append(graph.Graph[pay][get].List, offer)
	graph.Lock.Unlock()

}


func (graph *Graph) initGraph(pay string, get string) {
	graph.Lock.Lock()
	if graph.Graph[pay] == nil {
		graph.Graph[pay] = make(map[string]*TxList)
		txlist := make([]*Offer, capacityList)
		init := TxList{List: txlist}
		graph.Graph[pay][get] = &init
	} else if graph.Graph[pay][get] == nil {
		txlist := make([]*Offer, capacityList)
		init := TxList{List: txlist}
		graph.Graph[pay][get] = &init
	}
	graph.Lock.Unlock()
}



/*
func (graph *Graph) DeleteOffers(transactions []string){
	for _, tx := range transactions {
		graph.Lock.Lock()
		offer := graph.ActiveOffers[tx]
		for i, v  := range graph.Graph[offer.Pay][offer.Get].List {
			log.Println(i)
			if v.Hash == offer.Hash {
				v = nil
				log.Println("DELETED OFFER")
			}
		}
		if len(graph.Graph[offer.Pay][offer.Get].List) == 0 {
			graph.Graph[offer.Pay][offer.Get] = nil
		}
		graph.Lock.Unlock()
	}
}*/

