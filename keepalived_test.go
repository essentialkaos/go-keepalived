package keepalived

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2020 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"

	. "pkg.re/check.v1"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type KASuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&KASuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *KASuite) TestParsing(c *C) {
	ip, err := GetVirtualIP("testdata/1.conf")

	c.Assert(ip, Equals, "192.168.1.2")
	c.Assert(err, IsNil)
}

func (s *KASuite) TestParsingErrors(c *C) {
	ip, err := GetVirtualIP("testdata/unknown.conf")

	c.Assert(ip, Equals, "")
	c.Assert(err, NotNil)

	ip, err = GetVirtualIP("testdata/2.conf")

	c.Assert(ip, Equals, "")
	c.Assert(err, NotNil)

	ip, err = GetVirtualIP("testdata/3.conf")

	c.Assert(ip, Equals, "")
	c.Assert(err, NotNil)
}

func (s *KASuite) TestRoleChecker(c *C) {
	isMaster, err := IsMaster("192.168.0.100")

	c.Assert(isMaster, Equals, false)
	c.Assert(err, IsNil)

	isMaster, err = IsMaster("192.168.0")

	c.Assert(isMaster, Equals, false)
	c.Assert(err, NotNil)
}

func (s *KASuite) TestAux(c *C) {
	c.Assert(extractIP("  abcd  "), Equals, "")
	c.Assert(extractIP("  abcd abcd  "), Equals, "")
}
