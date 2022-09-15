package responses

type ToPartnerFunction struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			Customer                   string `json:"Customer"`
			SalesOrganization          string `json:"SalesOrganization"`
			DistributionChannel        string `json:"DistributionChannel"`
			Division                   string `json:"Division"`
			PartnerCounter             string `json:"PartnerCounter"`
			PartnerFunction            string `json:"PartnerFunction"`
			BPCustomerNumber           string `json:"BPCustomerNumber"`
			CustomerPartnerDescription string `json:"CustomerPartnerDescription"`
			DefaultPartner             bool   `json:"DefaultPartner"`
			Supplier                   string `json:"Supplier"`
			AuthorizationGroup         string `json:"AuthorizationGroup"`
		} `json:"results"`
	} `json:"d"`
}
