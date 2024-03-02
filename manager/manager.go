package manager

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func InstallModArchiver(files []string, destination string) error {
	return fmt.Errorf("not implemented")
}

func InstallMod(file, destination string) error {
	if _, err := os.Stat(destination); err == nil {
		// return fmt.Errorf("destination already exists: %s", destination)
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("could not check destination: %w", err)
	}

	destDir := filepath.Dir(destination)
	if err := os.MkdirAll(destDir, os.ModePerm); err != nil {
		return fmt.Errorf("could not create destination directory: %w", err)
	}

	srcFile, err := os.Open(file)
	if err != nil {
		return fmt.Errorf("could not open source file: %w", err)
	}
	defer srcFile.Close()

	destFile, err := os.Create(destination)
	if err != nil {
		return fmt.Errorf("could not create destination file: %w", err)
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return fmt.Errorf("could not copy file: %w", err)
	}

	return nil
}

func UninstallMod(mod string) error {
	if err := os.Remove(mod); err != nil {
		return fmt.Errorf("could not remove mod: %w", err)
	}

	return nil
}
