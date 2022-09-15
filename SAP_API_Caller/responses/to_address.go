package responses

type ToAddress struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			BusinessPartner   string `json:"BusinessPartner"`
			AddressID         string `json:"AddressID"`
			ValidityStartDate string `json:"ValidityStartDate"`
			ValidityEndDate   string `json:"ValidityEndDate"`
			Country           string `json:"Country"`
			Region            string `json:"Region"`
			StreetName        string `json:"StreetName"`
			CityName          string `json:"CityName"`
			PostalCode        string `json:"PostalCode"`
			Language          string `json:"Language"`
		} `json:"results"`
	} `json:"d"`
}
