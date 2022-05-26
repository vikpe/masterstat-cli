package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vikpe/udphelper"
)

type AppRunner struct {
	app string
}

func (ar AppRunner) run(appArgs string) string {
	argsStr := strings.TrimSpace(fmt.Sprintf("%s %s", ar.app, appArgs))
	argsArr := strings.Split(argsStr, " ")
	return captureOutput(func() { run(argsArr) })
}

var app = AppRunner{
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
		assert.Equal(t, app.run(""), helpText)
	})

	t.Run("Help", func(t *testing.T) {
		assert.Equal(t, app.run("help"), helpText)
		assert.Equal(t, app.run("--help"), helpText)
		assert.Equal(t, app.run("-h"), helpText)
	})
}

func TestServerAddresses(t *testing.T) {
	t.Run("UDP request error", func(t *testing.T) {
		output := app.run("foo:666")
		assert.Contains(t, output, "ERROR:")
	})

	t.Run("Get server addresses", func(t *testing.T) {
		go func() {
			responseBody := []byte{
				0xff, 0xff, 0xff, 0xff, 0x64, 0x0a, // header
				0x42, 0x45, 0x65, 0x94, 0x6b, 0x6c, //  server 1
				0xf5, 0x49, 0x6f, 0x6b, 0x6d, 0xc8, //  server 2
			}
			udphelper.New(":8000").Respond(responseBody)
		}()
		time.Sleep(10 * time.Millisecond)

		output := app.run(":8000")
		expectedServers := []string{
			"245.73.111.107:28104",
			"66.69.101.148:27500",
		}
		expectedOutput := strings.Join(expectedServers, "\n") + "\n"
		assert.Equal(t, expectedOutput, output)
	})
}
