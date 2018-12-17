package GoEnterpriseWechatAPI

/**
access token 管理
*/
import (
	"time"
)

type AccessToken struct {
	BaseResponse
	TokenStr    string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	ExpiresTime time.Time
}

func (this *AccessToken) SetExpiresTime() {
	// 过期时间，因网络延迟等，将实际过期时间提前10秒，以防止临界点
	this.ExpiresTime = time.Now().Add(time.Duration(this.ExpiresIn-10) * time.Second)
}

func (this *AccessToken) IsValid() bool {
	if this.ErrCode != 0 {
		return false
	}
	if this.TokenStr == "" {
		return false
	}
	if time.Now().After(this.ExpiresTime) {
		return false
	}
	return true
}

/*
https://work.weixin.qq.com/api/doc#90000/90135/91039
*/
func (this *Api) GetAccessToken() (token *AccessToken, err error) {
	url := CgiBinPrefix + "/gettoken?corpid=" + this.corpId + "&corpsecret=" + this.corpSecret
	token = new(AccessToken)
	err = this.httpClient.GetObject(url, token)
	if err != nil {
		return
	}
	token.SetExpiresTime()
	err = this.store.SaveToken(token)
	return
}

// 每次调用微信接口都会调用此方法
func (this *Api) EnsureAccessToken() (token *AccessToken, err error) {
	token, err = this.store.GetToken()
	if err != nil || token == nil || !token.IsValid() {
		return this.GetAccessToken()
	}
	return
}
