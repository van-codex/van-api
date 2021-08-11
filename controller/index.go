package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kainonly/go-bit"
	"lab-api/service"
)

type Index struct {
	admin *service.Admin
}

func NewIndex(admin *service.Admin, bit *bit.Bit) *Index {
	return &Index{admin}
}

func (x *Index) Index(c *gin.Context) {
	data, err := x.admin.FindByUsername("kain")
	if err != nil {
		c.AbortWithStatusJSON(500, err)
		return
	}
	c.JSON(200, gin.H{
		"version": "1.0",
		"data":    data,
	})
}
