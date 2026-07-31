package commands

import (
	"crypto/sha1"
	"fmt"
	"os"
)

func HashObject(filePath string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("read file %s: %w", filePath, err)
	}

	header := fmt.Sprintf("blob %d\x00", len(content))

	objectData := append([]byte(header), content...)

	hash := sha1.Sum(objectData)
	fmt.Printf("%x\n", hash)

	return nil
}
