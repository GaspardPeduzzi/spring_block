package data

type LedgerResponseExpanded struct {
	ID     int `json:"id"`
	Result struct {
		Ledger struct {
			Accepted            bool   `json:"accepted"`
			AccountHash         string `json:"account_hash"`
			CloseFlags          int    `json:"close_flags"`
			CloseTime           int    `json:"close_time"`
			CloseTimeHuman      string `json:"close_time_human"`
			CloseTimeResolution int    `json:"close_time_resolution"`
			Closed              bool   `json:"closed"`
			Hash                string `json:"hash"`
			LedgerHash          string `json:"ledger_hash"`
			LedgerIndex         string `json:"ledger_index"`
			ParentCloseTime     int    `json:"parent_close_time"`
			ParentHash          string `json:"parent_hash"`
			SeqNum              string `json:"seqNum"`
			TotalCoins          string `json:"totalCoins"`
			Total_Coins          string `json:"total_coins"`
			TransactionHash     string `json:"transaction_hash"`
			Transactions        []Transaction  `json:"transactions"`
		} `json:"ledger"`
		LedgerHash  string `json:"ledger_hash"`
		LedgerIndex int    `json:"ledger_index"`
		Validated   bool   `json:"validated"`
	} `json:"result"`
	Status string `json:"status"`
	Type   string `json:"type"`
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
		CreatedNode struct {
		LedgerEntryType string `json:"LedgerEntryType"`
		LedgerIndex     string `json:"LedgerIndex"`
		NewFields       struct {
		Account       string `json:"Account"`
		BookDirectory string `json:"BookDirectory"`
		Sequence      int    `json:"Sequence"`
		TakerGets     interface{}
		TakerPays interface{} `json:"TakerPays"`
	} `json:"NewFields"`
	} `json:"CreatedNode,omitempty"`
		ModifiedNode struct {
		FinalFields struct {
		Flags     int    `json:"Flags"`
		Owner     string `json:"Owner"`
		RootIndex string `json:"RootIndex"`
	} `json:"FinalFields"`
		LedgerEntryType string `json:"LedgerEntryType"`
		LedgerIndex     string `json:"LedgerIndex"`
	} `json:"ModifiedNode,omitempty"`
			DeletedNode `json:"DeletedNode,omitempty"`
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



type DeletedNode struct {
	FinalFields struct {
		Account           string `json:"Account"`
		BookDirectory     string `json:"BookDirectory"`
		BookNode          string `json:"BookNode"`
		Expiration        int    `json:"Expiration"`
		Flags             int    `json:"Flags"`
		OwnerNode         string `json:"OwnerNode"`
		PreviousTxnID     string `json:"PreviousTxnID"`
		PreviousTxnLgrSeq int    `json:"PreviousTxnLgrSeq"`
		Sequence          int    `json:"Sequence"`
		TakerGets         interface{} `json:"TakerGets"`
		TakerPays         interface{} `json:"TakerPays"`
	} `json:"FinalFields"`
	LedgerEntryType string `json:"LedgerEntryType"`
	LedgerIndex     string `json:"LedgerIndex"`
}

type LedgerSeqRequest struct {
	Id      int    `json:"id"`
	Command string `json:"command"`
}


type LedgerRequest struct {
	Id           int    `json:"id"`
	Command      string `json:"command"`
	LedgerIndex  string `json:"ledger_index"`
	Full         bool   `json:"full"`
	Accounts     bool   `json:"accounts"`
	Transactions bool   `json:"transactions"`
	Expand       bool   `json:"expand"`
	OwnerFunds   bool   `json:"owner_funds"`
}

type ResponseCurrent struct {
	ID     int `json:"id"`
	Result struct {
		LedgerCurrentIndex int `json:"ledger_current_index"`
	} `json:"result"`
	Status string `json:"status"`
	Type   string `json:"type"`
}



// Payload object
type message struct {
	Command string `json:"command"`
	//Taker     string    `json:"taker"`
	TakerGets takerGets `json:"taker_gets"`
	TakerPays takerPays `json:"taker_pays"`
	Limit     uint      `json:"limit"`
}

type takerGets struct {
	Currency string `json:"currency"`
}

type takerPays struct {
	Currency string `json:"currency"`
	Issuer   string `json:"issuer"`
}
