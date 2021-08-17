package dashboard

import (
	"pmo-test4.yz-intelligence.com/kit/component/dashboard/service"
)

// DashBoard DashBoard
type DashBoard struct {
	Service *service.Service
}

// New New
func New() *DashBoard {
	dashBoard := &DashBoard{}
	dashBoard.Service = service.New()
	return dashBoard
}
