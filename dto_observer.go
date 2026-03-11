package clicksignhandler

type ObserverData struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Attributes struct {
		Email             string `json:"email"`
		Kind              string `json:"kind"`
		CommunicateEvents struct {
			SignatureWatcherDocumentSent     string `json:"signature_watcher_document_sent"`
			SignatureWatcherDocumentSigned   string `json:"signature_watcher_document_signed"`
			SignatureWatcherDocumentDeadline string `json:"signature_watcher_document_deadline"`
			SignatureWatcherDocumentCanceled string `json:"signature_watcher_document_canceled"`
			SignatureWatcherEnvelopeClosed   string `json:"signature_watcher_envelope_closed"`
		} `json:"communicate_events"`
		AttachDocumentsEnabled bool   `json:"attach_documents_enabled"`
		Created                string `json:"created"`
		Modified               string `json:"modified"`
	} `json:"attributes"`
}

type ObserverCreate struct {
	Envelope *EnvelopeData
	Name     string
	Email    string
}
