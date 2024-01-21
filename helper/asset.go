package helper

import (
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
)

func LoadResourceFromPath(p string) (fyne.Resource, error) {
	data, err := os.ReadFile(p)
	if err != nil {
		return nil, err
	}
	return fyne.NewStaticResource(filepath.Base(p), data), nil
}
