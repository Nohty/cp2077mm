package archiver

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/mholt/archiver/v4"
)

var validDirs = []string{"archive", "bin", "r6", "red4ext", "engine"}

func ListArchiveContents(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open the archive file: %w", err)
	}
	defer file.Close()

	format, _, err := archiver.Identify(file.Name(), nil)
	if err != nil {
		return nil, fmt.Errorf("could not identify the archive format: %w", err)
	}

	extractor, ok := format.(archiver.Extractor)
	if !ok {
		return nil, fmt.Errorf("the archive format does not support extraction")
	}

	var files []string
	err = extractor.Extract(context.Background(), file, nil, func(ctx context.Context, f archiver.File) error {
		if !f.IsDir() {
			if filepath.Ext(f.NameInArchive) == ".archive" && !strings.Contains(f.NameInArchive, "/") {
				modifiedName := "/archive/pc/mod" + f.NameInArchive
				files = append(files, modifiedName)
			} else {
				pathComponents := strings.Split(f.NameInArchive, "/")
				if len(pathComponents) > 1 {
					parentDir := pathComponents[0]

					if !slices.Contains(validDirs, parentDir) {
						return fmt.Errorf("invalid directory: %s", parentDir)
					}

					files = append(files, f.NameInArchive)
				}
			}
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("could not extract the archive: %w", err)
	}

	return files, nil
}

// func ExtractArchive(path string) error {
// 	file, err := os.Open(path)
// 	if err != nil {
// 		return fmt.Errorf("could not open the archive file: %w", err)
// 	}
// 	defer file.Close()

// 	format, _, err := archiver.Identify(file.Name(), nil)
// 	if err != nil {
// 		return fmt.Errorf("could not identify the archive format: %w", err)
// 	}

// 	extractor, ok := format.(archiver.Extractor)
// 	if !ok {
// 		return fmt.Errorf("the archive format does not support extraction")
// 	}

// 	err = extractor.Extract(context.Background(), file, nil, func(ctx context.Context, f archiver.File) error {

// 		if f.IsDir() {
// 			if err := os.MkdirAll(dest, 0755); err != nil {
// 				return fmt.Errorf("could not create directory: %w", err)
// 			}
// 		} else {
// 			outFile, err := os.Create(dest)
// 			if err != nil {
// 				return fmt.Errorf("could not create file: %w", err)
// 			}
// 			defer outFile.Close()

// 			inFile, err := f.Open()
// 			if err != nil {
// 				return fmt.Errorf("could not open file: %w", err)
// 			}
// 			defer inFile.Close()

// 			if _, err := outFile.ReadFrom(inFile); err != nil {
// 				return fmt.Errorf("could not copy file: %w", err)
// 			}
// 		}

// 		events.SendLog(ctx, fmt.Sprintf("Extracted: %s", dest))

// 		return nil
// 	})

// 	if err != nil {
// 		return fmt.Errorf("could not extract the archive: %w", err)
// 	}

// 	return nil
// }
