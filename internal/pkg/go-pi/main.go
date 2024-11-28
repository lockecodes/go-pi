package gopi

import (
	"fmt"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"log"
	"os"
	"path/filepath"
)

// DefaultLayoutFile defines the default layout configuration file name used in the application.
//
// DefaultOutputDirectory specifies the default directory path for the application’s output.
//
// DirectoryMode sets the default permission mode for directories created by the application.
const (
	DefaultLayoutFile      = "configs/layout.yaml"
	DefaultOutputDirectory = "examples/_your_app_"
	DirectoryMode          = os.ModePerm
)

// Layout defines a structure representing a set of directories and files configuration.
// Directories is a map where each key is a directory name, and the value is a Directory struct.
// Files is a map where each key is a filename, and the value is a file path or content.
type Layout struct {
	Directories map[string]Directory `koanf:"directories"`
	Files       map[string]string    `koanf:"__files__,omitempty"`
}

// Directory represents a file system directory with possible subdirectories and files.
type Directory struct {
	Description    string               `koanf:"description"`
	Flags          []string             `koanf:"flags,omitempty"`
	SubDirectories map[string]Directory `koanf:"subDirectories,omitempty"`
	Files          map[string]string    `koanf:"__files__,omitempty"`
}

// LoadLayout reads a YAML file specified by filename and unmarshals its content into a Layout struct.
// It returns the loaded Layout and any error encountered during the file reading or unmarshalling process.
func LoadLayout(filename string) (Layout, error) {
	log.Printf("Loading layout from %s", filename)
	var k = koanf.New(".")
	log.Printf("Parsing YAML")
	if err := k.Load(file.Provider(filename), yaml.Parser()); err != nil {
		return Layout{}, fmt.Errorf("error reading %s: %w", filename, err)
	}
	var layout Layout
	log.Printf("Unmarshalling YAML")
	if err := k.Unmarshal("", &layout); err != nil {
		return Layout{}, fmt.Errorf("error parsing YAML: %w", err)
	}
	return layout, nil
}

// ShouldCreate determines if a directory should be created based on its flags and the input flags for web and kitchen.
// Returns true if the directory has a "required" flag, or if it matches the "web" or "kitchen" flags and the input flags.
func ShouldCreate(dir Directory, web, kitchen bool) bool {
	for _, flag := range dir.Flags {
		switch flag {
		case "required":
			return true
		case "web":
			if web {
				return true
			}
		case "kitchen":
			if kitchen {
				return true
			}
		}
	}
	return false
}

// CreateDirectory creates a directory with the specified name and populates it with files and subdirectories based on the provided Directory struct.
func CreateDirectory(name string, dir Directory) {
	log.Printf("Creating directory %s", name)
	if err := os.MkdirAll(name, DirectoryMode); err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}
	log.Printf("Writing README.md")
	readmePath := filepath.Join(name, "README.md")
	if err := os.WriteFile(readmePath, []byte(dir.Description), DirectoryMode); err != nil {
		fmt.Println("Error writing README.md:", err)
		return
	}
	CreateFiles(name, dir.Files)

	for subName, subDir := range dir.SubDirectories {
		CreateDirectory(filepath.Join(name, subName), subDir)
	}
}

// CreateFiles creates multiple files with specified content in a given directory. It ensures the directory exists before
// creating files. If the directory does not exist, it will be created with the default permissions. Each file's content
// is written to the specified file path within the directory.
func CreateFiles(dirPath string, files map[string]string) {
	log.Printf("Creating files in %s", dirPath)
	if err := os.MkdirAll(dirPath, DirectoryMode); err != nil {
		fmt.Println("Error ensuring directory exists:", err)
		return
	}
	for fileName, content := range files {
		CreateAndWriteFile(filepath.Join(dirPath, fileName), content)
	}
}

// CreateAndWriteFile writes the provided content to a file at the specified filePath.
// If the file does not exist, it will be created with default directory permissions.
// An error is logged if the file cannot be written.
func CreateAndWriteFile(filePath, content string) {
	log.Printf("Writing %s", filePath)
	if err := os.WriteFile(filePath, []byte(content), DirectoryMode); err != nil {
		fmt.Println("Error writing file:", filePath, "-", err)
	}
}
