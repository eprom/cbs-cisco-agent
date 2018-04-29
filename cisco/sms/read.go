package sms

import (
	"github.com/jakhog/cbs-cisco-agent/cisco/common"
	"github.com/jakhog/cbs-cisco-agent/cisco/sms/table"
)

func GetAllSMSes() ([]IncomingSMS, error) {
	smses, err := common.GetData("/cellular/0/lte/sms/view/all", func(data []byte) (interface{}, error) {
		option := table.AllowInvalidUTF8(true)
		return table.Parse("smslist", data, option)
	})
	if err != nil {
		return nil, err
	} else {
		return smses.([]IncomingSMS), nil
	}
}
