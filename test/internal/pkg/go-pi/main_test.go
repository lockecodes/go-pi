package go_pi

import (
	"github.com/stretchr/testify/assert"
	"gitlab.com/slocke716/go-pi/internal/pkg/go-pi"
	"io/ioutil"
	"os"
	"testing"
)

func Test_loadLayout(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Test Loading Valid File",
			args{"testdata/valid_layout.yaml"},
			false,
		},
		{
			"Test Loading Real File",
			args{"testdata/layout.yaml"},
			false,
		},
		{
			"Test Loading Invalid File",
			args{"testdata/invalid_layout.yaml"},
			true,
		},
		{
			"Test File Not Found",
			args{"testdata/non_existent.yaml"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := gopi.LoadLayout(tt.args.filename)
			hasErr := err != nil
			assert.Equal(t, tt.wantErr, hasErr)
		})
	}
}

func TestShouldCreate(t *testing.T) {
	testCases := []struct {
		dir             gopi.Directory
		web             bool
		kitchen         bool
		expected_create bool
	}{
		{
			gopi.Directory{Flags: []string{"required"}},
			false,
			false,
			true,
		},
		{
			gopi.Directory{Flags: []string{"web"}},
			true,
			false,
			true,
		},
		{
			gopi.Directory{Flags: []string{"kitchen"}},
			false,
			true,
			true,
		},
	}
	for _, testCase := range testCases {
		result := gopi.ShouldCreate(testCase.dir, testCase.web, testCase.kitchen)
		if result != testCase.expected_create {
			t.Errorf("ShouldCreate(%+v, %v, %v) = %v; want %v", testCase.dir, testCase.web, testCase.kitchen, result, testCase.expected_create)
		}
	}
}

func TestCreateDirectory(t *testing.T) {
	dir := gopi.Directory{
		Description: "Test Directory",
	}
	cleanUp := setupTestDirectory(t)
	gopi.CreateDirectory("testdir", dir)
	defer cleanUp("testdir")

	_, err := os.Stat("testdir")
	if os.IsNotExist(err) {
		t.Errorf("Directory was not created")
	}
}

func TestCreateFiles(t *testing.T) {
	files := map[string]string{
		"readme.md": "This is a test file",
	}

	cleanUp := setupTestDirectory(t)
	gopi.CreateFiles("testdir", files)
	defer cleanUp("testdir")

	for filename := range files {
		_, err := os.Stat("testdir/" + filename)
		if os.IsNotExist(err) {
			t.Errorf("File %v was not created", filename)
		}
	}
}

func TestCreateAndWriteFile(t *testing.T) {
	filepath := "testfile.txt"
	content := "This is a test file"

	gopi.CreateAndWriteFile(filepath, content)

	defer os.Remove(filepath)

	bytes, _ := ioutil.ReadFile(filepath)
	if string(bytes) != content {
		t.Errorf("Expected: %v, Got: %v", content, string(bytes))
	}
}

func setupTestDirectory(t *testing.T) func(dir string) {
	t.Helper()
	cleanup := func(dir string) {
		os.RemoveAll(dir)
	}
	return cleanup
}
