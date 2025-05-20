package tabs

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"ScanEvalApp/internal/config"
	"ScanEvalApp/internal/gui/themeUI"
	"ScanEvalApp/internal/gui/widgets"
	"ScanEvalApp/internal/logging"
	"log/slog"

	"github.com/sqweek/dialog"
)

type SettingTab struct {
	selectFolderBtn widget.Clickable
	selectedPath    string
}

func NewSettingTab(w *app.Window) *SettingTab {
	errorLogger := logging.GetErrorLogger()

	tab := &SettingTab{
		selectFolderBtn: widget.Clickable{},
	}

	if path, err := config.LoadLastPath(); err == nil {
		tab.selectedPath = path
	} else {
		errorLogger.Error("Nepodarilo sa načítať poslednú cestu", slog.String("error", err.Error()))
	}

	return tab
}

func (t *SettingTab) Layout(gtx layout.Context, th *themeUI.Theme, w *app.Window) layout.Dimensions {
	errorLogger := logging.GetErrorLogger()

	if t.selectFolderBtn.Clicked(gtx) {
		go func() {
			dir, err := dialog.Directory().Title("Vyber priečinok").Browse()
			if err != nil {
				errorLogger.Error("Chyba pri výbere priečinka", slog.String("error", err.Error()))
				return
			}
			t.selectedPath = dir

			if err := config.SaveLastPath(dir); err != nil {
				errorLogger.Error("Chyba pri ukladaní cesty", slog.String("error", err.Error()))
			}

			w.Invalidate()
		}()
	}
	return layout.Inset{
		Top:    unit.Dp(16),
		Left:   unit.Dp(16),
		Right:  unit.Dp(16),
		Bottom: unit.Dp(0),
	}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Stack{}.Layout(gtx,
			layout.Stacked(func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{
					Axis: layout.Horizontal,
				}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Inset{
							Top:   unit.Dp(5),
							Right: unit.Dp(8),
						}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return material.Label(th.Theme, unit.Sp(18), "Miesto ukladania súborov:").Layout(gtx)
						})
					}),
					layout.Flexed(0.1, func(gtx layout.Context) layout.Dimensions {
						text := "Žiadny priečinok nebol vybraný"
						if t.selectedPath != "" {
							text = t.selectedPath
						}
						return layout.Inset{
							Right: unit.Dp(8),
						}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return widgets.LabelBorder(gtx, th, unit.Sp(16), text)
						})
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Inset{
							Top:   unit.Dp(2),
							Right: unit.Dp(700),
						}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							btn := widgets.Button(th.Theme, &t.selectFolderBtn, widgets.FileFolderIcon, widgets.IconPositionStart, "Vybrať priečinok")
							btn.Background = themeUI.LightGreen
							btn.Color = themeUI.White
							return btn.Layout(gtx, th)
						})
					}),
				)
			}),
		)
	})

}
