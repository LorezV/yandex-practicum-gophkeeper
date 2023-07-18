package build

import "fmt"

var (
	BuildVersion string //nolint:gochecknoglobals // global buildVersion
	BuildDate    string //nolint:gochecknoglobals // global buildDate
	BuildCommit  string //nolint:gochecknoglobals // global buildCommit
)

const notGiven = "N/A"

func CheckBuild() {
	if BuildVersion == "" {
		BuildVersion = notGiven
	}

	if BuildDate == "" {
		BuildDate = notGiven
	}

	if BuildCommit == "" {
		BuildCommit = notGiven
	}
}

func PrintBulidInfo() {
	fmt.Printf("Build version: %s\n", BuildVersion) //nolint:forbidigo // print build info
	fmt.Printf("Build date: %s\n", BuildDate)       //nolint:forbidigo // print build info
	fmt.Printf("Build commit: %s\n", BuildCommit)   //nolint:forbidigo // print build info
}
