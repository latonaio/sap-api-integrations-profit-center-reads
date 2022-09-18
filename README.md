# sap-api-integrations-profit-center-reads
sap-api-integrations-profit-center-reads は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で 利益センタ を取得するマイクロサービスです。        
sap-api-integrations-profit-center-reads には、サンプルのAPI Json フォーマットが含まれています。       
sap-api-integrations-profit-center-reads は、オンプレミス版である（＝クラウド版ではない）SAPS4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。       
https://api.sap.com/api/OP_API_PROFITCENTER_SRV_0001/overview  

## 動作環境  

sap-api-integrations-profit-center-reads は、主にエッジコンピューティング環境における動作にフォーカスしています。  
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。  
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須）      

## クラウド環境での利用

sap-api-integrations-profit-center-reads は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。     

## 本レポジトリ が 対応する API サービス
sap-api-integrations-profit-center-reads が対応する APIサービス は、次のものです。  

* APIサービス概要説明 URL: https://api.sap.com/api/OP_API_PROFITCENTER_SRV_0001/overview    
* APIサービス名(=baseURL): API_PROFITCENTER_SRV   

## 本レポジトリ に 含まれる API名
sap-api-integrations-profit-center-reads には、次の API をコールするためのリソースが含まれています。  

* A_ProfitCenter（利益センタ - ヘッダ）※利益センタの詳細データを取得するために、ToCompanyCodeAssignmen、ToTextと合わせて利用されます。
* A_PrftCtrCompanyCodeAssignment（利益センタ - 会社コード割当）
* A_ProfitCenterText（利益センタ - テキスト）
* ToCompanyCodeAssignment（利益センタ - 会社コード割当 ※To）
* ToText（利益センタ - テキスト ※To）

## API への 値入力条件 の 初期値
sap-api-integrations-profit-center-reads において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inoutSDC.ProfitCenter.ControllingArea（管理領域）
* inoutSDC.ProfitCenter.ProfitCenter（利益センタ）
* inoutSDC.ProfitCenter.Text.Language（言語）
* inoutSDC.ProfitCenter.Text.ProfitCenterName（利益センタ名）

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"Header" が指定されています。    
  
```
	"api_schema": "SAPProfitCenterReads",
	"accepter": ["Header"],
	"profit_center_code": "100000",
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "SAPProfitCenterReads",
	"accepter": ["All"],
	"profit_center_code": "100000",
	"deleted": false
```


## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
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
```

## Output  
本マイクロサービスでは、[golang-logging-library-for-sap](https://github.com/latonaio/golang-logging-library-for-sap) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP 利益センタ の ヘッダデータ が取得された結果の JSON の例です。  
以下の項目のうち、"ControllingArea" ～ "to_Text" は、/SAP_API_Output_Formatter/type.go 内 の Type Header {} による出力結果です。"cursor" ～ "time"は、golang-logging-library-for-sap による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-profit-center-reads/SAP_API_Caller/caller.go#L65",
	"function": "sap-api-integrations-profit-center-reads/SAP_API_Caller.(*SAPAPICaller).Header",
	"level": "INFO",
	"message": [
		{
			"ControllingArea": "0001",
			"ProfitCenter": "100000",
			"ValidityEndDate": "9999-12-31",
			"ProfitCtrResponsiblePersonName": "Test",
			"CompanyCode": "",
			"ProfitCtrResponsibleUser": "",
			"ValidityStartDate": "1998-01-01",
			"Department": "",
			"ProfitCenterStandardHierarchy": "P0001",
			"Segment": "",
			"ProfitCenterIsBlocked": "",
			"FormulaPlanningTemplate": "",
			"FormOfAddress": "",
			"AddressName": "",
			"AdditionalName": "",
			"ProfitCenterAddrName3": "",
			"ProfitCenterAddrName4": "",
			"StreetAddressName": "",
			"POBox": "",
			"CityName": "",
			"PostalCode": "",
			"District": "",
			"Country": "",
			"Region": "",
			"TaxJurisdiction": "",
			"Language": "JA",
			"PhoneNumber1": "",
			"PhoneNumber2": "",
			"TeleboxNumber": "",
			"TelexNumber": "",
			"FaxNumber": "",
			"DataCommunicationPhoneNumber": "",
			"ProfitCenterPrinterName": "",
			"ProfitCenterCreatedByUser": "LTN0001",
			"ProfitCenterCreationDate": "2022-09-16",
			"YY1_LOB_PRC": "",
			"to_PrftCtrCompanyCodeAssignment": "http://100.21.57.120:8080/sap/opu/odata/sap/API_PROFITCENTER_SRV/A_ProfitCenter(ControllingArea='0001',ProfitCenter='100000',ValidityEndDate=datetime'9999-12-31T00%3A00%3A00')/to_PrftCtrCompanyCodeAssignment",
			"to_Text": "http://100.21.57.120:8080/sap/opu/odata/sap/API_PROFITCENTER_SRV/A_ProfitCenter(ControllingArea='0001',ProfitCenter='100000',ValidityEndDate=datetime'9999-12-31T00%3A00%3A00')/to_Text"
		}
	],
	"time": "2022-09-18T09:49:22+09:00"
}
```
