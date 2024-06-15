package main

import (
	"fmt"
	"io"
	"os"
	"plugin"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var drawFunc func()

func loadPlugin() error {
	tempFile, err := copyPlugin()
	if err != nil {
		panic(err)
	}

	defer os.Remove(tempFile)

	p, err := plugin.Open(tempFile)
	if err != nil {
		return err
	}

	sym, err := p.Lookup("Draw")
	if err != nil {
		return err
	}

	drawFunc = sym.(func())

	fmt.Println("Plugin loaded successfully")
	return nil
}

func main() {
	rl.InitWindow(800, 600, "Hot Reloading Example")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	loadPlugin()

	for !rl.WindowShouldClose() {
		if rl.IsKeyPressed(rl.KeyR) {
			loadPlugin()
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		if drawFunc != nil {
			drawFunc()
		}

		rl.EndDrawing()
	}
}

func copyPlugin() (string, error) {
	srcFile, err := os.Open("plugin.so")
	if err != nil {
		return "", err
	}
	defer srcFile.Close()

	tempFile, err := os.CreateTemp("", "plugin-*.so")
	if err != nil {
		return "", err
	}
	defer tempFile.Close()

	_, err = io.Copy(tempFile, srcFile)
	if err != nil {
		return "", err
	}

	return tempFile.Name(), nil
}
