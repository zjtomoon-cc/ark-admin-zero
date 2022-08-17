// Code generated by goctl. DO NOT EDIT.
package types

type LoginReq struct {
	CaptchaId  string `json:"captchaId"`
	VerifyCode string `json:"verifyCode"`
	Account    string `json:"account"`
	Password   string `json:"password"`
}

type LoginResp struct {
	Token string `json:"token"`
}

type UserInfoResp struct {
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

type UserProfileInfoResp struct {
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Gender   int64  `json:"gender"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	Remark   string `json:"remark"`
	Avatar   string `json:"avatar"`
}

type UpdateProfileReq struct {
	Username string `json:"username" validate:"required,min=2,max=12"`
	Nickname string `json:"nickname"`
	Gender   int64  `json:"gender" validate:"gte=0,lte=2"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	Remark   string `json:"remark"`
	Avatar   string `json:"avatar"`
}

type Menu struct {
	Id           int64  `json:"id"`
	ParentId     int64  `json:"parentId"`
	Name         string `json:"name"`
	Router       string `json:"router"`
	Type         int64  `json:"type"`
	Icon         string `json:"icon"`
	OrderNum     int64  `json:"orderNum"`
	ViewPath     string `json:"viewPath"`
	IsShow       int64  `json:"isShow"`
	ActiveRouter string `json:"activeRouter"`
}

type UserPermMenuResp struct {
	Menus []Menu   `json:"menus"`
	Perms []string `json:"perms"`
}

type UpdatePasswordReq struct {
	OldPassword string `json:"oldPassword" validate:"min=6,max=12"`
	NewPassword string `json:"newPassword" validate:"min=6,max=12"`
}

type LoginCaptchaResp struct {
	CaptchaId  string `json:"captchaId"`
	VerifyCode string `json:"verifyCode"`
}

type PermMenu struct {
	Id           int64    `json:"id"`
	ParentId     int64    `json:"parentId"`
	Name         string   `json:"name"`
	Router       string   `json:"router"`
	Perms        []string `json:"perms"`
	Type         int64    `json:"type"`
	Icon         string   `json:"icon"`
	OrderNum     int64    `json:"orderNum"`
	ViewPath     string   `json:"viewPath"`
	IsShow       int64    `json:"isShow"`
	ActiveRouter string   `json:"activeRouter"`
}

type SysPermMenuListResp struct {
	PermMenuList []PermMenu `json:"list"`
}

type AddSysPermMenuReq struct {
	ParentId     int64    `json:"parentId"`
	Name         string   `json:"name"`
	Router       string   `json:"router"`
	Perms        []string `json:"perms"`
	Type         int64    `json:"type"`
	Icon         string   `json:"icon"`
	OrderNum     int64    `json:"orderNum"`
	ViewPath     string   `json:"viewPath"`
	IsShow       int64    `json:"isShow"`
	ActiveRouter string   `json:"activeRouter"`
}

type DeleteSysPermMenuReq struct {
	Id int64 `json:"id"`
}

type UpdateSysPermMenuReq struct {
	Id           int64    `json:"id"`
	ParentId     int64    `json:"parentId"`
	Name         string   `json:"name"`
	Router       string   `json:"router"`
	Perms        []string `json:"perms"`
	Type         int64    `json:"type"`
	Icon         string   `json:"icon"`
	OrderNum     int64    `json:"orderNum"`
	ViewPath     string   `json:"viewPath"`
	IsShow       int64    `json:"isShow"`
	ActiveRouter string   `json:"activeRouter"`
}

type Role struct {
	Id        int64  `json:"id"`
	ParentId  int64  `json:"parentId"`
	Name      string `json:"name"`
	UniqueKey string `json:"uniqueKey"`
	Remark    string `json:"remark"`
	Status    int64  `json:"status"`
	OrderNum  int64  `json:"orderNum"`
}

type SysRoleListResp struct {
	RoleList []Role `json:"list"`
}

type AddSysRoleReq struct {
	ParentId  int64  `json:"parentId"`
	Name      string `json:"name"`
	UniqueKey string `json:"uniqueKey"`
	Remark    string `json:"remark"`
	Status    int64  `json:"status"`
	OrderNum  int64  `json:"orderNum"`
}

type DeleteSysRoleReq struct {
	Id int64 `json:"id"`
}

type UpdateSysRoleReq struct {
	Id        int64  `json:"id"`
	ParentId  int64  `json:"parentId"`
	Name      string `json:"name"`
	UniqueKey string `json:"uniqueKey"`
	Remark    string `json:"remark"`
	Status    int64  `json:"status"`
	OrderNum  int64  `json:"orderNum"`
}

type UpdateSysRolePermMenuReq struct {
	Id          int64   `json:"id"`
	PermMenuIds []int64 `json:"permMenuIds"`
}

type Dept struct {
	Id        int64  `json:"id"`
	ParentId  int64  `json:"parentId"`
	Name      string `json:"name"`
	FullName  string `json:"fullName"`
	UniqueKey string `json:"uniqueKey"`
	Status    int64  `json:"status"`
	OrderNum  int64  `json:"orderNum"`
	Remark    string `json:"remark"`
}

type SysDeptListResp struct {
	DeptList []Dept `json:"list"`
}

type AddSysDeptReq struct {
	ParentId  int64  `json:"parentId"`
	Name      string `json:"name"`
	FullName  string `json:"fullName"`
	UniqueKey string `json:"uniqueKey"`
	Status    int64  `json:"status"`
	OrderNum  int64  `json:"orderNum"`
	Remark    string `json:"remark"`
}

type DeleteSysDeptReq struct {
	Id int64 `json:"id"`
}

type UpdateSysDeptReq struct {
	Id        int64  `json:"id"`
	ParentId  int64  `json:"parentId"`
	Name      string `json:"name"`
	FullName  string `json:"fullName"`
	UniqueKey string `json:"uniqueKey"`
	Status    int64  `json:"status"`
	OrderNum  int64  `json:"orderNum"`
	Remark    string `json:"remark"`
}

type TransferSysDeptReq struct {
	Id       int64 `json:"id"`
	ParentId int64 `json:"parentId"`
}
