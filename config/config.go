package config

type configuration struct {
	Bold          int
	ColourOutline int
	ColourBomb    int
}

// https://www.tutorialspoint.com/how-to-output-colored-text-to-a-linux-terminal
var Config = configuration{
	Bold:          1,
	ColourOutline: 30,
	ColourBomb:    36,
}
