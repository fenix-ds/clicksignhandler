package clicksignhandler

import "time"

type EnvelopeCreate struct {
	Name           string
	DeadlineAt     *time.Time
	RemindInterval *uint
}

type EnvelopeData struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Attributes struct {
		Name              string      `json:"name"`
		Status            string      `json:"status"`
		DeadlineAt        string      `json:"deadline_at"`
		Locale            string      `json:"locale"`
		AutoClose         bool        `json:"auto_close"`
		RubricEnabled     bool        `json:"rubric_enabled"`
		RemindInterval    CSIntString `json:"remind_interval"`
		BlockAfterRefusal bool        `json:"block_after_refusal"`
		DefaultSubject    any         `json:"default_subject"`
		DefaultMessage    any         `json:"default_message"`
		Created           string      `json:"created"`
		Modified          string      `json:"modified"`
		Migrated          bool        `json:"migrated"`
		Metadata          struct {
		} `json:"metadata"`
	} `json:"attributes"`
	Relationships struct {
		Folder struct {
			Links struct {
				Self    string `json:"self"`
				Related string `json:"related"`
			} `json:"links"`
		} `json:"folder"`
		Documents struct {
			Links struct {
				Self    string `json:"self"`
				Related string `json:"related"`
			} `json:"links"`
		} `json:"documents"`
		Signers struct {
			Links struct {
				Self    string `json:"self"`
				Related string `json:"related"`
			} `json:"links"`
		} `json:"signers"`
		Requirements struct {
			Links struct {
				Self    string `json:"self"`
				Related string `json:"related"`
			} `json:"links"`
		} `json:"requirements"`
	} `json:"relationships"`
}

type EnvelopeEvent struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Name string `json:"name"`
		Data struct {
			Signer struct {
				Key                     string      `json:"key"`
				Email                   string      `json:"email"`
				Name                    string      `json:"name"`
				Birthday                interface{} `json:"birthday"`
				Documentation           interface{} `json:"documentation"`
				HasDocumentation        bool        `json:"has_documentation"`
				CreatedAt               string      `json:"created_at"`
				Auths                   []string    `json:"auths"`
				SelfieEnabled           bool        `json:"selfie_enabled"`
				HandwrittenEnabled      bool        `json:"handwritten_enabled"`
				OfficialDocumentEnabled bool        `json:"official_document_enabled"`
				LivenessEnabled         bool        `json:"liveness_enabled"`
				FacialBiometricsEnabled bool        `json:"facial_biometrics_enabled"`
				CommunicateBy           string      `json:"communicate_by"`
				Address                 string      `json:"address"`
			} `json:"signer"`
			Account struct {
				Key string `json:"key"`
			} `json:"account"`
		} `json:"data"`
		Created string `json:"created"`
	} `json:"attributes"`
}

type EnvelopeUpdate struct {
	Envelope       *EnvelopeData
	Name           *string
	RemindInterval *uint
	DeadlineAt     *time.Time
}
