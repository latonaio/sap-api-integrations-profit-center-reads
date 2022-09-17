package responses

type Text struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			Language             string `json:"Language"`
			ControllingArea      string `json:"ControllingArea"`
			ProfitCenter         string `json:"ProfitCenter"`
			ValidityEndDate      string `json:"ValidityEndDate"`
			ValidityStartDate    string `json:"ValidityStartDate"`
			ProfitCenterName     string `json:"ProfitCenterName"`
			ProfitCenterLongName string `json:"ProfitCenterLongName"`
		} `json:"results"`
	} `json:"d"`
}
