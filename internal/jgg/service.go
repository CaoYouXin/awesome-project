package jgg

import (
	"awesomeProject/internal/res"
	"errors"
	"fmt"
	"slices"
	"strconv"
	"unicode/utf8"

	"awesomeProject/internal/utils"

	"github.com/nosixtools/solarlunar"
)

var digitExist = map[rune]bool{
	'0': true, '1': true, '2': true, '3': true, '4': true, '5': true, '6': true, '7': true, '8': true, '9': true,
}

var digitNum = map[rune]int{
	'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9,
}

type Service struct {
	DAO
}

func (s Service) InitCounter(elem Element) map[rune]float64 {
	counter := map[rune]float64{
		'0': 0, '1': 0, '2': 0, '3': 0, '4': 0, '5': 0, '6': 0, '7': 0, '8': 0, '9': 0,
	}

	switch elem {
	case ELEMENT_UNKNOWN:
		return counter
	case ELEMENT_MEDAL:
		counter['1'] = 0.5
		counter['2'] = 0.5
		return counter
	case ELEMENT_WOOD:
		counter['4'] = 0.5
		counter['5'] = 0.5
		return counter
	case ELEMENT_WATER:
		counter['6'] = 0.5
		counter['9'] = 0.5
		return counter
	case ELEMENT_FIRE:
		counter['3'] = 0.5
		counter['0'] = 0.5
		return counter
	case ELEMENT_EARTH:
		counter['7'] = 0.5
		counter['8'] = 0.5
		return counter
	}

	return counter
}

func (s Service) Validate(p *Birthday) error {
	//if !p.Solar {
	//	return errors.New("暂不支持传入阴历生日！")
	//}

	if utf8.RuneCountInString(p.Date) != 8 {
		return errors.New("生日格式不正确：生日应该是八个数")
	}

	for i, ch := range p.Date {
		if !digitExist[ch] {
			return fmt.Errorf("生日格式不正确：第%d位不是数字，而是%c", i+1, ch)
		}
	}

	return nil
}

func (s Service) CalcCounter(v string, counter map[rune]float64) string {
	var result []rune

	for _, ch := range v {
		counter[ch] += 1
	}

	for ch, count := range counter {
		if count >= 3 {
			result = append(result, ch)
		}
	}

	slices.Sort(result)

	return string(result)
}

func (s Service) CalcSum(v string, counter map[rune]float64) string {
	result := ""
	value := v
	firstRound := true

	for {
		sum := 0

		for _, ch := range value {
			counter[ch] += 1
			sum += digitNum[ch]
		}

		if firstRound && sum < 11 {
			firstRound = false

			year := 0
			for _, ch := range value[0:4] {
				year += digitNum[ch]
			}

			month := 0
			for _, ch := range value[4:6] {
				month += digitNum[ch]
			}

			day := 0
			for _, ch := range value[6:8] {
				day += digitNum[ch]
			}

			value = strconv.Itoa(year) + strconv.Itoa(month) + strconv.Itoa(day)
			result += value
			continue
		}

		firstRound = false
		value = strconv.Itoa(sum)
		result += value

		if sum < 10 {
			counter[[]rune(value)[0]] += 1
			break
		}
	}

	return result
}

func (s Service) Calc(v string, elem Element) string {
	counter1 := s.InitCounter(ELEMENT_UNKNOWN)
	counter2 := s.InitCounter(elem)
	result := s.CalcCounter(v, counter1) + s.CalcSum(v, counter2)

	var (
		lacks     []rune
		halfLacks []rune
	)

	for ch, count := range counter2 {
		if count == 0 {
			lacks = append(lacks, ch)
		} else if count < 1 {
			halfLacks = append(halfLacks, ch)
		}
	}

	if len(lacks) > 0 {
		if counter1['0'] >= 3 {
			result += "/{"
		} else {
			result += "/"
		}

		slices.Sort(lacks)
		result += string(lacks)
	}

	if len(halfLacks) > 0 {
		if len(lacks) == 0 {
			if counter1['0'] >= 3 {
				result += "/{<"
			} else {
				result += "/<"
			}
		} else {
			result += "<"
		}

		slices.Sort(halfLacks)
		result += string(halfLacks) + ">"
	}

	if counter1['0'] >= 3 {
		result += "}"
	}

	return result
}

func (s Service) ConvertDate(b *Birthday) string {
	var value = b.Date[0:4] + "-" + b.Date[4:6] + "-" + b.Date[6:8]
	var converted string

	if b.Solar {
		converted = solarlunar.SolarToSimpleLuanr(value)
	} else {
		converted = solarlunar.LunarToSolar(value, b.LeapMonthFlag)
	}

	fmt.Println(converted)

	year := utils.SubString(converted, 0, 4)
	month := utils.SubString(converted, 5, 7)
	day := utils.SubString(converted, 8, 10)
	return year + month + day
}

func (s Service) ConvertHour(h int) Element {
	if h < 0 || h >= 24 {
		return ELEMENT_UNKNOWN
	}

	if h >= 21 || h < 1 {
		return ELEMENT_WATER
	}

	if h >= 3 && h < 7 {
		return ELEMENT_WOOD
	}

	if h >= 9 && h < 13 {
		return ELEMENT_FIRE
	}

	if h >= 15 && h < 19 {
		return ELEMENT_MEDAL
	}

	return ELEMENT_EARTH
}

func (s Service) SetBirthDay(params *Birthday) (*Ge, error) {
	var (
		solarGe   string
		lunarGe   string
		solarDate string
		lunarDate string
	)
	elem := s.ConvertHour(params.Hour)

	if params.Solar {
		solarDate = params.Date
		solarGe = s.Calc(solarDate, elem)
		lunarDate = s.ConvertDate(params)
		lunarGe = s.Calc(lunarDate, elem)
	} else {
		lunarDate = params.Date
		lunarGe = s.Calc(params.Date, elem)
		solarDate = s.ConvertDate(params)
		solarGe = s.Calc(solarDate, elem)
	}

	ge := &Ge{
		SolarDate: solarDate,
		LunarDate: lunarDate,
		LeapMonth: params.LeapMonthFlag,
		Hour:      params.Hour,
		Solar:     solarGe,
		Lunar:     lunarGe,
		Element:   elem,
	}

	err := s.DAO.AddGe(ge)
	if err != nil {
		return nil, res.DaoErr1(err)
	}

	return ge, nil
}
