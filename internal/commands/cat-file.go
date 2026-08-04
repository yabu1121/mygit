package commands

import "fmt"

type CatFileOption string
const (
	CatFilePretty CatFileOption = "-p"
	CatFileSize CatFileOption = "-s"
	CatFileType CatFileOption = "-t"
	CatFileExists CatFileOption = "-e"
)

func ParseCatFileOption(value string) (CatFileOption, error) {
	option := CatFileOption(value)

	switch option {
	case CatFilePretty, CatFileSize, CatFileType, CatFileExists:
		return option, nil
	default:
		return "", fmt.Errorf("unknown cat-file option: %s", value)
	}
}

func CatFile(option CatFileOption, objectID string) error {
	switch option {
	case CatFilePretty:
	case CatFileSize:
	case CatFileType:
	case CatFileExists:
	default:
		return fmt.Errorf("unknown cat-file options %s", option)
	}
	return nil
}
