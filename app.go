package main

import (
	"context"
	"cp2077mm/archiver"
	"cp2077mm/config"
	"cp2077mm/events"
	"cp2077mm/manager"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx    context.Context
	config *config.ConfigStore
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	var err error
	a.config, err = config.NewConfigStore()

	if err != nil {
		panic(err)
	}
}

func (a *App) GetMods() []config.ModConfig {
	mods, err := a.config.Mods()
	if err != nil {
		return []config.ModConfig{}
	}

	return mods
}

func (a *App) GetGameDir() string {
	c, err := a.config.Config()
	if err != nil {
		return ""
	}

	return c.GameDir
}

func (a *App) SetGameDir(dir string) {
	c, err := a.config.Config()
	if err != nil {
		events.SendError(a.ctx, fmt.Sprintf("Could not load config: %s", err))
		return
	}

	if dir == "" {
		events.SendError(a.ctx, "Game dir cannot be empty")
		return
	}

	c.GameDir = dir

	err = a.config.Save(c)
	if err != nil {
		events.SendError(a.ctx, fmt.Sprintf("Could not set game dir: %s", err))
		return
	}

	events.SendRefreshMods(a.ctx)
}

func (a *App) OpenFileDialog() string {
	result, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select File",
		Filters: []runtime.FileFilter{
			{DisplayName: "Cyberpunk 2077 Mod", Pattern: "*.zip;*.rar;*.7z;*.archive"},
			{DisplayName: "All Files", Pattern: "*"},
		},
	})

	if err != nil {
		return ""
	}

	return result
}

func (a *App) OpenFolderDialog() string {
	result, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select Folder",
	})

	if err != nil {
		return ""
	}

	return result
}

func (a *App) AddMod(name, path string) {
	c, err := a.config.Config()
	if err != nil {
		events.SendError(a.ctx, fmt.Sprintf("Could not load config: %s", err))
		return
	}

	if c.GameDir == "" {
		events.SendError(a.ctx, "Game dir is not set")
		return
	}

	if strings.HasSuffix(path, ".archive") {
		fileName := filepath.Base(path)
		file := fmt.Sprintf("archive/pc/mod/%s", fileName)
		destination := filepath.Join(c.GameDir, file)

		modConfig := config.ModConfig{Name: name, Files: []string{file}}

		err = a.config.Validate(modConfig)
		if err != nil {
			events.SendError(a.ctx, fmt.Sprintf("Mod is already installed: %s", err))
			return
		}

		err = manager.InstallMod(path, destination)
		if err != nil {
			events.SendError(a.ctx, fmt.Sprintf("Could not install mod: %s", err))
			return
		}

		err = a.config.AddMod(modConfig)
		if err != nil {
			events.SendError(a.ctx, fmt.Sprintf("Could not add mod to config: %s", err))
			return
		}

		events.SendLog(a.ctx, fmt.Sprintf("Created file: %s", destination))

		events.SendSuccess(a.ctx, fmt.Sprintf("Successfully installed mod %s", name))
		events.SendRefreshMods(a.ctx)
	} else {
		contents, err := archiver.ListArchiveContents(path)
		if err != nil {
			events.SendError(a.ctx, fmt.Sprintf("Could not list archive contents: %s", err))
			return
		}

		modConfig := config.ModConfig{Name: name, Files: contents}

		err = a.config.Validate(modConfig)
		if err != nil {
			events.SendError(a.ctx, fmt.Sprintf("Mod is already installed: %s", err))
			return
		}

		installedFiles := []string{}

		for _, file := range contents {
			destination := filepath.Join(c.GameDir, file)
			err = manager.InstallMod(path, destination)
			if err != nil {
				events.SendError(a.ctx, fmt.Sprintf("Could not install mod: %s", err))
				break
			}

			events.SendLog(a.ctx, fmt.Sprintf("Created file: %s", destination))

			installedFiles = append(installedFiles, destination)
		}

		if err != nil {
			events.SendError(a.ctx, fmt.Sprintf("Could not install mod: %s", err))

			for _, file := range installedFiles {
				err = manager.UninstallMod(file)
				if err != nil {
					events.SendError(a.ctx, fmt.Sprintf("Could not remove file: %s", err))
				}

				events.SendLog(a.ctx, fmt.Sprintf("Removed file: %s", file))
			}

			return
		}

		err = a.config.AddMod(modConfig)
		if err != nil {
			events.SendError(a.ctx, fmt.Sprintf("Could not add mod to config: %s", err))
			return
		}

		events.SendSuccess(a.ctx, fmt.Sprintf("Successfully installed mod %s", name))
		events.SendRefreshMods(a.ctx)
	}
}

func (a *App) RemoveMod(name string) {
	modConfig, err := a.config.Mod(name)
	if err != nil {
		events.SendError(a.ctx, fmt.Sprintf("Could not load mod: %s", err))
		return
	}

	c, err := a.config.Config()
	if err != nil {
		events.SendError(a.ctx, fmt.Sprintf("Could not load config: %s", err))
		return
	}

	for _, file := range modConfig.Files {
		err = manager.UninstallMod(filepath.Join(c.GameDir, file))
		if err != nil {
			events.SendError(a.ctx, fmt.Sprintf("Could not remove file: %s", err))
		}

		events.SendLog(a.ctx, fmt.Sprintf("Removed file: %s", file))
	}

	err = a.config.RemoveMod(name)
	if err != nil {
		events.SendError(a.ctx, fmt.Sprintf("Could not remove mod from config: %s", err))
		return
	}

	events.SendSuccess(a.ctx, fmt.Sprintf("Successfully removed mod %s", name))
	events.SendRefreshMods(a.ctx)
}
