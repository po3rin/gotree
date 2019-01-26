package gotree

// Level of depth
var (
	depthLimit    int
	dirColorFmt   = "%s"
	docColorFmt   = "%s"
	directoryOnly bool
)

func setLevel(level int) {
	depthLimit = level
}

func setDirectoryOnly() {
	directoryOnly = true
}

func setColor() {
	dirColorFmt = "\x1b[33m%s\x1b[0m"
	docColorFmt = "\x1b[36m%s\x1b[0m"
}
