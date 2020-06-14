package service

import (
	"fmt"
	"os"
	"testing"

	"github.com/holocycle/holo-back/pkg/test"
)

func TestMain(m *testing.M) {
	free, err := test.InitTestHelper()
	if err != nil {
		fmt.Printf("Failed to create TestHelper err=%+v\n", err)
		os.Exit(1)
	}
	defer free()

	os.Exit(m.Run())
}
