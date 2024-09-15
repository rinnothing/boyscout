package extension

import (
	"errors"
	"fmt"
	"github.com/rinnothing/boyscout"
)

func init() {
	scout := boyscout.DefaultGrandscout.GetScout("speakers")
	err := scout.Register("speaker_a", SpeakerA{})
	if err != nil {
		fmt.Println("Couldn't register SpeakerA:", err)
	}
}

type SpeakerA struct {
	i int
}

func (s *SpeakerA) Speak() (speech string, err error) {
	speech = "A"
	if s.i%2 == 1 {
		err = errors.New("uneven launch")
	} else {
		err = nil
	}
	return
}
