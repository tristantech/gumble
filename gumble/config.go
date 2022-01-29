package gumble

import (
	"time"
)

// Config holds the Mumble configuration used by Client. A single Config should
// not be shared between multiple Client instances.
type Config struct {
	// User name used when authenticating with the server.
	Username string
	// Password used when authenticating with the server. A password is not
	// usually required to connect to a server.
	Password string
	// The initial access tokens to the send to the server. Access tokens can be
	// resent to the server using:
	//  client.Send(config.Tokens)
	Tokens AccessTokens

	// AudioInterval is the interval at which audio packets are sent. Valid
	// values are: 10ms, 20ms, 40ms, and 60ms.
	AudioInterval time.Duration
	// AudioDataBytes is the number of bytes that an audio frame can use.
	AudioDataBytes int
	// Audio sample rate to use for audio I/O
	AudioSampleRate int

	// The event listeners used when client events are triggered.
	Listeners      Listeners
	AudioListeners AudioListeners
}

// NewConfig returns a new Config struct with default values set.
func NewConfig() *Config {
	return &Config{
		AudioInterval:  AudioDefaultInterval,
		AudioDataBytes: AudioDefaultDataBytes,
		AudioSampleRate: AudioDefaultSampleRate,
	}
}

// Attach is an alias of c.Listeners.Attach.
func (c *Config) Attach(l EventListener) Detacher {
	return c.Listeners.Attach(l)
}

// AttachAudio is an alias of c.AudioListeners.Attach.
func (c *Config) AttachAudio(l AudioListener) Detacher {
	return c.AudioListeners.Attach(l)
}

// AudioFrameSize returns the appropriate audio frame size, based off of the
// audio interval.
func (c *Config) AudioFrameSize() int {
	// AudioDefaultFrameSize is the number of audio frames that should be sent in
	// a 10ms window.
	AudioDefaultFrameSize := c.AudioSampleRate / 100
	return int(c.AudioInterval/AudioDefaultInterval) * AudioDefaultFrameSize
}

// Gets the maximum audio frame size from another user that will be processed.
func (c *Config) AudioMaximumFrameSize() int {
	return c.AudioSampleRate / 1000 * 60
}
