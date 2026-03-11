package clicksignhandler_test

import (
	"os"
	"testing"

	"github.com/fenix-ds/clicksignhandler"
	"github.com/joho/godotenv"
)

const (
	filePDFBase64 = "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aXR5LUgKL0Rlc2NlbmRhbnRGb250cyBbMTggMCBSXQovVG9Vbmljb2RlIDE5IDAgUj4+CmVuZG9iagp4cmVmCjAgMjAKMDAwMDAwMDAwMCA2NTUzNSBmIAowMDAwMDAwMDE1IDAwMDAwIG4gCjAwMDAwMDA2MTQgMDAwMDAgbiAKMDAwMDAwMDE3NSAwMDAwMCBuIAowMDAwMDA4OTEzIDAwMDAwIG4gCjAwMDAwMDAyMTIgMDAwMDAgbiAKMDAwMDAwMDI5NyAwMDAwMCBuIAowMDAwMDAwODQxIDAwMDAwIG4gCjAwMDAwMDEyNzkgMDAwMDAgbiAKMDAwMDAwMTEwMCAwMDAwMCBuIAowMDAwMDAwODk2IDAwMDAwIG4gCjAwMDAwMDA5NjQgMDAwMDAgbiAKMDAwMDAwMTAzMiAwMDAwMCBuIAowMDAwMDAxMTg1IDAwMDAwIG4gCjAwMDAwMDEyMjQgMDAwMDAgbiAKMDAwMDAwMTM2OSAwMDAwMCBuIAowMDAwMDAxNTY2IDAwMDAwIG4gCjAwMDAwMDgwNzEgMDAwMDAgbiAKMDAwMDAwODMyMyAwMDAwMCBuIAowMDAwMDA4NTg3IDAwMDAwIG4gCnRyYWlsZXIKPDwvU2l6ZSAyMAovUm9vdCAxNSAwIFIKL0luZm8gMSAwIFI+PgpzdGFydHhyZWYKOTA1OQolJUVPRgo="
)

var (
	clicksignHandler *clicksignhandler.ClicksignHandler
)

func TestNewClicksignHandler(t *testing.T) {
	err := loadfileEnv(t)
	if err != nil {
		return
	}

	if clicksignHandler, err = clicksignhandler.NewClicksignHandler(clicksignhandler.ClicksignParam{
		Environment: clicksignhandler.EnvSandbox,
		Key:         os.Getenv("ACCESS_TOKEN"),
	}); err != nil {
		t.Error(err)
	}
}

func Test_EnvelopeCreate(t *testing.T) {
	err := loadfileEnv(t)
	if err != nil {
		return
	}

	err = attributeVarClicksignHandler()
	if err != nil {
		return
	}

	if envelope, err := clicksignHandler.EnvelopeCreate(&clicksignhandler.EnvelopeCreate{
		Name: "Test",
	}); err != nil {
		t.Error(err)
		return
	} else {
		t.Log(envelope.Data.ID)
	}
}

func Test_EnvelopesGetFirstPage(t *testing.T) {
	err := loadfileEnv(t)
	if err != nil {
		return
	}

	err = attributeVarClicksignHandler()
	if err != nil {
		return
	}

	if envelope, err := clicksignHandler.EnvelopesGetFirstPage(clicksignhandler.EnvelopeGetFilters{}); err != nil {
		t.Error(err)
		return
	} else {
		t.Log(envelope.Meta.RecordCount)
	}
}

func Test_EnvelopesGetNextPage(t *testing.T) {
	err := loadfileEnv(t)
	if err != nil {
		return
	}

	err = attributeVarClicksignHandler()
	if err != nil {
		return
	}

	if envelope, err := clicksignHandler.EnvelopesGetFirstPage(clicksignhandler.EnvelopeGetFilters{}); err != nil {
		t.Error(err)
		return
	} else if envelopeNext, err := clicksignHandler.EnvelopesGetNextPage(&envelope.Links.Next); err != nil {
		t.Error(err)
		return
	} else {
		t.Log(envelopeNext.Meta.RecordCount)
	}
}

func Test_EnvelopeGetById(t *testing.T) {
	err := loadfileEnv(t)
	if err != nil {
		return
	}

	err = attributeVarClicksignHandler()
	if err != nil {
		return
	}

	if envelopes, err := clicksignHandler.EnvelopesGetFirstPage(clicksignhandler.EnvelopeGetFilters{}); err != nil {
		t.Error(err)
		return
	} else if envelopes.Meta.RecordCount == 0 {
		t.Log("Envelopes not found")
	} else if envelope, err := clicksignHandler.EnvelopeGetById(envelopes.Data[0].ID); err != nil {
		t.Error(err)
		return
	} else {
		t.Log(envelope)
	}
}

func Test_DocumentCreate(t *testing.T) {
	err := loadfileEnv(t)
	if err != nil {
		return
	}

	err = attributeVarClicksignHandler()
	if err != nil {
		return
	}

	if envelope, err := clicksignHandler.EnvelopeCreate(&clicksignhandler.EnvelopeCreate{
		Name: "Test",
	}); err != nil {
		t.Error(err)
		return
	} else if document, err := clicksignHandler.DocumentCreate(&clicksignhandler.DocumentCreate{
		Envelope:   &envelope.Data,
		FileType:   clicksignhandler.DFT_PDF,
		FileName:   "document.pdf",
		FileBase64: filePDFBase64,
	}); err != nil {
		t.Error(err)
		return
	} else {
		t.Log(document.Data.ID)
	}
}

func Test_DocumentGetById(t *testing.T) {
	err := loadfileEnv(t)
	if err != nil {
		return
	}

	err = attributeVarClicksignHandler()
	if err != nil {
		return
	}

	if envelope, err := clicksignHandler.EnvelopeCreate(&clicksignhandler.EnvelopeCreate{
		Name: "Test",
	}); err != nil {
		t.Error(err)
		return
	} else if document, err := clicksignHandler.DocumentCreate(&clicksignhandler.DocumentCreate{
		Envelope:   &envelope.Data,
		FileType:   clicksignhandler.DFT_PDF,
		FileName:   "document.pdf",
		FileBase64: filePDFBase64,
	}); err != nil {
		t.Error(err)
		return
	} else if result, err := clicksignHandler.DocumentGetById(envelope.Data.ID, document.Data.ID); err != nil {
		t.Error(err)
		return
	} else {
		t.Log(result.Data.Attributes)
	}
}

func Test_DocumentsGetFirstPage(t *testing.T) {
	err := loadfileEnv(t)
	if err != nil {
		return
	}

	err = attributeVarClicksignHandler()
	if err != nil {
		return
	}

	if envelope, err := clicksignHandler.EnvelopeCreate(&clicksignhandler.EnvelopeCreate{
		Name: "Test",
	}); err != nil {
		t.Error(err)
		return
	} else if _, err := clicksignHandler.DocumentCreate(&clicksignhandler.DocumentCreate{
		Envelope:   &envelope.Data,
		FileType:   clicksignhandler.DFT_PDF,
		FileName:   "document.pdf",
		FileBase64: filePDFBase64,
	}); err != nil {
		t.Error(err)
		return
	} else if result, err := clicksignHandler.DocumentsGetFirstPage(envelope.Data.ID); err != nil {
		t.Error(err)
		return
	} else {
		t.Log(result.Meta.RecordCount)
	}
}

func Test_DocumentsGetNextPage(t *testing.T) {
	err := loadfileEnv(t)
	if err != nil {
		return
	}

	err = attributeVarClicksignHandler()
	if err != nil {
		return
	}

	if envelopes, err := clicksignHandler.EnvelopesGetFirstPage(clicksignhandler.EnvelopeGetFilters{
		OnlyRunning: false,
		DeadlineAt:  nil,
	}); err != nil {
		t.Error(err)
		return
	} else if envelopes.Meta.RecordCount > 0 {
		if documentsFirstPage, err := clicksignHandler.DocumentsGetFirstPage(envelopes.Data[0].ID); err != nil {
			t.Error(err)
			return
		} else {
			t.Log(documentsFirstPage.Links)

			if result, err := clicksignHandler.DocumentsGetNextPage(&documentsFirstPage.Links.Next); err != nil {
				t.Error(err)
				return
			} else {
				t.Log(result)
			}
		}
	} else {
		if envelope, err := clicksignHandler.EnvelopeCreate(&clicksignhandler.EnvelopeCreate{
			Name: "Test",
		}); err != nil {
			t.Error(err)
			return
		} else if _, err := clicksignHandler.DocumentCreate(&clicksignhandler.DocumentCreate{
			Envelope:   &envelope.Data,
			FileType:   clicksignhandler.DFT_PDF,
			FileName:   "document.pdf",
			FileBase64: filePDFBase64,
		}); err != nil {
			t.Error(err)
			return
		} else if documentsFirstPage, err := clicksignHandler.DocumentsGetFirstPage(envelope.Data.ID); err != nil {
			t.Error(err)
			return
		} else {
			t.Log(documentsFirstPage.Links)

			if result, err := clicksignHandler.DocumentsGetNextPage(&documentsFirstPage.Links.Next); err != nil {
				t.Error(err)
				return
			} else {
				t.Log(result)
			}
		}
	}
}

func Test_SignerCreate(t *testing.T) {
	err := loadfileEnv(t)
	if err != nil {
		return
	}

	err = attributeVarClicksignHandler()
	if err != nil {
		return
	}

	if envelope, err := clicksignHandler.EnvelopeCreate(&clicksignhandler.EnvelopeCreate{
		Name: "Test",
	}); err != nil {
		t.Error(err)
		return
	} else if document, err := clicksignHandler.DocumentCreate(&clicksignhandler.DocumentCreate{
		Envelope:   &envelope.Data,
		FileType:   clicksignhandler.DFT_PDF,
		FileName:   "document.pdf",
		FileBase64: filePDFBase64,
	}); err != nil {
		t.Error(err)
		return
	} else if _, err := clicksignHandler.SignerCreate(&clicksignhandler.SignerCreate{
		Envelope: &envelope.Data,
		Document: &document.Data,
		Signer: &clicksignhandler.SignerPayload{
			Type:               clicksignhandler.SRT_WITNESS,
			AutomaticSignature: clicksignhandler.AUT_EMAIL,
			Name:               "Test Test Test",
			Email:              "test@test.com",
			HasDocumentation:   nil,
			CommunicateEvents: clicksignhandler.SignerCommunicateEvents{
				SignatureRequest:        clicksignhandler.SREQ_NONE,
				SignatureReminder:       clicksignhandler.SREM_NONE,
				SignatureDocumentSigned: clicksignhandler.SDS_EMAIL,
			},
		},
	}); err != nil {
		t.Error(err)
		return
	}

}

func Test_ObserverCreate(t *testing.T) {
	err := loadfileEnv(t)
	if err != nil {
		return
	}

	err = attributeVarClicksignHandler()
	if err != nil {
		return
	}

	if envelope, err := clicksignHandler.EnvelopeCreate(&clicksignhandler.EnvelopeCreate{
		Name: "Test",
	}); err != nil {
		t.Error(err)
		return
	} else if document, err := clicksignHandler.DocumentCreate(&clicksignhandler.DocumentCreate{
		Envelope:   &envelope.Data,
		FileType:   clicksignhandler.DFT_PDF,
		FileName:   "document.pdf",
		FileBase64: filePDFBase64,
	}); err != nil {
		t.Error(err)
		return
	} else if _, err := clicksignHandler.SignerCreate(&clicksignhandler.SignerCreate{
		Envelope: &envelope.Data,
		Document: &document.Data,
		Signer: &clicksignhandler.SignerPayload{
			Type:               clicksignhandler.SRT_WITNESS,
			AutomaticSignature: clicksignhandler.AUT_EMAIL,
			Name:               "Test Test Test",
			Email:              "test@test.com",
			HasDocumentation:   nil,
			CommunicateEvents: clicksignhandler.SignerCommunicateEvents{
				SignatureRequest:        clicksignhandler.SREQ_NONE,
				SignatureReminder:       clicksignhandler.SREM_NONE,
				SignatureDocumentSigned: clicksignhandler.SDS_EMAIL,
			},
		},
	}); err != nil {
		t.Error(err)
		return
	} else if _, err := clicksignHandler.ObserverCreate(&clicksignhandler.ObserverCreate{
		Envelope: &envelope.Data,
		Name:     "Observer Test",
		Email:    "observer@test.com",
	}); err != nil {
		t.Error(err)
		return
	}

}

func Test_EnvelopeDelete(t *testing.T) {
	err := loadfileEnv(t)
	if err != nil {
		return
	}

	err = attributeVarClicksignHandler()
	if err != nil {
		return
	}

	if envelope, err := clicksignHandler.EnvelopeCreate(&clicksignhandler.EnvelopeCreate{
		Name: "Test",
	}); err != nil {
		t.Error(err)
		return
	} else if document, err := clicksignHandler.DocumentCreate(&clicksignhandler.DocumentCreate{
		Envelope:   &envelope.Data,
		FileType:   clicksignhandler.DFT_PDF,
		FileName:   "document.pdf",
		FileBase64: filePDFBase64,
	}); err != nil {
		t.Error(err)
		return
	} else if _, err := clicksignHandler.SignerCreate(&clicksignhandler.SignerCreate{
		Envelope: &envelope.Data,
		Document: &document.Data,
		Signer: &clicksignhandler.SignerPayload{
			Type:               clicksignhandler.SRT_WITNESS,
			AutomaticSignature: clicksignhandler.AUT_EMAIL,
			Name:               "Test Test Test",
			Email:              "test@test.com",
			HasDocumentation:   nil,
			CommunicateEvents: clicksignhandler.SignerCommunicateEvents{
				SignatureRequest:        clicksignhandler.SREQ_NONE,
				SignatureReminder:       clicksignhandler.SREM_NONE,
				SignatureDocumentSigned: clicksignhandler.SDS_EMAIL,
			},
		},
	}); err != nil {
		t.Error(err)
		return
	} else if err := clicksignHandler.EnvelopeDelete(&envelope.Data); err != nil {
		t.Error(err)
		return
	}

}

func Test_EnvelopeActive(t *testing.T) {
	err := loadfileEnv(t)
	if err != nil {
		return
	}

	err = attributeVarClicksignHandler()
	if err != nil {
		return
	}

	if envelope, err := clicksignHandler.EnvelopeCreate(&clicksignhandler.EnvelopeCreate{
		Name: "Test",
	}); err != nil {
		t.Error(err)
		return
	} else if document, err := clicksignHandler.DocumentCreate(&clicksignhandler.DocumentCreate{
		Envelope:   &envelope.Data,
		FileType:   clicksignhandler.DFT_PDF,
		FileName:   "document.pdf",
		FileBase64: filePDFBase64,
	}); err != nil {
		t.Error(err)
		return
	} else if _, err := clicksignHandler.SignerCreate(&clicksignhandler.SignerCreate{
		Envelope: &envelope.Data,
		Document: &document.Data,
		Signer: &clicksignhandler.SignerPayload{
			Type:               clicksignhandler.SRT_WITNESS,
			AutomaticSignature: clicksignhandler.AUT_EMAIL,
			Name:               "Test Test Test",
			Email:              "test@test.com",
			HasDocumentation:   nil,
			CommunicateEvents: clicksignhandler.SignerCommunicateEvents{
				SignatureRequest:        clicksignhandler.SREQ_NONE,
				SignatureReminder:       clicksignhandler.SREM_NONE,
				SignatureDocumentSigned: clicksignhandler.SDS_EMAIL,
			},
		},
	}); err != nil {
		t.Error(err)
		return
	} else if _, err := clicksignHandler.EnvelopeActive(&envelope.Data); err != nil {
		t.Error(err)
		return
	}

}

func Test_DocumentCancel(t *testing.T) {
	err := loadfileEnv(t)
	if err != nil {
		return
	}

	err = attributeVarClicksignHandler()
	if err != nil {
		return
	}

	if envelope, err := clicksignHandler.EnvelopeCreate(&clicksignhandler.EnvelopeCreate{
		Name: "Test",
	}); err != nil {
		t.Error(err)
		return
	} else if document, err := clicksignHandler.DocumentCreate(&clicksignhandler.DocumentCreate{
		Envelope:   &envelope.Data,
		FileType:   clicksignhandler.DFT_PDF,
		FileName:   "document.pdf",
		FileBase64: filePDFBase64,
	}); err != nil {
		t.Error(err)
		return
	} else if _, err := clicksignHandler.SignerCreate(&clicksignhandler.SignerCreate{
		Envelope: &envelope.Data,
		Document: &document.Data,
		Signer: &clicksignhandler.SignerPayload{
			Type:               clicksignhandler.SRT_WITNESS,
			AutomaticSignature: clicksignhandler.AUT_EMAIL,
			Name:               "Test Test Test",
			Email:              "test@test.com",
			HasDocumentation:   nil,
			CommunicateEvents: clicksignhandler.SignerCommunicateEvents{
				SignatureRequest:        clicksignhandler.SREQ_NONE,
				SignatureReminder:       clicksignhandler.SREM_NONE,
				SignatureDocumentSigned: clicksignhandler.SDS_EMAIL,
			},
		},
	}); err != nil {
		t.Error(err)
		return
	} else if envelope, err := clicksignHandler.EnvelopeActive(&envelope.Data); err != nil {
		t.Error(err)
		return
	} else {
		if _, err := clicksignHandler.DocumentCancel(&envelope.Data, &document.Data.ID); err != nil {
			t.Error(err)
			return
		}
	}
}

func loadfileEnv(t *testing.T) error {
	if envs, err := godotenv.Read(".env"); err != nil {
		t.Error(err)
		return err
	} else {
		for key, value := range envs {
			t.Setenv(key, value)
		}
	}

	return nil
}

func attributeVarClicksignHandler() error {
	var err error
	if clicksignHandler, err = clicksignhandler.NewClicksignHandler(clicksignhandler.ClicksignParam{
		Environment: clicksignhandler.EnvSandbox,
		Key:         os.Getenv("ACCESS_TOKEN"),
	}); err != nil {
		return err
	}

	return nil
}
