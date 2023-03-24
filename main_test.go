package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestMainSuite(t *testing.T) {
	suite.Run(t, new(MainSuite))
}

type MainSuite struct {
	suite.Suite
}

func (s *MainSuite) TestMain() {
	s.Equal(1, 1)
	s.Equal(10000, 10000)
}
