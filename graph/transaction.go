package graph

import (
	"github.com/gaspardpeduzzi/spring_block/display"
	"log"
	"strconv"

	"github.com/gaspardpeduzzi/spring_block/data"
)




func (graph *Graph) parseOfferCancel (tx data.Transaction){
	for _, test := range tx.MetaData.AffectedNodes{
		log.Println(test.DeletedNode.FinalFields.PreviousTxnID)
	}
}

func (graph *Graph) ParseOfferCreate (tx data.Transaction) []Offer {
	//display.DisplayVerbose("====================================================================================")
	display.DisplayVerbose("Parsing tx", tx.Hash)
	resultingOffers := make([]Offer, 1)
	//rate :=
	for _ , v := range  tx.MetaData.AffectedNodes {

		created := v.CreatedNode.LedgerEntryType
		modified := v.ModifiedNode.LedgerEntryType
		deleted := v.DeletedNode.LedgerEntryType
		c := created != "" && modified == "" && deleted == ""
		m := created == "" && modified != "" && deleted == ""
		d := created == "" && modified == "" && deleted != ""

		if c {
			//display.DisplayVerbose("CREATED NODE")
			if v.CreatedNode.LedgerEntryType == "Offer" {

				//display.DisplayVerbose("CREATED new offer from", v.CreatedNode.NewFields.Account, "with seq #", v.CreatedNode.NewFields.Sequence)
				test := v.CreatedNode.NewFields.TakerGets
				test1 := v.CreatedNode.NewFields.TakerPays
				//Offering
				currency, amount, issuer := offerringCurrencyAndAmount(test)
				currency1, amount1, issuer1 := offerringCurrencyAndAmount(test1)
				//display.DisplayVerbose("I'm paying", amount, currency, "To receive", amount1, currency1)

				var actualIssuer string

				if (issuer == "") {
					actualIssuer = issuer1
				} else {
					actualIssuer = issuer
				}
				rate := amount1/amount



				newOffer := &Offer{
					XrpTx:          tx,
					TxHash:         tx.Hash,
					Account:        v.CreatedNode.NewFields.Account,
					SequenceNumber: v.CreatedNode.NewFields.Sequence,
					Rate:           rate,
					Quantity:         amount,
					CreatorWillPay:   currency,
					CreatorWillGet:   currency1,
					Issuer:			  actualIssuer,
				}
				//display.DisplayVerbose("ACTUAL ISSUER", actualIssuer)
				graph.insertNewOfferToAccount(newOffer)
				graph.insertNewOffer(newOffer)
			}

		} else if m {
			//display.DisplayVerbose("MODIFIED NODE", v.ModifiedNode.LedgerEntryType)
		} else if d {
			//display.DisplayVerbose("DELETED NODE", v.DeletedNode.LedgerEntryType)
			if v.DeletedNode.LedgerEntryType == "Offer" {
				// Would be great if we could check that it deletes actually the node here
				//delete previous offer

				/*
				toDelete := &Offer{
					XrpTx:          tx,
					TxHash:         tx.Hash,
					Account:        v.DeletedNode.FinalFields.Account,
					SequenceNumber: v.DeletedNode.FinalFields.Sequence,
					Rate:           0,
					Quantity:       0,
					CreatorWillPay:            "",
					CreatorWillGet:            "",
				}
				*/

				//graph.deleteOffer(toDelete)
				//graph.AccountRoots[v.DeletedNode.FinalFields.Account][v.DeletedNode.FinalFields.Sequence] = nil
				//display.DisplayVerbose("DELETED previous offer from", v.DeletedNode.FinalFields.Account, "with seq #", v.DeletedNode.FinalFields.Sequence)
			}
		}
		//resultingOffers = append(resultingOffers, offer)
	}

	//display.DisplayVerbose("====================================================================================")
	return resultingOffers
}











func offerringCurrencyAndAmount(takerPays interface{}) (string, float64, string) {
	var weWillPayInCurrency string
	var priceToPay string
	var issuer string
	switch object := takerPays.(type) {
		case map[string]interface{}:
			//We need to pay in a given currency
			//log.Print("TAKER PAYS currency ", object["currency"]," value ",object["value"])
			weWillPayInCurrency = object["currency"].(string)
			priceToPay = object["value"].(string)
			issuer = object["issuer"].(string)
		case string:
			//We need to pay with the native currency
			//log.Print("TAKER PAYS value ", DropToXrp(float64(price)), " XRP or ", DropToPriceInUSD(price), " USD" )
			weWillPayInCurrency = "XRP"
			priceToPay = object
		default:
			log.Println("unexpected type %T", object)
	}
	weWillPayAmount, err := strconv.ParseFloat(priceToPay, 64)
	if err != nil {
		log.Println("Error decoding", err)
	}

	return weWillPayInCurrency, weWillPayAmount, issuer
}
