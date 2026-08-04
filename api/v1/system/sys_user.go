package system

import (
	"github.com/WJX2001/gin-vue-admin-server/model/common/response"
	systemReq "github.com/WJX2001/gin-vue-admin-server/model/system/request"
	"github.com/WJX2001/gin-vue-admin-server/utils"
	"github.com/gin-gonic/gin"
)

func (b *BaseApi) Register(c *gin.Context) {
	var r systemReq.Register
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(r, utils.RegisterVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
}
