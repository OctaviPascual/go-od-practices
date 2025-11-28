package skip

// Code from https://github.com/OctaviPascual/twg/blob/master/skip/flag_test.go

import (
	"flag"
	"os"
	"testing"
)

var integration = false

func init() {
	flag.BoolVar(&integration, "integration", false, "run database integration tests")
}

// More info on https://pkg.go.dev/testing#hdr-Main
func TestMain(m *testing.M) {
	flag.Parse()
	if integration {
		// setup integration stuff if you need to
	}
	result := m.Run()
	if integration {
		// teardown integration stuff if you need to
	}
	os.Exit(result)
}

func TestWithFlag(t *testing.T) {
	if !integration {
		t.Skip()
	}
	t.Log("Running the integration test...")
}
