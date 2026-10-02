package control

import "testing"

func TestFirstApplyKeepsDuration(t *testing.T) {
	if got := New(10000, 3000).Apply("u1", "stun", 1000, 0); got != 1000 {
		t.Fatalf("第一次施加应该是满时长：%v", got)
	}
}

func TestCountStartsAtZero(t *testing.T) {
	if New(10000, 3000).Count("u1", "stun") != 0 {
		t.Fatal("计数初始应该是 0")
	}
}

func TestCountIncreasesAfterApply(t *testing.T) {
	tracker := New(10000, 3000)
	tracker.Apply("u1", "stun", 1000, 0)
	if tracker.Count("u1", "stun") == 0 {
		t.Fatal("施加之后计数应该增加")
	}
}

func TestAccessors(t *testing.T) {
	tracker := New(8000, 2500)
	if tracker.Window() != 8000 || tracker.Immune() != 2500 {
		t.Fatal("窗口与免疫时长应该能读回来")
	}
}

func TestScannedStartsAtZero(t *testing.T) {
	if New(10000, 3000).Scanned() != 0 {
		t.Fatal("Scanned 初始应该是 0")
	}
}
