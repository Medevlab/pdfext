package pdfext

// 表单
type UnitStyle struct {
	NameWidth  float64   // Name单元格的宽度
	ValueWidth float64   // Value单元格的宽度
	AutoFit    bool      // 自适应, 不生效
	NameStyle  CellStyle // Name单元格的样式
	ValueStyle CellStyle // Value单元格的样式
}

type FormUnit struct {
	UnitName  string
	UnitValue string
}

type FontStyle struct {
	FontSize  int    // 字体大小
	FontName  string // 字体名称, 使用AddFonts注册字体时的名称
	FontColor Color
}

type Form struct {
	Name      string     // 表单名称, 空则不显示名称
	NameStyle FontStyle  // 名称字体样式
	ColNum    int        //表单列数, Name+Value为一列
	Units     []FormUnit // Name+Value为Unit
	MaxWidth  float64    // 表单最大宽度
	Unitstyle UnitStyle  // Unit的样式
	YPosition float64    // 起始Y坐标
	RowHeight float64    // 每行高度
}

func NewForm() *Form {
	return &Form{}
}

func (f *Form) Draw(p *Pdf) {
	var bx float64 = 0
	var by float64 = f.YPosition

	left := (595.28 - f.MaxWidth) / 2
	bx = left

	var xy XY = XY{bx, by}
	var width float64 = 0

	width = (f.Unitstyle.NameWidth + f.Unitstyle.ValueWidth) * float64(f.ColNum)
	if width > f.MaxWidth {
		width = f.MaxWidth
	}
	margin := f.MaxWidth - width
	cell_margin := margin / float64(f.ColNum)

	xy = p.UpdateXY(xy)

	if f.Name != "" {
		p.SetFont(&FontStyle{f.NameStyle.FontSize, f.NameStyle.FontName, f.NameStyle.FontColor}).Cell(nil, f.Name)
		xy.Y += 20
		xy = p.UpdateXY(xy)
	}

	for i, u := range f.Units { //
		if i != 0 && i%f.ColNum == 0 {
			xy.Y += f.RowHeight
			xy.X = bx
			xy = p.UpdateXY(xy)
		}
		p.SetFont(&FontStyle{f.Unitstyle.NameStyle.FontSize, f.Unitstyle.NameStyle.Font, f.Unitstyle.NameStyle.FontColor})

		DrawRectCell(p.pdf, u.UnitName, int(f.Unitstyle.NameStyle.FontSize),
			xy.X, xy.Y, f.Unitstyle.NameWidth, f.RowHeight,
			f.Unitstyle.NameStyle.Background,
			f.Unitstyle.NameStyle.H_Align,
			f.Unitstyle.NameStyle.V_Align)
		xy.X += f.Unitstyle.NameWidth

		p.SetFont(&FontStyle{f.Unitstyle.ValueStyle.FontSize, f.Unitstyle.ValueStyle.Font, f.Unitstyle.ValueStyle.FontColor})
		DrawRectCell(p.pdf, u.UnitValue, int(f.Unitstyle.ValueStyle.FontSize),
			xy.X, xy.Y, f.Unitstyle.ValueWidth+cell_margin, f.RowHeight,
			f.Unitstyle.ValueStyle.Background,
			f.Unitstyle.ValueStyle.H_Align,
			f.Unitstyle.ValueStyle.V_Align)
		xy.X += (f.Unitstyle.ValueWidth + cell_margin)
		xy = p.UpdateXY(xy)
	}
	xy.Y += f.RowHeight
	xy = p.UpdateXY(xy)
	//f.pdf.WritePdf(f.FileName)
}
