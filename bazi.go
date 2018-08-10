package fate

import (
	"github.com/godcong/chronos"
)

var diIndex = map[string]int{
	"子": 0, "丑": 1, "寅": 2, "卯": 3, "辰": 4, "巳": 5, "午": 6, "未": 7, "申": 8, "酉": 9, "戌": 10, "亥": 11,
}

var tianIndex = map[string]int{
	"甲": 0,
	"乙": 1,
	"丙": 2,
	"丁": 3,
	"戊": 4,
	"己": 5,
	"庚": 6,
	"辛": 7,
	"壬": 8,
	"癸": 9,
}

var tiangan = [][]int{
	{1200, 1200, 1000, 1000, 1000, 1000, 1000, 1000, 1200, 1200},
	{1060, 1060, 1000, 1000, 1100, 1100, 1140, 1140, 1100, 1100},
	{1140, 1140, 1200, 1200, 1060, 1060, 1000, 1000, 1000, 1000},
	{1200, 1200, 1200, 1200, 1000, 1000, 1000, 1000, 1000, 1000},
	{1100, 1100, 1060, 1060, 1100, 1100, 1100, 1100, 1040, 1040},
	{1000, 1000, 1140, 1140, 1140, 1140, 1060, 1060, 1060, 1060},
	{1000, 1000, 1200, 1200, 1200, 1200, 1000, 1000, 1000, 1000},
	{1040, 1040, 1100, 1100, 1160, 1160, 1100, 1100, 1000, 1000},
	{1060, 1060, 1000, 1000, 1000, 1000, 1140, 1140, 1200, 1200},
	{1000, 1000, 1000, 1000, 1000, 1000, 1200, 1200, 1200, 1200},
	{1000, 1000, 1040, 1040, 1140, 1140, 1160, 1160, 1060, 1060},
	{1200, 1200, 1000, 1000, 1000, 1000, 1000, 1000, 1140, 1140},
}

//var zhizang = map[string]int{
//
//}

var dizhi = []map[string][]int{
	{
		"癸": {1200, 1100, 1000, 1000, 1040, 1060, 1000, 1000, 1200, 1200, 1060, 1140},
	}, {
		"癸": {360, 330, 300, 300, 312, 318, 300, 300, 360, 360, 318, 342},
		"辛": {200, 228, 200, 200, 230, 212, 200, 220, 228, 248, 232, 200},
		"己": {500, 550, 530, 500, 550, 570, 600, 580, 500, 500, 570, 500},
	}, {
		"丙": {300, 300, 360, 360, 318, 342, 360, 330, 300, 300, 342, 318},
		"甲": {840, 742, 798, 840, 770, 700, 700, 728, 742, 700, 700, 840},
	}, {
		"乙": {1200, 1060, 1140, 1200, 1100, 1000, 1000, 1040, 1060, 1000, 1000, 1200},
	}, {
		"乙": {360, 318, 342, 360, 330, 300, 300, 312, 318, 300, 300, 360},
		"癸": {240, 220, 200, 200, 208, 200, 200, 200, 240, 240, 212, 228},
		"戊": {500, 550, 530, 500, 550, 600, 600, 580, 500, 500, 570, 500},
	}, {
		"庚": {300, 342, 300, 300, 330, 300, 300, 330, 342, 360, 348, 300},
		"丙": {700, 700, 840, 840, 742, 840, 840, 798, 700, 700, 728, 742},
	}, {
		"丁": {1000, 1000, 1200, 1200, 1060, 1140, 1200, 1100, 1000, 1000, 1040, 1060},
	}, {
		"丁": {300, 300, 360, 360, 318, 342, 360, 330, 300, 300, 312, 318},
		"乙": {240, 212, 228, 240, 220, 200, 200, 208, 212, 200, 200, 240},
		"己": {500, 550, 530, 500, 550, 570, 600, 580, 500, 500, 570, 500},
	}, {
		"壬": {360, 330, 300, 300, 312, 318, 300, 300, 360, 360, 318, 342},
		"庚": {700, 798, 700, 700, 770, 742, 700, 770, 798, 840, 812, 700},
	}, {
		"辛": {1000, 1140, 1000, 1000, 1100, 1060, 1000, 1100, 1140, 1200, 1160, 1000},
	}, {
		"辛": {300, 342, 300, 300, 330, 318, 300, 330, 342, 360, 348, 300},
		"丁": {200, 200, 240, 240, 212, 228, 240, 220, 200, 200, 208, 212},
		"戊": {500, 550, 530, 500, 550, 570, 600, 580, 500, 500, 570, 500},
	}, {
		"甲": {360, 318, 342, 360, 330, 300, 300, 312, 318, 300, 300, 360},
		"壬": {840, 770, 700, 700, 728, 742, 700, 700, 840, 840, 724, 798},
	},
}

var wuXingTianGan = map[string]string{
	"甲": "木",
	"乙": "木",
	"丙": "火",
	"丁": "火",
	"戊": "土",
	"己": "土",
	"庚": "金",
	"辛": "金",
	"壬": "水",
	"癸": "水",
}

var wuXingDiZhi = map[string]string{
	"子": "水",
	"丑": "土",
	"寅": "木",
	"卯": "木",
	"辰": "土",
	"巳": "火",
	"午": "火",
	"未": "土",
	"申": "金",
	"酉": "金",
	"戌": "土",
	"亥": "水",
}

var sheng = []string{"木", "火", "土", "金", "水"}
var ke = []string{"木", "土", "水", "火", "金"}

func WuXingTianGan(s string) string {
	return wuXingTianGan[s]
}

func WuXingDiZhi(s string) string {
	return wuXingDiZhi[s]
}

type WuXingFen struct {
	Jin  int
	Mu   int
	Shui int
	Huo  int
	Tu   int
}

type XiYong struct {
	XiShen     string
	YongShen   string
	TongLei    []string
	TongLeiFen int
	YiLei      []string
	YiLeiFen   int
}

func (fen *WuXingFen) Add(s string, point int) {
	switch s {
	case "金":
		fen.Jin += point
	case "木":
		fen.Mu += point
	case "水":
		fen.Shui += point
	case "火":
		fen.Huo += point
	case "土":
		fen.Tu += point
	}
}

func (fen *WuXingFen) Get(s string) (point int) {
	switch s {
	case "金":
		point = fen.Jin
	case "木":
		point = fen.Mu
	case "水":
		point = fen.Shui
	case "火":
		point = fen.Huo
	case "土":
		point = fen.Tu
	}
	return point
}

type BaZi struct {
	baZi      []string
	wuXing    []string
	wuXingFen *WuXingFen
	xiyong    *XiYong
}

func NewBazi(calendar chronos.Calendar) *BaZi {
	ec := calendar.Lunar().EightCharacter()
	return &BaZi{
		baZi:      ec,
		wuXing:    baziToWuXing(ec),
		wuXingFen: point(ec),
		xiyong:    &XiYong{},
	}
}

func (z *BaZi) RiGan() string {
	return z.baZi[4]
}

func (z *BaZi) XiYong() *XiYong {
	z.xiyong = tong(z)
	z.xiyong = yi(z)
	z.xiyong = yongShen(z)
	z.xiyong = xiShen(z)
	return z.xiyong
}

func yongShen(z *BaZi) *XiYong {
	z.xiyong.YongShen = z.xiyong.TongLei[0]
	return z.xiyong
}

func (z *BaZi) QiangRuo() bool {
	if z.xiyong.TongLeiFen > z.xiyong.YiLeiFen {
		return true
	}
	return false
}

func xiShen(z *BaZi) (*XiYong) {
	rt := sheng
	if z.QiangRuo() {
		rt = ke
	}
	for i := range rt {
		if rt[i] == z.xiyong.YongShen {
			if i == len(rt)-1 {
				i = -1
			}
			z.xiyong.XiShen = rt[i+1]
			break
		}
	}
	return z.xiyong
}

func baziToWuXing(bazi []string) []string {
	var wx []string
	for idx, v := range bazi {
		if idx%2 == 0 {
			wx = append(wx, WuXingTianGan(v))
		} else {
			wx = append(wx, WuXingDiZhi(v))
		}
	}
	return wx
}

func point(bazi []string) *WuXingFen {
	di := diIndex[bazi[3]]
	var wxf WuXingFen
	for idx, v := range bazi {
		if idx%2 == 0 {
			wxf.Add(WuXingTianGan(v), tiangan[di][tianIndex[v]])
		} else {
			dz := dizhi[diIndex[v]]
			for k := range dz {
				wxf.Add(WuXingTianGan(k), dz[k][di])
			}
		}
	}
	return &wxf
}

func tong(z *BaZi) *XiYong {
	for i := range sheng {
		if wuXingTianGan[z.RiGan()] == sheng[i] {
			z.xiyong.TongLei = append(z.xiyong.TongLei, sheng[i])
			z.xiyong.TongLeiFen = z.wuXingFen.Get(sheng[i])
			if i == 0 {
				i = len(sheng) - 1
				z.xiyong.TongLei = append(z.xiyong.TongLei, sheng[i])
				z.xiyong.TongLeiFen += z.wuXingFen.Get(sheng[i])
			} else {
				z.xiyong.TongLei = append(z.xiyong.TongLei, sheng[i-1])
				z.xiyong.TongLeiFen += z.wuXingFen.Get(sheng[i-1])
			}
			break
		}
	}
	return z.xiyong
}

func yi(z *BaZi) *XiYong {
	for i := range sheng {
		for ti := range z.xiyong.TongLei {
			if z.xiyong.TongLei[ti] == sheng[i] {
				goto end
			}
		}
		z.xiyong.YiLei = append(z.xiyong.YiLei, sheng[i])
		z.xiyong.YiLeiFen += z.wuXingFen.Get(sheng[i])
	end:
		continue

	}
	return z.xiyong
}