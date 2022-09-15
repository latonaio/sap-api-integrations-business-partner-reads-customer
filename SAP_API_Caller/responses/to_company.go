package responses

type ToCompany struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			Customer                       string `json:"Customer"`
			CompanyCode                    string `json:"CompanyCode"`
			APARToleranceGroup             string `json:"APARToleranceGroup"`
			CustomerSupplierClearingIsUsed bool   `json:"CustomerSupplierClearingIsUsed"`
			HouseBank                      string `json:"HouseBank"`
			PaymentMethodsList             string `json:"PaymentMethodsList"`
			PaymentTerms                   string `json:"PaymentTerms"`
			ReconciliationAccount          string `json:"ReconciliationAccount"`
			DeletionIndicator              bool   `json:"DeletionIndicator"`
		} `json:"results"`
	} `json:"d"`
}
