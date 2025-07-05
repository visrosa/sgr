package sgr

import (
	"fmt"
	"strings"
)

// AnsiCode represents an ANSI SGR code and its name.
type AnsiCode struct {
	Code   string
	Name   string
	Abbr   string
	Symbol rune // Optional: leave empty if not set
}

func (a AnsiCode) Apply() string {
	return CSI(a.Code)
}

func (a AnsiCode) Render() string {
	return fmt.Sprintf("% 3s ", a.Code) + CSI(a.Code) + a.Abbr + string(a.Symbol) + a.Name + Reset.Apply()
}

func CSI(s ...string) string {
	switch s[0] {
	case "38", "48": //Set foreground/background.
		return "\x1b\x5b" + s[0] + ";5;" + s[1] + "m"
	default:
		return "\x1b\x5b" + strings.Join(s, ";") + "m"
	}
}

// In common terminal usage, bold is set via SGR 1 and faint by SGR 2. However, there is only one number to reset these attributes, SGR 22, which resets both. There is no way to reset one and not the other. kitty uses 221 and 222 to reset bold and faint independently.
var (
	Reset = AnsiCode{"0", "Reset", "", 0}
	// --- C0 control codes (ASCII 0x00-0x1F) as AnsiCode structs ---
	NUL = AnsiCode{"\x00", "NUL", "", 0}
	BEL = AnsiCode{"\x07", "BEL", "", 0}
	BS  = AnsiCode{"\x08", "BS", "", 0}
	HT  = AnsiCode{"\x09", "HT", "", 0}
	LF  = AnsiCode{"\x0A", "LF", "", 0}
	VT  = AnsiCode{"\x0B", "VT", "", 0}
	FF  = AnsiCode{"\x0C", "FF", "", 0}
	CR  = AnsiCode{"\x0D", "CR", "", 0}
	ESC = AnsiCode{"\x1B", "ESC", "", 0}

	ResetBold       = AnsiCode{"21", "ResetBold", "", 0}
	ResetDim        = AnsiCode{"22", "ResetDim", "", 0}
	ResetItalic     = AnsiCode{"23", "ResetItalic", "", 0}
	ResetUnderline  = AnsiCode{"24", "ResetUnderline", "", 0}
	ResetBlink      = AnsiCode{"25", "ResetBlink", "", 0}
	ResetReverse    = AnsiCode{"27", "ResetReverse", "", 0}
	ResetHidden     = AnsiCode{"28", "ResetHidden", "", 0}
	ResetStrike     = AnsiCode{"29", "ResetStrike", "", 0}
	FgBlack         = AnsiCode{"30", "FgBlack", "", 0}
	FgRed           = AnsiCode{"31", "FgRed", "", 0}
	FgGreen         = AnsiCode{"32", "FgGreen", "", 0}
	FgYellow        = AnsiCode{"33", "FgYellow", "", 0}
	FgBlue          = AnsiCode{"34", "FgBlue", "", 0}
	FgMagenta       = AnsiCode{"35", "FgMagenta", "", 0}
	FgCyan          = AnsiCode{"36", "FgCyan", "", 0}
	FgWhite         = AnsiCode{"37", "FgWhite", "", 0}
	FgDefault       = AnsiCode{"39", "FgDefault", "", 0}
	SetForeground   = AnsiCode{"38", "SetForeground", "", 0}
	BgBlack         = AnsiCode{"40", "BgBlack", "", 0}
	BgRed           = AnsiCode{"41", "BgRed", "", 0}
	BgGreen         = AnsiCode{"42", "BgGreen", "", 0}
	BgYellow        = AnsiCode{"43", "BgYellow", "", 0}
	BgBlue          = AnsiCode{"44", "BgBlue", "", 0}
	BgMagenta       = AnsiCode{"45", "BgMagenta", "", 0}
	BgCyan          = AnsiCode{"46", "BgCyan", "", 0}
	BgWhite         = AnsiCode{"47", "BgWhite", "", 0}
	BgDefault       = AnsiCode{"49", "BgDefault", "", 0}
	SetBackground   = AnsiCode{"48", "SetBackground", "", 0}
	FgBrightBlack   = AnsiCode{"90", "FgBrightBlack", "", 0}
	FgBrightRed     = AnsiCode{"91", "FgBrightRed", "", 0}
	FgBrightGreen   = AnsiCode{"92", "FgBrightGreen", "", 0}
	FgBrightYellow  = AnsiCode{"93", "FgBrightYellow", "", 0}
	FgBrightBlue    = AnsiCode{"94", "FgBrightBlue", "", 0}
	FgBrightMagenta = AnsiCode{"95", "FgBrightMagenta", "", 0}
	FgBrightCyan    = AnsiCode{"96", "FgBrightCyan", "", 0}
	FgBrightWhite   = AnsiCode{"97", "FgBrightWhite", "", 0}
	BgBrightBlack   = AnsiCode{"100", "BgBrightBlack", "", 0}
	BgBrightRed     = AnsiCode{"101", "BgBrightRed", "", 0}
	BgBrightGreen   = AnsiCode{"102", "BgBrightGreen", "", 0}
	BgBrightYellow  = AnsiCode{"103", "BgBrightYellow", "", 0}
	BgBrightBlue    = AnsiCode{"104", "BgBrightBlue", "", 0}
	BgBrightMagenta = AnsiCode{"105", "BgBrightMagenta", "", 0}
	BgBrightCyan    = AnsiCode{"106", "BgBrightCyan", "", 0}
	BgBrightWhite   = AnsiCode{"107", "BgBrightWhite", "", 0}
)

// SGR and control code helpers for terminal UI
// Usage: sgr.Bold(), sgr.Bold("off"), sgr.Underline(), sgr.Underline("off"), sgr.Fg(5), sgr.Bg(54), sgr.FgRGB(255,0,0), sgr.BgRGB(0,255,0), etc.

var Bold = struct {
	On    AnsiCode
	Off   AnsiCode
	Apply func(args ...string) AnsiCode
}{
	On:  AnsiCode{"1", "Bold", "", 0},
	Off: AnsiCode{"22", "ResetBold", "", 0},
	Apply: func(args ...string) AnsiCode {
		if len(args) == 0 || args[0] == "on" {
			return AnsiCode{"1", "Bold", "", 0}
		}
		return AnsiCode{"22", "ResetBold", "", 0}
	},
}

var Dim = struct {
	On    AnsiCode
	Off   AnsiCode
	Apply func(args ...string) AnsiCode
}{
	On:  AnsiCode{"2", "Dim", "", 0},
	Off: AnsiCode{"22", "ResetDim", "", 0},
	Apply: func(args ...string) AnsiCode {
		if len(args) == 0 || args[0] == "on" {
			return AnsiCode{"2", "Dim", "", 0}
		}
		return AnsiCode{"22", "ResetDim", "", 0}
	},
}

var Italic = struct {
	On    AnsiCode
	Off   AnsiCode
	Apply func(args ...string) AnsiCode
}{
	On:  AnsiCode{"3", "Italic", "", 0},
	Off: AnsiCode{"23", "ResetItalic", "", 0},
	Apply: func(args ...string) AnsiCode {
		if len(args) == 0 || args[0] == "on" {
			return AnsiCode{"3", "Italic", "", 0}
		}
		return AnsiCode{"23", "ResetItalic", "", 0}
	},
}

var Underline = struct {
	On    AnsiCode
	Off   AnsiCode
	Apply func(args ...string) AnsiCode
}{
	On:  AnsiCode{"4", "Underline", "", 0},
	Off: AnsiCode{"24", "ResetUnderline", "", 0},
	Apply: func(args ...string) AnsiCode {
		if len(args) == 0 || args[0] == "on" {
			return AnsiCode{"4", "Underline", "", 0}
		}
		return AnsiCode{"24", "ResetUnderline", "", 0}
	},
}

var Blink = struct {
	Slow  AnsiCode
	Rapid AnsiCode
	Off   AnsiCode
	Apply func(args ...string) AnsiCode
}{
	Slow:  AnsiCode{"5", "BlinkSlow", "", 0},
	Rapid: AnsiCode{"6", "BlinkRapid", "", 0},
	Off:   AnsiCode{"25", "ResetBlink", "", 0},
	Apply: func(args ...string) AnsiCode {
		if len(args) == 0 || args[0] == "on" || args[0] == "slow" {
			return AnsiCode{"5", "BlinkSlow", "", 0}
		}
		if args[0] == "rapid" {
			return AnsiCode{"6", "BlinkRapid", "", 0}
		}
		return AnsiCode{"25", "ResetBlink", "", 0}
	},
}

var Reverse = struct {
	On    AnsiCode
	Off   AnsiCode
	Apply func(args ...string) AnsiCode
}{
	On:  AnsiCode{"7", "Reverse", "", 0},
	Off: AnsiCode{"27", "ResetReverse", "", 0},
	Apply: func(args ...string) AnsiCode {
		if len(args) == 0 || args[0] == "on" {
			return AnsiCode{"7", "Reverse", "", 0}
		}
		return AnsiCode{"27", "ResetReverse", "", 0}
	},
}

var Hidden = struct {
	On    AnsiCode
	Off   AnsiCode
	Apply func(args ...string) AnsiCode
}{
	On:  AnsiCode{"8", "Hidden", "", 0},
	Off: AnsiCode{"28", "ResetHidden", "", 0},
	Apply: func(args ...string) AnsiCode {
		if len(args) == 0 || args[0] == "on" {
			return AnsiCode{"8", "Hidden", "", 0}
		}
		return AnsiCode{"28", "ResetHidden", "", 0}
	},
}

var Strike = struct {
	On    AnsiCode
	Off   AnsiCode
	Apply func(args ...string) AnsiCode
}{
	On:  AnsiCode{"9", "Strike", "", 0},
	Off: AnsiCode{"29", "ResetStrike", "", 0},
	Apply: func(args ...string) AnsiCode {
		if len(args) == 0 || args[0] == "on" {
			return AnsiCode{"9", "Strike", "", 0}
		}
		return AnsiCode{"29", "ResetStrike", "", 0}
	},
}

// Color helpers
var Fg = struct {
	Color func(n int) AnsiCode
	RGB   func(r, g, b int) AnsiCode
}{
	Color: func(n int) AnsiCode {
		return AnsiCode{fmt.Sprintf("38;5;%d", n), fmt.Sprintf("Fg256(%d)", n), "", 0}
	},
	RGB: func(r, g, b int) AnsiCode {
		return AnsiCode{fmt.Sprintf("38;2;%d;%d;%d", r, g, b), fmt.Sprintf("FgRGB(%d,%d,%d)", r, g, b), "", 0}
	},
}

var TextSize = func(size any) string {
	if size == "off" {
		return "\x07"
	}
	return fmt.Sprintf("\x1b]66;s=%v;", size)
}

var Bg = struct {
	Color func(n int) AnsiCode
	RGB   func(r, g, b int) AnsiCode
}{
	Color: func(n int) AnsiCode {
		return AnsiCode{fmt.Sprintf("48;5;%d", n), fmt.Sprintf("Bg256(%d)", n), "", 0}
	},
	RGB: func(r, g, b int) AnsiCode {
		return AnsiCode{fmt.Sprintf("48;2;%d;%d;%d", r, g, b), fmt.Sprintf("BgRGB(%d,%d,%d)", r, g, b), "", 0}
	},
}

// --- Cursor and screen control as struct-based helpers ---
var Cursor = struct {
	Up       func(n int) AnsiCode
	Down     func(n int) AnsiCode
	Forward  func(n int) AnsiCode
	Back     func(n int) AnsiCode
	NextLine func(n int) AnsiCode
	PrevLine func(n int) AnsiCode
	Column   func(n int) AnsiCode
	Position func(row, col int) AnsiCode
	Save     AnsiCode
	Restore  AnsiCode
	Hide     AnsiCode
	Show     AnsiCode
}{
	Up:       func(n int) AnsiCode { return AnsiCode{fmt.Sprintf("\x1b[%dA", n), "CursorUp", "", 0} },
	Down:     func(n int) AnsiCode { return AnsiCode{fmt.Sprintf("\x1b[%dB", n), "CursorDown", "", 0} },
	Forward:  func(n int) AnsiCode { return AnsiCode{fmt.Sprintf("\x1b[%dC", n), "CursorForward", "", 0} },
	Back:     func(n int) AnsiCode { return AnsiCode{fmt.Sprintf("\x1b[%dD", n), "CursorBack", "", 0} },
	NextLine: func(n int) AnsiCode { return AnsiCode{fmt.Sprintf("\x1b[%dE", n), "CursorNextLine", "", 0} },
	PrevLine: func(n int) AnsiCode { return AnsiCode{fmt.Sprintf("\x1b[%dF", n), "CursorPrevLine", "", 0} },
	Column:   func(n int) AnsiCode { return AnsiCode{fmt.Sprintf("\x1b[%dG", n), "CursorColumn", "", 0} },
	Position: func(row, col int) AnsiCode {
		return AnsiCode{fmt.Sprintf("\x1b[%d;%dH", row, col), "CursorPosition", "", 0}
	},
	Save:    AnsiCode{"\x1b[s", "SaveCursor", "", 0},
	Restore: AnsiCode{"\x1b[u", "RestoreCursor", "", 0},
	Hide:    AnsiCode{"\x1b[?25l", "HideCursor", "", 0},
	Show:    AnsiCode{"\x1b[?25h", "ShowCursor", "", 0},
}

// --- Erase/clear helpers as AnsiCode factories ---
func EraseDisplay(n int) AnsiCode { return AnsiCode{fmt.Sprintf("\x1b[%dJ", n), "EraseDisplay", "", 0} }
func EraseLine(n int) AnsiCode    { return AnsiCode{fmt.Sprintf("\x1b[%dK", n), "EraseLine", "", 0} }

// --- Device status as AnsiCode ---
func DeviceStatusReport() AnsiCode    { return AnsiCode{"\x1b[5n", "DeviceStatusReport", "", 0} }
func DeviceStatusReportCPR() AnsiCode { return AnsiCode{"\x1b[6n", "DeviceStatusReportCPR", "", 0} }

// --- Scroll as AnsiCode ---
func ScrollUp(n int) AnsiCode   { return AnsiCode{fmt.Sprintf("\x1b[%dS", n), "ScrollUp", "", 0} }
func ScrollDown(n int) AnsiCode { return AnsiCode{fmt.Sprintf("\x1b[%dT", n), "ScrollDown", "", 0} }

// --- Soft reset as AnsiCode ---
func SoftReset() AnsiCode { return AnsiCode{"\x1b[!p", "SoftReset", "", 0} }
