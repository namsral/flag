package flag

import (
	"testing"

	rzcheck "github.com/robert-zaremba/checkers"
	gocheck "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { gocheck.TestingT(t) }

type ExtraSuite struct{}

func init() {
	gocheck.Suite(&ExtraSuite{})
}

func (s *ExtraSuite) TestValidateDirExists(c *gocheck.C) {
	ff := Path{}
	err := ff.Set("/tmp")
	c.Assert(err, gocheck.IsNil)
	c.Assert(ff.String(), gocheck.Equals, "/tmp")
}

func (s *ExtraSuite) TestValidatePermissionDenied(c *gocheck.C) {
	ff := Path{}
	_ = ff.Set("/root/secret")
	err := ff.Check()
	c.Assert(err, gocheck.ErrorMatches, "stat /root/secret: permission denied")
}

func (s *ExtraSuite) TestValidateNoFile(c *gocheck.C) {
	ff := Path{}
	_ = ff.Set("")
	err := ff.Check()
	c.Assert(err, rzcheck.ErrorContains, "File path can't be empty")
}

func (s *ExtraSuite) TestHandleBadDefaultWithPanic(c *gocheck.C) {
	ff := Path{"hello-world"}
	err := ff.Check()
	c.Assert(err, gocheck.ErrorMatches, "stat hello-world: no such file or directory")
}

func (s *ExtraSuite) TestHandleDefault(c *gocheck.C) {
	ff := Path{"/tmp"}
	c.Assert(ff.String(), gocheck.Equals, "/tmp")
	ff = Path{"/"}
	c.Assert(ff.String(), gocheck.Equals, "/")
	err := ff.Set("/tmp")
	c.Assert(err, gocheck.IsNil)
	c.Assert(ff.String(), gocheck.Equals, "/tmp")
}
