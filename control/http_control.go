package control

import (
	"github.com/gin-gonic/gin"
	"github.com/vladshub/hydra-node/fsmanager"
	"github.com/vladshub/hydra-node/players"
)

type HttpControl struct {
	player    *players.Player
	fsmanager *fsmanager.FsManager
	engien    *gin.Engine
}

func NewHttpControl(player *players.Player, fsmanager *fsmanager.FsManager) {
	r := gin.Default()

	return HttpControl{
		player:    player,
		fsmanager: fsmanager,
		engien: r,
	}
}

func (ctl *HttpControl) Listen(addr string) error {
	return ctl.engien.Run(addr)
}
