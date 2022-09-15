package responses

type SalesArea struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			Customer                       string `json:"Customer"`
			SalesOrganization              string `json:"SalesOrganization"`
			DistributionChannel            string `json:"DistributionChannel"`
			Division                       string `json:"Division"`
			CompleteDeliveryIsDefined      bool   `json:"CompleteDeliveryIsDefined"`
			Currency                       string `json:"Currency"`
			CustomerAccountAssignmentGroup string `json:"CustomerAccountAssignmentGroup"`
			CustomerPaymentTerms           string `json:"CustomerPaymentTerms"`
			CustomerPriceGroup             string `json:"CustomerPriceGroup"`
			CustomerPricingProcedure       string `json:"CustomerPricingProcedure"`
			DeliveryPriority               string `json:"DeliveryPriority"`
			IncotermsClassification        string `json:"IncotermsClassification"`
			InvoiceDate                    string `json:"InvoiceDate"`
			OrderCombinationIsAllowed      bool   `json:"OrderCombinationIsAllowed"`
			PartialDeliveryIsAllowed       string `json:"PartialDeliveryIsAllowed"`
			PriceListType                  string `json:"PriceListType"`
			SalesGroup                     string `json:"SalesGroup"`
			SalesOffice                    string `json:"SalesOffice"`
			ShippingCondition              string `json:"ShippingCondition"`
			SupplyingPlant                 string `json:"SupplyingPlant"`
			SalesDistrict                  string `json:"SalesDistrict"`
			InvoiceListSchedule            string `json:"InvoiceListSchedule"`
			ExchangeRateType               string `json:"ExchangeRateType"`
			OrderIsBlockedForCustomer      string `json:"OrderIsBlockedForCustomer"`
			DeliveryIsBlockedForCustomer   string `json:"DeliveryIsBlockedForCustomer"`
			BillingIsBlockedForCustomer    string `json:"BillingIsBlockedForCustomer"`
			DeletionIndicator              bool   `json:"DeletionIndicator"`
			ToPartnerFunction              struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_PartnerFunction"`
		} `json:"results"`
	} `json:"d"`
}
