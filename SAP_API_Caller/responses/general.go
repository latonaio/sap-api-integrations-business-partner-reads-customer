package responses

type General struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			BusinessPartner                string      `json:"BusinessPartner"`
			Customer                       string      `json:"Customer"`
			Supplier                       string      `json:"Supplier"`
			AcademicTitle                  string      `json:"AcademicTitle"`
			AuthorizationGroup             string      `json:"AuthorizationGroup"`
			BusinessPartnerCategory        string      `json:"BusinessPartnerCategory"`
			BusinessPartnerFullName        string      `json:"BusinessPartnerFullName"`
			BusinessPartnerGrouping        string      `json:"BusinessPartnerGrouping"`
			BusinessPartnerName            string      `json:"BusinessPartnerName"`
			CorrespondenceLanguage         string      `json:"CorrespondenceLanguage"`
			CreationDate                   string      `json:"CreationDate"`
			CreationTime                   string      `json:"CreationTime"`
			FirstName                      string      `json:"FirstName"`
			Industry                       string      `json:"Industry"`
			IsFemale                       bool        `json:"IsFemale"`
			IsMale                         bool        `json:"IsMale"`
			IsNaturalPerson                string      `json:"IsNaturalPerson"`
			IsSexUnknown                   bool        `json:"IsSexUnknown"`
			GenderCodeName                 string      `json:"GenderCodeName"`
			Language                       string      `json:"Language"`
			LastChangeDate                 string      `json:"LastChangeDate"`
			LastChangeTime                 string      `json:"LastChangeTime"`
			LastName                       string      `json:"LastName"`
			OrganizationBPName1            string      `json:"OrganizationBPName1"`
			OrganizationBPName2            string      `json:"OrganizationBPName2"`
			OrganizationBPName3            string      `json:"OrganizationBPName3"`
			OrganizationBPName4            string      `json:"OrganizationBPName4"`
			OrganizationFoundationDate     string      `json:"OrganizationFoundationDate"`
			OrganizationLiquidationDate    string      `json:"OrganizationLiquidationDate"`
			SearchTerm1                    string      `json:"SearchTerm1"`
			SearchTerm2                    string      `json:"SearchTerm2"`
			AdditionalLastName             string      `json:"AdditionalLastName"`
			BirthDate                      string      `json:"BirthDate"`
			BusinessPartnerBirthplaceName  string      `json:"BusinessPartnerBirthplaceName"`
			BusinessPartnerDeathDate       string      `json:"BusinessPartnerDeathDate"`
			BusinessPartnerIsBlocked       bool        `json:"BusinessPartnerIsBlocked"`
			BusinessPartnerType            string      `json:"BusinessPartnerType"`
			GroupBusinessPartnerName1      string      `json:"GroupBusinessPartnerName1"`
			GroupBusinessPartnerName2      string      `json:"GroupBusinessPartnerName2"`
			IndependentAddressID           string      `json:"IndependentAddressID"`
			MiddleName                     string      `json:"MiddleName"`
			NameCountry                    string      `json:"NameCountry"`
			PersonFullName                 string      `json:"PersonFullName"`
			PersonNumber                   string      `json:"PersonNumber"`
			IsMarkedForArchiving           bool        `json:"IsMarkedForArchiving"`
			BusinessPartnerIDByExtSystem   string      `json:"BusinessPartnerIDByExtSystem"`
			TradingPartner                 string      `json:"TradingPartner"`
			ToRole struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_BusinessPartnerRole"`
			ToAddress struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_BusinessPartnerAddress"`
			ToBank struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_BusinessPartnerBank"`
			ToCustomer struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_Customer"`
		} `json:"results"`
	} `json:"d"`
}
