package gotest

import "testing"

// go test -bench=.
func BenchmarkCaseSumOne(b *testing.B)  {
	b.ResetTimer()
	var sum int
	CaseSumOne(&sum)
	b.StopTimer()
}

func BenchmarkCaseSumTwo(b *testing.B) {
	b.ResetTimer()
	var sum int
	CaseSumTwo(&sum)
	b.StopTimer()
}

func BenchmarkCaseSumThree(b *testing.B) {
	b.ResetTimer()
	var sum int
	CaseSumThree(&sum)
	b.StopTimer()
}

func BenchmarkCaseSumFour(b *testing.B) {
	b.ResetTimer()
	CaseSumFour()
	b.StopTimer()
}