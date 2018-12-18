package GoEnterpriseWechatAPI

import "strconv"

// ************************************* 创建部门 ****************************************
/**
https://work.weixin.qq.com/api/doc#90000/90135/90205
*/
type Department struct {
	Name     string `json:"name"`
	ParentId int    `json:"parentid"`
	Order    int    `json:"order"`
	Id       int    `json:"id"`
}

type DepartmentCreateResult struct {
	BaseResponse
	Id int `json:"id"`
}

func (this *Api) DepartmentCreate(req *Department) (result *DepartmentCreateResult, err error) {
	token, err := this.EnsureAccessToken()
	if err != nil {
		return
	}
	result = new(DepartmentCreateResult)
	url := CgiBinPrefix + "/department/create?access_token=" + token.TokenStr
	err = this.httpClient.PostObjectGetObject(url, req, result)
	return
}

// ************************************* 更新部门 ****************************************
/**
https://work.weixin.qq.com/api/doc#90000/90135/90206
*/
func (this *Api) DepartmentUpdate(req *Department) (result *BaseResponse, err error) {
	token, err := this.EnsureAccessToken()
	if err != nil {
		return
	}
	result = new(BaseResponse)
	url := CgiBinPrefix + "/department/update?access_token=" + token.TokenStr
	err = this.httpClient.PostObjectGetObject(url, req, result)
	return
}

// ************************************* 删除部门 ****************************************
/**
https://work.weixin.qq.com/api/doc#90000/90135/90207
*/
func (this *Api) DepartmentDelete(id int) (result *BaseResponse, err error) {
	token, err := this.EnsureAccessToken()
	if err != nil {
		return
	}
	result = new(BaseResponse)
	url := CgiBinPrefix + "/department/delete?access_token=" + token.TokenStr + "&id=" + strconv.Itoa(id)
	err = this.httpClient.GetObject(url, result)
	return
}

// ************************************* 获取部门列表 ****************************************
/**
https://work.weixin.qq.com/api/doc#90000/90135/90208
*/
type DepartmentListResult struct {
	BaseResponse
	Department []Department `json:"department"`
}

func (this *Api) DepartmentList(id int) (result *DepartmentListResult, err error) {
	token, err := this.EnsureAccessToken()
	if err != nil {
		return
	}
	result = new(DepartmentListResult)
	url := CgiBinPrefix + "/department/list?access_token=" + token.TokenStr + "&id=" + strconv.Itoa(id)
	err = this.httpClient.GetObject(url, result)
	return
}
