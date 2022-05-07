package pdfext

type Valign int32

// 文本垂直样式
const (
	ValignTop    Valign = 1
	ValignMiddle Valign = 2
	ValignBottom Valign = 3
)

type Halign int32

// 文本水平样式
const (
	HalignLeft   Halign = 1
	HalignCenter Halign = 2
	HalignRight  Halign = 3
)

type Color int32

// 颜色
const (
	Black     Color = 1
	White     Color = 2
	Gray      Color = 3
	Red       Color = 4
	Green     Color = 5
	Yellow    Color = 6
	Blue      Color = 7
	LightGray Color = 31
)

type rgb struct {
	r uint8
	g uint8
	b uint8
}

var colorMap = map[Color]rgb{
	Black:     {0, 0, 0},
	White:     {255, 255, 255},
	Gray:      {128, 128, 128},
	Red:       {255, 0, 0},
	Green:     {0, 255, 0},
	Yellow:    {255, 255, 0},
	Blue:      {0, 0, 255},
	LightGray: {224, 224, 224},
}

// x,y坐标
type XY struct {
	X float64
	Y float64
}
