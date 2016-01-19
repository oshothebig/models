package models

import (
	//"models"
	"arpd"
)

func ConvertarpdArpConfigObjToThrift(dbobj *ArpConfig, thriftobj *arpd.ArpConfig) {
	thriftobj.ArpConfigKey = string(dbobj.ArpConfigKey)
	thriftobj.Timeout = int32(dbobj.Timeout)
}

func ConvertThriftToarpdArpConfigObj(thriftobj *arpd.ArpConfig, dbobj *ArpConfig) {
	dbobj.ArpConfigKey = string(thriftobj.ArpConfigKey)
	dbobj.Timeout = int32(thriftobj.Timeout)
}
