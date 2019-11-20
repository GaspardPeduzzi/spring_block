package data

import "time"

type OfferCreate struct {
	Tx struct {
		Account            string `json:"Account"`
		Fee                string `json:"Fee"`
		Flags              int    `json:"Flags"`
		LastLedgerSequence int    `json:"LastLedgerSequence"`
		OfferSequence      int    `json:"OfferSequence"`
		Sequence           int    `json:"Sequence"`
		SigningPubKey      string `json:"SigningPubKey"`
		TakerGets          struct {
			Currency string `json:"currency"`
			Issuer   string `json:"issuer"`
			Value    string `json:"value"`
		} `json:"TakerGets"`
		TakerPays       string `json:"TakerPays"`
		TransactionType string `json:"TransactionType"`
		TxnSignature    string `json:"TxnSignature"`
	} `json:"tx"`
	Meta struct {
		AffectedNodes []struct {
			CreatedNode struct {
				LedgerEntryType string `json:"LedgerEntryType"`
				LedgerIndex     string `json:"LedgerIndex"`
				NewFields       struct {
					Account       string `json:"Account"`
					BookDirectory string `json:"BookDirectory"`
					Sequence      int    `json:"Sequence"`
					TakerGets     struct {
						Currency string `json:"currency"`
						Issuer   string `json:"issuer"`
						Value    string `json:"value"`
					} `json:"TakerGets"`
					TakerPays string `json:"TakerPays"`
				} `json:"NewFields"`
			} `json:"CreatedNode,omitempty"`
			DeletedNode struct {
				FinalFields struct {
					ExchangeRate      string `json:"ExchangeRate"`
					Flags             int    `json:"Flags"`
					RootIndex         string `json:"RootIndex"`
					TakerGetsCurrency string `json:"TakerGetsCurrency"`
					TakerGetsIssuer   string `json:"TakerGetsIssuer"`
					TakerPaysCurrency string `json:"TakerPaysCurrency"`
					TakerPaysIssuer   string `json:"TakerPaysIssuer"`
				} `json:"FinalFields"`
				LedgerEntryType string `json:"LedgerEntryType"`
				LedgerIndex     string `json:"LedgerIndex"`
			} `json:"DeletedNode,omitempty"`
			ModifiedNode struct {
				FinalFields struct {
					Account    string `json:"Account"`
					Balance    string `json:"Balance"`
					Flags      int    `json:"Flags"`
					OwnerCount int    `json:"OwnerCount"`
					Sequence   int    `json:"Sequence"`
				} `json:"FinalFields"`
				LedgerEntryType string `json:"LedgerEntryType"`
				LedgerIndex     string `json:"LedgerIndex"`
				PreviousFields  struct {
					Balance  string `json:"Balance"`
					Sequence int    `json:"Sequence"`
				} `json:"PreviousFields"`
				PreviousTxnID     string `json:"PreviousTxnID"`
				PreviousTxnLgrSeq int    `json:"PreviousTxnLgrSeq"`
			} `json:"ModifiedNode,omitempty"`
		} `json:"AffectedNodes"`
		TransactionIndex  int    `json:"TransactionIndex"`
		TransactionResult string `json:"TransactionResult"`
	} `json:"meta"`
	Hash        string    `json:"hash"`
	LedgerIndex int       `json:"ledger_index"`
	Date        time.Time `json:"date"`
}



type PayTransaction struct {
	Tx struct {
		Account string `json:"Account"`
		Amount  struct {
			Currency string `json:"currency"`
			Issuer   string `json:"issuer"`
			Value    string `json:"value"`
		} `json:"Amount"`
		Destination        string `json:"Destination"`
		Fee                string `json:"Fee"`
		Flags              int64  `json:"Flags"`
		LastLedgerSequence int    `json:"LastLedgerSequence"`
		SendMax            struct {
			Currency string `json:"currency"`
			Issuer   string `json:"issuer"`
			Value    string `json:"value"`
		} `json:"SendMax"`
		Sequence        int    `json:"Sequence"`
		SigningPubKey   string `json:"SigningPubKey"`
		TransactionType string `json:"TransactionType"`
		TxnSignature    string `json:"TxnSignature"`
	} `json:"tx"`
	Meta struct {
		AffectedNodes []struct {
			ModifiedNode ModifiedNode `json:"ModifiedNode"`
		} `json:"AffectedNodes"`
		TransactionIndex  int    `json:"TransactionIndex"`
		TransactionResult string `json:"TransactionResult"`
		DeliveredAmount   struct {
			Currency string `json:"currency"`
			Issuer   string `json:"issuer"`
			Value    string `json:"value"`
		} `json:"delivered_amount"`
	} `json:"meta"`
	Hash        string    `json:"hash"`
	LedgerIndex int       `json:"ledger_index"`
	Date        time.Time `json:"date"`
}


type Transaction struct {
	Account            string `json:"Account"`
	Fee                string `json:"Fee"`
	Flags              int    `json:"Flags"`
	LastLedgerSequence int    `json:"LastLedgerSequence,omitempty"`
	OfferSequence      int    `json:"OfferSequence,omitempty"`
	Sequence           int    `json:"Sequence"`
	SigningPubKey      string `json:"SigningPubKey"`
	TakerGets         interface{}
	TakerPays       interface{} `json:"TakerPays,omitempty"`
	TransactionType string `json:"TransactionType"`
	TxnSignature    string `json:"TxnSignature"`
	Hash            string `json:"hash"`
	MetaData        struct {
		AffectedNodes  []struct {
			CreatedNode	CreatedNode`json:"CreatedNode,omitempty"`
			ModifiedNode ModifiedNode`json:"ModifiedNode,omitempty"`
			DeletedNode DeletedNode`json:"DeletedNode,omitempty"`
		} `json:"AffectedNodes"`
		TransactionIndex  int    `json:"TransactionIndex"`
		TransactionResult string `json:"TransactionResult"`
	} `json:"metaData"`
	Amount interface{} `json:"Amount,omitempty"`
	Destination string `json:"Destination,omitempty"`
	Paths       [][]struct {
		Currency string `json:"currency"`
		Issuer   string `json:"issuer,omitempty"`
		Type     int    `json:"type"`
		TypeHex  string `json:"type_hex"`
	} `json:"Paths,omitempty"`
	SendMax    interface{} `json:"SendMax,omitempty"`
	Expiration int    `json:"Expiration,omitempty"`
	Memos      []struct {
		Memo struct {
			MemoData   string `json:"MemoData"`
			MemoFormat string `json:"MemoFormat"`
			MemoType   string `json:"MemoType"`
		} `json:"Memo"`
	} `json:"Memos,omitempty"`
	DestinationTag int `json:"DestinationTag,omitempty"`
}
