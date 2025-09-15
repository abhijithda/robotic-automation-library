package robot

/*
### Objective

Imagine you work in Thoughtful’s robotic automation factory, and your objective is to write a function for one of its robotic arms that will dispatch the packages to the correct stack according to their volume and mass.

### Rules

Sort the packages using the following criteria:

- A package is **bulky** if its volume (Width x Height x Length) is greater than or equal to 1,000,000 cm³ or when one of its dimensions is greater or equal to 150 cm.
- A package is **heavy** when its mass is greater or equal to 20 kg.

You must dispatch the packages in the following stacks:

- **STANDARD**: standard packages (those that are not bulky or heavy) can be handled normally.
- **SPECIAL**: packages that are either heavy or bulky can't be handled automatically.
- **REJECTED**: packages that are **both** heavy and bulky are rejected.

### Implementation

Implement the function **`sort(width, height, length, mass)`** (units are centimeters for the dimensions and kilogram for the mass). This function must return a string: the name of the stack where the package should go.

### Submission Guidance

1. **Time Management**: Allocate no more than 30 minutes to complete this challenge.
2. **Programming Language**: You may use any programming language you're comfortable with. This is an opportunity to showcase your skills in a language you are proficient in.
3. **Submission Format**:
    - **Option 1**: Submit a public GitHub repository with clear README instructions.
    - **Option 2 (Preferred)**: Host your solution on an online IDE like [Repl.it](http://repl.it/) or CodePen for immediate code review. Ensure the link is accessible for direct execution.
4. **Evaluation Criteria**: Submissions will be assessed on:
    - Correct sorting logic.
    - Code quality.
    - Handling edge cases and inputs.
    - Test coverage.
*/

/*
sort: sorts the packages according to their volume and mass, and

	returns the name of the stack where the package should go.

input:

	width: the width of the package
	height: the height of the package
	length: the length of the package
	mass: the mass of the package

output:

	the name of the stack where the package should go (STANDARD, SPECIAL, REJECTED)
*/

import (
	"flag"
	"log"
	"os"
	"path/filepath"
)

const (
	STANDARD = "STANDARD"
	SPECIAL  = "SPECIAL"
	REJECTED = "REJECTED"

	BulkyDimensionThreshold = 150
	BulkyVolumeThreshold    = 1000000
	HeavyMassThreshold      = 20
)

type cmdOptions struct {
	Width  int
	Height int
	Length int
	Mass   int

	LogFile string
}

var CmdOptions cmdOptions

func RegisterCmdOptions() {
	// Register the command line options for the sort command.
	flag.IntVar(&CmdOptions.Width, "width", 1, "the width of the package in centimeters")
	flag.IntVar(&CmdOptions.Height, "height", 1, "the height of the package in centimeters")
	flag.IntVar(&CmdOptions.Length, "length", 1, "the length of the package in centimeters")
	flag.IntVar(&CmdOptions.Mass, "mass", 1, "the mass of the package in kilograms")

	PROGRAM_NAME := filepath.Base(os.Args[0])
	// defaultLogFile := filepath.Dir(os.Args[0]) + string(os.PathSeparator) + PROGRAM_NAME + ".log"
	defaultLogFile := filepath.Join(os.TempDir(), PROGRAM_NAME+".log")
	flag.StringVar(&CmdOptions.LogFile, "log-file", defaultLogFile, "the log file to write logs to (default: stdout)")
}

func Sort(width, height, length, mass int) string {
	log.Printf("Entering Sort(width=%d, height=%d, length=%d, mass=%d)...", width, height, length, mass)
	res := STANDARD
	defer func() {
		log.Printf("Sort() returning %s.", res)
	}()

	bulky := isBulky(width, height, length)
	heavy := isHeavy(mass)
	log.Printf("isBulky=%t, isHeavy=%t", bulky, heavy)

	if bulky && heavy {
		res = REJECTED
	} else if bulky || heavy {
		res = SPECIAL
	}
	return res
}

func isBulky(width, height, length int) bool {
	return width >= BulkyDimensionThreshold ||
		height >= BulkyDimensionThreshold ||
		length >= BulkyDimensionThreshold ||
		width*height*length >= BulkyVolumeThreshold
}

func isHeavy(mass int) bool {
	return mass >= HeavyMassThreshold
}
