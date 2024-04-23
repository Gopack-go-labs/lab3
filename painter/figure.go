package painter

import (
	"image"
)

func FigureCoordinates(center image.Point) (image.Rectangle, image.Rectangle) {
	width := 300
	halfHeight := 150

	topLeft := center.Sub(image.Point{X: width / 2, Y: halfHeight})
	bottomRight := topLeft.Add(image.Point{X: width, Y: halfHeight})
	horizontal := image.Rectangle{Min: topLeft, Max: bottomRight}

	middleTop := center.Sub(image.Point{X: width / (3 * 2), Y: 0})
	middleBottom := middleTop.Add(image.Point{X: width / 3, Y: halfHeight})
	vertical := image.Rectangle{Min: middleTop, Max: middleBottom}

	return horizontal, vertical
}
