package jgg

import (
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

func initCounter(elem element) map[rune]float64 {
	res := map[rune]float64{
		'0': 0, '1': 0, '2': 0, '3': 0, '4': 0, '5': 0, '6': 0, '7': 0, '8': 0, '9': 0,
	}

	switch elem {
	case ELEMENT_UNKNOWN:
		return res
	case ELEMENT_MEDAL:
		res['1'] = 0.5
		res['2'] = 0.5
		return res
	case ELEMENT_WOOD:
		res['4'] = 0.5
		res['5'] = 0.5
		return res
	case ELEMENT_WATER:
		res['6'] = 0.5
		res['9'] = 0.5
		return res
	case ELEMENT_FIRE:
		res['3'] = 0.5
		res['0'] = 0.5
		return res
	case ELEMENT_EARTH:
		res['7'] = 0.5
		res['8'] = 0.5
		return res
	}

	return res
}

func validate(p *Birthday) error {
	if !p.Solar {
		return errors.New("暂不支持传入阴历生日！")
	}

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

func calcCounter(v string, counter map[rune]float64) string {
	var res []rune

	for _, ch := range v {
		counter[ch] += 1
	}

	for ch, count := range counter {
		if count >= 3 {
			res = append(res, ch)
		}
	}

	slices.Sort(res)

	return string(res)
}

func calcSum(v string, counter map[rune]float64) string {
	res := ""
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
			res += value
			continue
		}

		firstRound = false
		value = strconv.Itoa(sum)
		res += value

		if sum < 10 {
			counter[[]rune(value)[0]] += 1
			break
		}
	}

	return res
}

func calc(v string, elem element) string {
	counter1 := initCounter(ELEMENT_UNKNOWN)
	counter2 := initCounter(elem)
	res := calcCounter(v, counter1) + calcSum(v, counter2)

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
			res += "/{"
		} else {
			res += "/"
		}

		slices.Sort(lacks)
		res += string(lacks)
	}

	if len(halfLacks) > 0 {
		if len(lacks) == 0 {
			if counter1['0'] >= 3 {
				res += "/{<"
			} else {
				res += "/<"
			}
		} else {
			res += "<"
		}

		slices.Sort(halfLacks)
		res += string(halfLacks) + ">"
	}

	if counter1['0'] >= 3 {
		res += "}"
	}

	return res
}

func convertDate(b *Birthday) string {
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

func convertHour(h int) element {
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
