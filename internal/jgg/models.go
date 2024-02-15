package jgg

type Element string

const (
	ELEMENT_MEDAL   Element = "金"
	ELEMENT_WOOD    Element = "木"
	ELEMENT_WATER   Element = "水"
	ELEMENT_FIRE    Element = "火"
	ELEMENT_EARTH   Element = "土"
	ELEMENT_UNKNOWN Element = ""
)

type Birthday struct {
	Solar         bool   `json:"solar"`
	LeapMonthFlag bool   `json:"leapMonthFlag"`
	Date          string `json:"date"`
	Hour          int    `json:"hour"`
}

type Ge struct {
	Id        int     `json:"id"`
	SolarDate string  `json:"solarDate"`
	LunarDate string  `json:"lunarDate"`
	LeapMonth bool    `json:"leapMonthFlag"`
	Hour      int     `json:"hour"`
	Solar     string  `json:"solarGe"`
	Lunar     string  `json:"lunarGe"`
	Element   Element `json:"element"`
}
