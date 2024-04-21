package painter

import (
	"image"
	"image/color"

	"golang.org/x/exp/shiny/screen"
)

// Operation змінює вхідну текстуру.
type Operation interface {
	// Do виконує зміну операції, повертаючи true, якщо текстура вважається готовою для відображення.
	Do(t screen.Texture) (ready bool)
}

// OperationList групує список операції в одну.
type OperationList []Operation

func (ol OperationList) Do(t screen.Texture) (ready bool) {
	for _, o := range ol {
		ready = o.Do(t) || ready
	}
	return
}

// UpdateOp операція, яка не змінює текстуру, але сигналізує, що текстуру потрібно розглядати як готову.
var UpdateOp = updateOp{}

type updateOp struct{}

func (op updateOp) Do(t screen.Texture) bool { return true }

// OperationFunc використовується для перетворення функції оновлення текстури в Operation.
type OperationFunc func(t screen.Texture)

func (f OperationFunc) Do(t screen.Texture) bool {
	f(t)
	return false
}

// WhiteFill зафарбовує тестуру у білий колір. Може бути викоистана як Operation через OperationFunc(WhiteFill).
func WhiteFill(t screen.Texture) {
	t.Fill(t.Bounds(), color.White, screen.Src)
}

// GreenFill зафарбовує тестуру у зелений колір. Може бути викоистана як Operation через OperationFunc(GreenFill).
func GreenFill(t screen.Texture) {
	t.Fill(t.Bounds(), color.RGBA{G: 0xff, A: 0xff}, screen.Src)
}

func BgRect(rect image.Rectangle) OperationFunc {
	return func(t screen.Texture) {
		t.Fill(rect, color.Black, screen.Src)
	}
}

type Figure struct {
	center image.Point
	color  color.RGBA
}

func InitFigure(center image.Point, color color.RGBA) *Figure {
	return &Figure{center, color}
}

func (f *Figure) Do(t screen.Texture) bool {
	horizontal, vertical := FigureCoordinates(f.center)
	t.Fill(horizontal, f.color, screen.Src)
	t.Fill(vertical, f.color, screen.Src)
	return false
}

func Move(figures []Figure, vector image.Point) OperationFunc {
	return func(t screen.Texture) {
		for i := range figures {
			figures[i].center = figures[i].center.Add(vector)
			figures[i].Do(t)
		}
	}
}

func Reset() OperationFunc {
	return func(t screen.Texture) {
		t.Fill(t.Bounds(), color.Black, screen.Src)
	}
}
