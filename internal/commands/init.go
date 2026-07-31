package commands

import (
	"fmt"
	"os"
	"path/filepath"
)

func Init() error {
	gitDir := ".mygit"

	directories := []string{
		filepath.Join(gitDir, "objects"),
		filepath.Join(gitDir, "refs", "heads"),
	}

	for _, directory := range directories {
		if err := os.MkdirAll(directory, 0755); err != nil {
			return fmt.Errorf("create directory %s: %w", directory, err)
		}
	}

	headPath := filepath.Join(gitDir, "HEAD")
	headContent := []byte("ref: refs/heads/main\n")

	if err := os.WriteFile(headPath, headContent, 0644); err != nil {
		return fmt.Errorf("write HEAD: %w", err)
	}

	fmt.Printf("Initialized empty Mygit repository in %s\n", gitDir)

	return nil
}
