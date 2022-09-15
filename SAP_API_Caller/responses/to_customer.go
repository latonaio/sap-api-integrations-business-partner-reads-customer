package responses

type ToCustomer struct {
	D struct {
		Metadata struct {
			ID   string `json:"id"`
			URI  string `json:"uri"`
			Type string `json:"type"`
		} `json:"__metadata"`
		Customer                     string `json:"Customer"`
		AuthorizationGroup           string `json:"AuthorizationGroup"`
		BillingIsBlockedForCustomer  string `json:"BillingIsBlockedForCustomer"`
		CreationDate                 string `json:"CreationDate"`
		CustomerAccountGroup         string `json:"CustomerAccountGroup"`
		CustomerClassification       string `json:"CustomerClassification"`
		CustomerFullName             string `json:"CustomerFullName"`
		CustomerName                 string `json:"CustomerName"`
		DeliveryIsBlocked            string `json:"DeliveryIsBlocked"`
		OrderIsBlockedForCustomer    string `json:"OrderIsBlockedForCustomer"`
		PostingIsBlocked             bool   `json:"PostingIsBlocked"`
		Supplier                     string `json:"Supplier"`
		CustomerCorporateGroup       string `json:"CustomerCorporateGroup"`
		Industry                     string `json:"Industry"`
		TaxNumber1                   string `json:"TaxNumber1"`
		DeletionIndicator            bool   `json:"DeletionIndicator"`
		CityCode                     string `json:"CityCode"`
		County                       string `json:"County"`
		ToSalesArea         struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"to_CustomerSalesArea"`
		ToCompany           struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"to_CustomerCompany"`
	} `json:"d"`
}
