package lang

import (
	"fmt"
	"image"
	"image/color"
	"reflect"
	"strings"
	"testing"
	"golang.org/x/exp/shiny/screen"
	"github.com/roman-mazur/architecture-lab-3/painter"
)

func TestParser_Parse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []painter.Operation
	}{
		{
			name:  "Test white fill",
			input: "white",
			expected: []painter.Operation{painter.WhiteFillOp()},
		},
		{
			name:  "Test green fill",
			input: "green",
			expected: []painter.Operation{painter.GreenFillOp(),},
		},
		{
			name:  "Test update",
			input: "update",
			expected: []painter.Operation{updateOp,},
		},
		{
			name:  "Test bgrect",
			input: "bgrect 0.1 0.2 0.3 0.4",
			expected: []painter.Operation{painter.BgRect(image.Rect(80, 160, 240, 320)),},
		},
		{
			name:  "Test figure",
			input: "figure 0.1 0.2",
			expected: 
			[]painter.Operation{
				painter.FigureOp(image.Pt(160, 320), color.RGBA{R: 255, G: 241, B: 118, A: 255}),
			},
		},
		{
			name:  "Test move",
			input: "move 0.1 0.2",
			expected: []painter.Operation{painter.Move(nil, image.Pt(160, 320)),},
		},
		{
			name:  "Test reset",
			input: "reset",
			expected: []painter.Operation{painter.Reset(),},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			state := &State{}
			parser := &Parser{}
			input := strings.NewReader(test.input)
			actual, err := parser.Parse(input, state)

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if reflect.DeepEqual(actual, test.expected) {
				fmt.Println(actual)
				t.Errorf("expected %v, got %v", test.expected, actual)
			}
		})
	}
}

func TestConv(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"0.1", 80},
		{"0.2", 160},
		{"0.5", 400},
		{"0.9", 720},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			actual := conv(test.input)
			if actual != test.expected {
				t.Errorf("expected %d, got %d", test.expected, actual)
			}
		})
	}
}

var updateOp = UpdateOp{}

type UpdateOp struct{}

func (op UpdateOp) Do(t screen.Texture) bool { return true }
