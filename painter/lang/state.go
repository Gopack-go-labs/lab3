package lang

import (
	"github.com/roman-mazur/architecture-lab-3/painter"
	"slices"
)

type State struct {
	bg painter.Operation
	fg []painter.Operation
}

func (s *State) Reset() {
	s.bg = nil
	s.fg = nil
}

func (s *State) OperationList() painter.OperationList {
	var filtered []painter.Operation

	isBgRectFound := false
	for i := len(s.fg) - 1; i >= 0; i-- {
		op := s.fg[i]

		if _, isBgRectOp := op.(painter.BgRectOp); isBgRectOp {
			if isBgRectFound {
				continue
			}
			isBgRectFound = true
		}
		filtered = append(filtered, op)
	}

	if s.bg != nil {
		filtered = append(filtered, s.bg)
	}
	slices.Reverse(filtered)
	return filtered
}

func (s *State) AddOperation(op painter.Operation) {
	if _, isFillOp := op.(painter.FillOp); isFillOp {
		s.bg = op
	} else {
		s.fg = append(s.fg, op)
	}
}

func (s *State) GetFigures() []*painter.Figure {
	var figures []*painter.Figure

	for _, op := range s.fg {
		if f, isFigure := op.(*painter.Figure); isFigure {
			figures = append(figures, f)
		}
	}

	return figures
}
