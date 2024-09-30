package router

import (
	"github.com/deeptest-com/deeptest/cmd/server/v1/handler"
	"github.com/deeptest-com/deeptest/internal/pkg/core/module"
	"github.com/deeptest-com/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type UserModule struct {
	UserCtrl *handler.UserCtrl `inject:""`
}

// Party 用户
func (m *UserModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin())

		index.Get("/", m.UserCtrl.ListAll).Name = "用户列表"
		index.Get("/{id:uint}", m.UserCtrl.GetUser).Name = "用户详情"
		index.Post("/", m.UserCtrl.CreateUser).Name = "新建用户"
		index.Post("/{id:uint}", m.UserCtrl.UpdateUser).Name = "编辑用户"
		index.Delete("/{id:uint}", m.UserCtrl.DeleteUser).Name = "删除用户"

		index.Post("/invite", m.UserCtrl.Invite).Name = "邀请用户"
		index.Post("/updateEmail", m.UserCtrl.UpdateEmail).Name = "修改邮箱"
		index.Post("/updateName", m.UserCtrl.UpdateName).Name = "修改名称"
		index.Post("/updatePassword", m.UserCtrl.UpdatePassword).Name = "修改密码"

		index.Get("/profile", m.UserCtrl.Profile).Name = "个人信息"
		index.Get("/message", m.UserCtrl.Message).Name = "消息"
		index.Get("/logout", m.UserCtrl.Logout).Name = "退出"
		index.Get("/clear", m.UserCtrl.Clear).Name = "清空 token"
		index.Post("/change_avatar", m.UserCtrl.ChangeAvatar).Name = "修改头像"
		index.Get("/usersNotExistedInProject", m.UserCtrl.GetUsersNotExistedInProject).Name = "获取项目中没有的用户列表"
		index.Post("/changeUserSysRole", m.UserCtrl.ChangeUserSysRole).Name = "更新用户系统角色"

		// index.GetDetail("/expire", controllers.Expire).Title = "刷新 token"
	}
	return module.NewModule("/users", handler)
}
