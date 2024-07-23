package list

import "fmt"

const (
	ColorDefault = "\x1b[39m"

	ColorPurple   = "\x1b[38;5;128m" //#820AD1
	ColorLavender = "\x1b[38;5;183m" //#A67BF2
	ColorLime     = "\x1b[38;5;190m" //#C0D72F

)

func purple(s string) string {
	return fmt.Sprintf("%s%s%s", ColorPurple, s, ColorDefault)
}

func lavender(s string) string {
	return fmt.Sprintf("%s%s%s", ColorLavender, s, ColorDefault)
}

func lime(s string) string {
	return fmt.Sprintf("%s%s%s", ColorLime, s, ColorDefault)
}
