package main

import "testing"

func TestMain(t *testing.T) {
	{
		body := `
.0..
00..
....
...0`
		expected := 2
		if countConnected(body) != expected {
			t.Error("answer was mismatch, expected:", expected)
		}
	}
	{
		body := `
.0..
00..
..0.
...0`
		expected := 1
		if countConnected(body) != expected {
			t.Error("answer was mismatch, expected:", expected)
		}
	}
	{
		body := `
.0..0
00..0
.....
...0.`
		expected := 3
		if countConnected(body) != expected {
			t.Error("answer was mismatch, expected:", expected)
		}
	}
}

func BenchmarkMainSmall(b *testing.B) {
	body := `
.0.
00.`
	expected := 1
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if countConnected(body) != expected {
			b.Error("answer was mismatch, expected:", expected)
		}
	}
}

func BenchmarkMainBig(b *testing.B) {
	body := `
.0..0...................................................................................................
00..0...................................................................................................
..............................................................000000000.................................
...0....................................................................................................
........................................................................................................
........................................................................................................
........................................................................................................
........................................................................................................
........................................................................................................
....................................................................................00..................
....................................................................................00..................
.....................................................................................00.................
.................000.................................................................00.................
...................00................................................................0000...............
....................000................................................................00000000.........
........................................................................................................
........................................................................................................
........................................................................................................
........................................................................................................
........................................................................................................
........................................................................................................
........................................................................................................
........................................................................................................
........................................................................................................
......................................................................0.................................
......................................................................0.................................
......................................................................0.................................
......................................................................0.................................
......................................................................0.................................
........................................................................................................`
	expected := 7
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if countConnected(body) != expected {
			b.Error("answer was mismatch, expected:", expected)
		}
	}
}