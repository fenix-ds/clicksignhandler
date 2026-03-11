package clicksignhandler

import "fmt"

type ClicksignParam struct {
	Environment Environment
	Key         string
}

func (p ClicksignParam) checkData() error {
	if p.Environment != EnvSandbox && p.Environment != EnvProd {
		return fmt.Errorf("invalid environment")
	}

	if len(p.Key) == 0 {
		return fmt.Errorf("invalid key")
	}

	return nil
}
