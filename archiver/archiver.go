package archiver

import (
	"context"
	"cp2077mm/events"
	"cp2077mm/manager"
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

func ExtractArchive(ctx context.Context, path, destination string) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("could not open the archive file: %w", err)
	}
	defer file.Close()

	format, _, err := archiver.Identify(file.Name(), nil)
	if err != nil {
		return fmt.Errorf("could not identify the archive format: %w", err)
	}

	extractor, ok := format.(archiver.Extractor)
	if !ok {
		return fmt.Errorf("the archive format does not support extraction")
	}

	var files []string
	err = extractor.Extract(context.Background(), file, nil, func(c context.Context, f archiver.File) error {
		if !f.IsDir() {
			var modifiedName string
			if filepath.Ext(f.NameInArchive) == ".archive" && !strings.Contains(f.NameInArchive, "/") {
				modifiedName = filepath.Join(destination, "archive/pc/mod", f.NameInArchive)
			} else {
				pathComponents := strings.Split(f.NameInArchive, "/")

				if len(pathComponents) > 1 {
					parentDir := pathComponents[0]

					if !slices.Contains(validDirs, parentDir) {
						return fmt.Errorf("invalid directory: %s", parentDir)
					}

					modifiedName = filepath.Join(destination, f.NameInArchive)
				}
			}

			if modifiedName != "" {
				reader, err := f.Open()
				if err != nil {
					return fmt.Errorf("could not open file: %w", err)
				}
				defer reader.Close()

				err = manager.InstallModArchiver(reader, modifiedName)
				if err != nil {
					return fmt.Errorf("could not install mod: %w", err)
				}

				events.SendLog(ctx, fmt.Sprintf("Created file: %s", modifiedName))

				files = append(files, modifiedName)
			}
		}

		return nil
	})

	if err != nil {
		for _, file := range files {
			err = manager.UninstallMod(file)

			if err != nil {
				return fmt.Errorf("could not uninstall mod: %w", err)
			}
		}
	}

	return nil
}
