package initialize

import (
	//email "github.com/flipped-aurora/gva-plug-email"   // 在线仓库模式
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/email" // 本地插件仓库地址模式
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/example_plugin"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/notify"
	"net/http"

	_ "github.com/flipped-aurora/gin-vue-admin/server/docs"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/router"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// 初始化总路由

func Routers() *gin.Engine {
	var Router = gin.Default()
	Router.StaticFS(global.GVA_CONFIG.Local.Path, http.Dir(global.GVA_CONFIG.Local.Path)) // 为用户头像和文件提供静态地址
	// Router.Use(middleware.LoadTls())  // 打开就能玩https了
	global.GVA_LOG.Info("use middleware logger")
	// 跨域
	//Router.Use(middleware.Cors()) // 如需跨域可以打开
	global.GVA_LOG.Info("use middleware cors")
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.GVA_LOG.Info("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用

	//获取路由组实例
	systemRouter := router.RouterGroupApp.System
	exampleRouter := router.RouterGroupApp.Example
	autocodeRouter := router.RouterGroupApp.Autocode
	PublicGroup := Router.Group("")
	{
		systemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
		systemRouter.InitInitRouter(PublicGroup) // 自动初始化相关
	}
	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		systemRouter.InitApiRouter(PrivateGroup)                    // 注册功能api路由
		systemRouter.InitJwtRouter(PrivateGroup)                    // jwt相关路由
		systemRouter.InitUserRouter(PrivateGroup)                   // 注册用户路由
		systemRouter.InitMenuRouter(PrivateGroup)                   // 注册menu路由
		systemRouter.InitSystemRouter(PrivateGroup)                 // system相关路由
		systemRouter.InitCasbinRouter(PrivateGroup)                 // 权限相关路由
		systemRouter.InitAutoCodeRouter(PrivateGroup)               // 创建自动化代码
		systemRouter.InitAuthorityRouter(PrivateGroup)              // 注册角色路由
		systemRouter.InitSysDictionaryRouter(PrivateGroup)          // 字典管理
		systemRouter.InitSysOperationRecordRouter(PrivateGroup)     // 操作记录
		systemRouter.InitSysDictionaryDetailRouter(PrivateGroup)    // 字典详情管理
		exampleRouter.InitFileUploadAndDownloadRouter(PrivateGroup) // 文件上传下载功能路由
		exampleRouter.InitExcelRouter(PrivateGroup)                 // 表格导入导出
		exampleRouter.InitSimpleUploaderRouter(PrivateGroup)        // 断点续传（插件版）
		exampleRouter.InitCustomerRouter(PrivateGroup)              // 客户路由

		// Code generated by github.com/flipped-aurora/gin-vue-admin/server Begin; DO NOT EDIT.
		autocodeRouter.InitSysAutoCodeExampleRouter(PrivateGroup)
		// Code generated by github.com/flipped-aurora/gin-vue-admin/server End; DO NOT EDIT.
	}

	//  添加开放权限的插件 示例
	PluginInit(PublicGroup, example_plugin.ExamplePlugin)

	//  钉钉通知，暂时开放权限
	PluginInit(PublicGroup, notify.CreateDDPlug(
		global.GVA_CONFIG.DingDing.Url,
		global.GVA_CONFIG.DingDing.Secret,
		global.GVA_CONFIG.DingDing.Token,
	))

	//  添加跟角色挂钩权限的插件 示例 本地示例模式于在线仓库模式注意上方的import 可以自行切换 效果相同
	PluginInit(PrivateGroup, email.CreateEmailPlug(
		global.GVA_CONFIG.Email.To,
		global.GVA_CONFIG.Email.From,
		global.GVA_CONFIG.Email.Host,
		global.GVA_CONFIG.Email.Secret,
		global.GVA_CONFIG.Email.Nickname,
		global.GVA_CONFIG.Email.Port,
		global.GVA_CONFIG.Email.IsSSL,
	))

	global.GVA_LOG.Info("router register success")
	return Router
}
