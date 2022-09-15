package main

import (
	sap_api_caller "sap-api-integrations-business-partner-reads-customer/SAP_API_Caller"
	sap_api_input_reader "sap-api-integrations-business-partner-reads-customer/SAP_API_Input_Reader"
	"sap-api-integrations-business-partner-reads-customer/config"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	sap_api_request_client_header_setup "github.com/latonaio/sap-api-request-client-header-setup"
	sap_api_time_value_converter "github.com/latonaio/sap-api-time-value-converter"
)

func main() {
	l := logger.NewLogger()
	conf := config.NewConf()
	fr := sap_api_input_reader.NewFileReader()
	gc := sap_api_request_client_header_setup.NewSAPRequestClientWithOption(conf.SAP)
	caller := sap_api_caller.NewSAPAPICaller(
		conf.SAP.BaseURL(),
		"100",
		gc,
		l,
	)
	inputSDC := fr.ReadSDC("./Inputs//SDC_Business_Partner_Customer_Sales_Area_sample.json")
	sap_api_time_value_converter.ChangeTimeFormatToSAPFormatStruct(&inputSDC)
	accepter := inputSDC.Accepter
	if len(accepter) == 0 || accepter[0] == "All" {
		accepter = []string{
			"General", "Role", "Address", "Bank", "BPName", "Customer", "SalesArea", "Company",
		}
	}

	caller.AsyncGetBPCustomer(
		inputSDC.BusinessPartner.BusinessPartner,
		inputSDC.BusinessPartner.Role.BusinessPartnerRole,
		inputSDC.BusinessPartner.Address.AddressID,
		inputSDC.BusinessPartner.Bank.BankCountryKey,
		inputSDC.BusinessPartner.Bank.BankNumber,
		inputSDC.BusinessPartner.BusinessPartnerName,
		inputSDC.BusinessPartner.CustomerData.Customer,
		inputSDC.BusinessPartner.CustomerData.SalesArea.SalesOrganization,
		inputSDC.BusinessPartner.CustomerData.SalesArea.DistributionChannel,
		inputSDC.BusinessPartner.CustomerData.SalesArea.Division,
		inputSDC.BusinessPartner.CustomerData.Company.CompanyCode,
		accepter,
	)
}
