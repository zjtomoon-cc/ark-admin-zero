// Code generated by goctl. DO NOT EDIT.
package types

type PageReq struct {
	Page  uint64 `form:"page"  validate:"number,gte=1"  label:"页数"`
	Limit uint64 `form:"limit" validate:"number,gte=1"  label:"条数"`
}

type Pagination struct {
	Page  uint64 `json:"page"`
	Limit uint64 `json:"limit"`
	Total uint64 `json:"total"`
}

type LoginReq struct {
	CaptchaId  string `json:"captchaId"   label:"验证码id"`
	VerifyCode string `json:"verifyCode"  label:"验证码"`
	Account    string `json:"account"     label:"账号"`
	Password   string `json:"password"    label:"密码"`
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
	Gender   uint64 `json:"gender"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	Remark   string `json:"remark"`
	Avatar   string `json:"avatar"`
}

type UpdateProfileReq struct {
	Username string `json:"username"  validate:"required,min=2,max=12"   label:"姓名"`
	Nickname string `json:"nickname"  validate:"omitempty,min=2,max=12"  label:"昵称"`
	Gender   uint64 `json:"gender"    validate:"gte=0,lte=2"             label:"性别"`
	Email    string `json:"email"     validate:"omitempty,email"         label:"邮箱"`
	Mobile   string `json:"mobile"    validate:"omitempty,len=11"        label:"手机号"`
	Avatar   string `json:"avatar"    validate:"required,url"            label:"头像"`
}

type Menu struct {
	Id           uint64 `json:"id"`
	ParentId     uint64 `json:"parentId"`
	Name         string `json:"name"`
	Router       string `json:"router"`
	Type         uint64 `json:"type"`
	Icon         string `json:"icon"`
	OrderNum     uint64 `json:"orderNum"`
	ViewPath     string `json:"viewPath"`
	IsShow       uint64 `json:"isShow"`
	ActiveRouter string `json:"activeRouter"`
}

type UserPermMenuResp struct {
	Menus []Menu   `json:"menus"`
	Perms []string `json:"perms"`
}

type UpdatePasswordReq struct {
	OldPassword string `json:"oldPassword"  validate:"min=6,max=12"  label:"旧密码"`
	NewPassword string `json:"newPassword"  validate:"min=6,max=12"  label:"新密码"`
}

type LoginCaptchaResp struct {
	CaptchaId  string `json:"captchaId"`
	VerifyCode string `json:"verifyCode"`
}

type GenerateAvatarResp struct {
	AvatarUrl string `json:"avatarUrl"`
}

type PermMenu struct {
	Id           uint64   `json:"id"`
	ParentId     uint64   `json:"parentId"`
	Name         string   `json:"name"`
	Router       string   `json:"router"`
	Perms        []string `json:"perms"`
	Type         uint64   `json:"type"`
	Icon         string   `json:"icon"`
	OrderNum     uint64   `json:"orderNum"`
	ViewPath     string   `json:"viewPath"`
	IsShow       uint64   `json:"isShow"`
	ActiveRouter string   `json:"activeRouter"`
	Has          uint64   `json:"has"`
}

type SysPermMenuListResp struct {
	PermMenuList []PermMenu `json:"list"`
}

type AddSysPermMenuReq struct {
	ParentId     uint64   `json:"parentId"      validate:"number,gte=0"           label:"父级菜单id"`
	Name         string   `json:"name"          validate:"min=2,max=50"           label:"菜单名称"`
	Router       string   `json:"router"        validate:"omitempty,max=1024"     label:"路由"`
	Perms        []string `json:"perms"         validate:"omitempty,unique"       label:"权限"`
	Type         uint64   `json:"type"          validate:"number,gte=0,lte=2"     label:"类型"`
	Icon         string   `json:"icon"          validate:"omitempty,max=200"      label:"图标"`
	OrderNum     uint64   `json:"orderNum"      validate:"number,gte=0,lte=9999"  label:"排序"`
	ViewPath     string   `json:"viewPath"      validate:"omitempty,max=1024"     label:"视图路径"`
	IsShow       uint64   `json:"isShow"        validate:"number,gte=0,lte=1"     label:"显示状态"`
	ActiveRouter string   `json:"activeRouter"  validate:"omitempty,max=1024"     label:"激活路由"`
}

type DeleteSysPermMenuReq struct {
	Id uint64 `json:"id"  validate:"number,gte=1" label:"菜单id"`
}

type UpdateSysPermMenuReq struct {
	Id           uint64   `json:"id"            validate:"number,gte=1"           label:"菜单id"`
	ParentId     uint64   `json:"parentId"      validate:"number,gte=0"           label:"父级菜单id"`
	Name         string   `json:"name"          validate:"min=2,max=50"           label:"菜单名称"`
	Router       string   `json:"router"        validate:"omitempty,max=1024"     label:"路由"`
	Perms        []string `json:"perms"         validate:"omitempty,unique"       label:"权限"`
	Type         uint64   `json:"type"          validate:"number,gte=0,lte=2"     label:"类型"`
	Icon         string   `json:"icon"          validate:"omitempty,max=200"      label:"图标"`
	OrderNum     uint64   `json:"orderNum"      validate:"number,gte=0,lte=9999"  label:"排序"`
	ViewPath     string   `json:"viewPath"      validate:"omitempty,max=1024"     label:"视图路径"`
	IsShow       uint64   `json:"isShow"        validate:"number,gte=0,lte=1"     label:"显示状态"`
	ActiveRouter string   `json:"activeRouter"  validate:"omitempty,max=1024"     label:"激活路由"`
}

type Role struct {
	Id          uint64   `json:"id"`
	ParentId    uint64   `json:"parentId"`
	Name        string   `json:"name"`
	UniqueKey   string   `json:"uniqueKey"`
	PermMenuIds []uint64 `json:"permMenuIds"`
	Remark      string   `json:"remark"`
	Status      uint64   `json:"status"`
	OrderNum    uint64   `json:"orderNum"`
}

type SysRoleListResp struct {
	RoleList []Role `json:"list"`
}

type AddSysRoleReq struct {
	ParentId    uint64   `json:"parentId"     validate:"number,gte=0"          label:"父级角色id"`
	Name        string   `json:"name"         validate:"min=2,max=50"          label:"角色名称"`
	UniqueKey   string   `json:"uniqueKey"    validate:"min=2,max=50"          label:"角色标识"`
	PermMenuIds []uint64 `json:"permMenuIds"  validate:"omitempty,unique"      label:"权限ids"`
	Remark      string   `json:"remark"       validate:"max=200"               label:"备注"`
	Status      uint64   `json:"status"       validate:"number,gte=0,lte=1"    label:"状态"`
	OrderNum    uint64   `json:"orderNum"     validate:"number,gte=0,lte=9999" label:"排序"`
}

type DeleteSysRoleReq struct {
	Id uint64 `json:"id"  validate:"number,gte=2" label:"角色id"`
}

type UpdateSysRoleReq struct {
	Id          uint64   `json:"id"           validate:"number,gte=1"           label:"角色id"`
	ParentId    uint64   `json:"parentId"     validate:"number,gte=0"           label:"父级角色id"`
	Name        string   `json:"name"         validate:"min=2,max=50"           label:"角色名称"`
	UniqueKey   string   `json:"uniqueKey"    validate:"min=2,max=50"           label:"角色标识"`
	PermMenuIds []uint64 `json:"permMenuIds"  validate:"omitempty,unique"       label:"权限ids"`
	Remark      string   `json:"remark"       validate:"max=200"                label:"备注"`
	Status      uint64   `json:"status"       validate:"number,gte=0,lte=1"     label:"状态"`
	OrderNum    uint64   `json:"orderNum"     validate:"number,gte=0,lte=9999"  label:"排序"`
}

type Dept struct {
	Id        uint64 `json:"id"`
	ParentId  uint64 `json:"parentId"`
	Name      string `json:"name"`
	FullName  string `json:"fullName"`
	UniqueKey string `json:"uniqueKey"`
	Type      uint64 `json:"type"`
	Status    uint64 `json:"status"`
	OrderNum  uint64 `json:"orderNum"`
	Remark    string `json:"remark"`
}

type SysDeptListResp struct {
	DeptList []Dept `json:"list"`
}

type AddSysDeptReq struct {
	ParentId  uint64 `json:"parentId"   validate:"number,gte=0"            label:"父级部门id"`
	Name      string `json:"name"       validate:"min=2,max=50"            label:"部门名称"`
	FullName  string `json:"fullName"   validate:"omitempty,min=2,max=50"  label:"部门全称"`
	UniqueKey string `json:"uniqueKey"  validate:"min=2,max=50"            label:"部门标识"`
	Type      uint64 `json:"type"       validate:"number,gte=1,lte=3"      label:"部门类型"`
	Status    uint64 `json:"status"     validate:"number,gte=0,lte=1"      label:"状态"`
	OrderNum  uint64 `json:"orderNum"   validate:"number,gte=0,lte=9999"   label:"排序"`
	Remark    string `json:"remark"     validate:"max=200"                 label:"备注"`
}

type DeleteSysDeptReq struct {
	Id uint64 `json:"id"  validate:"number,gte=1" label:"部门id"`
}

type UpdateSysDeptReq struct {
	Id        uint64 `json:"id"         validate:"number,gte=1"            label:"部门id"`
	ParentId  uint64 `json:"parentId"   validate:"number,gte=0"            label:"父级部门id"`
	Name      string `json:"name"       validate:"min=2,max=50"            label:"部门名称"`
	FullName  string `json:"fullName"   validate:"omitempty,min=2,max=50"  label:"部门全称"`
	UniqueKey string `json:"uniqueKey"  validate:"min=2,max=50"            label:"部门标识"`
	Type      uint64 `json:"type"       validate:"number,gte=1,lte=3"      label:"部门类型"`
	Status    uint64 `json:"status"     validate:"number,gte=0,lte=1"      label:"状态"`
	OrderNum  uint64 `json:"orderNum"   validate:"number,gte=0,lte=9999"   label:"排序"`
	Remark    string `json:"remark"     validate:"max=200"                 label:"备注"`
}

type Job struct {
	Id       uint64 `json:"id"`
	Name     string `json:"name"`
	Status   uint64 `json:"status"`
	OrderNum uint64 `json:"orderNum"`
}

type SysJobPageReq struct {
	PageReq
}

type SysJobPageResp struct {
	JobList    []Job      `json:"list"`
	Pagination Pagination `json:"pagination"`
}

type AddSysJobReq struct {
	Name     string `json:"name"      validate:"min=2,max=50"           label:"岗位名称"`
	Status   uint64 `json:"status"    validate:"number,gte=0,lte=1"     label:"状态"`
	OrderNum uint64 `json:"orderNum"  validate:"number,gte=0,lte=9999"  label:"排序"`
}

type DeleteSysJobReq struct {
	Id uint64 `json:"id"  validate:"number,gte=1" label:"岗位id"`
}

type UpdateSysJobReq struct {
	Id       uint64 `json:"id"        validate:"number,gte=1"           label:"岗位id"`
	Name     string `json:"name"      validate:"min=2,max=50"           label:"岗位名称"`
	Status   uint64 `json:"status"    validate:"number,gte=0,lte=1"     label:"状态"`
	OrderNum uint64 `json:"orderNum"  validate:"number,gte=0,lte=9999"  label:"排序"`
}

type Profession struct {
	Id       uint64 `json:"id"`
	Name     string `json:"name"`
	Status   uint64 `json:"status"`
	OrderNum uint64 `json:"orderNum"`
}

type SysProfessionPageReq struct {
	PageReq
}

type SysProfessionPageResp struct {
	ProfessionList []Profession `json:"list"`
	Pagination     Pagination   `json:"pagination"`
}

type AddSysProfessionReq struct {
	Name     string `json:"name"      validate:"min=2,max=50"           label:"职称"`
	Status   uint64 `json:"status"    validate:"number,gte=0,lte=1"     label:"状态"`
	OrderNum uint64 `json:"orderNum"  validate:"number,gte=0,lte=9999"  label:"排序"`
}

type DeleteSysProfessionReq struct {
	Id uint64 `json:"id"  validate:"number,gte=1" label:"职称id"`
}

type UpdateSysProfessionReq struct {
	Id       uint64 `json:"id"        validate:"number,gte=1"           label:"职称id"`
	Name     string `json:"name"      validate:"min=2,max=50"           label:"职称"`
	Status   uint64 `json:"status"    validate:"number,gte=0,lte=1"     label:"状态"`
	OrderNum uint64 `json:"orderNum"  validate:"number,gte=0,lte=9999"  label:"排序"`
}

type UserProfession struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

type UserJob struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

type UserDept struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

type UserRole struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

type User struct {
	Id         uint64         `json:"id"`
	Account    string         `json:"account"`
	Username   string         `json:"username"`
	Nickname   string         `json:"nickname"`
	Gender     uint64         `json:"gender"`
	Email      string         `json:"email"`
	Mobile     string         `json:"mobile"`
	Profession UserProfession `json:"profession"`
	Job        UserJob        `json:"job"`
	Dept       UserDept       `json:"dept"`
	Roles      []UserRole     `json:"roles"`
	Status     uint64         `json:"status"`
	OrderNum   uint64         `json:"orderNum"`
	Remark     string         `json:"remark"`
}

type SysUserPageReq struct {
	PageReq
	DeptId uint64 `form:"deptId"  validate:"number,gte=0" label:"部门id"`
}

type SysUserPageResp struct {
	UserList   []User     `json:"list"`
	Pagination Pagination `json:"pagination"`
}

type AddSysUserReq struct {
	Account      string   `json:"account"       validate:"min=4,max=50"            label:"账号"`
	Username     string   `json:"username"      validate:"min=2,max=50"            label:"姓名"`
	Nickname     string   `json:"nickname"      validate:"omitempty,min=2,max=50"  label:"昵称"`
	Gender       uint64   `json:"gender"        validate:"number,gte=0,lte=2"      label:"性别"`
	Email        string   `json:"email"         validate:"omitempty,email"         label:"邮箱"`
	Mobile       string   `json:"mobile"        validate:"omitempty,min=11"        label:"手机号"`
	ProfessionId uint64   `json:"professionId"  validate:"number,gte=1"            label:"职称id"`
	JobId        uint64   `json:"jobId"         validate:"number,gte=1"            label:"岗位id"`
	DeptId       uint64   `json:"deptId"        validate:"number,gte=1"            label:"部门id"`
	RoleIds      []uint64 `json:"roleIds"       validate:"unique"                  label:"角色ids"`
	Status       uint64   `json:"status"        validate:"number,gte=0,lte=1"      label:"状态"`
	OrderNum     uint64   `json:"orderNum"      validate:"number,gte=0,lte=9999"   label:"排序"`
	Remark       string   `json:"remark"        validate:"max=200"                 label:"备注"`
}

type DeleteSysUserReq struct {
	Id uint64 `json:"id"  validate:"number,gte=2" label:"用户id"`
}

type UpdateSysUserReq struct {
	Id           uint64   `json:"id"            validate:"number,gte=2"            label:"用户id"`
	Username     string   `json:"username"      validate:"min=2,max=50"            label:"姓名"`
	Nickname     string   `json:"nickname"      validate:"omitempty,min=2,max=50"  label:"昵称"`
	Gender       uint64   `json:"gender"        validate:"number,gte=0,lte=2"      label:"性别"`
	Email        string   `json:"email"         validate:"omitempty,email"         label:"邮箱"`
	Mobile       string   `json:"mobile"        validate:"omitempty,min=11"        label:"手机号"`
	ProfessionId uint64   `json:"professionId"  validate:"number,gte=1"            label:"职称id"`
	JobId        uint64   `json:"jobId"         validate:"number,gte=1"            label:"岗位id"`
	DeptId       uint64   `json:"deptId"        validate:"number,gte=1"            label:"部门id"`
	RoleIds      []uint64 `json:"roleIds"       validate:"unique"                  label:"角色ids"`
	Status       uint64   `json:"status"        validate:"number,gte=0,lte=1"      label:"状态"`
	OrderNum     uint64   `json:"orderNum"      validate:"number,gte=0,lte=9999"   label:"排序"`
	Remark       string   `json:"remark"        validate:"max=200"                 label:"备注"`
}

type UpdateSysUserPasswordReq struct {
	Id       uint64 `json:"id"        validate:"number,gte=2"  label:"用户id"`
	Password string `json:"password"  validate:"min=6,max=12"  label:"密码"`
}

type GetSysUserRdpjInfoReq struct {
	UserId uint64 `form:"userId"  validate:"number,gte=0"  label:"用户id"`
}

type Rdpj struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

type DeptTree struct {
	Id       uint64 `json:"id"`
	ParentId uint64 `json:"parentId"`
	Name     string `json:"name"`
}

type RoleTree struct {
	Id       uint64 `json:"id"`
	ParentId uint64 `json:"parentId"`
	Name     string `json:"name"`
	Has      uint64 `json:"has"`
}

type GetSysUserRdpjInfoResp struct {
	Role       []RoleTree `json:"role"`
	Dept       []DeptTree `json:"dept"`
	Profession []Rdpj     `json:"profession"`
	Job        []Rdpj     `json:"job"`
}

type ConfigDict struct {
	Id        uint64 `json:"id"`
	ParentId  uint64 `json:"parentId"`
	Name      string `json:"name"`
	Type      uint64 `json:"type"`
	UniqueKey string `json:"uniqueKey"`
	Value     string `json:"value"`
	OrderNum  uint64 `json:"orderNum"`
	Remark    string `json:"remark"`
	Status    uint64 `json:"status"`
}

type ConfigDictListResp struct {
	DictList []ConfigDict `json:"list"`
}

type ConfigDictPageReq struct {
	PageReq
	ParentId uint64 `form:"parentId"  validate:"number,gte=0" label:"字典集id"`
}

type ConfigDictPageResp struct {
	ConfigDictList []ConfigDict `json:"list"`
	Pagination     Pagination   `json:"pagination"`
}

type AddConfigDictReq struct {
	ParentId  uint64 `json:"parentId"   validate:"number,gte=0"         label:"字典集id"`
	Name      string `json:"name"       validate:"min=2,max=50"         label:"名称"`
	Type      uint64 `json:"type"       validate:"number,gte=1,lte=12"  label:"类型"`
	UniqueKey string `json:"uniqueKey"  validate:"min=2,max=50"         label:"标识"`
	Value     string `json:"value"      validate:"max=2048"             label:"字典项值"`
	OrderNum  uint64 `json:"orderNum"   validate:"gte=0,lte=9999"       label:"排序"`
	Remark    string `json:"remark"     validate:"max=200"              label:"备注"`
	Status    uint64 `json:"status"     validate:"number,gte=0,lte=1"   label:"状态"`
}

type DeleteConfigDictReq struct {
	Id uint64 `json:"id"  validate:"number,gte=1" label:"字典id"`
}

type UpdateConfigDictReq struct {
	Id       uint64 `json:"id"         validate:"number,gte=1"         label:"字典id"`
	ParentId uint64 `json:"parentId"   validate:"number,gte=0"         label:"字典集id"`
	Name     string `json:"name"       validate:"min=2,max=50"         label:"名称"`
	Type     uint64 `json:"type"       validate:"number,gte=1,lte=12"  label:"类型"`
	Value    string `json:"value"      validate:"max=2048"             label:"字典项值"`
	OrderNum uint64 `json:"orderNum"   validate:"gte=0,lte=9999"       label:"排序"`
	Remark   string `json:"remark"     validate:"max=200"              label:"备注"`
	Status   uint64 `json:"status"     validate:"number,gte=0,lte=1"   label:"状态"`
}

type LogLogin struct {
	Id         uint64 `json:"id"`
	Account    string `json:"account"`
	Ip         string `json:"ip"`
	Uri        string `json:"uri"`
	Status     uint64 `json:"status"`
	CreateTime string `json:"createTime"`
}

type LogLoginPageReq struct {
	PageReq
}

type LogLoginPageResp struct {
	LogLoginList []LogLogin `json:"list"`
	Pagination   Pagination `json:"pagination"`
}
