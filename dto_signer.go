package clicksignhandler

import (
	"encoding/json"
	"fmt"
	"time"
)

type SignerData struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Attributes struct {
		Name                    string `json:"name"`
		Birthday                any    `json:"birthday"`
		Email                   string `json:"email"`
		PhoneNumber             any    `json:"phone_number"`
		LocationRequiredEnabled bool   `json:"location_required_enabled"`
		HasDocumentation        bool   `json:"has_documentation"`
		Documentation           any    `json:"documentation"`
		Refusable               bool   `json:"refusable"`
		Group                   int    `json:"group"`
		CommunicateEvents       struct {
			DocumentSigned    string `json:"document_signed"`
			SignatureRequest  string `json:"signature_request"`
			SignatureReminder string `json:"signature_reminder"`
		} `json:"communicate_events"`
		SignatureHost struct {
		} `json:"signature_host"`
		Created  string `json:"created"`
		Modified string `json:"modified"`
	} `json:"attributes"`
}

type SignerCreate struct {
	Envelope *EnvelopeData
	Document *DocumentData
	Signer   *SignerPayload
}

type SignerPayload struct {
	Type               SignerType
	AutomaticSignature AuthenticationType
	Name               string
	Email              string
	HasDocumentation   *SignerDocumentation
	CommunicateEvents  SignerCommunicateEvents
}

type SignerDocumentation struct {
	Doc       *string
	DateBirth *time.Time
}

type SignerCommunicateEvents struct {
	SignatureRequest        SignatureRequest
	SignatureReminder       SignatureReminder
	SignatureDocumentSigned SignatureDocumentSigned
}

func (p *SignerCreate) checkdata() error {
	if p.Envelope == nil || p.Document == nil {
		return fmt.Errorf("envelope and/or document data not found")
	}

	if err := p.Signer.checkdata(); err != nil {
		return err
	}

	return nil
}

func (p *SignerPayload) checkdata() error {
	if len(string(p.Type)) == 0 || len(string(p.AutomaticSignature)) == 0 || len(p.Name) == 0 || len(p.Email) == 0 {
		return fmt.Errorf("one or more required fields (type, name, email) were not found")
	} else if p.AutomaticSignature == AUT_AUTHSIGNATURE && p.HasDocumentation == nil {
		return fmt.Errorf("automatic signature activated but signer documents not found")
	} else if err := p.CommunicateEvents.checkdata(); err != nil {
		return err
	}

	return nil
}

func (p *SignerPayload) createPayload() ([]byte, error) {
	var payload map[string]interface{}

	if p.HasDocumentation != nil {
		payload = map[string]interface{}{
			"data": map[string]interface{}{
				"type": "signers",
				"attributes": map[string]interface{}{
					"name":              p.Name,
					"email":             p.Email,
					"birthday":          p.HasDocumentation.DateBirth.Format("2006-01-02"),
					"documentation":     p.HasDocumentation.Doc,
					"has_documentation": true,
					"communicate_events": map[string]interface{}{
						"signature_request":  p.CommunicateEvents.SignatureRequest,
						"signature_reminder": p.CommunicateEvents.SignatureReminder,
						"document_signed":    p.CommunicateEvents.SignatureDocumentSigned,
					},
				},
			},
		}
	} else {
		payload = map[string]interface{}{
			"data": map[string]interface{}{
				"type": "signers",
				"attributes": map[string]interface{}{
					"name":              p.Name,
					"email":             p.Email,
					"has_documentation": false,
					"communicate_events": map[string]interface{}{
						"signature_request":  p.CommunicateEvents.SignatureRequest,
						"signature_reminder": p.CommunicateEvents.SignatureReminder,
						"document_signed":    p.CommunicateEvents.SignatureDocumentSigned,
					},
				},
			},
		}
	}

	return json.Marshal(payload)
}

func (p *SignerCommunicateEvents) checkdata() error {
	if len(string(p.SignatureRequest)) == 0 || len(string(p.SignatureReminder)) == 0 || len(string(p.SignatureDocumentSigned)) == 0 {
		return fmt.Errorf("no data found for event communication.")
	}

	return nil
}

type _SignerRequerimentAutenticationCreate struct {
	EnvelopeId         string
	DocumentId         string
	SignerId           string
	AuthenticationType AuthenticationType
}

type _SignerRequerimentQualificationCreate struct {
	EnvelopeId string
	DocumentId string
	SignerId   string
	SignerType SignerType
}

type _SignerBulkRequirementsCreate struct {
	EnvelopeId         string
	DocumentId         string
	SignerId           string
	AuthenticationType AuthenticationType
	SignerType         SignerType
}
