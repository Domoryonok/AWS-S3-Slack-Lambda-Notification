package slack

const (
	yellow   = "#F3F636"
	blue     = "#5FCCF4"
	pink     = "#F78CE0"
	darkBlue = "#3776ab"
	gray     = "#BABDBF"
)

var colorOfExtension = map[string]string{
	".js":  yellow,
	".go":  blue,
	".cpp": pink,
	".c":   pink,
	".cc":  pink,
	".py":  darkBlue,
}

func ExtensionToColor(extension string) (color string) {
	if color, ok := colorOfExtension[extension]; ok {
		return color
	} else {
		return gray
	}
}