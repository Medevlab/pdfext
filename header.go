package pdfext

type Header struct {
	LeftText    string
	CenterText  string
	RightText   string
	HeaderStyle FontStyle
}

func NewHeader() *Header {
	return &Header{}
}

func (h *Header) Draw(p *Pdf) {
	var bx float64 = 0
	var by float64 = 35
	var y float64 = by

	left := (595.28 - p.MaxWidth) / 2
	bx = left
	var x float64 = bx
	//xstep := make(map[int]float64)
	//var ystep float64 = 14

	// for i := 0; i < h.ColNum; i++ {
	// 	xstep[i] = u.Width + cell_margin
	// }

	p.SetFont(&FontStyle{int(h.HeaderStyle.FontSize), h.HeaderStyle.FontName, h.HeaderStyle.FontColor})
	p.pdf.SetY(y)

	if h.LeftText != "" {
		p.pdf.SetX(x)
		p.pdf.Cell(nil, h.LeftText)
	}

	if h.CenterText != "" {
		textw, _ := p.pdf.MeasureTextWidth(h.CenterText)
		x = bx + (p.MaxWidth / 2) - (textw / 2)
		p.pdf.SetX(x)
		p.pdf.Cell(nil, h.CenterText)
	}

	if h.RightText != "" {
		textw, _ := p.pdf.MeasureTextWidth(h.RightText)
		x = bx + p.MaxWidth - textw
		p.pdf.SetX(x)
		p.pdf.Cell(nil, h.RightText)
	}
	y += float64(h.HeaderStyle.FontSize) + 2
	p.pdf.SetLineWidth(0.1)
	// p.pdf.SetLineType("dashed")
	//pdf.SetLineType("dotted")
	p.pdf.Line(bx, y, bx+p.MaxWidth, y)
	y += 20
	p.CurrentXY = XY{X: x, Y: y}
}
