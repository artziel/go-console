package GoConsole

type SelectComponent struct {
	Options []string
}

func (c *SelectComponent) AddOption(label string) *SelectComponent {
	if c.Options == nil {
		c.Options = []string{label}
	} else {
		c.Options = append(c.Options, label)

	}
	return c
}

func (c *SelectComponent) printOptions() {

	for i, op := range c.Options {
		Printf("<yellow>%v)</yellow> %v\n", i+1, op)
	}
}

func (c *SelectComponent) Request(message string) int {

	c.printOptions()
	Printf(message)

	for {
		v := ScanInt()
		if v <= len(c.Options) && v > 0 {
			return v
		}
	}
}

func NewSelect(options []string) SelectComponent {
	c := SelectComponent{
		Options: options,
	}

	return c
}
