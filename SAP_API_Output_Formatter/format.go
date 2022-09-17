package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-profit-center-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) ([]Header, error) {
	pm := &responses.Header{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	header := make([]Header, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		header = append(header, Header{
			ControllingArea:                data.ControllingArea,
			ProfitCenter:                   data.ProfitCenter,
			ValidityEndDate:                data.ValidityEndDate,
			ProfitCtrResponsiblePersonName: data.ProfitCtrResponsiblePersonName,
			CompanyCode:                    data.CompanyCode,
			ProfitCtrResponsibleUser:       data.ProfitCtrResponsibleUser,
			ValidityStartDate:              data.ValidityStartDate,
			Department:                     data.Department,
			ProfitCenterStandardHierarchy:  data.ProfitCenterStandardHierarchy,
			Segment:                        data.Segment,
			ProfitCenterIsBlocked:          data.ProfitCenterIsBlocked,
			FormulaPlanningTemplate:        data.FormulaPlanningTemplate,
			FormOfAddress:                  data.FormOfAddress,
			AddressName:                    data.AddressName,
			AdditionalName:                 data.AdditionalName,
			ProfitCenterAddrName3:          data.ProfitCenterAddrName3,
			ProfitCenterAddrName4:          data.ProfitCenterAddrName4,
			StreetAddressName:              data.StreetAddressName,
			POBox:                          data.POBox,
			CityName:                       data.CityName,
			PostalCode:                     data.PostalCode,
			District:                       data.District,
			Country:                        data.Country,
			Region:                         data.Region,
			TaxJurisdiction:                data.TaxJurisdiction,
			Language:                       data.Language,
			PhoneNumber1:                   data.PhoneNumber1,
			PhoneNumber2:                   data.PhoneNumber2,
			TeleboxNumber:                  data.TeleboxNumber,
			TelexNumber:                    data.TelexNumber,
			FaxNumber:                      data.FaxNumber,
			DataCommunicationPhoneNumber:   data.DataCommunicationPhoneNumber,
			ProfitCenterPrinterName:        data.ProfitCenterPrinterName,
			ProfitCenterCreatedByUser:      data.ProfitCenterCreatedByUser,
			ProfitCenterCreationDate:       data.ProfitCenterCreationDate,
			Yy1LobPrc:                      data.Yy1LobPrc,
			ToCompanyCodeAssignment:        data.ToCompanyCodeAssignment.Deferred.URI,
			ToText:                         data.ToText.Deferred.URI,
		})
	}

	return header, nil
}

func ConvertToCompanyCodeAssignment(raw []byte, l *logger.Logger) ([]CompanyCodeAssignment, error) {
	pm := &responses.CompanyCodeAssignment{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to CompanyCodeAssignment. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	companyCodeAssignment := make([]CompanyCodeAssignment, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		companyCodeAssignment = append(companyCodeAssignment, CompanyCodeAssignment{
			ControllingArea: data.ControllingArea,
			ProfitCenter:    data.ProfitCenter,
			CompanyCode:     data.CompanyCode,
		})
	}

	return companyCodeAssignment, nil
}

func ConvertToText(raw []byte, l *logger.Logger) ([]Text, error) {
	pm := &responses.Text{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Text. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	text := make([]Text, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		text = append(text, Text{
			Language:             data.Language,
			ControllingArea:      data.ControllingArea,
			ProfitCenter:         data.ProfitCenter,
			ValidityEndDate:      data.ValidityEndDate,
			ValidityStartDate:    data.ValidityStartDate,
			ProfitCenterName:     data.ProfitCenterName,
			ProfitCenterLongName: data.ProfitCenterLongName,
		})
	}

	return text, nil
}

func ConvertToToCompanyCodeAssignment(raw []byte, l *logger.Logger) ([]ToCompanyCodeAssignment, error) {
	pm := &responses.ToCompanyCodeAssignment{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToCompanyCodeAssignment. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	toCompanyCodeAssignment := make([]ToCompanyCodeAssignment, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toCompanyCodeAssignment = append(toCompanyCodeAssignment, ToCompanyCodeAssignment{
			ControllingArea: data.ControllingArea,
			ProfitCenter:    data.ProfitCenter,
			CompanyCode:     data.CompanyCode,
		})
	}

	return toCompanyCodeAssignment, nil
}

func ConvertToToText(raw []byte, l *logger.Logger) ([]ToText, error) {
	pm := &responses.Text{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToText. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	toText := make([]ToText, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toText = append(toText, ToText{
			Language:             data.Language,
			ControllingArea:      data.ControllingArea,
			ProfitCenter:         data.ProfitCenter,
			ValidityEndDate:      data.ValidityEndDate,
			ValidityStartDate:    data.ValidityStartDate,
			ProfitCenterName:     data.ProfitCenterName,
			ProfitCenterLongName: data.ProfitCenterLongName,
		})
	}

	return toText, nil
}
