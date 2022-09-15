package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-business-partner-reads-customer/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToGeneral(raw []byte, l *logger.Logger) ([]General, error) {
	pm := &responses.General{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to General. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	general := make([]General, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		general = append(general, General{
			BusinessPartner:               data.BusinessPartner,
			Customer:                      data.Customer,
			Supplier:                      data.Supplier,
			AcademicTitle:                 data.AcademicTitle,
			AuthorizationGroup:            data.AuthorizationGroup,
			BusinessPartnerCategory:       data.BusinessPartnerCategory,
			BusinessPartnerFullName:       data.BusinessPartnerFullName,
			BusinessPartnerGrouping:       data.BusinessPartnerGrouping,
			BusinessPartnerName:           data.BusinessPartnerName,
			CorrespondenceLanguage:        data.CorrespondenceLanguage,
			CreationDate:                  data.CreationDate,
			CreationTime:                  data.CreationTime,
			FirstName:                     data.FirstName,
			Industry:                      data.Industry,
			IsFemale:                      data.IsFemale,
			IsMale:                        data.IsMale,
			IsNaturalPerson:               data.IsNaturalPerson,
			IsSexUnknown:                  data.IsSexUnknown,
			GenderCodeName:                data.GenderCodeName,
			Language:                      data.Language,
			LastChangeDate:                data.LastChangeDate,
			LastChangeTime:                data.LastChangeTime,
			LastName:                      data.LastName,
			OrganizationBPName1:           data.OrganizationBPName1,
			OrganizationBPName2:           data.OrganizationBPName2,
			OrganizationBPName3:           data.OrganizationBPName3,
			OrganizationBPName4:           data.OrganizationBPName4,
			OrganizationFoundationDate:    data.OrganizationFoundationDate,
			OrganizationLiquidationDate:   data.OrganizationLiquidationDate,
			SearchTerm1:                   data.SearchTerm1,
			SearchTerm2:                   data.SearchTerm2,
			AdditionalLastName:            data.AdditionalLastName,
			BirthDate:                     data.BirthDate,
			BusinessPartnerBirthplaceName: data.BusinessPartnerBirthplaceName,
			BusinessPartnerDeathDate:      data.BusinessPartnerDeathDate,
			BusinessPartnerIsBlocked:      data.BusinessPartnerIsBlocked,
			BusinessPartnerType:           data.BusinessPartnerType,
			GroupBusinessPartnerName1:     data.GroupBusinessPartnerName1,
			GroupBusinessPartnerName2:     data.GroupBusinessPartnerName2,
			IndependentAddressID:          data.IndependentAddressID,
			MiddleName:                    data.MiddleName,
			NameCountry:                   data.NameCountry,
			PersonFullName:                data.PersonFullName,
			PersonNumber:                  data.PersonNumber,
			IsMarkedForArchiving:          data.IsMarkedForArchiving,
			BusinessPartnerIDByExtSystem:  data.BusinessPartnerIDByExtSystem,
			TradingPartner:                data.TradingPartner,
			ToRole:                        data.ToRole.Deferred.URI,
			ToAddress:                     data.ToAddress.Deferred.URI,
			ToBank:                        data.ToBank.Deferred.URI,
			ToCustomer:                    data.ToCustomer.Deferred.URI,
		})
	}

	return general, nil
}

func ConvertToRole(raw []byte, l *logger.Logger) ([]Role, error) {
	pm := &responses.Role{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Role. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	role := make([]Role, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		role = append(role, Role{
			BusinessPartner:     data.BusinessPartner,
			BusinessPartnerRole: data.BusinessPartnerRole,
			ValidFrom:           data.ValidFrom,
			ValidTo:             data.ValidTo,
		})
	}

	return role, nil
}

func ConvertToAddress(raw []byte, l *logger.Logger) ([]Address, error) {
	pm := &responses.Address{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Address. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	address := make([]Address, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		address = append(address, Address{
			BusinessPartner:   data.BusinessPartner,
			AddressID:         data.AddressID,
			ValidityStartDate: data.ValidityStartDate,
			ValidityEndDate:   data.ValidityEndDate,
			Country:           data.Country,
			Region:            data.Region,
			StreetName:        data.StreetName,
			CityName:          data.CityName,
			PostalCode:        data.PostalCode,
			Language:          data.Language,
		})
	}

	return address, nil
}

func ConvertToBank(raw []byte, l *logger.Logger) ([]Bank, error) {
	pm := &responses.Bank{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Bank. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	bank := make([]Bank, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		bank = append(bank, Bank{
			BusinessPartner:          data.BusinessPartner,
			BankIdentification:       data.BankIdentification,
			BankCountryKey:           data.BankCountryKey,
			BankName:                 data.BankName,
			BankNumber:               data.BankNumber,
			SWIFTCode:                data.SWIFTCode,
			BankControlKey:           data.BankControlKey,
			BankAccountHolderName:    data.BankAccountHolderName,
			BankAccountName:          data.BankAccountName,
			ValidityStartDate:        data.ValidityStartDate,
			ValidityEndDate:          data.ValidityEndDate,
			IBAN:                     data.IBAN,
			IBANValidityStartDate:    data.IBANValidityStartDate,
			BankAccount:              data.BankAccount,
			BankAccountReferenceText: data.BankAccountReferenceText,
			CollectionAuthInd:        data.CollectionAuthInd,
			CityName:                 data.CityName,
			AuthorizationGroup:       data.AuthorizationGroup,
		})
	}

	return bank, nil
}

func ConvertToCustomer(raw []byte, l *logger.Logger) ([]Customer, error) {
	pm := &responses.Customer{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Customer. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	customer := make([]Customer, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		customer = append(customer, Customer{
			Customer:                    data.Customer,
			AuthorizationGroup:          data.AuthorizationGroup,
			BillingIsBlockedForCustomer: data.BillingIsBlockedForCustomer,
			CreationDate:                data.CreationDate,
			CustomerAccountGroup:        data.CustomerAccountGroup,
			CustomerClassification:      data.CustomerClassification,
			CustomerFullName:            data.CustomerFullName,
			CustomerName:                data.CustomerName,
			DeliveryIsBlocked:           data.DeliveryIsBlocked,
			OrderIsBlockedForCustomer:   data.OrderIsBlockedForCustomer,
			PostingIsBlocked:            data.PostingIsBlocked,
			Supplier:                    data.Supplier,
			CustomerCorporateGroup:      data.CustomerCorporateGroup,
			Industry:                    data.Industry,
			TaxNumber1:                  data.TaxNumber1,
			DeletionIndicator:           data.DeletionIndicator,
			CityCode:                    data.CityCode,
			County:                      data.County,
			ToSalesArea:                 data.ToSalesArea.Deferred.URI,
			ToCompany:                   data.ToCompany.Deferred.URI,
		})
	}

	return customer, nil
}

func ConvertToSalesArea(raw []byte, l *logger.Logger) ([]SalesArea, error) {
	pm := &responses.SalesArea{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to SalesArea. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	salesArea := make([]SalesArea, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		salesArea = append(salesArea, SalesArea{
			Customer:                       data.Customer,
			SalesOrganization:              data.SalesOrganization,
			DistributionChannel:            data.DistributionChannel,
			Division:                       data.Division,
			CompleteDeliveryIsDefined:      data.CompleteDeliveryIsDefined,
			Currency:                       data.Currency,
			CustomerAccountAssignmentGroup: data.CustomerAccountAssignmentGroup,
			CustomerPaymentTerms:           data.CustomerPaymentTerms,
			CustomerPriceGroup:             data.CustomerPriceGroup,
			CustomerPricingProcedure:       data.CustomerPricingProcedure,
			DeliveryPriority:               data.DeliveryPriority,
			IncotermsClassification:        data.IncotermsClassification,
			InvoiceDate:                    data.InvoiceDate,
			OrderCombinationIsAllowed:      data.OrderCombinationIsAllowed,
			PartialDeliveryIsAllowed:       data.PartialDeliveryIsAllowed,
			PriceListType:                  data.PriceListType,
			SalesGroup:                     data.SalesGroup,
			SalesOffice:                    data.SalesOffice,
			ShippingCondition:              data.ShippingCondition,
			SupplyingPlant:                 data.SupplyingPlant,
			SalesDistrict:                  data.SalesDistrict,
			InvoiceListSchedule:            data.InvoiceListSchedule,
			ExchangeRateType:               data.ExchangeRateType,
			OrderIsBlockedForCustomer:      data.OrderIsBlockedForCustomer,
			DeliveryIsBlockedForCustomer:   data.DeliveryIsBlockedForCustomer,
			BillingIsBlockedForCustomer:    data.BillingIsBlockedForCustomer,
			DeletionIndicator:              data.DeletionIndicator,
			ToPartnerFunction:              data.ToPartnerFunction.Deferred.URI,
		})
	}

	return salesArea, nil
}

func ConvertToCompany(raw []byte, l *logger.Logger) ([]Company, error) {
	pm := &responses.Company{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Company. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	company := make([]Company, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		company = append(company, Company{
			Customer:                       data.Customer,
			CompanyCode:                    data.CompanyCode,
			APARToleranceGroup:             data.APARToleranceGroup,
			CustomerSupplierClearingIsUsed: data.CustomerSupplierClearingIsUsed,
			HouseBank:                      data.HouseBank,
			PaymentMethodsList:             data.PaymentMethodsList,
			PaymentTerms:                   data.PaymentTerms,
			ReconciliationAccount:          data.ReconciliationAccount,
			DeletionIndicator:              data.DeletionIndicator,
		})
	}

	return company, nil
}

func ConvertToToRole(raw []byte, l *logger.Logger) ([]ToRole, error) {
	pm := &responses.ToRole{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToRole. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	toRole := make([]ToRole, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toRole = append(toRole, ToRole{
			BusinessPartner:     data.BusinessPartner,
			BusinessPartnerRole: data.BusinessPartnerRole,
			ValidFrom:           data.ValidFrom,
			ValidTo:             data.ValidTo,
		})
	}

	return toRole, nil
}

func ConvertToToAddress(raw []byte, l *logger.Logger) ([]ToAddress, error) {
	pm := &responses.ToAddress{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToAddress. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	toAddress := make([]ToAddress, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toAddress = append(toAddress, ToAddress{
			BusinessPartner:   data.BusinessPartner,
			AddressID:         data.AddressID,
			ValidityStartDate: data.ValidityStartDate,
			ValidityEndDate:   data.ValidityEndDate,
			Country:           data.Country,
			Region:            data.Region,
			StreetName:        data.StreetName,
			CityName:          data.CityName,
			PostalCode:        data.PostalCode,
			Language:          data.Language,
		})
	}

	return toAddress, nil
}

func ConvertToToBank(raw []byte, l *logger.Logger) ([]ToBank, error) {
	pm := &responses.ToBank{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToBank. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	toBank := make([]ToBank, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toBank = append(toBank, ToBank{
			BusinessPartner:          data.BusinessPartner,
			BankIdentification:       data.BankIdentification,
			BankCountryKey:           data.BankCountryKey,
			BankName:                 data.BankName,
			BankNumber:               data.BankNumber,
			SWIFTCode:                data.SWIFTCode,
			BankControlKey:           data.BankControlKey,
			BankAccountHolderName:    data.BankAccountHolderName,
			BankAccountName:          data.BankAccountName,
			ValidityStartDate:        data.ValidityStartDate,
			ValidityEndDate:          data.ValidityEndDate,
			IBAN:                     data.IBAN,
			IBANValidityStartDate:    data.IBANValidityStartDate,
			BankAccount:              data.BankAccount,
			BankAccountReferenceText: data.BankAccountReferenceText,
			CollectionAuthInd:        data.CollectionAuthInd,
			CityName:                 data.CityName,
			AuthorizationGroup:       data.AuthorizationGroup,
		})
	}

	return toBank, nil
}

func ConvertToToCustomer(raw []byte, l *logger.Logger) (*ToCustomer, error) {
	pm := &responses.ToCustomer{}

	err := json.Unmarshal(raw, &pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToCustomer. unmarshal error: %w", err)
	}

	return &ToCustomer{
		Customer:                    pm.D.Customer,
		AuthorizationGroup:          pm.D.AuthorizationGroup,
		BillingIsBlockedForCustomer: pm.D.BillingIsBlockedForCustomer,
		CreationDate:                pm.D.CreationDate,
		CustomerAccountGroup:        pm.D.CustomerAccountGroup,
		CustomerClassification:      pm.D.CustomerClassification,
		CustomerFullName:            pm.D.CustomerFullName,
		CustomerName:                pm.D.CustomerName,
		DeliveryIsBlocked:           pm.D.DeliveryIsBlocked,
		OrderIsBlockedForCustomer:   pm.D.OrderIsBlockedForCustomer,
		PostingIsBlocked:            pm.D.PostingIsBlocked,
		Supplier:                    pm.D.Supplier,
		CustomerCorporateGroup:      pm.D.CustomerCorporateGroup,
		Industry:                    pm.D.Industry,
		TaxNumber1:                  pm.D.TaxNumber1,
		DeletionIndicator:           pm.D.DeletionIndicator,
		CityCode:                    pm.D.CityCode,
		County:                      pm.D.County,
		ToSalesArea:                 pm.D.ToSalesArea.Deferred.URI,
		ToCompany:                   pm.D.ToCompany.Deferred.URI,
	}, nil
}

func ConvertToToSalesArea(raw []byte, l *logger.Logger) ([]ToSalesArea, error) {
	pm := &responses.ToSalesArea{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToSalesArea. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	toSalesArea := make([]ToSalesArea, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toSalesArea = append(toSalesArea, ToSalesArea{
			Customer:                       data.Customer,
			SalesOrganization:              data.SalesOrganization,
			DistributionChannel:            data.DistributionChannel,
			Division:                       data.Division,
			CompleteDeliveryIsDefined:      data.CompleteDeliveryIsDefined,
			Currency:                       data.Currency,
			CustomerAccountAssignmentGroup: data.CustomerAccountAssignmentGroup,
			CustomerPaymentTerms:           data.CustomerPaymentTerms,
			CustomerPriceGroup:             data.CustomerPriceGroup,
			CustomerPricingProcedure:       data.CustomerPricingProcedure,
			DeliveryPriority:               data.DeliveryPriority,
			IncotermsClassification:        data.IncotermsClassification,
			InvoiceDate:                    data.InvoiceDate,
			OrderCombinationIsAllowed:      data.OrderCombinationIsAllowed,
			PartialDeliveryIsAllowed:       data.PartialDeliveryIsAllowed,
			PriceListType:                  data.PriceListType,
			SalesGroup:                     data.SalesGroup,
			SalesOffice:                    data.SalesOffice,
			ShippingCondition:              data.ShippingCondition,
			SupplyingPlant:                 data.SupplyingPlant,
			SalesDistrict:                  data.SalesDistrict,
			InvoiceListSchedule:            data.InvoiceListSchedule,
			ExchangeRateType:               data.ExchangeRateType,
			OrderIsBlockedForCustomer:      data.OrderIsBlockedForCustomer,
			DeliveryIsBlockedForCustomer:   data.DeliveryIsBlockedForCustomer,
			BillingIsBlockedForCustomer:    data.BillingIsBlockedForCustomer,
			DeletionIndicator:              data.DeletionIndicator,
			ToPartnerFunction:              data.ToPartnerFunction.Deferred.URI,
		})
	}

	return toSalesArea, nil
}

func ConvertToToPartnerFunction(raw []byte, l *logger.Logger) ([]ToPartnerFunction, error) {
	pm := &responses.ToPartnerFunction{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToPartnerFunction. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	toPartnerFunction := make([]ToPartnerFunction, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toPartnerFunction = append(toPartnerFunction, ToPartnerFunction{
			Customer:                   data.Customer,
			SalesOrganization:          data.SalesOrganization,
			DistributionChannel:        data.DistributionChannel,
			Division:                   data.Division,
			PartnerCounter:             data.PartnerCounter,
			PartnerFunction:            data.PartnerFunction,
			BPCustomerNumber:           data.BPCustomerNumber,
			CustomerPartnerDescription: data.CustomerPartnerDescription,
			DefaultPartner:             data.DefaultPartner,
			Supplier:                   data.Supplier,
			AuthorizationGroup:         data.AuthorizationGroup,
		})
	}

	return toPartnerFunction, nil
}

func ConvertToToCompany(raw []byte, l *logger.Logger) ([]ToCompany, error) {
	pm := &responses.ToCompany{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToCompany. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	toCompany := make([]ToCompany, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toCompany = append(toCompany, ToCompany{
			Customer:                       data.Customer,
			CompanyCode:                    data.CompanyCode,
			APARToleranceGroup:             data.APARToleranceGroup,
			CustomerSupplierClearingIsUsed: data.CustomerSupplierClearingIsUsed,
			HouseBank:                      data.HouseBank,
			PaymentMethodsList:             data.PaymentMethodsList,
			PaymentTerms:                   data.PaymentTerms,
			ReconciliationAccount:          data.ReconciliationAccount,
			DeletionIndicator:              data.DeletionIndicator,
		})
	}

	return toCompany, nil
}
