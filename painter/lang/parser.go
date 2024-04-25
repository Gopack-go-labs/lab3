package lang

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/roman-mazur/architecture-lab-3/painter"
)

// Parser уміє прочитати дані з вхідного io.Reader та повернути список операцій представлені вхідним скриптом.
type Parser struct {
} 

// Парсинг команд
func (p *Parser) Parse(in io.Reader, state *State) ([]painter.Operation, error) {
	var res []painter.Operation

	scanner := bufio.NewScanner(in)

	for scanner.Scan() {
		commandLine := scanner.Text()
		op := p.parse(commandLine, state) // parse the line to get Operation
		state.AddOperation(op)
	}

	res = state.OperationList()

	return res, nil
}

func (p *Parser) parse(commandLine string, state *State) painter.Operation {
	var op painter.Operation

	args := strings.Split(commandLine, " ")
	command := args[0]

	switch command {
	case "white":
		op = painter.WhiteFillOp()

	case "green":
		op = painter.GreenFillOp()
	
	case "update":
		op = painter.UpdateOp
	
	case "bgrect":
		var rect image.Rectangle

		rect.Min.X = conv(args[1])
		rect.Min.Y = conv(args[2])
		rect.Max.X = conv(args[3])
		rect.Max.Y = conv(args[4])
		
		op = painter.BgRect(rect)
	
	case "figure":
		figureColor := color.RGBA{0xff, 0xf1, 0x76, 0xff}
		var center image.Point
		
		center.X = conv(args[1])
		center.Y = conv(args[2])
		
		op = painter.FigureOp(center, figureColor)
	
	case "move":
		var center image.Point

		center.X = conv(args[1])
		center.Y = conv(args[2])

		op = painter.Move(state.OperationList(), center)

	case "reset":
		state.fg = []painter.Operation{}
		op = painter.Reset()
	
	default:
		fmt.Println("Warning: Unknown operation")
	}

	return op
}

func conv(f string) int {
	wSize := 800
	num, err := strconv.ParseFloat(f, 64)

	if err != nil {
		log.Printf("Cannot convert %s", f)
	}

	res := int(float64(wSize) * num)
	return res
}
