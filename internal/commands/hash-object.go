package commands

import (
	"compress/zlib"
	"crypto/sha1"
	"fmt"
	"os"
	"path/filepath"
)

func HashObject(filePath string, write bool) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("read file %s: %w", filePath, err)
	}

	header := fmt.Sprintf("blob %d\x00", len(content))

	objectData := append([]byte(header), content...)

	hash := sha1.Sum(objectData)
	hashString := fmt.Sprintf("%x", hash)

	if write {

		dirName := hashString[:2]
		fileName := hashString[2:]

		objectDir := filepath.Join(".mygit", "objects", dirName)
		objectPath := filepath.Join(objectDir, fileName)

		if err := os.MkdirAll(objectDir, 0755); err != nil {
			return fmt.Errorf("create object directory: %w", err)
		}

		file, err := os.Create(objectPath)
		if err != nil {
			return fmt.Errorf("create object file: %w", err)
		}
		defer file.Close()

		zw := zlib.NewWriter(file)
		if _, err := zw.Write(objectData); err != nil {
			return fmt.Errorf("compress object data: %w", err)
		}

		if err := zw.Close(); err != nil {
			return fmt.Errorf("close zlib writer: %w", err)
		}
	}
	fmt.Println(hashString)
	return nil
}
