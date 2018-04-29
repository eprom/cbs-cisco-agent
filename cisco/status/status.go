package status

import (
	"github.com/jakhog/cbs-cisco-agent/cisco/common"
)

func GetGeneralStatus(cfg common.Config) (Status, error) {
	stat, err := common.GetData(cfg, "/show/cellular/0/all", func(data []byte) (interface{}, error) {
		return Parse("status", data)
	})
	if err != nil {
		return Status{}, err
	} else {
		return stat.(Status), nil
	}
}
