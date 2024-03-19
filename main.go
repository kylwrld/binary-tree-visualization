package main

import (
	"fmt"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Circle struct {
	posX, posY int32
	radius     float32
	color      rl.Color
	value      int
	left       *Circle
	right      *Circle
}

func (b *Circle) append(value int) {
	if value < b.value {
		if b.left != nil {
			b.left.append(value)
		} else {
			b.left = &Circle{value: value, posX: b.posX - 50, posY: b.posY + 50, color: rl.Blue, radius: 25}

		}
	} else {
		if b.right != nil {
			b.right.append(value)
		} else {
			b.right = &Circle{value: value, posX: b.posX + 50, posY: b.posY + 50, color: rl.Blue, radius: 25}
		}
	}
}

func Is_numeric_key_pressed() int {
	numeric_key := []int32{rl.KeyZero, rl.KeyOne, rl.KeyTwo, rl.KeyThree, rl.KeyFour, rl.KeyFive, rl.KeySix,
		rl.KeySeven, rl.KeyEight, rl.KeyNine}

	for i, key := range numeric_key {
		if rl.IsKeyPressed(key) {
			return i
		}
	}
	return -1
}

func NewCircle(posX int32, posY int32, value int, color rl.Color) *Circle {
	var c Circle = Circle{}
	c.posX = posX
	c.posY = posY
	c.value = value
	c.color = color
	c.radius = 25

	return &c
}

func (c *Circle) DrawChildren(parent *Circle) {
	rl.DrawCircle(c.posX, c.posY, float32(c.radius), c.color)
	rl.DrawLineEx(rl.Vector2{X: float32(parent.posX) + 2, Y: float32(parent.posY) + 2}, rl.Vector2{X: float32(c.posX) - 2, Y: float32(c.posY) - 2}, 5.0, rl.Black)
	value := strconv.Itoa(c.value)
	HandleText(&value, c)

	if c.left != nil {
		c.left.DrawChildren(c)
	}
	if c.right != nil {
		c.right.DrawChildren(c)
	}
}

func (c *Circle) Draw() {
	rl.DrawCircle(c.posX, c.posY, float32(c.radius), c.color)

	value := strconv.Itoa(c.value)
	HandleText(&value, c)

	if c.left != nil {
		c.left.DrawChildren(c)
	}
	if c.right != nil {
		c.right.DrawChildren(c)
	}
}

func HandleText(value *string, c *Circle) {
	if len(*value) < 2 {
		rl.DrawText(*value, c.posX-4, c.posY-10, 25, rl.White)
	} else {
		rl.DrawText(*value, c.posX-10, c.posY-10, 25, rl.White)
	}
}

func GetNumber(typing_mode bool, number *string) {
	if typing_mode {
		if key := Is_numeric_key_pressed(); key != -1 {
			if len(*number) < 2 {
				*number += strconv.Itoa(key)
			} else {
				r := []rune(*number)
				key := strconv.Itoa(key)
				new_key := []rune(key)
				r[1] = rune(new_key[0])
				*number = string(r)
			}
		}
	}
}

func GetValue(number *string) int {
	value, err := strconv.Atoi(*number)
	if err != nil {
		fmt.Println(" \"\" invalid ")
	}
	return value
}

func HandleEnterKey(first_circle_created *bool, typing_mode *bool, circle *Circle, number *string, width int32, height int32) {
	if rl.IsKeyPressed(rl.KeyEnter) {
		if *typing_mode {
			*typing_mode = false
			value := GetValue(number)
			if !*first_circle_created {
				circle.value = value
				circle.posX = (width / 2) - int32(circle.radius)
				circle.posY = (height / 2) - int32(circle.radius)
				*first_circle_created = true
			} else {
				circle.append(value)
			}
			*number = ""
		} else {
			*typing_mode = true
		}
	}

	if *typing_mode {
		number_width := rl.MeasureText(*number, 25)
		if *number == "" {
			rl.DrawText("_", width/2, height/5, 25, rl.Black)
		} else {
			rl.DrawText(*number, (width/2)-number_width, height/5, 25, rl.Black)
		}
	}

}

func main() {
	var width int32 = 1600
	var height int32 = 900
	rl.InitWindow(width, height, "Balancer")
	defer rl.CloseWindow()

	var first_circle_created bool = false
	var circle Circle = Circle{color: rl.Blue, radius: 30, posX: -20, posY: -20}
	var typing_mode bool = false
	var number string

	rl.SetTargetFPS(240)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		var fps string = strconv.Itoa(int(rl.GetFPS()))
		var mouseX string = strconv.Itoa(int(rl.GetMouseX()))
		var mouseY string = strconv.Itoa(int(rl.GetMouseY()))
		rl.DrawText("FPS: "+fps, 1400, 10, 30, rl.Black)
		rl.DrawText("MouseX: "+mouseX, 1400, 40, 30, rl.Black)
		rl.DrawText("MouseY: "+mouseY, 1400, 70, 30, rl.Black)

		HandleEnterKey(&first_circle_created, &typing_mode, &circle, &number, width, height)
		GetNumber(typing_mode, &number)

		(&circle).Draw()

		rl.EndDrawing()
	}

}
