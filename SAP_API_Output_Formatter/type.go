package sap_api_output_formatter

type BusinessPartner struct {
	ConnectionKey   string `json:"connection_key"`
	Result          bool   `json:"result"`
	RedisKey        string `json:"redis_key"`
	Filepath        string `json:"filepath"`
	Product         string `json:"Product"`
	APISchema       string `json:"api_schema"`
	BusinessPartner string `json:"business_partner_code"`
	Deleted         bool   `json:"deleted"`
}

type General struct {
	BusinessPartner               string `json:"BusinessPartner"`
	Customer                      string `json:"Customer"`
	Supplier                      string `json:"Supplier"`
	AcademicTitle                 string `json:"AcademicTitle"`
	AuthorizationGroup            string `json:"AuthorizationGroup"`
	BusinessPartnerCategory       string `json:"BusinessPartnerCategory"`
	BusinessPartnerFullName       string `json:"BusinessPartnerFullName"`
	BusinessPartnerGrouping       string `json:"BusinessPartnerGrouping"`
	BusinessPartnerName           string `json:"BusinessPartnerName"`
	CorrespondenceLanguage        string `json:"CorrespondenceLanguage"`
	CreationDate                  string `json:"CreationDate"`
	CreationTime                  string `json:"CreationTime"`
	FirstName                     string `json:"FirstName"`
	Industry                      string `json:"Industry"`
	IsFemale                      bool   `json:"IsFemale"`
	IsMale                        bool   `json:"IsMale"`
	IsNaturalPerson               string `json:"IsNaturalPerson"`
	IsSexUnknown                  bool   `json:"IsSexUnknown"`
	GenderCodeName                string `json:"GenderCodeName"`
	Language                      string `json:"Language"`
	LastChangeDate                string `json:"LastChangeDate"`
	LastChangeTime                string `json:"LastChangeTime"`
	LastName                      string `json:"LastName"`
	OrganizationBPName1           string `json:"OrganizationBPName1"`
	OrganizationBPName2           string `json:"OrganizationBPName2"`
	OrganizationBPName3           string `json:"OrganizationBPName3"`
	OrganizationBPName4           string `json:"OrganizationBPName4"`
	OrganizationFoundationDate    string `json:"OrganizationFoundationDate"`
	OrganizationLiquidationDate   string `json:"OrganizationLiquidationDate"`
	SearchTerm1                   string `json:"SearchTerm1"`
	SearchTerm2                   string `json:"SearchTerm2"`
	AdditionalLastName            string `json:"AdditionalLastName"`
	BirthDate                     string `json:"BirthDate"`
	BusinessPartnerBirthplaceName string `json:"BusinessPartnerBirthplaceName"`
	BusinessPartnerDeathDate      string `json:"BusinessPartnerDeathDate"`
	BusinessPartnerIsBlocked      bool   `json:"BusinessPartnerIsBlocked"`
	BusinessPartnerType           string `json:"BusinessPartnerType"`
	GroupBusinessPartnerName1     string `json:"GroupBusinessPartnerName1"`
	GroupBusinessPartnerName2     string `json:"GroupBusinessPartnerName2"`
	IndependentAddressID          string `json:"IndependentAddressID"`
	MiddleName                    string `json:"MiddleName"`
	NameCountry                   string `json:"NameCountry"`
	PersonFullName                string `json:"PersonFullName"`
	PersonNumber                  string `json:"PersonNumber"`
	IsMarkedForArchiving          bool   `json:"IsMarkedForArchiving"`
	BusinessPartnerIDByExtSystem  string `json:"BusinessPartnerIDByExtSystem"`
	TradingPartner                string `json:"TradingPartner"`
	ToRole                        string `json:"to_BusinessPartnerRole"`
	ToAddress                     string `json:"to_BusinessPartnerAddress"`
	ToBank                        string `json:"to_BusinessPartnerBank"`
	ToCustomer                    string `json:"to_Customer"`
}

type Role struct {
	BusinessPartner     string `json:"BusinessPartner"`
	BusinessPartnerRole string `json:"BusinessPartnerRole"`
	ValidFrom           string `json:"ValidFrom"`
	ValidTo             string `json:"ValidTo"`
}

type Address struct {
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
}

type Bank struct {
	BusinessPartner          string `json:"BusinessPartner"`
	BankIdentification       string `json:"BankIdentification"`
	BankCountryKey           string `json:"BankCountryKey"`
	BankName                 string `json:"BankName"`
	BankNumber               string `json:"BankNumber"`
	SWIFTCode                string `json:"SWIFTCode"`
	BankControlKey           string `json:"BankControlKey"`
	BankAccountHolderName    string `json:"BankAccountHolderName"`
	BankAccountName          string `json:"BankAccountName"`
	ValidityStartDate        string `json:"ValidityStartDate"`
	ValidityEndDate          string `json:"ValidityEndDate"`
	IBAN                     string `json:"IBAN"`
	IBANValidityStartDate    string `json:"IBANValidityStartDate"`
	BankAccount              string `json:"BankAccount"`
	BankAccountReferenceText string `json:"BankAccountReferenceText"`
	CollectionAuthInd        bool   `json:"CollectionAuthInd"`
	CityName                 string `json:"CityName"`
	AuthorizationGroup       string `json:"AuthorizationGroup"`
}

type Customer struct {
	Customer                    string `json:"Customer"`
	AuthorizationGroup          string `json:"AuthorizationGroup"`
	BillingIsBlockedForCustomer string `json:"BillingIsBlockedForCustomer"`
	CreationDate                string `json:"CreationDate"`
	CustomerAccountGroup        string `json:"CustomerAccountGroup"`
	CustomerClassification      string `json:"CustomerClassification"`
	CustomerFullName            string `json:"CustomerFullName"`
	CustomerName                string `json:"CustomerName"`
	DeliveryIsBlocked           string `json:"DeliveryIsBlocked"`
	OrderIsBlockedForCustomer   string `json:"OrderIsBlockedForCustomer"`
	PostingIsBlocked            bool   `json:"PostingIsBlocked"`
	Supplier                    string `json:"Supplier"`
	CustomerCorporateGroup      string `json:"CustomerCorporateGroup"`
	Industry                    string `json:"Industry"`
	TaxNumber1                  string `json:"TaxNumber1"`
	DeletionIndicator           bool   `json:"DeletionIndicator"`
	CityCode                    string `json:"CityCode"`
	County                      string `json:"County"`
	ToSalesArea                 string `json:"to_CustomerSalesArea"`
	ToCompany                   string `json:"to_CustomerCompany"`
}

type SalesArea struct {
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
	ToPartnerFunction              string `json:"to_PartnerFunction"`
}

type Company struct {
	Customer                       string `json:"Customer"`
	CompanyCode                    string `json:"CompanyCode"`
	APARToleranceGroup             string `json:"APARToleranceGroup"`
	CustomerSupplierClearingIsUsed bool   `json:"CustomerSupplierClearingIsUsed"`
	HouseBank                      string `json:"HouseBank"`
	PaymentMethodsList             string `json:"PaymentMethodsList"`
	PaymentTerms                   string `json:"PaymentTerms"`
	ReconciliationAccount          string `json:"ReconciliationAccount"`
	DeletionIndicator              bool   `json:"DeletionIndicator"`
}

type ToRole struct {
	BusinessPartner     string `json:"BusinessPartner"`
	BusinessPartnerRole string `json:"BusinessPartnerRole"`
	ValidFrom           string `json:"ValidFrom"`
	ValidTo             string `json:"ValidTo"`
}

type ToAddress struct {
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
}

type ToBank struct {
	BusinessPartner          string `json:"BusinessPartner"`
	BankIdentification       string `json:"BankIdentification"`
	BankCountryKey           string `json:"BankCountryKey"`
	BankName                 string `json:"BankName"`
	BankNumber               string `json:"BankNumber"`
	SWIFTCode                string `json:"SWIFTCode"`
	BankControlKey           string `json:"BankControlKey"`
	BankAccountHolderName    string `json:"BankAccountHolderName"`
	BankAccountName          string `json:"BankAccountName"`
	ValidityStartDate        string `json:"ValidityStartDate"`
	ValidityEndDate          string `json:"ValidityEndDate"`
	IBAN                     string `json:"IBAN"`
	IBANValidityStartDate    string `json:"IBANValidityStartDate"`
	BankAccount              string `json:"BankAccount"`
	BankAccountReferenceText string `json:"BankAccountReferenceText"`
	CollectionAuthInd        bool   `json:"CollectionAuthInd"`
	CityName                 string `json:"CityName"`
	AuthorizationGroup       string `json:"AuthorizationGroup"`
}

type ToCustomer struct {
	Customer                    string `json:"Customer"`
	AuthorizationGroup          string `json:"AuthorizationGroup"`
	BillingIsBlockedForCustomer string `json:"BillingIsBlockedForCustomer"`
	CreationDate                string `json:"CreationDate"`
	CustomerAccountGroup        string `json:"CustomerAccountGroup"`
	CustomerClassification      string `json:"CustomerClassification"`
	CustomerFullName            string `json:"CustomerFullName"`
	CustomerName                string `json:"CustomerName"`
	DeliveryIsBlocked           string `json:"DeliveryIsBlocked"`
	OrderIsBlockedForCustomer   string `json:"OrderIsBlockedForCustomer"`
	PostingIsBlocked            bool   `json:"PostingIsBlocked"`
	Supplier                    string `json:"Supplier"`
	CustomerCorporateGroup      string `json:"CustomerCorporateGroup"`
	Industry                    string `json:"Industry"`
	TaxNumber1                  string `json:"TaxNumber1"`
	DeletionIndicator           bool   `json:"DeletionIndicator"`
	CityCode                    string `json:"CityCode"`
	County                      string `json:"County"`
	ToSalesArea                 string `json:"to_CustomerSalesArea"`
	ToCompany                   string `json:"to_CustomerCompany"`
}

type ToSalesArea struct {
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
	ToPartnerFunction              string `json:"to_PartnerFunction"`
}

type ToPartnerFunction struct {
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
}

type ToCompany struct {
	Customer                       string `json:"Customer"`
	CompanyCode                    string `json:"CompanyCode"`
	APARToleranceGroup             string `json:"APARToleranceGroup"`
	CustomerSupplierClearingIsUsed bool   `json:"CustomerSupplierClearingIsUsed"`
	HouseBank                      string `json:"HouseBank"`
	PaymentMethodsList             string `json:"PaymentMethodsList"`
	PaymentTerms                   string `json:"PaymentTerms"`
	ReconciliationAccount          string `json:"ReconciliationAccount"`
	DeletionIndicator              bool   `json:"DeletionIndicator"`
}
