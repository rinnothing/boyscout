package extension

import (
	"fmt"
	"github.com/rinnothing/boyscout"
)

func init() {
	scout := boyscout.DefaultGrandscout.GetScout("speakers")
	err := scout.Register("speaker_b", SpeakerB{})
	if err != nil {
		fmt.Println("Couldn't register SpeakerB:", err)
	}
}

type SpeakerB struct {
	i int
}

func (s *SpeakerB) Speak() (speech string, err error) {
	speech = "B"
	err = nil
	return
}
