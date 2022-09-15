package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	sap_api_output_formatter "sap-api-integrations-business-partner-reads-customer/SAP_API_Output_Formatter"
	"strings"
	"sync"

	sap_api_request_client_header_setup "github.com/latonaio/sap-api-request-client-header-setup"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
)

type SAPAPICaller struct {
	baseURL         string
	sapClientNumber string
	requestClient   *sap_api_request_client_header_setup.SAPRequestClient
	log             *logger.Logger
}

func NewSAPAPICaller(baseUrl, sapClientNumber string, requestClient *sap_api_request_client_header_setup.SAPRequestClient, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:         baseUrl,
		requestClient:   requestClient,
		sapClientNumber: sapClientNumber,
		log:             l,
	}
}

func (c *SAPAPICaller) AsyncGetBPCustomer(businessPartner, businessPartnerRole, addressID, bankCountryKey, bankNumber, bPName, customer, salesOrganization, distributionChannel, division, companyCode string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "General":
			func() {
				c.General(businessPartner)
				wg.Done()
			}()
		case "Role":
			func() {
				c.Role(businessPartner, businessPartnerRole)
				wg.Done()
			}()
		case "Address":
			func() {
				c.Address(businessPartner, addressID)
				wg.Done()
			}()
		case "Bank":
			func() {
				c.Bank(businessPartner, bankCountryKey, bankNumber)
				wg.Done()
			}()
		case "BPName":
			func() {
				c.BPName(bPName)
				wg.Done()
			}()
		case "Customer":
			func() {
				c.Customer(customer)
				wg.Done()
			}()
		case "SalesArea":
			func() {
				c.SalesArea(customer, salesOrganization, distributionChannel, division)
				wg.Done()
			}()
		case "Company":
			func() {
				c.Company(customer, companyCode)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) General(businessPartner string) {
	generalData, err := c.callBPCustomerSrvAPIRequirementGeneral("A_BusinessPartner", businessPartner)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(generalData)

	roleData, err := c.callToRole(generalData[0].ToRole)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(roleData)

	addressData, err := c.callToAddress(generalData[0].ToAddress)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(addressData)

	bankData, err := c.callToBank(generalData[0].ToBank)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(bankData)

	customerData, err := c.callToCustomer(generalData[0].ToCustomer)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(customerData)

	salesAreaData, err := c.callToSalesArea(customerData.ToSalesArea)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(salesAreaData)

	partnerFunctionData, err := c.callToPartnerFunction(salesAreaData[0].ToPartnerFunction)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(partnerFunctionData)

	companyData, err := c.callToCompany(customerData.ToCompany)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(companyData)
}

func (c *SAPAPICaller) callBPCustomerSrvAPIRequirementGeneral(api, businessPartner string) ([]sap_api_output_formatter.General, error) {
	url := strings.Join([]string{c.baseURL, "API_BUSINESS_PARTNER", api}, "/")
	param := c.getQueryWithGeneral(map[string]string{}, businessPartner)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToGeneral(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToRole(url string) ([]sap_api_output_formatter.ToRole, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToRole(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToAddress(url string) ([]sap_api_output_formatter.ToAddress, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToAddress(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToBank(url string) ([]sap_api_output_formatter.ToBank, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToBank(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToCustomer(url string) (*sap_api_output_formatter.ToCustomer, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToCustomer(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToSalesArea(url string) ([]sap_api_output_formatter.ToSalesArea, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToSalesArea(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToPartnerFunction(url string) ([]sap_api_output_formatter.ToPartnerFunction, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToPartnerFunction(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToCompany(url string) ([]sap_api_output_formatter.ToCompany, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToCompany(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Role(businessPartner, businessPartnerRole string) {
	data, err := c.callBPCustomerSrvAPIRequirementRole("A_BusinessPartnerRole", businessPartner, businessPartnerRole)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callBPCustomerSrvAPIRequirementRole(api, businessPartner, businessPartnerRole string) ([]sap_api_output_formatter.Role, error) {
	url := strings.Join([]string{c.baseURL, "API_BUSINESS_PARTNER", api}, "/")

	param := c.getQueryWithRole(map[string]string{}, businessPartner, businessPartnerRole)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToRole(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Address(businessPartner, addressID string) {
	data, err := c.callBPCustomerSrvAPIRequirementAddress("A_BusinessPartnerAddress", businessPartner, addressID)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callBPCustomerSrvAPIRequirementAddress(api, businessPartner, addressID string) ([]sap_api_output_formatter.Address, error) {
	url := strings.Join([]string{c.baseURL, "API_BUSINESS_PARTNER", api}, "/")

	param := c.getQueryWithAddress(map[string]string{}, businessPartner, addressID)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToAddress(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Bank(businessPartner, bankCountryKey, bankNumber string) {
	data, err := c.callBPCustomerSrvAPIRequirementBank("A_BusinessPartnerBank", businessPartner, bankCountryKey, bankNumber)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callBPCustomerSrvAPIRequirementBank(api, businessPartner, bankCountryKey, bankNumber string) ([]sap_api_output_formatter.Bank, error) {
	url := strings.Join([]string{c.baseURL, "API_BUSINESS_PARTNER", api}, "/")

	param := c.getQueryWithBank(map[string]string{}, businessPartner, bankCountryKey, bankNumber)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToBank(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) BPName(bPName string) {
	data, err := c.callBPCustomerSrvAPIRequirementBPName("A_BusinessPartner", bPName)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callBPCustomerSrvAPIRequirementBPName(api, bPName string) ([]sap_api_output_formatter.General, error) {
	url := strings.Join([]string{c.baseURL, "API_BUSINESS_PARTNER", api}, "/")

	param := c.getQueryWithBPName(map[string]string{}, bPName)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToGeneral(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Customer(customer string) {
	customerData, err := c.callBPCustomerSrvAPIRequirementCustomer("A_Customer", customer)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(customerData)

	salesAreaData, err := c.callToSalesArea(customerData[0].ToSalesArea)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(salesAreaData)

	partnerFunctionData, err := c.callToPartnerFunction(salesAreaData[0].ToPartnerFunction)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(partnerFunctionData)

	companyData, err := c.callToCompany(customerData[0].ToCompany)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(companyData)
}

func (c *SAPAPICaller) callBPCustomerSrvAPIRequirementCustomer(api, customer string) ([]sap_api_output_formatter.Customer, error) {
	url := strings.Join([]string{c.baseURL, "API_BUSINESS_PARTNER", api}, "/")

	param := c.getQueryWithCustomer(map[string]string{}, customer)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToCustomer(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) SalesArea(customer, salesOrganization, distributionChannel, division string) {
	salesAreaData, err := c.callBPCustomerSrvAPIRequirementSalesArea("A_CustomerSalesArea", customer, salesOrganization, distributionChannel, division)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(salesAreaData)

	partnerFunctionData, err := c.callToPartnerFunction(salesAreaData[0].ToPartnerFunction)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(partnerFunctionData)
}

func (c *SAPAPICaller) callBPCustomerSrvAPIRequirementSalesArea(api, customer, salesOrganization, distributionChannel, division string) ([]sap_api_output_formatter.SalesArea, error) {
	url := strings.Join([]string{c.baseURL, "API_BUSINESS_PARTNER", api}, "/")

	param := c.getQueryWithSalesArea(map[string]string{}, customer, salesOrganization, distributionChannel, division)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToSalesArea(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Company(customer, companyCode string) {
	data, err := c.callBPCustomerSrvAPIRequirementCompany("A_CustomerCompany", customer, companyCode)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callBPCustomerSrvAPIRequirementCompany(api, customer, companyCode string) ([]sap_api_output_formatter.Company, error) {
	url := strings.Join([]string{c.baseURL, "API_BUSINESS_PARTNER", api}, "/")

	param := c.getQueryWithCompany(map[string]string{}, customer, companyCode)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToCompany(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) getQueryWithGeneral(params map[string]string, businessPartner string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("BusinessPartner eq '%s'", businessPartner)
	return params
}

func (c *SAPAPICaller) getQueryWithRole(params map[string]string, businessPartner, businessPartnerRole string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("BusinessPartner eq '%s' and BusinessPartnerRole eq '%s'", businessPartner, businessPartnerRole)
	return params
}

func (c *SAPAPICaller) getQueryWithAddress(params map[string]string, businessPartner, addressID string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("BusinessPartner eq '%s' and AddressID eq '%s'", businessPartner, addressID)
	return params
}

func (c *SAPAPICaller) getQueryWithBank(params map[string]string, businessPartner, bankCountryKey, bankNumber string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("BusinessPartner eq '%s' and BankCountryKey eq '%s' and BankNumber eq '%s'", businessPartner, bankCountryKey, bankNumber)
	return params
}

func (c *SAPAPICaller) getQueryWithBPName(params map[string]string, businessPartnerName string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("substringof('%s', BusinessPartnerName)", businessPartnerName)
	return params
}

func (c *SAPAPICaller) getQueryWithCustomer(params map[string]string, customer string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("Customer eq '%s'", customer)
	return params
}

func (c *SAPAPICaller) getQueryWithSalesArea(params map[string]string, customer, salesOrganization, distributionChannel, division string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("Customer eq '%s' and SalesOrganization eq '%s' and DistributionChannel eq '%s' and Division eq '%s'", customer, salesOrganization, distributionChannel, division)
	return params
}

func (c *SAPAPICaller) getQueryWithCompany(params map[string]string, customer, companyCode string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("Customer eq '%s' and CompanyCode eq '%s'", customer, companyCode)
	return params
}
