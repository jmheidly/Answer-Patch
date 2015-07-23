package settings

import (
	"os"
	"testing"
)

func TestGetRedisDSN(t *testing.T) {
	os.Setenv("GO_ENV", "preproduction")
	GetRedisDSN()
}
