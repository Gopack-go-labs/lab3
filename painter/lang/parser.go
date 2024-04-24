package lang

import (
	"bufio"
	"image"
	"image/color"
	"io"
	"strconv"
	"strings"

	"github.com/roman-mazur/architecture-lab-3/painter"
)

// Parser уміє прочитати дані з вхідного io.Reader та повернути список операцій представлені вхідним скриптом.
type Parser struct {
} 

// Парсинг команд
func (p *Parser) Parse(in io.Reader) ([]painter.Operation, error) {
	var res []painter.Operation
	state := &State{}

	scanner := bufio.NewScanner(in)

	for scanner.Scan() {
		commandLine := scanner.Text()
		op := p.parse(commandLine) // parse the line to get Operation
		state.AddOperation(op)
	}

	res = state.OperationList()

	return res, nil
}

func (p *Parser) parse(commandLine string) painter.Operation {
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

		rect.Min.X, _ = strconv.Atoi(args[1])
		rect.Min.Y, _ = strconv.Atoi(args[2])
		rect.Max.X, _ = strconv.Atoi(args[3])
		rect.Max.Y, _ = strconv.Atoi(args[4])
		
		op = painter.BgRect(rect)
	
	case "figure":
		figureColor := color.RGBA{0xff, 0xf1, 0x76, 0xff}
		var center image.Point
		center.X, _ = strconv.Atoi(args[1])
		center.Y, _ = strconv.Atoi(args[2])
		op = painter.FigureOp(center, figureColor)
	
	case "move":

	case "reset":
	}

	return op
}
