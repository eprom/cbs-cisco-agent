package sms

import (
	"github.com/jakhog/cbs-cisco-agent/cisco/common"
	"github.com/jakhog/cbs-cisco-agent/cisco/sms/status"
)

func GetSMSStatus() (Status, error) {
	stat, err := common.GetData("/show/cellular/0/sms", func(data []byte) (interface{}, error) {
		return status.Parse("smsstatus", data)
	})
	if err != nil {
		return Status{}, err
	} else {
		return stat.(Status), nil
	}
}
