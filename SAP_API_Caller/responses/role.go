package responses

type Role struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			BusinessPartner     string `json:"BusinessPartner"`
			BusinessPartnerRole string `json:"BusinessPartnerRole"`
			ValidFrom           string `json:"ValidFrom"`
			ValidTo             string `json:"ValidTo"`
		} `json:"results"`
	} `json:"d"`
}
