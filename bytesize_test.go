// Copyright (c) 2022 golang-boy
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package bytesize

import (
	"fmt"
	"testing"
)

func TestPretty(t *testing.T) {

	args := []struct {
		expected string
		bytes    float64
	}{
		{"2.00KB", 2048},
		{"2.00MB", 2048 * 1024},
		{"2.00GB", 2048 * 1024 * 1024},
		{"2.00TB", 2048 * 1024 * 1024 * 1024},
		{"2.00PB", 2048 * 1024 * 1024 * 1024 * 1024},
		{"2.00EB", 2048 * 1024 * 1024 * 1024 * 1024 * 1024},
		{"2.00ZB", 2048 * 1024 * 1024 * 1024 * 1024 * 1024 * 1024},
		{"2.00YB", 2048 * 1024 * 1024 * 1024 * 1024 * 1024 * 1024 * 1024},
	}

	for _, arg := range args {

		if arg.expected != Pretty(arg.bytes).String() {
			t.Errorf("bad format: %s", Pretty(arg.bytes))
			return
		}
	}
}

func ExamplePretty() {

	kb := 2048.00000
	fmt.Printf("%s", Pretty(kb))
	// Output:
	// 2.00KB

}
