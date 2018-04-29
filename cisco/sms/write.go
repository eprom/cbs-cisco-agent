package sms

import (
	"errors"
	"fmt"
	"github.com/jakhog/cbs-cisco-agent/cisco/common"
	"github.com/mozillazg/go-unidecode"
	"net/url"
)

func cleanString(original string) string {
	// Replace any special Unicode characters as best we can
	sevenBits := unidecode.Unidecode(original)
	// Take out the rest that doesn't work on the Cisco
	clean := make([]rune, 0, len(sevenBits))
	for _, r := range sevenBits {
		if r < ' ' || r > '~' {
			// Characters outside this space doesn't make sense in text
			continue
		}
		switch r {
		case '?':
			clean = append(clean, '.')
			// The few remaining characters should work
		default:
			clean = append(clean, r)
		}
	}
	// This should be good
	return string(clean)
}

func SendSMS(sms OutgoingSMS) error {
	text := cleanString(sms.Text) // The Cisco is very strict with its characters
	if len(text) > 160 {
		return errors.New("SMS too long")
	}

	path := fmt.Sprintf("/cellular/0/lte/sms/send/%v/%v", url.QueryEscape(sms.To), url.QueryEscape(text))
	return common.PostCommandSimple(path)
}
