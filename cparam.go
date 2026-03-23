package clicksignhandler

import (
	"fmt"
	"time"
)

type ClicksignParam struct {
	Environment Environment
	Key         string
	DefautUTC   *time.Location
}

func (p ClicksignParam) checkData() error {
	if p.Environment != EnvSandbox && p.Environment != EnvProd {
		return fmt.Errorf("invalid environment")
	}

	if len(p.Key) == 0 {
		return fmt.Errorf("invalid key")
	}

	if p.DefautUTC == nil {
		return fmt.Errorf("Added timezone for communication configuration with clicksing.")
	}

	return nil
}
