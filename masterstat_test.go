package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestRunner struct {
	app string
}

func (ar TestRunner) run(appArgs string) string {
	argsStr := strings.TrimSpace(fmt.Sprintf("%s %s", ar.app, appArgs))
	argsArr := strings.Split(argsStr, " ")
	return captureOutput(func() { run(argsArr) })
}

var runner = TestRunner{
	app: "masterstat",
}

func captureOutput(f func()) string {
	rescueStderr := os.Stderr
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	os.Stderr = w

	f()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stderr = rescueStderr
	os.Stdout = rescueStdout

	return string(out)
}

func TestHelp(t *testing.T) {
	helpText := `masterstat [__VERSION__]
Fetch server addresses from QuakeWorld master servers.

  Usage:   masterstat [<address> ...]
Example:   masterstat master.quakeworld.nu:27000 qwmaster.ocrana.de:27000
`

	t.Run("No args", func(t *testing.T) {
		assert.Equal(t, runner.run(""), helpText)
	})

	t.Run("Help", func(t *testing.T) {
		assert.Equal(t, runner.run("help"), helpText)
		assert.Equal(t, runner.run("--help"), helpText)
		assert.Equal(t, runner.run("-h"), helpText)
	})
}

func TestError(t *testing.T) {
	output := runner.run("foo:666")
	assert.Contains(t, output, "ERROR:")
}
