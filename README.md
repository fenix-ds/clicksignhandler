# Clicksign Handler
Golang package for communication with the Clicksign 3.0 API. This repository was created by an independent developer and is not affiliated with Clicksign or its associates.

## Technologies

### Compiler version
<div class="highlight highlight-source-shell">
    <pre>
        Golang (Go) 1.26.0
     </pre>
</div>

### Libraries
<div class="highlight highlight-source-shell">
    <pre>
        github.com/joho/godotenv v1.5.1
    </pre>
</div>

## Installation
Add to the Go project using the command

```bash
    go get github.com/fenix-ds/clicksignhandler
```
### Quick Start
```
    clicksignHandler, err := clicksignhandler.NewClicksignHandler(clicksignhandler.ClicksignParam{
		Environment: clicksignhandler.EnvSandbox,
		Key:         os.Getenv("ACCESS_TOKEN"),
	})
    
    if err != nil {
		panic(err)
	}
```
Obs: Use the value 'EnvSandbox' for testing in the Clicksign sandbox environment. For production and legal validity, use the value 'EnvProd'.

After that, you will have access to all the methods of the Clicksign API. Use the API documentation (developers.clicksign.com/reference/comece-agora) to understand how the functions work.

### Available functions
01. EnvelopeCreate: Function to create new envelopes, initial status 'Draft'.
```
    if envelope, err := clicksignHandler.EnvelopeCreate(&clicksignhandler.EnvelopeCreate{
		Name: "Test",
	}); err != nil {
		panic(err)
	} else {
		fmt.Print(envelope.Data.ID)
	}
```
2.  EnvelopeGetById: Search for envelope by ID
```
    if envelope, err := clicksignHandler.EnvelopeGetById("0123-015656894-2561asd-adsf"); err != nil {
		panic(err)
	} else {
		fmt.Print(envelope)
	}
```
3.  EnvelopesGetFirstPage: Find the first page with the envelopes.
```
    if envelope, err := clicksignHandler.EnvelopesGetFirstPage(clicksignhandler.EnvelopeGetFilters{}); err != nil {
		panic(err)
	} else {
		fmt.Print(envelope)
	}
```
4.  EnvelopesGetNextPage: Find the next page of the envelopes.
```
    if envelope, err := clicksignHandler.EnvelopesGetFirstPage(clicksignhandler.EnvelopeGetFilters{}); err != nil {
		panic(err)
	} else if envelopeNext, err := clicksignHandler.EnvelopesGetNextPage(&envelope.Links.Next); err != nil {
		panic(err)
	} else {
        fmt.Print(envelopeNext)
	}
```
5.  EnvelopeGetSigners: Search for envelope signers
```
```
6.  EnvelopeGetEvents: Fetch envelope events
```
```
7.  EnvelopeActive: Enable envelope
```
    if envelope, err := clicksignHandler.EnvelopeCreate(&clicksignhandler.EnvelopeCreate{
		Name: "Test",
	}); err != nil {
		panic(err)
	} else if document, err := clicksignHandler.DocumentCreate(&clicksignhandler.DocumentCreate{
		Envelope:   &envelope.Data,
		FileType:   clicksignhandler.DFT_PDF,
		FileName:   "document.pdf",
		FileBase64: filePDFBase64,
	}); err != nil {
		panic(err)
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
		panic(err)
	} else if _, err := clicksignHandler.EnvelopeActive(&envelope.Data); err != nil {
		panic(err)
	}
```
8.  EnvelopesUpdate: Update envelope data
```
```
9.  EnvelopeDelete: Delete envelope that has 'Draft' status.
```
```
10. DocumentCreate: Send document (only *.pdf) in base64
```
    if envelope, err := clicksignHandler.EnvelopeCreate(&clicksignhandler.EnvelopeCreate{
		Name: "Test",
	}); err != nil {
		panic(err)
	} else if document, err := clicksignHandler.DocumentCreate(&clicksignhandler.DocumentCreate{
		Envelope:   &envelope.Data,
		FileType:   clicksignhandler.DFT_PDF,
		FileName:   "document.pdf",
		FileBase64: filePDFBase64,
	}); err != nil {
		panic(err)
	} else {
		fmt.Print(document.Data.ID)
	}
```
11. DocumentGetById: Search for a document in an envelope by ID.
```
    if envelope, err := clicksignHandler.EnvelopeCreate(&clicksignhandler.EnvelopeCreate{
		Name: "Test",
	}); err != nil {
		panic(err)
	} else if document, err := clicksignHandler.DocumentCreate(&clicksignhandler.DocumentCreate{
		Envelope:   &envelope.Data,
		FileType:   clicksignhandler.DFT_PDF,
		FileName:   "document.pdf",
		FileBase64: filePDFBase64,
	}); err != nil {
		panic(err)
	} else if result, err := clicksignHandler.DocumentGetById(envelope.Data.ID, document.Data.ID); err != nil {
		panic(err)
	} else {
		fmt.Print(result.Data.Attributes)
	}
```
12. DocumentsGetFirstPage: Search for documents in an envelope. Only the first page.
```
```
13. DocumentsGetNextPage: Search for documents in an envelope. Only the first page next.
```
```
14. DocumentGetEvents_AddSigner: Search for events of type 'AddSigner' for a document.
```
```
15. DocumentGetEvents_Sign: Search for 'Sign' type events for a document.
```
```
16. DocumentCancel: Cancel a document from an envelope. If there is only one envelope, the envelope is cancelled.
```
    if envelope, err := clicksignHandler.EnvelopeCreate(&clicksignhandler.EnvelopeCreate{
		Name: "Test",
	}); err != nil {
		panic(err)
	} else if document, err := clicksignHandler.DocumentCreate(&clicksignhandler.DocumentCreate{
		Envelope:   &envelope.Data,
		FileType:   clicksignhandler.DFT_PDF,
		FileName:   "document.pdf",
		FileBase64: filePDFBase64,
	}); err != nil {
		panic(err)
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
		panic(err)
	} else if envelope, err := clicksignHandler.EnvelopeActive(&envelope.Data); err != nil {
		panic(err)
	} else {
		if _, err := clicksignHandler.DocumentCancel(&envelope.Data, &document.Data.ID); err != nil {
			panic(err)
		}
	}
```
17. SignerCreate: Create a signer on a document in an envelope.
```
    if envelope, err := clicksignHandler.EnvelopeCreate(&clicksignhandler.EnvelopeCreate{
		Name: "Test",
	}); err != nil {
		panic(err)
	} else if document, err := clicksignHandler.DocumentCreate(&clicksignhandler.DocumentCreate{
		Envelope:   &envelope.Data,
		FileType:   clicksignhandler.DFT_PDF,
		FileName:   "document.pdf",
		FileBase64: filePDFBase64,
	}); err != nil {
		panic(err)
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
		panic(err)
	}
```
18. SignerDelete: Remove a signer from a document on an envelope.
```
```
19. ObserverCreate: Add a observer to a document in an envelope.
```
    if envelope, err := clicksignHandler.EnvelopeCreate(&clicksignhandler.EnvelopeCreate{
		Name: "Test",
	}); err != nil {
		panic(err)
	} else if document, err := clicksignHandler.DocumentCreate(&clicksignhandler.DocumentCreate{
		Envelope:   &envelope.Data,
		FileType:   clicksignhandler.DFT_PDF,
		FileName:   "document.pdf",
		FileBase64: filePDFBase64,
	}); err != nil {
		panic(err)
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
		panic(err)
	} else if _, err := clicksignHandler.ObserverCreate(&clicksignhandler.ObserverCreate{
		Envelope: &envelope.Data,
		Name:     "Observer Test",
		Email:    "observer@test.com",
	}); err != nil {
		panic(err)
	}
```
