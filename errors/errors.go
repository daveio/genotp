package errors

import (
	"fmt"
	"github.com/daveio/gotp/verbose"
)

var (
	V = verbose.V
)

func (e *UidError) Error() string {
	return fmt.Sprintf("UID %s not found for site %s", e.Site, e.Uid)
}

func (e *SiteError) Error() string {
	return fmt.Sprintf("Site %s not found", e.Site)
}
