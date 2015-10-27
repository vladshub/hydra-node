package players

import (
	"io"
	"os"
	"os/exec"
	"regexp"
)

type OmxPlayer struct {
	supportedTypes *regexp.Regexp
	binPath        string

	currentFile string
	command     chan string
	omxIn       io.WriteCloser

	player *exec.Cmd

	// OMXPlayer control commands, these are piped via STDIN to omxplayer process
	commands map[string]string
}

func NewOmxPlayer() (PlayerI, error) {
	player := &OmxPlayer{
		supportedTypes: regexp.MustCompile(".(avi|mpg|mov|flv|wmv|asf|mpeg|m4v|divx|mp4|mkv)$"),
		binPath:        "/usr/bin/omxplayer",
		commands: map[string]string{
			"pause":             "p",            // Pause/continue playback
			"stop":              "q",            // Stop playback and exit
			"volume_up":         "+",            // Change volume by +3dB
			"volume_down":       "-",            // Change volume by -3dB
			"subtitles":         "s",            // Enable/disable subtitles
			"seek_back":         "\x1b\x5b\x44", // Seek -30 seconds
			"seek_back_fast":    "\x1b\x5b\x42", // Seek -600 second
			"seek_forward":      "\x1b\x5b\x43", // Seek +30 second
			"seek_forward_fast": "\x1b\x5b\x41", // Seek +600 seconds
		},
	}
	player.omxListen()
	return player, nil
}

func (omx OmxPlayer) omxListen() error {
	omx.command = make(chan string)

	for {
		command := <-omx.command

		// Skip command handling of omx player is not active
		if omx.player == nil {
			continue
		}

		// Send command to the player
		omx.omxWrite(command)

		// Attempt to kill the process if stop command is requested
		if command == "stop" {
			omx.player.Process.Kill()
		}
	}
}

func (omx OmxPlayer) omxWrite(command string) {
	if omx.omxIn != nil {
		io.WriteString(omx.omxIn, omx.commands[command])
	}
}

func (omx OmxPlayer) Play(path string) (bool, error) {
	omx.player = exec.Command(
		omx.binPath, // path to omxplayer executable
		"--refresh", // adjust framerate/resolution to video
		"--blank",   // set background to black
		"--adev",    // audio out device
		"hdmi",      // using hdmi for audio/video
		path,        // path to video file
	)

	stdin, err := omx.player.StdinPipe()
	if err != nil {
		return false, err
	}

	defer stdin.Close()

	// Redirect output for debugging purposes
	omx.player.Stdout = os.Stdout

	// Start omxplayer execution.
	// If successful, something will appear on HDMI display.
	err = omx.player.Start()
	if err != nil {
		return false, err
	}

	// Make child's STDIN globally available
	omx.omxIn = stdin

	// Wait until child process is finished
	err = omx.player.Wait()
	if err != nil {
		return false, err
	}

	omx.Reset()
	return true, nil
}

// Reset internal state and stop any running processes
func (omx OmxPlayer) Reset() (error) {
	omx.player = nil
	omx.omxIn = nil
	omx.omxKill()
	return nil
}

// Terminate any running omxplayer processes. Fixes random hangs.
func (omx OmxPlayer) omxKill() {
	exec.Command("killall", "omxplayer.bin").Output()
	exec.Command("killall", "omxplayer").Output()
}

// CanPlay returns true or false if the file can be played
func (omx OmxPlayer) CanPlay(path string) bool {
	if omx.supportedTypes.Match([]byte(path)) {
		return true
	}
	return false
}

func (omx OmxPlayer) Pause() (bool, error) {
	return false, nil
}
func (omx OmxPlayer) Stop() (bool, error) {
	return false, nil
}

//Seek Controlls
func (omx OmxPlayer) SeekBack() (bool, error) {
	return false, nil
}
func (omx OmxPlayer) SeekBackFast() (bool, error) {
	return false, nil
}
func (omx OmxPlayer) SeekForward() (bool, error) {
	return false, nil
}
func (omx OmxPlayer) SeekForwardFast() (bool, error) {
	return false, nil
}

//Volume Controlls
func (omx OmxPlayer) VolumeUp(amount int) (bool, error) {
	return false, nil
}
func (omx OmxPlayer) VolumeDown(amount int) (bool, error) {
	return false, nil
}

//Subtitle Controlls
func (omx OmxPlayer) SubtitlesOn() (bool, error) {
	return false, nil
}
func (omx OmxPlayer) SubtitlesOff() (bool, error) {
	return false, nil
}
func (omx OmxPlayer) SubtitlesUse(path string) (bool, error) {
	return false, nil
}
