package go_pi

import (
	"github.com/urfave/cli/v2"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	gopi "gitlab.com/slocke716/go-pi/internal/pkg/go-pi"
)

func TestMain(m *testing.M) {
	// Perform setup if necessary, e.g., initializing resources

	// Run the tests
	exitVal := m.Run()

	// Perform teardown if necessary, e.g., releasing resources

	// Exit with the appropriate exit code
	os.Exit(exitVal)
}

func newTestApp() *cli.App {
	app := cli.NewApp()
	app.Name = "Go Project Interface"
	app.Usage = "Setup Go Projects with a simple YAML layout"
	app.Commands = []*cli.Command{
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
		},
	}
	return app
}

func TestNewApp(t *testing.T) {
	t.Run("NewApp", func(t *testing.T) {
		app := newTestApp()
		require.NotNil(t, app)
		assert.Equal(t, "Go Project Interface", app.Name)
		assert.Equal(t, "Setup Go Projects with a simple YAML layout", app.Usage)
		assert.NotNil(t, app.Commands)
		assert.Equal(t, 1, len(app.Commands))
		testCommand := app.Commands[0]
		assert.NotNil(t, testCommand)
		assert.Equal(t, "generate", testCommand.Name)
		assert.Equal(t, "Generate directories and files", testCommand.Usage)
	})
}

func TestCommandFlags(t *testing.T) {
	t.Run("CommandFlags", func(t *testing.T) {
		app := newTestApp()
		require.NotNil(t, app)
		testCommand := app.Commands[0]
		require.NotNil(t, testCommand)

		expectedFlags := []cli.Flag{
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
		}

		for i, expectedFlag := range expectedFlags {
			actualFlag := testCommand.Flags[i]

			assert.Equal(t, expectedFlag.Names(), actualFlag.Names())

			switch ef := expectedFlag.(type) {
			case *cli.BoolFlag:
				af := actualFlag.(*cli.BoolFlag)
				assert.Equal(t, ef.Usage, af.Usage)
			case *cli.StringFlag:
				af := actualFlag.(*cli.StringFlag)
				assert.Equal(t, ef.Usage, af.Usage)
				assert.Equal(t, ef.Value, af.Value)
			}
		}
	})
}
