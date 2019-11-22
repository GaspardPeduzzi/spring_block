package graph

import (
	"log"

	"github.com/gaspardpeduzzi/spring_block/data"
	"github.com/gaspardpeduzzi/spring_block/display"
)

func (graph *Graph) PaymentTransactionParse(tx data.Transaction) (newOffers []Offer) {

	if len(tx.Paths) > 0 && tx.MetaData.TransactionResult != "tesSUCCESS" {
		display.DisplayAnalysis("====================================================================================")
		display.DisplayAnalysis("Parsing PAYMENT tx", tx.Hash, "transaction status", tx.MetaData.TransactionResult)
		display.DisplayAnalysis("FOUND PATHS")
		display.DisplayAnalysis("PATHS in transactions", len(tx.Paths))
		for k, v := range tx.Paths {
			for i, j := range v {
				display.DisplayAnalysis(k, i, v, j)
			}
		}

		display.DisplayAnalysis(tx.TakerGets)
		display.DisplayAnalysis(tx.TakerPays)
		display.DisplayAnalysis("====================================================================================")
	}

	resultingOffers := make([]Offer, 1)

	for _, v := range tx.MetaData.AffectedNodes {
		created := v.CreatedNode.LedgerEntryType
		modified := v.ModifiedNode.LedgerEntryType
		deleted := v.DeletedNode.LedgerEntryType
		c := created != "" && modified == "" && deleted == ""
		m := created == "" && modified != "" && deleted == ""
		d := created == "" && modified == "" && deleted != ""

		if c {

			if v.CreatedNode.LedgerEntryType == "Offer" {
				//display.DisplayVerbose("CREATED NODE", tx.Hash)
				//display.DisplayVerbose("CREATED new offer from", v.CreatedNode.NewFields.Account, "with seq #", v.CreatedNode.NewFields.Sequence)

				test := v.CreatedNode.NewFields.TakerGets
				test1 := v.CreatedNode.NewFields.TakerPays

				//Offering
				currency, amount, issuer := CurrencyAmountAndIssuer(test)
				currency1, amount1, issuer1 := CurrencyAmountAndIssuer(test1)

				//display.DisplayVerbose("I'm paying", amount, currency, "To receive", amount1, currency1)
				var actualIssuer string
				if issuer == "" {
					actualIssuer = issuer1
				} else {
					actualIssuer = issuer
				}
				rate := amount / amount1
				//display.DisplayVerbose("RATE", rate)

				newOffer := &Offer{
					XrpTx:          tx,
					TxHash:         tx.Hash,
					Account:        v.CreatedNode.NewFields.Account,
					SequenceNumber: v.CreatedNode.NewFields.Sequence,
					Rate:           rate,
					Quantity:       amount,
					CreatorWillPay: currency,
					CreatorWillGet: currency1,
					Issuer:         actualIssuer,
				}
				if !graph.Issuers[newOffer.Issuer] {
					graph.Issuers[newOffer.Issuer] = true
					display.DisplayAnalysis("NEW ISSUER: ", newOffer.Issuer, "TRACK ISSUERS: ", len(graph.Issuers))
				}

				graph.insertNewOffer(newOffer)
			}

		} else if m {

			if v.DeletedNode.LedgerEntryType == "Offer" {
				display.DisplayVerbose("MODIFIED NODE OFFER", v.ModifiedNode.LedgerEntryType)

			}
		} else if d {
			//display.DisplayVerbose("DELETED NODE", v.DeletedNode.LedgerEntryType)
			if v.DeletedNode.LedgerEntryType == "Offer" && tx.MetaData.TransactionResult == "tesSUCCESS" {

				//graph BTC ETH donne offerCreate pour obtenir ETH en payant BTC
				//[TG][TP]
				account := v.DeletedNode.FinalFields.Account
				seq := v.DeletedNode.FinalFields.Sequence
				tp := v.DeletedNode.FinalFields.TakerPays
				tg := v.DeletedNode.FinalFields.TakerGets

				currencyTP, _, _ := CurrencyAmountAndIssuer(tp)
				currencyTG, _, issuerTG := CurrencyAmountAndIssuer(tg)

				//display.DisplayVerbose("DELETED", account, seq, "ORDERBOOK", currencyTG, "/", currencyTP)

				offer := &Offer{
					XrpTx:          tx,
					TxHash:         tx.Hash,
					Account:        account,
					SequenceNumber: seq,
					Rate:           0,
					Quantity:       0,
					CreatorWillPay: currencyTG,
					CreatorWillGet: currencyTP,
					Issuer:         issuerTG,
				}

				//deletedOffers = append(deletedOffers, offer)
				if graph.Graph[offer.CreatorWillPay][offer.CreatorWillGet] != nil {
					for k, v := range graph.Graph[offer.CreatorWillPay][offer.CreatorWillGet].List {
						if v.Account == account && v.SequenceNumber == seq {
							removeOffer(graph.Graph[offer.CreatorWillPay][offer.CreatorWillGet].List, k)
							log.Println("Order book adjusted")

						}
					}
				}

				//display.DisplayVerbose("Taker gets", currencyTG, amountTG, issuerTG)
				//display.DisplayVerbose("Taker pays", currencyTP, amountTP, issuerTP)

			}

		}
	}
	return resultingOffers

}
