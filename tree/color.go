package cmd

var (
	dirColorFmt = "%s"
	docColorFmt = "%s"
)

func setColor() {
	dirColorFmt = "\x1b[33m%s\x1b[0m"
	docColorFmt = "\x1b[36m%s\x1b[0m"
}
