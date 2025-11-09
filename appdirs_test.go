package appdirs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAppDir(t *testing.T) {
	_, err := NewAppDirs("", "")
	assert.NotNil(t, err)
	assert.Errorf(t, err, "app name cannot be an empty string")
}

func TestAppDirPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("this code should have recovered")
		}
	}()

	appdirs := AppDirs{}
	// this should panic
	appdirs.UserAppDir()
}
