# Enterprise Wechat API for Golang

企业微信API

## 功能列表
* 成员管理

## Installation
````
go get github.com/yacen/GoEnterpriseWechatAPI.git
````

## Usage
````
api := NewApi("corpid", "secret")
result, err := api.UserGet("zhangsan")
````

## cluster
````
type RedisStore struct{}

func (this *RedisStore) SaveToken(token *AccessToken) (err error) {
	r := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	defer r.Close()
	data, err := json.Marshal(token)
	if err != nil {
		return
	}
	return r.Set("enterprise:accessToken", string(data), time.Second*time.Duration(token.ExpiresIn)).Err()
}

func (this *RedisStore) GetToken() (token *AccessToken, err error) {
	r := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	defer r.Close()
	data, err := r.Get("enterprise:accessToken").Bytes()
	if err != nil {
		return
	}
	err = json.Unmarshal(data, token)
	return
}


api := NewApi(
    "wxdd725338566d6ffe",
    "vQT_03RsfVA3uE6J5dofR7hJeOdiXUvccqV8mDgLdLI",
    &RedisStore{},
)
result, err := api.UserGet("zhangsan")

````
## 捐赠
如果觉得此库对你有帮助，减轻你的工作量，欢迎请作者喝杯奶茶

![qrcode](./qrcode.png)
