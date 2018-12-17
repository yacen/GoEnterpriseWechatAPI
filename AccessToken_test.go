package GoEnterpriseWechatAPI

import (
	"encoding/json"
	"github.com/go-redis/redis"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

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

func TestAccessToken(t *testing.T) {

	tokenStr := "cOph-q9E7HDP2IDJRoVRKWv5tqnaWFWpWB0HiueUk8UKlytawFsItS6B5yj7koiGtIDw_IrAenO60KAhjefO9UgNZwHtWtbdJ_fawjtPvkXxGBoq3fOxXrAnxw2epe78mNsnfgL5vi0pj3LF8hcSOwEWBngtY29aezrW2e7-76RtmicmY81bWSBGx3ox1fd-_s0UKR9fF8tr7GJJZjILEQ"

	Convey("token is valid", t, func() {
		token := AccessToken{
			TokenStr:  tokenStr,
			ExpiresIn: 7200}
		token.SetExpiresTime()
		So(token.IsValid(), ShouldBeTrue)
	})
	Convey("token is not valid", t, func() {
		token := AccessToken{
			ExpiresIn: 7200}
		token.SetExpiresTime()
		So(token.IsValid(), ShouldBeFalse)
	})
	Convey("token is not valid", t, func() {
		token := AccessToken{
			TokenStr:  tokenStr,
			ExpiresIn: -1}
		token.SetExpiresTime()
		So(token.IsValid(), ShouldBeFalse)
	})
}

func TestGetAccessToken(t *testing.T) {

	Convey("get access token should fail", t, func() {
		api := NewApi("wx6207bq381a8d25410", "94w878c76bda145802980f0b34b5e4751")
		token, err := api.GetAccessToken()
		So(err, ShouldBeNil)
		So(token, ShouldNotBeNil)
		So(token.ErrCode, ShouldEqual, 40013)
		So(token.IsValid(), ShouldBeFalse)

	})

	Convey("get access token should success", t, func() {
		api := NewApi("wxdd725338566d6ffe", "vQT_03RsfVA3uE6J5dofR7hJeOdiXUvccqV8mDgLdLI")
		token, err := api.GetAccessToken()
		So(err, ShouldBeNil)
		So(token, ShouldNotBeNil)
		So(token.ErrCode, ShouldEqual, 0)
		So(token.ErrMsg, ShouldEqual, "ok")
		So(token.TokenStr, ShouldNotBeBlank)
		So(token.ExpiresIn, ShouldEqual, 7200)
		So(token.IsValid(), ShouldBeTrue)
	})

	Convey("save token to redis", t, func() {
		api := NewApi(
			"wxdd725338566d6ffe",
			"vQT_03RsfVA3uE6J5dofR7hJeOdiXUvccqV8mDgLdLI",
			&RedisStore{},
		)
		token, err := api.GetAccessToken()
		So(err, ShouldBeNil)
		So(token, ShouldNotBeNil)
		So(token.ErrCode, ShouldEqual, 0)
		So(token.ErrMsg, ShouldEqual, "ok")
		So(token.TokenStr, ShouldNotBeBlank)
		So(token.ExpiresIn, ShouldEqual, 7200)
		So(token.IsValid(), ShouldBeTrue)

	})
}

func TestEnsureAccessToken(t *testing.T) {

	Convey("ensure access token fail", t, func() {
		api := NewApi("wx6207bq381a8d25410", "94w878c76bda145802980f0b34b5e4751")
		token, err := api.EnsureAccessToken()
		So(err, ShouldBeNil)
		So(token, ShouldNotBeNil)
		So(token.ErrCode, ShouldEqual, 40013)
		So(token.IsValid(), ShouldBeFalse)
	})

	Convey("ensure access token  success", t, func() {
		api := NewApi("ww7138020fc27495ae", "NEg9fKNZcV4NJhSmqfHAdf_PBOyDsNXwwFHNGMkVacM")
		token, err := api.GetAccessToken()
		So(err, ShouldBeNil)
		So(token, ShouldNotBeNil)
		So(token.ErrCode, ShouldEqual, 0)
		So(token.ErrMsg, ShouldEqual, "ok")
		So(token.TokenStr, ShouldNotBeBlank)
		So(token.ExpiresIn, ShouldEqual, 7200)
		So(token.IsValid(), ShouldBeTrue)
	})

	Convey("test token expire success", t, func() {
		api := NewApi("ww7138020fc27495ae", "NEg9fKNZcV4NJhSmqfHAdf_PBOyDsNXwwFHNGMkVacM")
		token, err := api.GetAccessToken()
		So(err, ShouldBeNil)
		So(token, ShouldNotBeNil)
		So(token.ErrCode, ShouldEqual, 0)
		So(token.ErrMsg, ShouldEqual, "ok")
		So(token.TokenStr, ShouldNotBeBlank)
		So(token.ExpiresIn, ShouldEqual, 7200)
		So(token.IsValid(), ShouldBeTrue)

		token.ExpiresIn = -1
		token.SetExpiresTime()

		token2, err := api.EnsureAccessToken()
		So(err, ShouldBeNil)
		So(token2, ShouldNotBeNil)
		So(token2.ErrCode, ShouldEqual, 0)
		So(token2.ErrMsg, ShouldEqual, "ok")
		So(token2.TokenStr, ShouldNotBeBlank)
		So(token2.ExpiresIn, ShouldEqual, 7200)
		So(token2.IsValid(), ShouldBeTrue)

		So(token, ShouldNotEqual, token2)

	})
}
