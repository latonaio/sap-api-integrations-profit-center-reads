package sap_api_input_reader

type EC_MC struct {
	ConnectionKey   string `json:"connection_key"`
	Result          bool   `json:"result"`
	RedisKey        string `json:"redis_key"`
	Filepath        string `json:"filepath"`
	ProductionOrder struct {
		ManufacturingOrder             string `json:"document_no"`
		DeliverTo                      string `json:"deliver_to"`
		MfgOrderItemPlannedTotalQty    string `json:"quantity"`
		MfgOrderItemActualDeviationQty string `json:"picked_quantity"`
		Price                          string `json:"price"`
		Batch                          string `json:"batch"`
	} `json:"document"`
	ManufacturingOrder struct {
		ManufacturingOrder             string `json:"document_no"`
		StatusCode                     string `json:"status"`
		DeliverTo                      string `json:"deliver_to"`
		MfgOrderItemPlannedTotalQty    string `json:"quantity"`
		MfgOrderItemActualDeviationQty string `json:"completed_quantity"`
		PlannedStartDate               string `json:"planned_start_date"`
		MfgOrderItemPlndDeliveryDate   string `json:"planned_validated_date"`
		ActualStartDate                string `json:"actual_start_date"`
		MfgOrderItemActualDeliveryDate string `json:"actual_validated_date"`
		Batch                          string `json:"batch"`
		BillOfOperations               struct {
			OrderIntBillOfOperationsItem string `json:"work_no"`
			OpPlannedTotalQuantity       string `json:"quantity"`
			OpTotalConfirmedYieldQty     string `json:"completed_quantity"`
			ErroredQuantity              string `json:"errored_quantity"`
			Component                    string `json:"component"`
			MaterialCompOriginalQuantity string `json:"planned_component_quantity"`
			OpErlstSchedldExecStrtDte    string `json:"planned_start_date"`
			OpErlstSchedldExecStrtTme    string `json:"planned_start_time"`
			OpErlstSchedldExecEndDte     string `json:"planned_validated_date"`
			OpErlstSchedldExecEndTme     string `json:"planned_validated_time"`
			OpActualExecutionStartDate   string `json:"actual_start_date"`
			OpActualExecutionStartTime   string `json:"actual_start_time"`
			OpActualExecutionEndDate     string `json:"actual_validated_date"`
			OpActualExecutionEndTime     string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"production_order"`
	APISchema                      string `json:"api_schema"`
	MaterialCode                   string `json:"material_code"`
	ProductionPlant                string `json:"plant/supplier"`
	Stock                          string `json:"stock"`
	ManufacturingOrderType         string `json:"document_type"`
	ManufacturingOrderNo           string `json:"document_no"`
	MfgOrderItemPlndDeliveryDate   string `json:"planned_date"`
	MfgOrderItemActualDeliveryDate string `json:"validated_date"`
	Deleted                        bool   `json:"deleted"`
}

type SDC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	ProfitCenter  struct {
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
		Text                           struct {
			Language             string `json:"Language"`
			ValidityEndDate      string `json:"ValidityEndDate"`
			ValidityStartDate    string `json:"ValidityStartDate"`
			ProfitCenterName     string `json:"ProfitCenterName"`
			ProfitCenterLongName string `json:"ProfitCenterLongName"`
		} `json:"Text"`
		CompanyCodeAssignment struct {
			CompanyCode string `json:"CompanyCode"`
		} `json:"CompanyCodeAssignment"`
	} `json:"ProfitCenter"`
	APISchema        string   `json:"api_schema"`
	Accepter         []string `json:"accepter"`
	ProfitCenterCode string   `json:"profit_center_code"`
	Deleted          bool     `json:"deleted"`
}
