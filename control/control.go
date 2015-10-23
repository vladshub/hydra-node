package control

import (
	"github.com/vladshub/hydra-node/fsmanager"
	"github.com/vladshub/hydra-node/players"
	"errors"
)

type Control interface {
	Listen(addr string) error
}

func ControlFactory(ctlType string,
	player *players.Player,
	fsmanager *fsmanager.FsManager) (Control, error) {
	switch ctlType {
	case "HttpControl":
		return NewHttpControl(player, fsmanager), nil
	}

	return nil, errors.New("Unknown Control")
}
