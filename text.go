package pdfext

// 文本
type Text struct {
	Content string // 文本内容
	//MaxWidth  float64   // 文本最大宽度
	TextStyle CellStyle // 文本样式
	YPosition float64   // Y坐标
}

func NewText() *Text {
	return &Text{}
}

func (t *Text) Draw(p *Pdf) {
	var bx float64 = 0
	var by float64 = t.YPosition
	var y float64 = by

	left := (595.28 - p.MaxWidth) / 2
	bx = left
	var x float64 = bx

	p.SetFont(FontStyle{int(t.TextStyle.FontSize), t.TextStyle.Font, t.TextStyle.FontColor})
	if t.TextStyle.H_Align == HalignCenter {
		textw, _ := p.pdf.MeasureTextWidth(t.Content)
		x = x + (p.MaxWidth / 2) - (textw / 2)
	} else if t.TextStyle.H_Align == HalignRight {
		textw, _ := p.pdf.MeasureTextWidth(t.Content)
		x = x + p.MaxWidth - textw
	}
	p.pdf.SetX(x)
	p.pdf.SetY(y)
	p.pdf.Cell(nil, t.Content)

	y += 20
	p.CurrentXY = XY{X: x, Y: y}
}
