package system

import (
	"github.com/WJX2001/gin-vue-admin-server/global"
	"github.com/google/uuid"
	"time"
)

type SysUser struct {
	global.GVA_MODEL
	UUID      uuid.UUID `json:"uuid" gorm:"index;comment:用户UUID"`
	Username  string    `json:"userName" gorm:"index;comment:用户登陆名"`
	Password  string    `json:"-"  gorm:"comment:用户登录密码"`
	NickName  string    `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`
	HeaderImg string    `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"`
	//AuthorityId uint      `json:"authorityId" gorm:"default:888;comment:用户角色ID"`
	//Authority   SysAuthority   `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"` // 用户角色
	//Authorities []SysAuthority `json:"authorities" gorm:"many2many:sys_user_authority;"`
	//DeptId             uint            `json:"deptId" gorm:"comment:主部门ID(数据归属/盖章)"`                                                               // 主部门ID(数据归属)
	//Dept               SysDepartment   `json:"dept" form:"-" gorm:"foreignKey:DeptId;references:ID;comment:主部门"`                                   // 主部门;form:"-" 阻断 gin 绑定递归(与 SysDepartment.Leader 成环会栈溢出)
	//Departments        []SysDepartment `json:"departments" gorm:"many2many:sys_user_departments;"`                                                 // 多部门归属(数据可见范围)
	//Positions          []SysPosition   `json:"positions" gorm:"many2many:sys_user_positions;"`
	Phone  string `json:"phone"  gorm:"comment:用户手机号"`                     // 用户手机号
	Email  string `json:"email"  gorm:"comment:用户邮箱"`                      // 用户邮箱
	Enable int    `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"` //用户是否被冻结 1正常 2冻结
	//OriginSetting      common.JSONMap `json:"originSetting" form:"originSetting" gorm:"type:text;default:null;column:origin_setting;comment:配置;"` //配置
	PasswordUpdatedAt  *time.Time `json:"passwordUpdatedAt" gorm:"comment:密码最后修改时间"` //密码最后修改时间
	MustChangePassword bool       `json:"-" gorm:"default:false;comment:是否必须修改初始密码"` //是否必须修改初始密码
}
