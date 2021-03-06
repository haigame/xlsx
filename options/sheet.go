package options

import (
	"github.com/plandem/xlsx/types"
)

type sheetOption func(co *SheetOptions)

//SheetOptions is a helper type to simplify process of settings options for sheet
type SheetOptions struct {
	Visibility types.VisibilityType
}

//Sheet is a 'namespace' for all possible options for sheet
//
// Possible options are:
// Visibility
var Sheet sheetOption

//NewSheetOptions create and returns option set for sheet
func NewSheetOptions(options ...sheetOption) *SheetOptions {
	s := &SheetOptions{}
	s.Set(options...)
	return s
}

//Set sets new options for option set
func (co *SheetOptions) Set(options ...sheetOption) {
	for _, o := range options {
		o(co)
	}
}

//Visibility sets flag indicating if the affected column are hidden on this worksheet.
func (o *sheetOption) Visibility(visibility types.VisibilityType) sheetOption {
	return func(co *SheetOptions) {
		co.Visibility = visibility
	}
}
