// arrays_test.go
package common

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	num := []int{1, 2, 3, 4, 5}
	reversed := []int{5, 4, 3, 2, 1}
	var rev_num = Reverse(num)
	if !reflect.DeepEqual(rev_num, reversed) {
		t.Error("Arrays don't match")
	}
	t.Log("Finished")
}
