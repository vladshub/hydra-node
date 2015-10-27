package players

import "errors"

type PlayerI interface {
	//Play Controlls
	Play(path string) (bool, error)
	CanPlay(path string) bool
	Pause() (bool, error)
	Stop() (bool, error)
	Reset() error

	//Seek Controlls
	SeekBack() (bool, error)
	SeekBackFast() (bool, error)
	SeekForward() (bool, error)
	SeekForwardFast() (bool, error)

	//Volume Controlls
	VolumeUp(amount int) (bool, error)
	VolumeDown(amount int) (bool, error)

	//Subtitle Controlls
	SubtitlesOn() (bool, error)
	SubtitlesOff() (bool, error)
	SubtitlesUse(path string) (bool, error)
}

var activePlayers = map[string]*PlayerI{}

func PlayerFactory(player_type string) (*PlayerI, error) {
	if activePlayers == nil {
		activePlayers = make(map[string]*PlayerI)
	} else {
		player, ok := activePlayers[player_type]
		if ok {
			return player, nil
		}
	}
	switch player_type {
	case "OmxPlayer":
		player, err := NewOmxPlayer()
		activePlayers[player_type] = &player
		return &player, err
	}
	return nil, errors.New("Unknown player")
}
