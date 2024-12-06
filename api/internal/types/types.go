// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package types

type AddPolicyReq struct {
	Rules [][]string `json:"rules"`
}

type AddRoleReq struct {
	RoleName string `json:"roleName"`
	ParentId int64  `json:"parentId"`
}

type AddRoleResp struct {
	Role *Role `json:"role"`
}

type ChangePasswordReq struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

type ChangePasswordResp struct {
}

type DeleteUserReq struct {
	UserId int64 `json:"userId"`
}

type DeleteUserResp struct {
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResp struct {
	Token string `json:"token"`
}

type QuerySelfInfoReq struct {
	ID uint64 `json:"id,optional"`
}

type QuerySelfInfoResp struct {
	User QueryUserDetailResp `json:"user"`
}

type QueryUserDetailReq struct {
	Id int64 `path:"id"`
}

type QueryUserDetailResp struct {
	Id        int64  `json:"id"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
	DeletedAt int64  `json:"deletedAt"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Avatar    string `json:"avatar"`
	Role      int64  `json:"role"`
	Status    int32  `json:"status"`
	Nickname  string `json:"nickname"`
	Gender    int32  `json:"gender"`
	Major     string `json:"major"`
	Grade     string `json:"grade"`
	College   string `json:"college"`
	Realname  string `json:"realname"`
	Class     string `json:"class"`
}

type QueryUserListReq struct {
	Page     int32  `json:"page"`
	PageSize int32  `json:"pageSize"`
	Username string `json:"username,optional"`
	Email    string `json:"email,optional"`
	Phone    string `json:"phone,optional"`
	Nickname string `json:"nickname,optional"`
	Role     int64  `json:"role,optional"`
	Status   int32  `json:"status,optional"`
	Major    string `json:"major,optional"`
	Grade    string `json:"grade,optional"`
	College  string `json:"college,optional"`
	Realname string `json:"realname,optional"`
	Class    string `json:"class,optional"`
	Gender   int32  `json:"gender,optional"`
	OrderBy  string `json:"orderBy,optional"`
}

type QueryUserListResp struct {
	Total    int64                  `json:"total"`
	List     []*QueryUserDetailResp `json:"list"`
	Page     int32                  `json:"page"`
	PageSize int32                  `json:"pageSize"`
}

type RegisterReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email,optional"`
	Phone    string `json:"phone,optional"`
	Nickname string `json:"nickname,optional"`
	Role     int64  `json:"role"`
	Gender   int32  `json:"gender"`
	Major    string `json:"major"`
	Grade    string `json:"grade"`
	College  string `json:"college"`
	Realname string `json:"realname"`
	Class    string `json:"class"`
}

type RegisterResp struct {
	User *QueryUserDetailResp `json:"user"`
}

type ResetPasswordReq struct {
	UserId int64 `json:"userId"`
}

type ResetPasswordResp struct {
}

type Role struct {
	RoleName  string `json:"roleName"`
	ParentId  int64  `json:"parentId"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
	ID        int64  `json:"id"`
}

type RoleTree struct {
	Role     *Role       `json:"role"`
	Children []*RoleTree `json:"children"`
}

type RoleTreeReq struct {
	RoleName string `json:"roleName"`
	Page     int32  `json:"page"`
	PageSize int32  `json:"pageSize"`
}

type RoleTreeResp struct {
	Roles    []*RoleTree `json:"roles"`
	Total    int64       `json:"total"`
	Page     int32       `json:"page"`
	PageSize int32       `json:"pageSize"`
}

type SetSelfInfoReq struct {
	ID       uint64 `json:"id,optional"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Avatar   string `json:"avatar"`
	Role     int64  `json:"role"`
	Status   int32  `json:"status"`
	Nickname string `json:"nickname"`
	Gender   int32  `json:"gender"`
	Major    string `json:"major"`
	Grade    string `json:"grade"`
	College  string `json:"college"`
	Realname string `json:"realname"`
	Class    string `json:"class"`
}

type SetSelfInfoResp struct {
}

type SetUserInfoReq struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Avatar   string `json:"avatar"`
	Role     int64  `json:"role"`
	Status   int32  `json:"status"`
	Nickname string `json:"nickname"`
	Gender   int32  `json:"gender"`
	Major    string `json:"major"`
	Grade    string `json:"grade"`
	College  string `json:"college"`
	Realname string `json:"realname"`
	Class    string `json:"class"`
}

type SetUserInfoResp struct {
}

type SetUserRoleReq struct {
	UserId  int64   `json:"userId"`
	RoleIds []int64 `json:"roleIds"`
}

type SetUserRoleResp struct {
}
