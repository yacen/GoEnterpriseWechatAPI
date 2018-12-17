package GoEnterpriseWechatAPI

/**
成员管理
*/

type BaseUser struct {
	UserId           string          `json:"userid"`
	Name             string          `json:"name"`
	Alias            string          `json:"alias"`
	Mobile           string          `json:"mobile"`
	Department       []int           `json:"department"`
	Order            []int           `json:"order"`
	Position         string          `json:"position"`
	Gender           string          `json:"gender"`
	Email            string          `json:"email"`
	IsLeader         int             `json:"isleader"`
	Telephone        string          `json:"telephone"`
	Enable           int             `json:"enable"`
	ExtAttr          ExtAttr         `json:"extattr"`
	QrCode           string          `json:"qr_code"`
	ExternalPosition string          `json:"external_position"`
	ExternalProfile  ExternalProfile `json:"external_profile"`
}

/*
成员对外信息
https://work.weixin.qq.com/api/doc#90000/90135/90223
*/
type ExternalProfile struct {
	ExternalAttr []ExternalAttrItem `json:"external_attr"`
}

type ExternalAttrItem struct {
	Type        int                 `json:"type"`
	Name        string              `json:"name"`
	Text        ExternalText        `json:"text"`
	Web         ExternalWeb         `json:"web"`
	MiniProgram ExternalMiniProgram `json:"miniprogram"`
}

type ExternalText struct {
	Value string `json:"value"`
}

type ExternalWeb struct {
	Url   string `json:"url"`
	Title string `json:"title"`
}

type ExternalMiniProgram struct {
	AppId    string `json:"appid"`
	Title    string `json:"title"`
	PagePath string `json:"pagepath"`
}

type ExtAttr struct {
	Attrs []Attrs `json:"attrs"`
}

type Attrs struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// ************************************* 创建成员 ****************************************
/*
创建成员请求数据
https://work.weixin.qq.com/api/doc#90000/90135/90195
*/
type UserCreateRequest struct {
	BaseUser
	AvatarMediaId string `json:"avatar_mediaid"`
	ToInvite      bool   `json:"to_invite"`
}

func (this *Api) UserCreate(user *UserCreateRequest) (result *BaseResponse, err error) {
	token, err := this.EnsureAccessToken()
	if err != nil {
		return
	}
	result = new(BaseResponse)
	url := CgiBinPrefix + "/user/create?access_token=" + token.TokenStr
	err = this.httpClient.PostObjectGetObject(url, user, result)
	return
}

// ************************************* 读取成员 ****************************************
/**
https://work.weixin.qq.com/api/doc#90000/90135/90196
*/
type UserGetResult struct {
	BaseResponse
	BaseUser
	Avatar string `json:"avatar"`
	Status string `json:"status"`
}

func (this *Api) UserGet(userId string) (result *UserGetResult, err error) {
	token, err := this.EnsureAccessToken()
	if err != nil {
		return
	}
	result = new(UserGetResult)
	url := CgiBinPrefix + "/user/get?access_token=" + token.TokenStr + "&userid=" + userId
	err = this.httpClient.GetObject(url, result)
	return
}

// ************************************* 更新成员 ****************************************
/**
https://work.weixin.qq.com/api/doc#90000/90135/90197
*/
type UserUpdateRequest struct {
	BaseUser
	AvatarMediaId string `json:"avatar_mediaid"`
}

func (this *Api) UserUpdate(user *UserUpdateRequest) (result *BaseResponse, err error) {
	token, err := this.EnsureAccessToken()
	if err != nil {
		return
	}
	result = new(BaseResponse)
	url := CgiBinPrefix + "/user/update?access_token=" + token.TokenStr
	err = this.httpClient.PostObjectGetObject(url, user, result)
	return
}

// ************************************* 删除成员 ****************************************
/**
https://work.weixin.qq.com/api/doc#90000/90135/90198
*/
func (this *Api) UserDelete(userId string) (result *BaseResponse, err error) {
	token, err := this.EnsureAccessToken()
	if err != nil {
		return
	}
	result = new(BaseResponse)
	url := CgiBinPrefix + "/user/delete?access_token=" + token.TokenStr + "&userid=" + userId
	err = this.httpClient.GetObject(url, result)
	return
}

// ************************************* 批量删除成员 ****************************************
/**
https://work.weixin.qq.com/api/doc#90000/90135/90199
*/
func (this *Api) UserBatchDelete(userIds []string) (result *BaseResponse, err error) {
	token, err := this.EnsureAccessToken()
	if err != nil {
		return
	}
	result = new(BaseResponse)
	url := CgiBinPrefix + "/user/delete?access_token=" + token.TokenStr
	err = this.httpClient.PostObjectGetObject(url, userIds, result)
	return
}

// ************************************* 获取部门成员 ****************************************
/**
https://work.weixin.qq.com/api/doc#90000/90135/90200
*/
type UserSimpleListResult struct {
	BaseResponse
	UserList []UserSimpleItem `json:"userlist"`
}

type UserSimpleItem struct {
	UserId     string   `json:"userid"`
	Name       string   `json:"name"`
	Department []string `json:"department"`
}

func (this *Api) UserSimpleList(departmentId string, fetchChild string) (result *UserSimpleListResult, err error) {
	token, err := this.EnsureAccessToken()
	if err != nil {
		return
	}
	result = new(UserSimpleListResult)
	url := CgiBinPrefix + "/user/simplelist?access_token=" + token.TokenStr + "&department_id=" + departmentId + "&fetch_child=" + fetchChild
	err = this.httpClient.GetObject(url, result)
	return
}

// ************************************* 获取部门成员详情 ****************************************
/**
https://work.weixin.qq.com/api/doc#90000/90135/90201
*/
type UserListResult struct {
	BaseResponse
	UserList []UserItem `json:"userlist"`
}

type UserItem struct {
	BaseUser
	Avatar string `json:"avatar"`
	Status string `json:"status"`
}

func (this *Api) UserList(departmentId string, fetchChild string) (result *UserListResult, err error) {
	token, err := this.EnsureAccessToken()
	if err != nil {
		return
	}
	result = new(UserListResult)
	url := CgiBinPrefix + "/user/list?access_token=" + token.TokenStr + "&department_id=" + departmentId + "&fetch_child=" + fetchChild
	err = this.httpClient.GetObject(url, result)
	return
}
