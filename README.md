# MusGo_int64
Provides serialization of the Golang's `int64` type.

# How to use
Simply cast `int64` to `musgo_int64.Int64`:
```go
	var n int64 = 5
	// Marshal
	buf := make([]byte, musgo_int64.Int64(n).SizeMUS())
	musgo_int64.Int64(n).MarshalMUS(buf)
	// Unmarshal
	_, err := (*musgo_int64.Int64)(&n).UnmarshalMUS(buf)
	if err != nil {
		panic(err)
	}
```

# More info
You can find at [github.com/ymz-ncnk/musgo](https://github.com/ymz-ncnk/musgo).

