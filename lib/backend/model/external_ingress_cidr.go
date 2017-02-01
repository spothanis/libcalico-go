package model

import (
	cnet "github.com/projectcalico/libcalico-go/lib/net"
)

type AddressSet struct {
	Kind string        `json:omitempty`
	Sets []*cnet.IPNet `json:sets`
	Ref  string        `json:omitempty` // this should be the name of the reference object of the format ns.name
	Name string        `json:omitempty`
	Desc string        `json:omitempty`
}
