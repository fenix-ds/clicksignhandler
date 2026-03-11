package clicksignhandler

type ResultData[T any] struct {
	Data T `json:"data"`
}

type ResultList[T any] struct {
	Data []T `json:"data"`
	Meta struct {
		RecordCount float64 `json:"record_count"`
	} `json:"meta"`
	Links struct {
		First string `json:"first"`
		Next  string `json:"next"`
		Last  string `json:"last"`
	}
}
