package responses

type ToCompanyCodeAssignment struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			ControllingArea string `json:"ControllingArea"`
			ProfitCenter    string `json:"ProfitCenter"`
			CompanyCode     string `json:"CompanyCode"`
		} `json:"results"`
	} `json:"d"`
}
