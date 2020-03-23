package flag

import (
	"flag"
	"net/url"
	"os"

	"github.com/robert-zaremba/errstack"
	bat "github.com/robert-zaremba/go-bat"
)

// URL is a structure containing network connection details to a service
// The value should have the following structure:
//     //username:password@host:port/directory
type URL struct {
	url.URL
}

// Set implements github.com/robert-zaremba/flag Value interface
func (a *URL) Set(value string) error {
	u, err := url.Parse(value)
	if err != nil {
		return errstack.WrapAsReq(err, "Can't parse network address config")
	}
	a.URL = *u
	return nil
}

// Path represents a file in a filesystem
type Path struct {
	Path string
}

// Set implements github.com/robert-zaremba/flag Value interface
func (a *Path) Set(filePath string) error {
	a.Path = filePath
	return a.Check()
}

// String implements github.com/robert-zaremba/flag Value interface
func (a *Path) String() string {
	return a.Path
}

// Check returns an error if it can't find the file
func (a Path) Check() error {
	if a.Path == "" {
		return errstack.NewReq("File path can't be empty")
	}
	_, err := os.Stat(a.Path)
	return err
}

// SrvFlags represents common server flags
type SrvFlags struct {
	Production *bool
	Port       *string
}

// NewSrvFlags setups common server flags
func NewSrvFlags() SrvFlags {
	return SrvFlags{
		flag.Bool("production", false, "Run in production mode"),
		flag.String("port", "8000", "The HTTP listening port"),
	}
}

// Check validates the flags. It may panic!
func (f SrvFlags) Check() error {
	errb := errstack.NewBuilder()
	bat.Atoi64Errp(*f.Port, errb.Putter("port"))
	return errb.ToReqErr()
}
