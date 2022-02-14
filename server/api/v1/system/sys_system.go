package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SystemApi struct{}

// @Tags System
// @Summary 获取配置文件内容
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {object} response.Response{data=systemRes.SysConfigResponse,msg=string} "获取配置文件内容,返回包括系统配置"
// @Router /system/getSystemConfig [post]
func (s *SystemApi) GetSystemConfig(c *gin.Context) {
	if err, config := systemConfigService.GetSystemConfig(); err != nil {
		global.GVA_LOG.Error(global.Translate("general.getDataFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.getDataFailErr"), c)
	} else {
		response.OkWithDetailed(systemRes.SysConfigResponse{Config: config}, global.Translate("general.getDataSuccess"), c)
	}
}

// @Tags System
// @Summary 设置配置文件内容
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body system.System true "设置配置文件内容"
// @Success 200 {object} response.Response{data=string} "设置配置文件内容"
// @Router /system/setSystemConfig [post]
func (s *SystemApi) SetSystemConfig(c *gin.Context) {
	var sys system.System
	_ = c.ShouldBindJSON(&sys)
	if err := systemConfigService.SetSystemConfig(sys); err != nil {
		global.GVA_LOG.Error(global.Translate("general.setupFailErr"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.setupFail"), c)
	} else {
		response.OkWithData(global.Translate("general.setupSuccess"), c)
	}
}

// @Tags System
// @Summary 重启系统
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {object} response.Response{msg=string} "重启系统"
// @Router /system/reloadSystem [post]
func (s *SystemApi) ReloadSystem(c *gin.Context) {
	err := utils.Reload()
	if err != nil {
		global.GVA_LOG.Error(global.Translate("sys_system.rebootFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("sys_system.rebootFailErr"), c)
	} else {
		response.OkWithMessage(global.Translate("sys_system.rebootSuccess"), c)
	}
}

// @Tags System
// @Summary 获取服务器信息
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "获取服务器信息"
// @Router /system/getServerInfo [post]
func (s *SystemApi) GetServerInfo(c *gin.Context) {
	if server, err := systemConfigService.GetServerInfo(); err != nil {
		global.GVA_LOG.Error(global.Translate("general.getDataFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.getDataFailErr"), c)
	} else {
		response.OkWithDetailed(gin.H{"server": server}, global.Translate("general.getDataSuccess"), c)
	}
}
