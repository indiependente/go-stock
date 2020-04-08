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
	numPoints := 80
	if err := ui.Init(); err != nil {
		return fmt.Errorf("failed to initialize termui: %w", err)
	}
	defer ui.Close()
	data := ts.Price()
	lendata := len(data)
	if lendata != len(ts.Labels) {
		return errors.New("num labels mismatches data points")
	}
	lbls := ts.Labels[lendata-numPoints:]
	p := widgets.NewParagraph()
	p.Title = fmt.Sprintf("Stock Price for %s", symbol)
	p.Text = strconv.FormatFloat(currentPrice, 'f', -1, 64)
	p.SetRect(0, 0, 50, 5)
	p.TextStyle.Fg = ui.ColorWhite
	p.BorderStyle.Fg = ui.ColorCyan

	bc := widgets.NewBarChart()
	bc.Title = "Bar Chart"
	bc.SetRect(0, 5, 80, 30)
	bc.Labels = lbls
	bc.BarColors[0] = ui.ColorGreen
	bc.NumStyles[0] = ui.NewStyle(ui.ColorBlack)
	bc.Data = data[lendata-numPoints:]
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
