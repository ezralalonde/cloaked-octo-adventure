package swap

import (
	"testing"
)

func tSwap(in []int, out int, t *testing.T) {
	f := Swap
    mod := make([]int, len(in))
    copy(mod, in)
	ans := f(mod)
	if ans != out {
		t.Errorf("Swap(%v) = %v, wanted %v", in, ans, out)
	}
}

func TestSwap1(t *testing.T) {
	data := []int{1, 3, 2}
	tSwap(data, 1, t)
}
func TestSwap2(t *testing.T) {
	data := []int{4, 3, 2, 1}
	tSwap(data, 6, t)
}
func TestSwap3(t *testing.T) {
	data := []int{2, 1}
	tSwap(data, 1, t)
}
