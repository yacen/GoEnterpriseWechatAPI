package GoEnterpriseWechatAPI

/**
网页授权登录
https://work.weixin.qq.com/api/doc#90000/90135/91020
*/

/**
扫玛授权登录
https://work.weixin.qq.com/api/doc#90000/90135/90988
*/

/*
获取访问用户身份
https://work.weixin.qq.com/api/doc#90000/90135/91023
*/
type UserGetUserInfoResult struct {
	BaseUser
	UserId   string `json:"UserId"`
	DeviceId string `json:"DeviceId"`

	OpenId string `json:"OpenId"`
}

func (this *Api) UserGetUserInfo(code string) (result *UserGetUserInfoResult, err error) {
	token, err := this.EnsureAccessToken()
	if err != nil {
		return
	}
	result = new(UserGetUserInfoResult)
	url := CgiBinPrefix + "/user/create?access_token=" + token.TokenStr + "&code=" + code
	err = this.httpClient.GetObject(url, result)
	return
}
