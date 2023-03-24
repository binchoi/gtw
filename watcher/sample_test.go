package watcher

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestMainSuiteTwo(t *testing.T) {
	suite.Run(t, new(MainSuiteTwo))
}

type MainSuiteTwo struct {
	suite.Suite
}

func (s *MainSuiteTwo) TestMainTwo() {
	s.Equal(1, 1)
}

func TestAdd(t *testing.T) {
	got := 4 + 6
	want := 10

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
