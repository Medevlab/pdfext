package pdfext

import (
	"encoding/json"
	"log"
	"math"

	"github.com/signintech/gopdf"
)

// pdf
type Pdf struct {
	FileName    string //文件名
	pdf         *gopdf.GoPdf
	Fonts       []string //已注册字体
	CurrentFont string   // 当前字体
	CurrentXY   XY       //当前位置
	MaxWidth    float64  // 内容区域最大宽度
}

func NewPdf(filename string) *Pdf {
	p := &Pdf{
		FileName:  filename,
		MaxWidth:  595.28,
		pdf:       &gopdf.GoPdf{},
		CurrentXY: XY{0, 0},
	}
	p.pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 595.28, H: 841.89}}) //595.28, 841.89 = A4
	p.pdf.AddPage()

	// 默认字段, 待优化
	p.AddFont("HDZB_5", "../ttf/wts11.ttf")
	p.SetFont(FontStyle{12, "HDZB_5", Black})

	return p
}

// 注册字体
func (p *Pdf) AddFont(fontname, fontpath string) error {
	err := p.pdf.AddTTFFont(fontname, fontpath)
	if err != nil {
		log.Print(err.Error())
		return err
	}
	p.Fonts = append(p.Fonts, fontname)
	return nil
}

// 注册字体
func (p *Pdf) AddFonts(fonts map[string]string) error {
	for fontname, fontpath := range fonts {
		err := p.pdf.AddTTFFont(fontname, fontpath)
		if err != nil {
			log.Print(err.Error())
			return err
		}
		p.Fonts = append(p.Fonts, fontname)
	}

	return nil
}

func (p *Pdf) AddForm(form *Form) error {
	form.Draw(p)
	return nil
}
func (p *Pdf) AddTable(table *Table) error {
	table.Draw(p)
	return nil
}

func (p *Pdf) AddText(text *Text) error {
	text.Draw(p)
	return nil
}
func (p *Pdf) AddHeader(header *Header) error {
	header.Draw(p)
	return nil
}

func (p *Pdf) Save() error {
	return p.pdf.WritePdf(p.FileName)
}

////////////////////////////////////////////////////
func (p *Pdf) SetFont(style FontStyle) *gopdf.GoPdf {

	color := colorMap[style.FontColor]

	p.pdf.SetFont(style.FontName, "", style.FontSize)
	p.pdf.SetFillColor(color.r, color.g, color.b)
	//f.pdf.SetGrayFill(0.5)
	return p.pdf
}

func DrawRectCell(pdf *gopdf.GoPdf,
	text string,
	fontSize int,
	x, y, w, h float64,
	color Color,
	align Halign, valign Valign,
) {

	c := colorMap[color]
	pdf.SetLineWidth(0.1)
	pdf.SetFillColor(c.r, c.g, c.b) //setup fill color
	pdf.RectFromUpperLeftWithStyle(x, y, w, h, "FD")
	pdf.SetFillColor(0, 0, 0)

	if align == HalignCenter {
		textw, _ := pdf.MeasureTextWidth(text)
		x = x + (w / 2) - (textw / 2)
	} else if align == HalignRight {
		textw, _ := pdf.MeasureTextWidth(text)
		x = x + w - textw
	}

	pdf.SetX(x)

	if valign == ValignMiddle {
		y = y + (h / 2) - (float64(fontSize) / 2)
	} else if valign == ValignBottom {
		y = y + h - float64(fontSize)
	}

	pdf.SetY(y)
	pdf.Cell(nil, text)
}

func DrawRectMultiCell(pdf *gopdf.GoPdf,
	text string,
	fontSize int,
	x, y, w, h float64,
	rowheight float64,
	color Color,
	align Halign, valign Valign,
) {

	c := colorMap[color]
	pdf.SetLineWidth(0.1)
	pdf.SetFillColor(c.r, c.g, c.b) //setup fill color
	pdf.RectFromUpperLeftWithStyle(x, y, w, h, "FD")
	pdf.SetFillColor(0, 0, 0)
	textw, _ := pdf.MeasureTextWidth(text)
	row := math.Ceil(textw / w)

	rh := rowheight
	if row == 1 {
		nx := x
		if align == HalignCenter {
			nx = x + (w / 2) - (textw / 2)
		} else if align == HalignRight {
			nx = x + w - textw
		}
		if nx < x {
			nx = x
		} else {
			x = nx
		}
	}

	if h > rowheight {
		rh = float64(fontSize) * row
	}
	pdf.SetX(x)

	if row == 1 {
		if valign == ValignMiddle {
			y = y + (h / 2) - (float64(fontSize) / 2)
		} else if valign == ValignBottom {
			y = y + h - float64(fontSize)
		}
	}

	pdf.SetY(y)
	pdf.MultiCell(&gopdf.Rect{W: w, H: rh}, text)
}

func ToTableData(data interface{}) []map[string]string {
	tdata := []map[string]string{}
	jdata, err := json.Marshal(data)
	if err != nil {
		return nil
	}
	json.Unmarshal(jdata, &tdata)
	return tdata
}
