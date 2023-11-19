package jett

import (
	"fmt"
	"os"

	"github.com/gookit/color"
)

// Sends to console an error log
// If fatal is equal to `true` then `os.Exit(1)` is called
func Error(str string, fatal ...bool) {
	info := color.Hex("0xFF4245", true).Sprintf
	text := color.Gray.Render

	fmt.Println(info(" ERROR ") + " " + text(str))

	if fatal[0] {
		os.Exit(1)
	}
}

// Sends to console a warn log
func Warn(str string) {
	info := color.Hex("0xFACA16", true).Sprintf
	text := color.Gray.Render

	fmt.Println(info(" WARN ") + " " + text(str))
}

// Sends to console an info log
func Info(str string) {
	info := color.Hex("0x1167B1", true).Sprintf
	text := color.Gray.Render

	fmt.Println(info(" INFO ") + " " + text(str))
}

// Sends to console a success log
func Success(str string) {
	info := color.Hex("0x06A94D", true).Sprintf
	text := color.Gray.Render

	fmt.Println(info(" SUCCESS ") + " " + text(str))
}
