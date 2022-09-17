package sap_api_output_formatter

type ProfitCenter struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Product       string `json:"Product"`
	APISchema     string `json:"api_schema"`
	ProfitCenter  string `json:"profit_center_code"`
	Deleted       bool   `json:"deleted"`
}

type Header struct {
	ControllingArea                string `json:"ControllingArea"`
	ProfitCenter                   string `json:"ProfitCenter"`
	ValidityEndDate                string `json:"ValidityEndDate"`
	ProfitCtrResponsiblePersonName string `json:"ProfitCtrResponsiblePersonName"`
	CompanyCode                    string `json:"CompanyCode"`
	ProfitCtrResponsibleUser       string `json:"ProfitCtrResponsibleUser"`
	ValidityStartDate              string `json:"ValidityStartDate"`
	Department                     string `json:"Department"`
	ProfitCenterStandardHierarchy  string `json:"ProfitCenterStandardHierarchy"`
	Segment                        string `json:"Segment"`
	ProfitCenterIsBlocked          string `json:"ProfitCenterIsBlocked"`
	FormulaPlanningTemplate        string `json:"FormulaPlanningTemplate"`
	FormOfAddress                  string `json:"FormOfAddress"`
	AddressName                    string `json:"AddressName"`
	AdditionalName                 string `json:"AdditionalName"`
	ProfitCenterAddrName3          string `json:"ProfitCenterAddrName3"`
	ProfitCenterAddrName4          string `json:"ProfitCenterAddrName4"`
	StreetAddressName              string `json:"StreetAddressName"`
	POBox                          string `json:"POBox"`
	CityName                       string `json:"CityName"`
	PostalCode                     string `json:"PostalCode"`
	District                       string `json:"District"`
	Country                        string `json:"Country"`
	Region                         string `json:"Region"`
	TaxJurisdiction                string `json:"TaxJurisdiction"`
	Language                       string `json:"Language"`
	PhoneNumber1                   string `json:"PhoneNumber1"`
	PhoneNumber2                   string `json:"PhoneNumber2"`
	TeleboxNumber                  string `json:"TeleboxNumber"`
	TelexNumber                    string `json:"TelexNumber"`
	FaxNumber                      string `json:"FaxNumber"`
	DataCommunicationPhoneNumber   string `json:"DataCommunicationPhoneNumber"`
	ProfitCenterPrinterName        string `json:"ProfitCenterPrinterName"`
	ProfitCenterCreatedByUser      string `json:"ProfitCenterCreatedByUser"`
	ProfitCenterCreationDate       string `json:"ProfitCenterCreationDate"`
	Yy1LobPrc                      string `json:"YY1_LOB_PRC"`
	ToCompanyCodeAssignment        string `json:"to_PrftCtrCompanyCodeAssignment"`
	ToText                         string `json:"to_Text"`
}

type CompanyCodeAssignment struct {
	ControllingArea string `json:"ControllingArea"`
	ProfitCenter    string `json:"ProfitCenter"`
	CompanyCode     string `json:"CompanyCode"`
}

type Text struct {
	Language             string `json:"Language"`
	ControllingArea      string `json:"ControllingArea"`
	ProfitCenter         string `json:"ProfitCenter"`
	ValidityEndDate      string `json:"ValidityEndDate"`
	ValidityStartDate    string `json:"ValidityStartDate"`
	ProfitCenterName     string `json:"ProfitCenterName"`
	ProfitCenterLongName string `json:"ProfitCenterLongName"`
}

type ToCompanyCodeAssignment struct {
	ControllingArea string `json:"ControllingArea"`
	ProfitCenter    string `json:"ProfitCenter"`
	CompanyCode     string `json:"CompanyCode"`
}

type ToText struct {
	Language             string `json:"Language"`
	ControllingArea      string `json:"ControllingArea"`
	ProfitCenter         string `json:"ProfitCenter"`
	ValidityEndDate      string `json:"ValidityEndDate"`
	ValidityStartDate    string `json:"ValidityStartDate"`
	ProfitCenterName     string `json:"ProfitCenterName"`
	ProfitCenterLongName string `json:"ProfitCenterLongName"`
}
