package clicksignhandler

import "fmt"

func NewClicksignHandler(param ClicksignParam) (*ClicksignHandler, error) {
	if err := param.checkData(); err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://%s.clicksign.com/api/v3", string(param.Environment))

	return &ClicksignHandler{
		url:         &url,
		accesstoken: &param.Key,
	}, nil
}
