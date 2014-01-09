package main

import (
	. "github.com/matttproud/gocheck"
	"testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestHelloWorld(c *C) {
	c.Check(42, Equals, 42)
	c.Check(42, Equals, 42)
}
