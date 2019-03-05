package otputils

import (
	"github.com/daveio/gotp/verbose"
	"github.com/pquerna/otp/totp"
	"time"
)

var (
	V = verbose.V
)

func GenerateOTP(key string) (otp string, error error) {
	res, err := totp.GenerateCode(key, time.Now())
	return res, err
}
// TODO URI parsing logic
/* func ParseURI(totpURI string) (site string, uid string, key string, error error) {
	meta := regexp.QuoteMeta("?")
	expr, err := regexp.Compile("^otpauth://totp/(.*):(.*)" + meta + "(.*)$")
	if err == nil {
		fmt.Println(expr.FindStringSubmatch(totpURI))
	} else {
		return "", "", "", err
	}
} */
