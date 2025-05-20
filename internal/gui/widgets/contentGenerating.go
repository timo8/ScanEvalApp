package widgets

import (
	"gioui.org/layout"

	"gioui.org/unit"

	"gioui.org/widget/material"
	"ScanEvalApp/internal/gui/themeUI"
)

func ContentGenerating(th *themeUI.Theme, isGenerating *bool, generatedPath *string, message *string) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical, Alignment: layout.Middle}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					var text string
					if *isGenerating {
						text = "Generujem..."
					} else if *generatedPath == "" {
						text = "Chyba pri generovaní: " + *message
					} else {
						text = "Úspešne vygenerované:\n" + *generatedPath
					}
					lbl := material.Label(th.Material(), unit.Sp(20), text)
					return lbl.Layout(gtx)
				}),
			)
		})
	}
}
