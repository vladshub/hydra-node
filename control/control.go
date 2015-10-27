package control

import (
	"errors"
	"github.com/vladshub/hydra-node/fsmanager"
	"github.com/vladshub/hydra-node/players"
)

type Control interface {
	Listen(addr string) error
}

func ControlFactory(ctlType string,
	player *players.PlayerI,
	fsmanager *fsmanager.FileSystemManagerI) (Control, error) {
	switch ctlType {
	case "HttpControl":
		return NewHttpControl(player, fsmanager), nil
	}

	return nil, errors.New("Unknown Control")
}
