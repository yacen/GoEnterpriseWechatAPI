package GoEnterpriseWechatAPI

import (
	"errors"
	"github.com/yacen/gotosee"
)

// Token保存策略，默认是存在成员变量中，
// 考虑分布式环境，可以存在数据库、redis等，这样才能在集群模式下使用。
type TokenStore interface {
	SaveToken(token *AccessToken) error
	GetToken() (token *AccessToken, err error)
}

type defaultTokenStore struct {
	token *AccessToken
}

func (this *defaultTokenStore) SaveToken(token *AccessToken) error {
	this.token = token
	return nil
}

func (this *defaultTokenStore) GetToken() (token *AccessToken, err error) {
	if this.token == nil {
		err = errors.New("TokenStr is not exists")
		return
	}
	return this.token, nil
}

type Api struct {
	corpId     string
	corpSecret string
	store      TokenStore
	httpClient *gotosee.GoToSee
}

func NewApi(corpId, corpSecret string, store ...TokenStore) (api *Api) {
	api = &Api{
		corpId:     corpId,
		corpSecret: corpSecret,
		httpClient: gotosee.NewGoToSee(),
	}
	if len(store) == 0 {
		api.store = &defaultTokenStore{}
	} else {
		api.store = store[0]
	}
	return
}

func (this *Api) saveToken(token *AccessToken) error {
	if this.store != nil {
	}
	return this.store.SaveToken(token)
}

func (this *Api) getToken() (token *AccessToken, err error) {
	return this.store.GetToken()
}
