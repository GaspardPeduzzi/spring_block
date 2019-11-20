package main

import (
	"flag"
	display "github.com/gaspardpeduzzi/spring_block/display"
)

func main() {

	//go s.LaunchServer()

	var addr = flag.String("addr", "s1.ripple.com:51233", "http service address")
	var verb = flag.Bool("verb", false, "Display more information")
	flag.Parse()
	display.VERBOSE = *verb
	display.Init()

	c := make(chan int)



	liquidOptimizer := NewOptimizer(*addr, c)
	liquidOptimizer.NConstructTxGraph()






		/*
		for {
			display.DisplayVerbose("waiting for next block...")
			<-c
			allOffers, cycle := liquidOptimizer.Graph.GetProfitableOffers()
			//returns map[int][]Offer, []string

			if allOffers != nil {
				//Should never be displayed in verbose mode :)
				log.Println("Found profitable cycle:", cycle)
				log.Println("====================================================================================")
				for i, offers := range allOffers {
					for _, offer := range offers {
						log.Println(cycle[i], "->", cycle[(i+1)%len(cycle)], offer.Rate, "OfferCreate Hash:", offer.Hash, "Volume:", offer.Volume)
					}
				}
				log.Println("====================================================================================")
				//return
			}
		}*/

}
