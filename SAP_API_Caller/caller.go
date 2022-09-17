package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	sap_api_output_formatter "sap-api-integrations-profit-center-reads/SAP_API_Output_Formatter"
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

func (c *SAPAPICaller) AsyncGetProfitCenter(controllingArea, profitCenter, language, profitCenterName string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(controllingArea, profitCenter)
				wg.Done()
			}()
		case "CompanyCodeAssignment":
			func() {
				c.CompanyCodeAssignment(controllingArea, profitCenter)
				wg.Done()
			}()
		case "ProfitCenterName":
			func() {
				c.ProfitCenterName(language, profitCenterName)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) Header(controllingArea, profitCenter string) {
	headerData, err := c.callProfitCenterSrvAPIRequirementHeader("A_ProfitCenter", controllingArea, profitCenter)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(headerData)

	companyCodeAssignmentData, err := c.callToCompanyCodeAssignment(headerData[0].ToCompanyCodeAssignment)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(companyCodeAssignmentData)

	textData, err := c.callToText(headerData[0].ToText)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(textData)
}

func (c *SAPAPICaller) callProfitCenterSrvAPIRequirementHeader(api, controllingArea, profitCenter string) ([]sap_api_output_formatter.Header, error) {
	url := strings.Join([]string{c.baseURL, "API_PROFITCENTER_SRV", api}, "/")
	param := c.getQueryWithHeader(map[string]string{}, controllingArea, profitCenter)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToHeader(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToCompanyCodeAssignment(url string) ([]sap_api_output_formatter.ToCompanyCodeAssignment, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToCompanyCodeAssignment(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToText(url string) ([]sap_api_output_formatter.ToText, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToText(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) CompanyCodeAssignment(controllingArea, profitCenter string) {
	data, err := c.callProfitCenterSrvAPIRequirementCompanyCodeAssignment("A_PrftCtrCompanyCodeAssignment", controllingArea, profitCenter)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callProfitCenterSrvAPIRequirementCompanyCodeAssignment(api, controllingArea, profitCenter string) ([]sap_api_output_formatter.CompanyCodeAssignment, error) {
	url := strings.Join([]string{c.baseURL, "API_PROFITCENTER_SRV", api}, "/")

	param := c.getQueryWithCompanyCodeAssignment(map[string]string{}, controllingArea, profitCenter)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToCompanyCodeAssignment(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) ProfitCenterName(language, profitCenterName string) {
	data, err := c.callProfitCenterSrvAPIRequirementProfitCenterName("A_ProfitCenterText", language, profitCenterName)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)

}

func (c *SAPAPICaller) callProfitCenterSrvAPIRequirementProfitCenterName(api, language, profitCenterName string) ([]sap_api_output_formatter.Text, error) {
	url := strings.Join([]string{c.baseURL, "API_PROFITCENTER_SRV", api}, "/")

	param := c.getQueryWithProfitCenterName(map[string]string{}, language, profitCenterName)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToText(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) getQueryWithHeader(params map[string]string, controllingArea, profitCenter string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("ControllingArea eq '%s' and ProfitCenter eq '%s'", controllingArea, profitCenter)
	return params
}

func (c *SAPAPICaller) getQueryWithCompanyCodeAssignment(params map[string]string, controllingArea, profitCenter string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("ControllingArea eq '%s' and ProfitCenter eq '%s'", controllingArea, profitCenter)
	return params
}

func (c *SAPAPICaller) getQueryWithProfitCenterName(params map[string]string, language, profitCenterName string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("Language eq '%s' and ProfitCenterName eq '%s'", language, profitCenterName)
	return params
}
