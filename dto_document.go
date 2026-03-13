package clicksignhandler

type DocumentData struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Links struct {
		Self  string `json:"self"`
		Files struct {
			Original string  `json:"original"`
			Signed   *string `json:"signed"`
		} `json:"files"`
	} `json:"links"`
	Attributes struct {
		Status   string `json:"status"`
		Filename string `json:"filename"`
		Template any    `json:"template"`
		Metadata struct {
			PositionSignFields []any `json:"position_sign_fields"`
		} `json:"metadata"`
		Migrated bool   `json:"migrated"`
		Created  string `json:"created"`
		Modified string `json:"modified"`
	} `json:"attributes"`
}

type DocumentCreate struct {
	Envelope   *EnvelopeData
	FileType   DocumentFileType
	FileName   string
	FileBase64 string
}

type DocumentEvent_AddSigner struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Name string `json:"name"`
		Data struct {
			User struct {
				Email string `json:"email"`
				Name  string `json:"name"`
			} `json:"user"`
			Account struct {
				Key string `json:"key"`
			} `json:"account"`
			Signers []struct {
				SignAs                  string      `json:"sign_as"`
				ListKey                 string      `json:"list_key"`
				RequestedSealPages      interface{} `json:"requested_seal_pages"`
				Key                     string      `json:"key"`
				Email                   string      `json:"email"`
				Name                    string      `json:"name"`
				Birthday                string      `json:"birthday"`
				CreatedAt               string      `json:"created_at"`
				Documentation           string      `json:"documentation"`
				HasDocumentation        bool        `json:"has_documentation"`
				Auths                   []string    `json:"auths"`
				SelfieEnabled           bool        `json:"selfie_enabled"`
				HandwrittenEnabled      bool        `json:"handwritten_enabled"`
				OfficialDocumentEnabled bool        `json:"official_document_enabled"`
				LivenessEnabled         bool        `json:"liveness_enabled"`
				FacialBiometricsEnabled bool        `json:"facial_biometrics_enabled"`
				CommunicateBy           string      `json:"communicate_by"`
				PhoneNumber             interface{} `json:"phone_number"`
				PhoneNumberHash         interface{} `json:"phone_number_hash"`
				FederalDataValidation   interface{} `json:"federal_data_validation"`
				URL                     string      `json:"url"`
			} `json:"signers"`
		} `json:"data"`
		Created string `json:"created"`
	} `json:"attributes"`
}

type DocumentEvent_Sign struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Name string `json:"name"`
		Data struct {
			Signer struct {
				SignAs                  string      `json:"sign_as"`
				ListKey                 string      `json:"list_key"`
				Key                     string      `json:"key"`
				Email                   string      `json:"email"`
				Name                    string      `json:"name"`
				Birthday                interface{} `json:"birthday"`
				CreatedAt               string      `json:"created_at"`
				Documentation           interface{} `json:"documentation"`
				HasDocumentation        bool        `json:"has_documentation"`
				Auths                   []string    `json:"auths"`
				SelfieEnabled           bool        `json:"selfie_enabled"`
				HandwrittenEnabled      bool        `json:"handwritten_enabled"`
				OfficialDocumentEnabled bool        `json:"official_document_enabled"`
				LivenessEnabled         bool        `json:"liveness_enabled"`
				FacialBiometricsEnabled bool        `json:"facial_biometrics_enabled"`
				CommunicateBy           string      `json:"communicate_by"`
				PhoneNumber             interface{} `json:"phone_number"`
				PhoneNumberHash         interface{} `json:"phone_number_hash"`
				FederalDataValidation   interface{} `json:"federal_data_validation"`
				URL                     string      `json:"url"`
				Address                 string      `json:"address"`
				Longitude               interface{} `json:"longitude"`
				Latitude                interface{} `json:"latitude"`
			} `json:"signer"`
			SecretHmac interface{} `json:"secret_hmac"`
			Account    struct {
				Key                             string `json:"key"`
				TimestampSignatureFunctionality bool   `json:"timestamp_signature_functionality"`
			} `json:"account"`
			LogVersion string `json:"log_version"`
			URL        string `json:"url"`
		} `json:"data"`
		Created string `json:"created"`
	} `json:"attributes"`
}
