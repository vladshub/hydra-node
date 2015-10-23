package players
import "errors"

type Player interface {
	//Play Controlls
	Play(path string) (bool, error)
	CanPlay(path string) bool
	Pause() (bool, error)
	Stop() (bool, error)

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


func PlayerFactory(player_type string ) (Player, error) {
	switch player_type {
	case "OmxPlayer":
		return NewOmxPlayer()
	}
	return nil, errors.New("Unknown player")
}