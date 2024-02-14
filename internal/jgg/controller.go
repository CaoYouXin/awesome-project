package jgg

import (
	"awesomeProject/internal/res"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Controller struct {
	Service
}

func (jgg Controller) SetBirthDay(c *gin.Context) {
	params := &Birthday{}

	if err := c.ShouldBind(params); err != nil {
		_ = c.Error(res.BadRequestErr2(err, "要求生日类型JSON"))
		return
	}

	if err := jgg.Service.Validate(params); err != nil {
		_ = c.Error(res.BadRequestErr2(err, "参数校验不通过"))
		return
	}

	ge, err := jgg.Service.SetBirthDay(params)
	if err != nil {
		_ = c.Error(res.InternalErr1(res.ServiceErr1(err)))
		return
	}

	res.Success(c, ge)
}

func (jgg Controller) ListGe(c *gin.Context) {
	list, err := jgg.DAO.ListGe()
	if err != nil {
		_ = c.Error(res.InternalErr1(res.DaoErr1(err)))
		return
	}

	res.Success(c, list)
}

func (jgg Controller) DelGe(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		_ = c.Error(res.BadRequestErr2(err, "URL 参数不正确"))
		return
	}

	suc, err := jgg.DAO.DelGe(id)
	if err != nil {
		_ = c.Error(res.BadRequestErr1(res.DaoErr1(err)))
		return
	}

	res.Success(c, suc)
}
