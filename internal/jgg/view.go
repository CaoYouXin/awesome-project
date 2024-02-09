package jgg

type element string

const (
	ELEMENT_MEDAL   element = "金"
	ELEMENT_WOOD    element = "木"
	ELEMENT_WATER   element = "水"
	ELEMENT_FIRE    element = "火"
	ELEMENT_EARTH   element = "土"
	ELEMENT_UNKNOWN element = ""
)

type Birthday struct {
	Solar         bool   `json:"solar"`
	LeapMonthFlag bool   `json:"leapMonthFlag"`
	Date          string `json:"date"`
	Hour          int    `json:"hour"`
}

type Ge struct {
	Solar     string  `json:"solarGe"`
	Lunar     string  `json:"lunarGe"`
	SolarDate string  `json:"solarDate"`
	LunarDate string  `json:"lunarDate"`
	Element   element `json:"element"`
}
