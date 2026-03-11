package clicksignhandler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *ClicksignHandler) _SignerAddRequerimentAutentication(param _SignerRequerimentAutenticationCreate) error {
	url := fmt.Sprintf("%s/envelopes/%s/requirements?access_token=%s", *c.url, param.EnvelopeId, *c.accesstoken)

	payload := map[string]interface{}{
		"data": map[string]interface{}{
			"type": "requirements",
			"attributes": map[string]interface{}{
				"action": "provide_evidence",
				"auth":   string(param.AuthenticationType),
			},
			"relationships": map[string]interface{}{
				"document": map[string]interface{}{
					"data": map[string]interface{}{
						"id":   param.DocumentId,
						"type": "documents",
					},
				},
				"signer": map[string]interface{}{
					"data": map[string]interface{}{
						"id":   param.SignerId,
						"type": "signers",
					},
				},
			},
		},
	}

	payloadJson, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payloadJson))
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/vnd.api+json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode >= 300 {
		b, err := io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("error adding authentication requirement: %s", err.Error())
		}

		return fmt.Errorf("error adding authentication requirement: %s", string(b))
	}

	return nil
}

func (c *ClicksignHandler) _SignerAddRequerimentQualification(param _SignerRequerimentQualificationCreate) error {
	url := fmt.Sprintf("%s/envelopes/%s/requirements?access_token=%s", *c.url, param.EnvelopeId, *c.accesstoken)

	payload := map[string]interface{}{
		"data": map[string]interface{}{
			"type": "requirements",
			"attributes": map[string]interface{}{
				"action": "agree",
				"role":   param.SignerType,
			},
			"relationships": map[string]interface{}{
				"document": map[string]interface{}{
					"data": map[string]interface{}{
						"id":   param.DocumentId,
						"type": "documents",
					},
				},
				"signer": map[string]interface{}{
					"data": map[string]interface{}{
						"id":   param.SignerId,
						"type": "signers",
					},
				},
			},
		},
	}

	payloadJson, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payloadJson))
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/vnd.api+json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode >= 300 {
		b, err := io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("error adding authentication qualification: %s", err.Error())
		}

		return fmt.Errorf("error adding authentication qualification: %s", string(b))
	}

	return nil
}

func _DocumentGetEvents[T any](url string) (*ResultList[T], error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Set("Content-Type", "application/vnd.api+json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var response ResultList[T]
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, err
}
