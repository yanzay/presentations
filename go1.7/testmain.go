import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// initialize
	res := m.Run()
	// tear-down
	os.Exit(res)
}
