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
	Id        int     `json:"id" db:"id"`
	SolarDate string  `json:"solarDate" db:"solar_date"`
	LunarDate string  `json:"lunarDate" db:"lunar_date"`
	LeapMonth bool    `json:"leapMonthFlag" db:"leap_month"`
	Hour      int     `json:"hour" db:"hour"`
	Solar     string  `json:"solarGe" db:"solar_ge"`
	Lunar     string  `json:"lunarGe" db:"lunar_ge"`
	Element   Element `json:"Element" db:"Element"`
}
