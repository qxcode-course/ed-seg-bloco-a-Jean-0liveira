package main

import (
	"fmt"
	"math/rand"
)

func randInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func main() {
	pen := NewPen(500, 500)   // cria um canvas de 500 de largura por 500 de altura
	pen.SetRGB(0, 0, 0)     // muda a cor do pincel para vermelho
	pen.SetPosition(250, 500) // move o pincel para x 250, y 500
	pen.SetHeading(90)        // coloca o pincel apontando para cima
	pen.Walk(200)             // anda 100 pixels
	pen.Left(30)
	pen.Walk(100)
	pen.Down()
	pen.Walk(100)
	pen.Up()
	pen.Right(30)
	pen.Walk(100)
	for range 10 {
		pen.Up()
		pen.Walk(30) // anda sem riscar
		pen.Down()

	}

	pen.SavePNG("tree.png")
	fmt.Println("PNG file created successfully.")
}
