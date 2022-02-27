package package_testing

import (
	"fmt"
	"testing"
)

func IntSum( a, b int ) int {
	return a + b
}

func TestIntSum( t *testing.T ) {
	var tests = []struct {
		a, b int
		want int
	} {
		{ 1, 1, 2 },
		{ 2, 3, 5 },
		{ -1, 1, 0 },
		{ 5, 5, 10 },
		{ -1, -1, -2 },
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
	
		t.Run( testname, func( t *testing.T ) {
			ans := IntSum( tt.a, tt.b )

			if ans != tt.want {
				t.Errorf( "got %d, want %d", ans, tt.want )
			}
		})
	}
}