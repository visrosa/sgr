package sgr

import (
	"strings"
	"testing"
)

func TestBoldStruct(t *testing.T) {
	if Bold.On.Apply() != "\x1b[1m" {
		t.Errorf("Bold.On.Apply() = %q, want %q", Bold.On.Apply(), "\x1b[1m")
	}
	if Bold.Off.Apply() != "\x1b[22m" {
		t.Errorf("Bold.Off.Apply() = %q, want %q", Bold.Off.Apply(), "\x1b[22m")
	}
	if Bold.Apply().Apply() != Bold.On.Apply() {
		t.Errorf("Bold.Apply() = %q, want %q", Bold.Apply().Apply(), Bold.On.Apply())
	}
}

func TestUnderlineStruct(t *testing.T) {
	if Underline.On.Apply() != "\x1b[4m" {
		t.Errorf("Underline.On.Apply() = %q, want %q", Underline.On.Apply(), "\x1b[4m")
	}
	if Underline.Off.Apply() != "\x1b[24m" {
		t.Errorf("Underline.Off.Apply() = %q, want %q", Underline.Off.Apply(), "\x1b[24m")
	}
	if Underline.Apply().Apply() != Underline.On.Apply() {
		t.Errorf("Underline.Apply() = %q, want %q", Underline.Apply().Apply(), Underline.On.Apply())
	}
}

func TestBlinkStruct(t *testing.T) {
	if Blink.Slow.Apply() != "\x1b[5m" {
		t.Errorf("Blink.Slow.Apply() = %q, want %q", Blink.Slow.Apply(), "\x1b[5m")
	}
	if Blink.Rapid.Apply() != "\x1b[6m" {
		t.Errorf("Blink.Rapid.Apply() = %q, want %q", Blink.Rapid.Apply(), "\x1b[6m")
	}
	if Blink.Off.Apply() != "\x1b[25m" {
		t.Errorf("Blink.Off.Apply() = %q, want %q", Blink.Off.Apply(), "\x1b[25m")
	}
	if Blink.Apply().Apply() != Blink.Slow.Apply() {
		t.Errorf("Blink.Apply() = %q, want %q", Blink.Apply().Apply(), Blink.Slow.Apply())
	}
}

func TestFgBgColorHelpers(t *testing.T) {
	fg := Fg.Color(123)
	if fg.Apply() != "\x1b[38;5;123m" {
		t.Errorf("Fg.Color(123).Apply() = %q, want %q", fg.Apply(), "\x1b[38;5;123m")
	}
	bg := Bg.Color(200)
	if bg.Apply() != "\x1b[48;5;200m" {
		t.Errorf("Bg.Color(200).Apply() = %q, want %q", bg.Apply(), "\x1b[48;5;200m")
	}
}

func TestFgBgRGBHelpers(t *testing.T) {
	fg := Fg.RGB(1, 2, 3)
	if fg.Apply() != "\x1b[38;2;1;2;3m" {
		t.Errorf("Fg.RGB(1,2,3).Apply() = %q, want %q", fg.Apply(), "\x1b[38;2;1;2;3m")
	}
	bg := Bg.RGB(4, 5, 6)
	if bg.Apply() != "\x1b[48;2;4;5;6m" {
		t.Errorf("Bg.RGB(4,5,6).Apply() = %q, want %q", bg.Apply(), "\x1b[48;2;4;5;6m")
	}
}

func TestCSI(t *testing.T) {
	if CSI("31") != "\x1b[31m" {
		t.Errorf("CSI(\"31\") = %q, want %q", CSI("31"), "\x1b[31m")
	}
	if CSI("38", "123") != "\x1b[38;5;123m" {
		t.Errorf("CSI(\"38\",\"5\",\"123\") = %q, want %q", CSI("38", "5", "123"), "\x1b[38;5;123m")
	}
	if CSI("48", "200") != "\x1b[48;5;200m" {
		t.Errorf("CSI(\"48\",\"5\",\"200\") = %q, want %q", CSI("48", "5", "200"), "\x1b[48;5;200m")
	}
}

func TestRender(t *testing.T) {
	r := Bold.On.Render()
	if !strings.Contains(r, "Bold") || !strings.Contains(r, "\x1b[1m") {
		t.Errorf("Bold.On.Render() = %q, want code and name", r)
	}
}
