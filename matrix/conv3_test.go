package matrix

import (
	"reflect"
	"testing"
)

func TestConv3(t *testing.T) {
	in := [][][]int64{{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, {{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, {{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}}
	kernel := [][][]int64{{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, {{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, {{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}}

	out := Conv3(in, kernel, [3]int{1, 1, 1})
	expected := [][][]int64{{{8, 12, 8}, {12, 18, 12}, {8, 12, 8}}, {{12, 18, 12}, {18, 27, 18}, {12, 18, 12}}, {{8, 12, 8}, {12, 18, 12}, {8, 12, 8}}}
	if !reflect.DeepEqual(out, expected) {
		t.Errorf("Expected: %v\nGot: %v\n", expected, out)
	}
}
