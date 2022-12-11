package musgo_int64

import "testing"

func TestMusgoInt64(t *testing.T) {
	var n int64 = 5
	buf := make([]byte, Int64(n).SizeMUS())
	Int64(n).MarshalMUS(buf)

	var an int64
	(*Int64)(&an).UnmarshalMUS(buf)

	if n != an {
		t.Fatal("something went wrong")
	}
}
