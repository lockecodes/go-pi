package main

import (
	"fmt"
	gopi "gitlab.com/slocke716/go-pi/internal/pkg/go-pi"
	"log"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

// main is the entry point of the application. It sets up the CLI interface with app configuration, commands, and flags.
func main() {
	app := &cli.App{
		Name:  "Go Project Interface",
		Usage: "Setup Go Projects with a simple YAML layout",
		Commands: []*cli.Command{
			{
				Name:  "generate",
				Usage: "Generate directories and files",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:  "web",
						Usage: "Include directories flagged with 'web'",
					},
					&cli.BoolFlag{
						Name:  "kitchen",
						Usage: "Include directories flagged with 'kitchen'",
					},
					&cli.StringFlag{
						Name:    "layout",
						Aliases: []string{"l"},
						Usage:   "Path to the layout YAML file",
						Value:   gopi.DefaultLayoutFile,
					},
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Usage:   "Directory to output the generated structure",
						Value:   gopi.DefaultOutputDirectory,
					},
				},
				Action: func(c *cli.Context) error {
					layoutFile, outputDirectory := c.String("layout"), c.String("output")
					layout, err := gopi.LoadLayout(layoutFile)
					if err != nil {
						return fmt.Errorf("error loading layout: %w", err)
					}
					web, kitchen := c.Bool("web"), c.Bool("kitchen")

					if err := os.MkdirAll(outputDirectory, gopi.DirectoryMode); err != nil {
						return fmt.Errorf("error creating output directory: %w", err)
					}

					gopi.CreateFiles(outputDirectory, layout.Files)

					for dirName, dir := range layout.Directories {
						log.Printf("Checking flags for %s", dirName)
						if gopi.ShouldCreate(dir, web, kitchen) {
							gopi.CreateDirectory(filepath.Join(outputDirectory, dirName), dir)
						}
					}
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
