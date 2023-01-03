package main

import "testing"

func BenchmarkCompile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetSegments("./file.m3u8")
	}
}
