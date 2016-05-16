package models

var ConfigObjectMap = map[string]ConfigObj{
	"User":                 &User{},
	"UserState":            &UserState{},
	"SystemStatusState":    &SystemStatusState{},
	"SystemSwVersionState": &SystemSwVersionState{},
	"Daemon":               &Daemon{},
	"ArpDeleteByIPv4Addr":  &ArpDeleteByIPv4Addr{},
	"ArpDeleteByIfName":    &ArpDeleteByIfName{},
	"ArpRefreshByIPv4Addr": &ArpRefreshByIPv4Addr{},
	"ArpRefreshByIfName":   &ArpRefreshByIfName{},
}
