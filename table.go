package pdfext

import (
	"math"
)

// 格子样式
type CellStyle struct {
	Background Color
	Font       string
	FontColor  Color
	FontSize   int
	H_Align    Halign
	V_Align    Valign
}

// 列
type Column struct {
	//Size       int
	Width   float64 // 宽度
	AutoFit bool    // 自适应,未生效
	Name    string  // 列名
}

// 表格
type Table struct {
	Name       string              // 表格名字,为空不显示
	NameStyle  FontStyle           // 名字样式
	Colums     []Column            // 列
	Width      float64             // 表格宽度
	Data       []map[string]string // 数据, 可通过ToTableData()转换
	YPosition  float64             // 起始Y坐标
	RowHeight  float64             // 行高
	TitleStyle CellStyle           // 表头样式
	DataStyle  CellStyle           // 表格数据样式
}

func NewTable() *Table {
	return &Table{}
}

func (t *Table) Draw(p *Pdf) {
	var bx float64 = 0
	var by float64 = t.YPosition

	left := (595.28 - t.Width) / 2
	bx = left
	var x float64 = bx
	var y float64 = by
	xstep := make(map[string]float64)

	var width float64 = 0
	colnum := len(t.Colums)
	for _, c := range t.Colums { //
		width += c.Width
	}
	if width > t.Width {
		width = t.Width
	}
	margin := t.Width - width
	col_margin := margin / float64(colnum)

	p.pdf.SetX(x)
	p.pdf.SetY(y)
	if t.Name != "" {
		p.SetFont(FontStyle{14, "HDZB_5", Black}).Cell(nil, t.Name)
		//f.pdf.SetFillColor(r uint8, g uint8, b uint8)
		p.SetFont(FontStyle{10, "Roboto-Regular", Black})
		y += 20
	}

	for _, c := range t.Colums { // draw title
		xstep[c.Name] = c.Width + col_margin
		DrawRectCell(p.pdf, c.Name, int(t.TitleStyle.FontSize),
			x, y, xstep[c.Name], t.RowHeight,
			t.TitleStyle.Background, t.TitleStyle.H_Align, t.TitleStyle.V_Align)
		x += (c.Width + col_margin)
	}
	x = bx
	y += t.RowHeight
	for _, d := range t.Data { // draw title
		maxrow := 1.0
		for _, c := range t.Colums {
			text := d[c.Name]
			cw, _ := p.pdf.MeasureTextWidth(text)
			row := math.Ceil(cw / xstep[c.Name])
			if row > maxrow {
				maxrow = row
			}
		}

		height := t.RowHeight * maxrow
		if maxrow > 1 {
			height = float64(t.DataStyle.FontSize) * maxrow
		}
		for _, c := range t.Colums { // draw title
			DrawRectMultiCell(p.pdf, d[c.Name], int(t.DataStyle.FontSize),
				x, y, xstep[c.Name], height, t.RowHeight,
				t.DataStyle.Background, t.DataStyle.H_Align, t.DataStyle.V_Align)
			x += (c.Width + col_margin)
		}
		x = bx
		y += height

	}
	p.CurrentXY = XY{X: x, Y: y}
}
