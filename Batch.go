package GoEnterpriseWechatAPI

/**
异步批量接口
https://work.weixin.qq.com/api/doc#90000/90135/90978
*/

// ************************************* 增量更新成员 ****************************************
/**
https://work.weixin.qq.com/api/doc#90000/90135/90980
*/
type BatchSyncReplaceUserRequest struct {
	MediaId  string        `json:"media_id"`
	ToInvite bool          `json:"to_invite"`
	Callback BatchCallback `json:"callback"`
}
type BatchCallback struct {
	Url            string `json:"url"`
	Token          string `json:"token"`
	EncodingAeskey string `json:"encodingaeskey"`
}
type BatchSyncReplaceUserPartyResult struct {
	BaseResponse
	JobId string `json:"jobid"`
}

func (this *Api) BatchSyncUser(req *BatchSyncReplaceUserRequest) (result *BatchSyncReplaceUserPartyResult, err error) {
	token, err := this.EnsureAccessToken()
	if err != nil {
		return
	}
	result = new(BatchSyncReplaceUserPartyResult)
	url := CgiBinPrefix + "/batch/syncuser?access_token=" + token.TokenStr
	err = this.httpClient.PostObjectGetObject(url, req, result)
	return
}

// ************************************* 全量覆盖成员 ****************************************
/**
https://work.weixin.qq.com/api/doc#90000/90135/90980
*/
func (this *Api) BatchReplaceUser(req *BatchSyncReplaceUserRequest) (result *BatchSyncReplaceUserPartyResult, err error) {
	token, err := this.EnsureAccessToken()
	if err != nil {
		return
	}
	result = new(BatchSyncReplaceUserPartyResult)
	url := CgiBinPrefix + "/batch/replaceuser?access_token=" + token.TokenStr
	err = this.httpClient.PostObjectGetObject(url, req, result)
	return
}

// ************************************* 全量覆盖部门 ****************************************
/**
https://work.weixin.qq.com/api/doc#90000/90135/90982
*/
type BatchReplacePartyRequest struct {
	MediaId  string        `json:"media_id"`
	Callback BatchCallback `json:"callback"`
}

func (this *Api) BatchReplaceParty(req *BatchReplacePartyRequest) (result *BatchSyncReplaceUserPartyResult, err error) {
	token, err := this.EnsureAccessToken()
	if err != nil {
		return
	}
	result = new(BatchSyncReplaceUserPartyResult)
	url := CgiBinPrefix + "/batch/replaceparty?access_token=" + token.TokenStr
	err = this.httpClient.PostObjectGetObject(url, req, result)
	return
}

// ************************************* 获取异步任务结果 ****************************************
/**
https://work.weixin.qq.com/api/doc#90000/90135/90983
*/
type BatchGetResultResult struct {
	BaseResponse
	Status     int            `json:"status"`
	Type       string         `json:"type"`
	Total      int            `json:"total"`
	Percentage int            `json:"percentage"`
	Result     []ResultDetail `json:"result"`
}

type ResultDetail struct {
	UserId  string `json:"userid"`
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`

	Action  int `json:"action"`
	Partyid int `json:"partyid"`
}

func (this *Api) BatchGetResult(jobId string) (result *BatchGetResultResult, err error) {
	token, err := this.EnsureAccessToken()
	if err != nil {
		return
	}
	result = new(BatchGetResultResult)
	url := CgiBinPrefix + "/batch/getresult?access_token=" + token.TokenStr + "&jobid=" + jobId
	err = this.httpClient.GetObject(url, result)
	return
}
