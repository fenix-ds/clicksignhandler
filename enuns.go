package clicksignhandler

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type Environment string

const (
	EnvSandbox Environment = "sandbox"
	EnvProd    Environment = "app"
)

func StrToEnvironment(value string) (*Environment, error) {
	switch strings.ToLower(value) {
	case "sandbox":
		return EnvSandbox.ToPoint(), nil
	case "app":
		return EnvProd.ToPoint(), nil
	default:
		return nil, fmt.Errorf("invalid value")
	}
}

func (e Environment) ToPoint() *Environment {
	return &e
}

type CSIntString int

func (i *CSIntString) UnmarshalJSON(b []byte) error {
	var n int
	if err := json.Unmarshal(b, &n); err == nil {
		*i = CSIntString(n)
		return nil
	}

	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	v, err := strconv.Atoi(s)
	if err != nil {
		return err
	}

	*i = CSIntString(v)
	return nil
}

type DocumentFileType string

const (
	DFT_PDF DocumentFileType = "pdf"
)

type SignerType string

const (
	SRT_CONTRACTEE SignerType = "contractee"
	SRT_WITNESS    SignerType = "witness"
	SRT_CONTRACTOR SignerType = "contractor"
)

type AuthenticationType string

const (
	AUT_EMAIL         AuthenticationType = "email"
	AUT_AUTHSIGNATURE AuthenticationType = "auto_signature"
)

type SignatureRequest string

const (
	SREQ_NONE     SignatureRequest = "none"
	SREQ_EMAIL    SignatureRequest = "email"
	SREQ_WHATSAPP SignatureRequest = "whatsapp"
	SREQ_SMS      SignatureRequest = "sms"
)

type SignatureReminder string

const (
	SREM_NONE  SignatureReminder = "none"
	SREM_EMAIL SignatureReminder = "email"
)

type SignatureDocumentSigned string

const (
	SDS_EMAIL    SignatureDocumentSigned = "email"
	SDS_WHATSAPP SignatureDocumentSigned = "whatsapp"
)
