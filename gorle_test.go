package gorle

import (
	"testing"
)

func TestEncode(t *testing.T) {
	data := []byte("1111113320000000000")

	t.Logf("%v", data)
	t.Logf("%v", Encode(data))
	t.Logf("%v", Decode(Encode(data)))
}

func BenchmarkEncode(b *testing.B) {
	data := []byte("11111133200000000001111111111111112222222222222222233333333333333331231231231231111113333")
	for i := 0; i < b.N; i++ {
		Encode(data)
	}
}

func BenchmarkDecode(b *testing.B) {
	data := Encode([]byte("11111133200000000001111111111111112222222222222222233333333333333331231231231231111113333"))
	for i := 0; i < b.N; i++ {
		Decode(data)
	}
}
