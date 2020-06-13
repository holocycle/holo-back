package test

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	free, err := InitTestHelper()
	if err != nil {
		fmt.Printf("Failed to create TestHelper err=%+v\n", err)
		os.Exit(1)
	}
	defer free()

	os.Exit(m.Run())
}
