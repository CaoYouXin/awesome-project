package jgg

import (
	"awesomeProject/internal/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

func SetBirthDay(c *gin.Context) {
	params := &Birthday{}

	if err := c.ShouldBind(params); err != nil {
		_ = c.Error(fmt.Errorf("%w: %s(%s)", res.BadRequestErr, "要求生日类型JSON", err.Error()))
		return
	}

	if err := validate(params); err != nil {
		_ = c.Error(fmt.Errorf("%w: %s(%s)", res.BadRequestErr, "参数校验不通过", err.Error()))
		return
	}

	var (
		solarGe   string
		lunarGe   string
		solarDate string
		lunarDate string
	)
	elem := convertHour(params.Hour)

	if params.Solar {
		solarDate = params.Date
		solarGe = calc(solarDate, elem)
		lunarDate = convertDate(params)
		lunarGe = calc(lunarDate, elem)
	} else {
		lunarDate = params.Date
		lunarGe = calc(params.Date, elem)
		solarDate = convertDate(params)
		solarGe = calc(solarDate, elem)
	}

	res.Success(c, &Ge{
		Solar:     solarGe,
		Lunar:     lunarGe,
		SolarDate: solarDate,
		LunarDate: lunarDate,
		Element:   elem,
	})
}
