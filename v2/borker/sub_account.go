package broker

// SubAccount represent sub account entity
type SubAccount struct {
	SubAccountID string `json:"subaccountId"`
	Email        string `json:"email"`
	Tag          string `json:"tag"`
}

// EnableFuturesToSubAccountResponse is response struct
type EnableFuturesToSubAccountResponse struct {
	SubAccountID  string `json:"subaccountId"`
	EnableFutures bool   `json:"enableFutures"`
	UpdateTime    int64  `json:"updateTime"`
}

// EnableMarginToSubAccountResponse is response struct
type EnableMarginToSubAccountResponse struct {
	SubAccountID string `json:"subaccountId"`
	EnableMargin bool   `json:"enableMargin"`
	UpdateTime   int64  `json:"updateTime"`
}

// CreateApiKeyRequest is a request struct
type CreateApiKeyRequest struct {
	SubAccountID string
	CanTrade     bool
	MarginTrade  bool
	FuturesTrade bool
}

// CreateApiKeyResponse is response struct
type CreateApiKeyResponse struct {
	SubAccountID string `json:"subaccountId"`
	ApiKey       string `json:"apiKey"`
	SecretKey    string `json:"secretKey"`
	CatTrade     bool   `json:"catTrade"`
	MarginTrade  bool   `json:"marginTrade"`
	FuturesTrade bool   `json:"futuresTrade"`
}

// DeleteSubApiKeyRequest is a request struct
type DeleteSubApiKeyRequest struct {
	SubAccountID string
	ApiKey       string
}

// ChangeApiPermissionRequest is a request struct
type ChangeApiPermissionRequest struct {
	SubAccountID     string `json:"subaccountId"`
	SubAccountApiKey string `json:"subAccountApiKey"`
	CatTrade         bool   `json:"catTrade"`
	MarginTrade      bool   `json:"marginTrade"`
	FuturesTrade     bool   `json:"futuresTrade"`
}

// ChangeApiPermissionResponse is a response struct
type ChangeApiPermissionResponse struct {
	SubAccountID string `json:"subaccountId"`
	ApiKey       string `json:"apiKey"`
	CatTrade     bool   `json:"catTrade"`
	MarginTrade  bool   `json:"marginTrade"`
	FuturesTrade bool   `json:"futuresTrade"`
}

type ChangeSubAccountCommissionRequest struct {
	SubAccountID          string
	MakerCommission       float32
	TakerCommission       float32
	MarginMakerCommission float32
	MarginTakerCommission float32
}

// ChangeSubAccountCommissionResponse is response struct
type ChangeSubAccountCommissionResponse struct {
	SubAccountID          string  `json:"subaccountId"`
	MakerCommission       float32 `json:"makerCommission"`
	TakerCommission       float32 `json:"takerCommission"`
	MarginMakerCommission float32 `json:"marginMakerCommission"`
	TakerMakerCommission  float32 `json:"marginTakerCommission"`
}

// SubAccountTransferRequest is a request struct
type SubAccountTransferRequest struct {
	FromID           string
	ToID             string
	ClientTransferID string
	Asset            string
	Amount           float64
}

// SubAccountTransferResponse is a response struct
type SubAccountTransferResponse struct {
	TxnID        string `json:"txnId"`
	ClientTranID string `json:"clientTranId"`
}
