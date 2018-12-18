package GoEnterpriseWechatAPI

import "strconv"

type Tag struct {
	TagName string `json:"tagname"`
	TagId   int    `json:"tagid"`
}

// ************************************* 创建标签 ****************************************
/**
https://work.weixin.qq.com/api/doc#90000/90135/90210
*/
type TagCreateResult struct {
	BaseResponse
	TagId int `json:"tagid"`
}

func (this *Api) TagCreate(tagName string, tagId int) (result *TagCreateResult, err error) {
	token, err := this.EnsureAccessToken()
	if err != nil {
		return
	}
	req := &Tag{TagName: tagName, TagId: tagId}
	result = new(TagCreateResult)
	url := CgiBinPrefix + "/tag/create?access_token=" + token.TokenStr
	err = this.httpClient.PostObjectGetObject(url, req, result)
	return
}

// ************************************* 更新标签名字 ****************************************
/**
https://work.weixin.qq.com/api/doc#90000/90135/90211
*/
func (this *Api) TagUpdate(tagName string, tagId int) (result *BaseResponse, err error) {
	token, err := this.EnsureAccessToken()
	if err != nil {
		return
	}
	req := &Tag{TagName: tagName, TagId: tagId}
	result = new(BaseResponse)
	url := CgiBinPrefix + "/tag/update?access_token=" + token.TokenStr
	err = this.httpClient.PostObjectGetObject(url, req, result)
	return
}

// ************************************* 删除标签 ****************************************
/**
https://work.weixin.qq.com/api/doc#90000/90135/90212
*/
func (this *Api) TagDelete(tagId int) (result *BaseResponse, err error) {
	token, err := this.EnsureAccessToken()
	if err != nil {
		return
	}
	result = new(BaseResponse)
	url := CgiBinPrefix + "/tag/delete?access_token=" + token.TokenStr + "&tagid=" + strconv.Itoa(tagId)
	err = this.httpClient.GetObject(url, result)
	return
}

// ************************************* 获取标签成员 ****************************************
/**
https://work.weixin.qq.com/api/doc#90000/90135/90213
*/

type TagGet struct {
	BaseResponse
	TagName  string    `json:"tagname"`
	UserList []TagUser `json:"userlist"`
}

type TagUser struct {
	UserId string `json:"userid"`
	Name   string `json:"name"`
}

func (this *Api) TagGet(tagId int) (result *TagGet, err error) {
	token, err := this.EnsureAccessToken()
	if err != nil {
		return
	}
	result = new(TagGet)
	url := CgiBinPrefix + "/tag/get?access_token=" + token.TokenStr + "&tagid=" + strconv.Itoa(tagId)
	err = this.httpClient.GetObject(url, result)
	return
}

// ************************************* 增加标签成员 ****************************************
/**
https://work.weixin.qq.com/api/doc#90000/90135/90214
*/

type TagAddDelTagUserRequest struct {
	TagId     int      `json:"tagid"`
	UserList  []string `json:"userlist"`
	PartyList []int    `json:"partylist"`
}

type TagAddDelTagUserResult struct {
	BaseResponse
	InvalidList  string `json:"invalidlist"`
	InvalidParty []int  `json:"invalidparty"`
}

func (this *Api) TagAddTagUser(req *TagAddDelTagUserRequest) (result *TagAddDelTagUserResult, err error) {
	token, err := this.EnsureAccessToken()
	if err != nil {
		return
	}
	result = new(TagAddDelTagUserResult)
	url := CgiBinPrefix + "/tag/addtagusers?access_token=" + token.TokenStr
	err = this.httpClient.PostObjectGetObject(url, req, result)
	return
}

// ************************************* 删除标签成员 ****************************************
/**
https://work.weixin.qq.com/api/doc#90000/90135/90215
*/
func (this *Api) TagDelTagUser(req *TagAddDelTagUserRequest) (result *TagAddDelTagUserResult, err error) {
	token, err := this.EnsureAccessToken()
	if err != nil {
		return
	}
	result = new(TagAddDelTagUserResult)
	url := CgiBinPrefix + "/tag/deltagusers?access_token=" + token.TokenStr
	err = this.httpClient.PostObjectGetObject(url, req, result)
	return
}

// ************************************* 获取标签列表 ****************************************
/**
https://work.weixin.qq.com/api/doc#90000/90135/90216
*/
type TagListResult struct {
	BaseResponse
	TagList []Tag `json:"taglist"`
}

func (this *Api) TagList() (result *TagListResult, err error) {
	token, err := this.EnsureAccessToken()
	if err != nil {
		return
	}
	result = new(TagListResult)
	url := CgiBinPrefix + "/tag/list?access_token=" + token.TokenStr
	err = this.httpClient.GetObject(url, result)
	return
}
