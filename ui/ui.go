package ui

import (
	"errors"
	"fmt"
	"strconv"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/indiependente/go-stock/models"
)

func HandleData(symbol string, currentPrice float64, ts models.TimeSeries) error {
	//x, y := ui.TerminalDimensions()
	x, y := 150, 40
	numPoints := x
	if err := ui.Init(); err != nil {
		return fmt.Errorf("failed to initialize termui: %w", err)
	}
	defer ui.Close()
	data := ts.High()
	lendata := len(data)
	if lendata != len(ts.Labels) {
		return errors.New("num labels mismatches data points")
	}

	lbls := ts.Labels
	if lendata > numPoints {
		lbls = ts.Labels[lendata-numPoints:]
		data = data[lendata-numPoints:]
	}

	p := widgets.NewParagraph()
	p.Title = fmt.Sprintf("Stock Price for %s", symbol)
	p.Text = strconv.FormatFloat(currentPrice, 'f', -1, 64)

	p.SetRect(0, 0, 50, 5)
	p.TextStyle.Fg = ui.ColorWhite
	p.BorderStyle.Fg = ui.ColorCyan
	p.TextStyle.Modifier = ui.ModifierBold

	bc := widgets.NewBarChart()
	bc.Title = "Bar Chart"
	bc.SetRect(0, 5, x, y-5)
	bc.BarGap = 0
	bc.Labels = lbls
	bc.BarColors[0] = ui.ColorGreen
	bc.NumStyles[0] = ui.NewStyle(ui.ColorBlack)
	bc.Data = data
	ui.Render(p, bc)

	uiEvents := ui.PollEvents()
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return nil
			}
		}
	}
}
