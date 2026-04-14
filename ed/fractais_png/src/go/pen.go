package main

import (
	"math"

	"github.com/fogleman/gg"
)

type Pen struct {
	x, y    float64
	angle   float64
	penDown bool
	dc      *gg.Context
}

func NewPen(width, height int) *Pen {
	dc := gg.NewContext(width, height)
	dc.SetRGB(0, 0, 0)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	dc.SetLineWidth(1)
	return &Pen{
		x: float64(width) / 2, y: float64(height) / 2,
		angle: 0, penDown: true, dc: dc,
	}
}

func (t *Pen) Walk(dist float64) {
	newX := t.x + dist*math.Cos(t.angle*math.Pi/180)
	newY := t.y - dist*math.Sin(t.angle*math.Pi/180)
	if t.penDown {
		t.dc.DrawLine(t.x, t.y, newX, newY)
		t.dc.Stroke()
	}
	t.x, t.y = newX, newY
}

func (t *Pen) Left(deg float64)  { t.angle += deg }
func (t *Pen) Right(deg float64) { t.angle -= deg }

func (t *Pen) Up()   { t.penDown = false }
func (t *Pen) Down() { t.penDown = true }

func (t *Pen) Goto(x, y float64) {
	if t.penDown {
		t.dc.DrawLine(t.x, t.y, x, y)
		t.dc.Stroke()
	}
	t.x, t.y = x, y
}

func (t *Pen) SetPosition(x, y float64) {
	t.x = x
	t.y = y
}

func (t *Pen) SetHeading(angle float64) {
	t.angle = angle
}

func (t *Pen) DrawCircle(radius float64) {
	if t.penDown {
		t.dc.DrawCircle(t.x, t.y, radius)
		t.dc.Stroke()
	}
}

func (t *Pen) DrawRect(w, h float64) {
	if t.penDown {
		t.dc.DrawRectangle(t.x, t.y, w, h)
		t.dc.Stroke()
	}
}

func (t *Pen) FillCircle(radius float64) {
	if t.penDown {
		t.dc.DrawCircle(t.x, t.y, radius)
		t.dc.Fill()
	}
}

func (t *Pen) FillSquare(w, h float64) {
	t.dc.DrawRectangle(t.x, t.y, w, h)
	t.dc.Fill()
}

func (t *Pen) SetRGB(r, g, b float64) {
	t.dc.SetRGB(r/255, g/255, b/255)
}

func (t *Pen) SetLineWidth(w float64) {
	t.dc.SetLineWidth(w)
}

func (t *Pen) SavePNG(path string) {
	t.dc.SavePNG(path)
}

func ramo (pen *Pen, size float64){
	if size < 2 {
		return
	}

	pen.Walk(size)

	subDiv := 5.0
	tamSubDiv := size / subDiv

	x := pen.x
	y := pen.y
	angulo := pen.angle

	pen.SetPosition(x, y)
	pen.SetHeading(angulo)

	for i := 0; i < int(subDiv); i++{
		pen.Walk(tamSubDiv)

		rx := x
		ry := y
		rang := angulo

		pen.Left(25)
		ramo(pen, size * 0.4)

		pen.SetPosition(rx, ry)
		pen.SetHeading(rang)

		pen.Right(25)
		ramo(pen, size * 0.4)

		pen.SetPosition(rx, ry)
		pen.SetHeading(rang)
	}

	
}
