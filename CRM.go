package GoEnterpriseWechatAPI

/**
外部联系人管理
https://work.weixin.qq.com/api/doc#90000/90135/90221
*/

// ************************************* 离职成员的外部联系人再分配 ****************************************
/**
https://work.weixin.qq.com/api/doc#90000/90135/90222
*/

type CRMTransferExternalContact struct {
	ExternalUserId string `json:"external_userid"`
	HandoverUserId string `json:"handover_userid"`
	TakeoverUserId string `json:"takeover_userid"`
}

func (this *Api) CRMTransferExternalContact(req *CRMTransferExternalContact) (result *BaseResponse, err error) {
	token, err := this.EnsureAccessToken()
	if err != nil {
		return
	}
	result = new(BaseResponse)
	url := CgiBinPrefix + "/crm/transfer_external_contact?access_token=" + token.TokenStr
	err = this.httpClient.PostObjectGetObject(url, req, result)
	return
}

// ************************************* 获取外部联系人详情 ****************************************
/**
https://work.weixin.qq.com/api/doc#90000/90135/90224
*/
type CRMGetExternalContactResult struct {
	BaseResponse
	ExternalUserId  string          `json:"external_userid"`
	Name            string          `json:"name"`
	Position        string          `json:"position"`
	Avatar          string          `json:"avatar"`
	CorpName        string          `json:"corp_name"`
	CorpFullName    string          `json:"corp_full_name"`
	Type            int             `json:"type"`
	Gender          int             `json:"gender"`
	UnionId         string          `json:"unionid"`
	ExternalProfile ExternalProfile `json:"external_profile"`
	FollowUser      []FollowUser    `json:"follow_user"`
}

type FollowUser struct {
	UserId      string `json:"userid"`
	Remark      string `json:"remark"`
	Description string `json:"description"`
	CreateTime  int    `json:"createtime"`
}

func (this *Api) CRMGetExternalContact(externalUserId string) (result *CRMGetExternalContactResult, err error) {
	token, err := this.EnsureAccessToken()
	if err != nil {
		return
	}
	result = new(CRMGetExternalContactResult)
	url := CgiBinPrefix + "/crm/get_external_contact?access_token=" + token.TokenStr + "&external_userid=" + externalUserId
	err = this.httpClient.GetObject(url, result)
	return
}
