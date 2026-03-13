package clicksignhandler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type ClicksignHandler struct {
	url         *string
	accesstoken *string
}

type EnvelopeGetFilters struct {
	OnlyRunning bool
	DeadlineAt  *EnvelopeFilterDeadlineAt
}

type EnvelopeFilterDeadlineAt struct {
	Begin time.Time
	End   *time.Time
}

func (c *ClicksignHandler) EnvelopeCreate(param *EnvelopeCreate) (*ResultData[EnvelopeData], error) {
	if len(param.Name) == 0 {
		return nil, fmt.Errorf("envelope name not found")
	}

	url := fmt.Sprintf("%s/envelopes?access_token=%s", *c.url, *c.accesstoken)

	var remindInterval uint = 1
	if param.RemindInterval != nil {
		remindInterval = *param.RemindInterval
	}

	deadlineAt := time.Now().AddDate(0, 0, 30).Format(time.RFC3339)
	if param.DeadlineAt != nil {
		deadlineAt = param.DeadlineAt.Format(time.RFC3339)
	}

	payload := map[string]interface{}{
		"data": map[string]interface{}{
			"type": "envelopes",
			"attributes": map[string]interface{}{
				"name":            param.Name,
				"remind_interval": remindInterval,
				"deadline_at":     deadlineAt,
			},
		},
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
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

	if res.StatusCode >= 300 {
		b, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("error create envelope. details: StatusCode= %d | Message= %s", res.StatusCode, string(b))
	}

	var result ResultData[EnvelopeData]
	json.NewDecoder(res.Body).Decode(&result)

	return &result, nil
}

func (c *ClicksignHandler) EnvelopeGetById(envelopeId string) (*ResultData[EnvelopeData], error) {
	url := fmt.Sprintf("%s/envelopes/%s?access_token=%s", *c.url, envelopeId, *c.accesstoken)

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

	if res.StatusCode >= 300 {
		b, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("error find envelope. details: StatusCode= %d | Message= %s", res.StatusCode, string(b))
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response ResultData[EnvelopeData]
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *ClicksignHandler) EnvelopesGetFirstPage(param EnvelopeGetFilters) (*ResultList[EnvelopeData], error) {
	url := fmt.Sprintf("%s/envelopes?access_token=%s", *c.url, *c.accesstoken)

	if param.OnlyRunning {
		url += "&filter[status]=running"
	}

	if param.DeadlineAt != nil {
		yB, mB, dB := param.DeadlineAt.Begin.Date()
		beginUTC := time.Date(yB, mB, dB, 0, 0, 0, 0, time.UTC)
		endUTC := time.Date(yB, mB, dB, 23, 59, 59, 999, time.UTC)

		if param.DeadlineAt.End != nil {
			yE, mE, dE := param.DeadlineAt.End.Date()
			endUTC = time.Date(yE, mE, dE, 23, 59, 59, 999, time.UTC)
		}

		url += fmt.Sprintf("&filter[deadline_at]=%s,%s", beginUTC.Format(time.RFC3339), endUTC.Format(time.RFC3339))
	}

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

	if res.StatusCode >= 300 {
		b, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("error find envelope first page. details: StatusCode= %d | Message= %s", res.StatusCode, string(b))
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response ResultList[EnvelopeData]
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *ClicksignHandler) EnvelopesGetNextPage(urlLinkMoreData *string) (*ResultList[EnvelopeData], error) {
	if urlLinkMoreData == nil {
		return nil, fmt.Errorf("url not sent for more items.")
	} else if len(*urlLinkMoreData) == 0 {
		return nil, fmt.Errorf("there is no url for more items.")
	}

	url := fmt.Sprintf("%s&access_token=%s", *urlLinkMoreData, *c.accesstoken)

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

	if res.StatusCode >= 300 {
		b, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("error find envelope next page. details: URLNextPage=%s | StatusCode= %d | Message= %s", *urlLinkMoreData, res.StatusCode, string(b))
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response ResultList[EnvelopeData]
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *ClicksignHandler) EnvelopeGetSigners(envelopeId string) (*ResultList[SignerData], error) {
	url := fmt.Sprintf("%s/envelopes/%s/signers?access_token=%s", *c.url, envelopeId, *c.accesstoken)

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

	if res.StatusCode >= 300 {
		b, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("error find envelope events. details: StatusCode= %d | Message= %s", res.StatusCode, string(b))
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response ResultList[SignerData]
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *ClicksignHandler) EnvelopeGetEvents(envelopeId string) (*ResultList[EnvelopeEvent], error) {
	url := fmt.Sprintf("%s/envelopes/%s/events?access_token=%s", *c.url, envelopeId, *c.accesstoken)

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

	if res.StatusCode >= 300 {
		b, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("error find envelope events. details: StatusCode= %d | Message= %s", res.StatusCode, string(b))
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response ResultList[EnvelopeEvent]
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *ClicksignHandler) EnvelopeActive(envelope *EnvelopeData) (*ResultData[EnvelopeData], error) {
	url := fmt.Sprintf("%s/envelopes/%s?access_token=%s", *c.url, envelope.ID, *c.accesstoken)

	payload := map[string]interface{}{
		"data": map[string]interface{}{
			"id":   envelope.ID,
			"type": "envelopes",
			"attributes": map[string]interface{}{
				"status": "running",
			},
		},
	}

	payloadJson, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(payloadJson))
	if err != nil {
		if errDel := c.EnvelopeDelete(envelope); errDel != nil {
			return nil, fmt.Errorf("error activating envelope. it's in drafts. details: %s", errDel)
		}

		return nil, fmt.Errorf("error activating envelope. details: %s", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/vnd.api+json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		if errDel := c.EnvelopeDelete(envelope); errDel != nil {
			return nil, fmt.Errorf("error activating envelope. it's in drafts. details: %s", errDel)
		}

		return nil, fmt.Errorf("error activating envelope. details: %s", err)
	}
	defer res.Body.Close()

	if res.StatusCode >= 300 {
		errDel := c.EnvelopeDelete(envelope)

		b, err := io.ReadAll(res.Body)
		if err != nil {
			if errDel != nil {
				return nil, fmt.Errorf("error activating envelope: %s | it's in drafts. details: %s", err, errDel)
			}

			return nil, fmt.Errorf("error activating envelope: %s", err.Error())
		}

		if errDel != nil {
			return nil, fmt.Errorf("error activating envelope: %s | it's in drafts. details: %s", string(b), errDel)
		}

		return nil, fmt.Errorf("error activating envelope: %s", string(b))
	}

	var result ResultData[EnvelopeData]
	json.NewDecoder(res.Body).Decode(&result)

	return &result, nil
}

func (c *ClicksignHandler) EnvelopesUpdate(param EnvelopeUpdate) (*ResultData[EnvelopeData], error) {
	if param.Envelope == nil {
		return nil, fmt.Errorf("")
	}

	url := fmt.Sprintf("%s/envelopes/%s?access_token=%s", *c.url, param.Envelope.ID, *c.accesstoken)

	name := param.Envelope.Attributes.Name
	if param.Name != nil {
		name = *param.Name
	}

	remindInterval := param.Envelope.Attributes.RemindInterval
	if param.RemindInterval != nil {
		remindInterval = CSIntString(*param.RemindInterval)
	}

	deadlineAt := param.Envelope.Attributes.DeadlineAt
	if param.DeadlineAt != nil {
		deadlineAt = param.DeadlineAt.Format(time.RFC3339)
	}

	payload := map[string]interface{}{
		"data": map[string]interface{}{
			"id":   param.Envelope.ID,
			"type": "envelopes",
			"attributes": map[string]interface{}{
				"name":            name,
				"deadline_at":     deadlineAt,
				"remind_interval": remindInterval,
			},
		},
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/vnd.api+json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode >= 300 {
		b, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("error adding updated envelope: %s", string(b))
	}

	var result ResultData[EnvelopeData]
	json.NewDecoder(res.Body).Decode(&result)

	return &result, nil
}

func (c *ClicksignHandler) EnvelopeDelete(envelope *EnvelopeData) error {
	if envelope == nil {
		return fmt.Errorf("envelope not sent")
	} else if envelope.Attributes.Status != "draft" {
		return fmt.Errorf("this envelope cannot be deleted. status: draft")
	}

	url := fmt.Sprintf("%s/envelopes/%s?access_token=%s", *c.url, envelope.ID, *c.accesstoken)

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Set("Content-Type", "application/vnd.api+json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode >= 300 {
		b, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("error delete envelope. details: StatusCode= %d | Message= %s", res.StatusCode, string(b))
	}

	defer res.Body.Close()

	return nil
}

func (c *ClicksignHandler) DocumentCreate(param *DocumentCreate) (*ResultData[DocumentData], error) {
	url := fmt.Sprintf("%s/envelopes/%s/documents?access_token=%s",
		*c.url, param.Envelope.ID, *c.accesstoken)

	body := map[string]interface{}{
		"data": map[string]interface{}{
			"type": "documents",
			"attributes": map[string]string{
				"filename":       param.FileName,
				"content_base64": fmt.Sprintf("data:application/%s;base64,%s", string(param.FileType), param.FileBase64),
			},
		},
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("error when assembling JSON: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(string(bodyBytes)))
	if err != nil {
		return nil, err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Content-type", "application/vnd.api+json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode >= 300 {
		b, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("error upload document. details: StatusCode= %d | Message= %s", res.StatusCode, string(b))
	}

	var result ResultData[DocumentData]
	json.NewDecoder(res.Body).Decode(&result)

	return &result, nil
}

func (c *ClicksignHandler) DocumentGetById(envelopeId, documentId string) (*ResultData[DocumentData], error) {
	url := fmt.Sprintf("%s/envelopes/%s/documents/%s?access_token=%s",
		*c.url, envelopeId, documentId, *c.accesstoken)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Content-type", "application/vnd.api+json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode >= 300 {
		b, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("error upload document. details: StatusCode= %d | Message= %s", res.StatusCode, string(b))
	}

	var result ResultData[DocumentData]
	json.NewDecoder(res.Body).Decode(&result)

	return &result, nil
}

func (c *ClicksignHandler) DocumentsGetFirstPage(envelopeId string) (*ResultList[DocumentData], error) {
	url := fmt.Sprintf("%s/envelopes/%s/documents?access_token=%s", *c.url, envelopeId, *c.accesstoken)

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

	if res.StatusCode >= 300 {
		b, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("error find document first page. details: StatusCode= %d | Message= %s", res.StatusCode, string(b))
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var result ResultList[DocumentData]
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println(string(body))
		return nil, err
	}

	return &result, nil
}

func (c *ClicksignHandler) DocumentsGetNextPage(urlLinkMoreData *string) (*ResultList[DocumentData], error) {
	if urlLinkMoreData == nil {
		return nil, fmt.Errorf("url not sent for more items.")
	} else if len(*urlLinkMoreData) == 0 {
		return nil, fmt.Errorf("there is no url for more items.")
	}

	url := fmt.Sprintf("%s&access_token=%s", *urlLinkMoreData, *c.accesstoken)

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

	if res.StatusCode >= 300 {
		b, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("error find documents next page. details: URLNextPage=%s | StatusCode= %d | Message= %s", *urlLinkMoreData, res.StatusCode, string(b))
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var result ResultList[DocumentData]
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *ClicksignHandler) DocumentGetEvents_AddSigner(envelopeId, documentId string) (*ResultList[DocumentEvent_AddSigner], error) {
	url := fmt.Sprintf("%s/envelopes/%s/documents/%s/events?access_token=%s",
		*c.url, envelopeId, documentId, *c.accesstoken)

	return _DocumentGetEvents[DocumentEvent_AddSigner](url)
}

func (c *ClicksignHandler) DocumentGetEvents_Sign(envelopeId, documentId string) (*ResultList[DocumentEvent_Sign], error) {
	url := fmt.Sprintf("%s/envelopes/%s/documents/%s/events?access_token=%s",
		*c.url, envelopeId, documentId, *c.accesstoken)

	return _DocumentGetEvents[DocumentEvent_Sign](url)
}

func (c *ClicksignHandler) DocumentCancel(envelope *EnvelopeData, documentId *string) (*ResultData[DocumentData], error) {
	if envelope == nil {
		return nil, fmt.Errorf("envelope not sent")
	} else if envelope.Attributes.Status != "running" {
		return nil, fmt.Errorf("it is only possible to change the status of running envelopes")
	} else if documentId == nil {
		return nil, fmt.Errorf("document id not sent")
	} else if len(*documentId) == 0 {
		return nil, fmt.Errorf("invalid document id")
	}

	url := fmt.Sprintf("%s/envelopes/%s/documents/%s?access_token=%s",
		*c.url, envelope.ID, *documentId, *c.accesstoken)

	payload, err := json.Marshal(map[string]interface{}{
		"data": map[string]interface{}{
			"type": "documents",
			"id":   *documentId,
			"attributes": map[string]interface{}{
				"status": "canceled",
			},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("error when assembling JSON: %w", err)
	}

	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Content-type", "application/vnd.api+json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode >= 300 {
		b, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("error cancel document. details: StatusCode= %d | Message= %s", res.StatusCode, string(b))
	}

	var result ResultData[DocumentData]
	json.NewDecoder(res.Body).Decode(&result)

	return &result, nil
}

func (c *ClicksignHandler) SignerCreate(param *SignerCreate) (*ResultData[SignerData], error) {
	if err := param.checkdata(); err != nil {
		return nil, fmt.Errorf("001: %s", err.Error())
	}

	url := fmt.Sprintf("%s/envelopes/%s/signers?access_token=%s",
		*c.url, param.Envelope.ID, *c.accesstoken)

	payload, err := param.Signer.createPayload()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Content-type", "application/vnd.api+json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode >= 300 {
		b, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("error create signer in document. details: StatusCode= %d | Message= %s", res.StatusCode, string(b))
	}

	var result ResultData[SignerData]
	json.NewDecoder(res.Body).Decode(&result)

	switch strings.ToLower(param.Envelope.Attributes.Status) {
	case "draft":
		if err := c._SignerAddRequerimentAutentication(_SignerRequerimentAutenticationCreate{
			EnvelopeId:         param.Envelope.ID,
			DocumentId:         param.Document.ID,
			SignerId:           result.Data.ID,
			AuthenticationType: param.Signer.AutomaticSignature,
		}); err != nil {
			return nil, err
		}

		if err := c._SignerAddRequerimentQualification(_SignerRequerimentQualificationCreate{
			EnvelopeId: param.Envelope.ID,
			DocumentId: param.Document.ID,
			SignerId:   result.Data.ID,
			SignerType: param.Signer.Type,
		}); err != nil {
			return nil, err
		}
	case "running":
		if err := c._SignerBulkRequirementsAutenticationAndQualification(_SignerBulkRequirementsCreate{
			EnvelopeId:         param.Envelope.ID,
			DocumentId:         param.Document.ID,
			SignerId:           result.Data.ID,
			SignerType:         param.Signer.Type,
			AuthenticationType: param.Signer.AutomaticSignature,
		}); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("envelope status is invalid.")
	}

	return &result, nil
}

func (c *ClicksignHandler) SignerDelete(envelopeId, signerId string) error {
	url := fmt.Sprintf("%s/envelopes/%s/signers/%s?access_token=%s", *c.url, envelopeId, signerId, *c.accesstoken)

	req, err := http.NewRequest(http.MethodDelete, url, nil)

	if err != nil {
		return err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Set("Content-Type", "application/vnd.api+json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	if res.StatusCode != 204 {
		return fmt.Errorf("%s", string(body))
	}

	return nil
}

func (c *ClicksignHandler) ObserverCreate(param *ObserverCreate) (*ResultData[ObserverData], error) {
	url := fmt.Sprintf("%s/envelopes/%s/signature_watchers?access_token=%s", *c.url, param.Envelope.ID, *c.accesstoken)

	payload := map[string]interface{}{
		"data": map[string]interface{}{
			"type": "signature_watchers",
			"attributes": map[string]interface{}{
				"email":                    param.Email,
				"kind":                     "on_finished",
				"attach_documents_enabled": true,
			},
		},
	}

	payloadJson, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payloadJson))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/vnd.api+json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode >= 300 {
		b, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("error add observer: %s", err.Error())
		}

		return nil, fmt.Errorf("error add observer: %s", string(b))
	}

	var result ResultData[ObserverData]
	json.NewDecoder(res.Body).Decode(&result)

	return &result, nil
}
