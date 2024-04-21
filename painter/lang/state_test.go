package lang

import (
	"github.com/roman-mazur/architecture-lab-3/painter"
	"golang.org/x/exp/shiny/materialdesign/colornames"
	"image"
	"testing"
)

func TestStateShouldRemoveDuplicateBgFillOp(t *testing.T) {
	s := State{}
	ops := []painter.Operation{
		painter.WhiteFillOp(),
		painter.BgRect(image.Rect(0, 0, 10, 10)),
		painter.FigureOp(image.Pt(0, 0), colornames.Black),
		painter.BgRect(image.Rect(10, 10, 20, 20)),
		painter.BgRect(image.Rect(10, 10, 20, 20)),
		painter.FigureOp(image.Pt(0, 0), colornames.Black),
		painter.GreenFillOp(),
	}

	for _, op := range ops {
		s.AddOperation(op)
	}

	actual := s.OperationList()

	if len(actual) != 4 {
		t.Fatalf("Expected %v, got %v", 4, len(actual))
	}
	if _, ok := actual[0].(painter.FillOp); !ok {
		t.Fatalf("Expected first FillOp, got %T", actual[0])
	}
	if _, ok := actual[1].(*painter.Figure); !ok {
		t.Fatalf("Expected second figre, got %T", actual[1])
	}
	if _, ok := actual[2].(painter.BgRectOp); !ok {
		t.Fatalf("Expected third BgRecFill, got %T", actual[2])
	}
	if _, ok := actual[3].(*painter.Figure); !ok {
		t.Fatalf("Expected fourth figre, got %T", actual[3])
	}
}
