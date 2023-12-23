//go:build !linux

package device

import (
	"github.com/maxnasonov/wireguard-go/conn"
	"github.com/maxnasonov/wireguard-go/rwcancel"
)

func (device *Device) startRouteListener(bind conn.Bind) (*rwcancel.RWCancel, error) {
	return nil, nil
}
