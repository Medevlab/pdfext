package main

import (
	"github.com/Medevlab/pdfext"
)

type AccountInfo struct {
	Name        string `json:"Name"`
	AccountNo   string `json:"Account No"`
	IDNumber    string `json:"ID number"`
	Nationality string `json:"Nationality"`
	RegAddress  string `json:"Registered Address"`
}

func main() {

	pdft := pdfext.NewPdf("./example.pdf")
	pdft.MaxWidth = 500
	pdft.AddFonts(
		map[string]string{
			"Roboto-Regular": "../ttf/Roboto-Regular.ttf",
			"HDZB_5":         "../ttf/wts11.ttf"})

	// ---------------页眉---------------
	header := pdfext.NewHeader()
	header.LeftText = "leftest "
	header.CenterText = "test"
	header.RightText = "Confidential"
	header.HeaderStyle = pdfext.FontStyle{
		FontSize:  8,
		FontName:  "Roboto-Regular",
		FontColor: pdfext.Black,
	}

	pdft.AddHeader(header)

	// ---------------标题---------------
	txt := pdfext.NewText()
	txt.Content = "Account Closure Report"
	txt.YPosition = pdft.CurrentXY.Y + 15
	txtsty := pdfext.CellStyle{
		Background: pdfext.White,
		Font:       "HDZB_5",
		FontColor:  pdfext.Black,
		FontSize:   14,
		H_Align:    pdfext.HalignCenter,
	}
	txt.TextStyle = txtsty
	pdft.AddText(txt)

	// ---------------表单---------------
	form := pdfext.NewForm()
	namesty := pdfext.CellStyle{
		Background: pdfext.LightGray,
		Font:       "Roboto-Regular",
		FontColor:  pdfext.Black,
		FontSize:   10,
		H_Align:    pdfext.HalignCenter,
		V_Align:    pdfext.ValignMiddle,
	}

	valuesty := pdfext.CellStyle{
		Background: pdfext.White,
		Font:       "Roboto-Regular",
		FontColor:  pdfext.Black,
		FontSize:   10,
		H_Align:    pdfext.HalignCenter,
		V_Align:    pdfext.ValignMiddle,
	}

	unitstyle := pdfext.UnitStyle{
		NameWidth:  120,
		ValueWidth: 120,
		AutoFit:    true,
		NameStyle:  namesty,
		ValueStyle: valuesty,
	}
	form.ColNum = 2
	form.Unitstyle = unitstyle
	form.Units = []pdfext.FormUnit{
		{"Name", "zhegadfq"},
		{"Account NO", "123432153"},
		{"Type of", "212343"},
		{"ID number", "2123152312355"},
	}
	form.Name = "Account Basic Information"
	form.NameStyle = pdfext.FontStyle{
		FontName:  "Roboto-Regular",
		FontColor: pdfext.Black,
		FontSize:  14,
	}
	form.MaxWidth = 500
	form.RowHeight = 16
	form.YPosition = pdft.CurrentXY.Y + 15

	pdft.AddForm(form)

	// ----------------表单补充部分---------------
	form2 := pdfext.NewForm()
	form2.ColNum = 1
	form2.Unitstyle = unitstyle
	form2.Units = []pdfext.FormUnit{
		{"Correspondence Address", "zhegadfq"},
		{"Type of Identification", "123432153"},
	}
	//form2.Name = ""
	form2.MaxWidth = 500
	form2.RowHeight = 16
	form2.YPosition = pdft.CurrentXY.Y //与上一部份拼接
	pdft.AddForm(form2)

	// ---------------普通文本---------------
	txt1 := pdfext.NewText()
	txt1.Content = "Account Closure Reason: Unresolved Screening Hits"
	txt1.YPosition = pdft.CurrentXY.Y + 15
	txt1sty := pdfext.CellStyle{
		Background: pdfext.White,
		Font:       "HDZB_5",
		FontColor:  pdfext.Black,
		FontSize:   14,
		H_Align:    pdfext.HalignLeft,
	}
	txt1.TextStyle = txt1sty
	pdft.AddText(txt1)

	// ---------------表格---------------
	tbl := pdfext.NewTable()
	ttitlesty := pdfext.CellStyle{
		Background: pdfext.LightGray,
		//	Font       :,
		FontColor: pdfext.Black,
		FontSize:  14,
		H_Align:   pdfext.HalignCenter,
		V_Align:   pdfext.ValignMiddle,
	}

	tdatasty := pdfext.CellStyle{
		Background: pdfext.White,
		//	Font       :,
		FontColor: pdfext.Black,
		FontSize:  14,
		H_Align:   pdfext.HalignCenter,
		V_Align:   pdfext.ValignMiddle,
	}

	Columns := []pdfext.Column{
		{
			Width:   60,
			AutoFit: true,
			Name:    "Name",
		},
		{
			Width:   60,
			AutoFit: true,
			Name:    "Account No",
		},
		{
			Width:   60,
			AutoFit: true,
			Name:    "ID number",
		},
		{
			Width:   60,
			AutoFit: true,
			Name:    "Nationality",
		},
		{
			Width:   60,
			AutoFit: true,
			Name:    "Registered Address",
		},
	}
	tbl.Name = "Account Closure Application Checklist"
	tbl.NameStyle = pdfext.FontStyle{
		FontName:  "Roboto-Regular",
		FontColor: pdfext.Black,
		FontSize:  14,
	}
	tbl.Colums = Columns
	tbl.Width = 500

	tbl.RowHeight = 20
	tbl.YPosition = pdft.CurrentXY.Y + 15
	tbl.TitleStyle = ttitlesty
	tbl.DataStyle = tdatasty
	tbl.Data = pdfext.ToTableData([]AccountInfo{
		{
			Name:        "21234",
			AccountNo:   "21234",
			IDNumber:    "21234",
			Nationality: "21234",
			RegAddress:  "21234",
		},
		{
			Name:        "dsadfewqewqw",
			AccountNo:   "dsadfewqewqw",
			IDNumber:    "dsadfewqewqw",
			Nationality: "dsadfewqewqw",
			RegAddress:  "dsadfewqewqw",
		}, {
			Name:        "2dddddddddddd234",
			AccountNo:   "2dddddddddddd234",
			IDNumber:    "2dddddddddddd234",
			Nationality: `3ii. Has the Money Laundering Reporting Officer or Deputy Officer(MLRO or DMLRO) responded and confirmed that the suspect has been cleared and that there is no need to report to the FinancialReporting Authority?`,
			RegAddress:  "2dddddddddddd234",
		}, {
			Name:        "21234",
			AccountNo:   "21234",
			IDNumber:    "21234",
			Nationality: "21234",
			RegAddress:  "21234",
		},
	})

	pdft.AddTable(tbl)

	// ---------------写文件---------------
	pdft.Save()
}
