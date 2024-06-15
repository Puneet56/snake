package main

import rl "github.com/gen2brain/raylib-go/raylib"

var x int32 = 0
var y int32 = 0

func Draw() {
	rl.DrawRectangle(200, 200, 200, 200, rl.Pink)
	rl.DrawText("Hello, World!", 250, 250, 20, rl.Black)

	rl.DrawCircle(x, y, 50, rl.Orange)

	x++
	y++

	if x > 800 {
		x = 0
	}

	if y > 600 {
		y = 0
	}
}
