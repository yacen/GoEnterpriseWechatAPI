package GoEnterpriseWechatAPI

import (
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestUser(t *testing.T) {

	Convey("init api should success", t, func() {

		api := NewApi("ww7138020fc27495ae", "NEg9fKNZcV4NJhSmqfHAdf_PBOyDsNXwwFHNGMkVacM")
		So(api, ShouldNotBeNil)

		Convey("create user success", func() {
			userStr := `{"userid":"zhangsan","name":"张三","alias":"jackzhang","mobile":"15913215421","department":[1,2],"order":[10,40],"position":"产品经理","gender":"1","email":"zhangsan@gzdev.com","isleader":1,"enable":1,"avatar_mediaid":"2-G6nrLmr5EC3MNb_-zL1dDdzkd0p7cNliYu9V5w7o8K0","telephone":"020-123456","extattr":{"attrs":[{"name":"爱好","value":"旅游"},{"name":"卡号","value":"1234567234"}]},"to_invite":false,"external_position":"高级产品经理","external_profile":{"external_attr":[{"type":0,"name":"文本名称","text":{"value":"文本"}},{"type":1,"name":"网页名称","web":{"url":"http://www.test.com","title":"标题"}},{"type":2,"name":"测试app","miniprogram":{"appid":"wx8bd80126147df384","pagepath":"/index","title":"miniprogram"}}]}}`
			ur := new(UserCreateRequest)
			err := json.Unmarshal([]byte(userStr), ur)
			So(err, ShouldBeNil)

			result, err := api.UserCreate(ur)
			So(err, ShouldBeNil)
			So(result, ShouldNotBeNil)
		})

		Convey("get user success", func() {
			result, err := api.UserGet("zhangsan")
			So(err, ShouldBeNil)
			So(result, ShouldNotBeNil)
		})

		Convey("update user success", func() {
			userStr := `{"userid":"zhangsan","name":"李四","department":[1],"order":[10],"position":"后台工程师","mobile":"15913215421","gender":"1","email":"zhangsan@gzdev.com","isleader":0,"enable":1,"avatar_mediaid":"2-G6nrLmr5EC3MNb_-zL1dDdzkd0p7cNliYu9V5w7o8K0","telephone":"020-123456","alias":"jackzhang","extattr":{"attrs":[{"name":"爱好","value":"旅游"},{"name":"卡号","value":"1234567234"}]},"external_position":"工程师","external_profile":{"external_attr":[{"type":0,"name":"文本名称","text":{"value":"文本"}},{"type":1,"name":"网页名称","web":{"url":"http://www.test.com","title":"标题"}},{"type":2,"name":"测试app","miniprogram":{"appid":"wx8bd80126147df384","pagepath":"/index","title":"my miniprogram"}}]}}`
			ur := new(UserUpdateRequest)
			err := json.Unmarshal([]byte(userStr), ur)
			So(err, ShouldBeNil)

			result, err := api.UserUpdate(ur)
			So(err, ShouldBeNil)
			So(result, ShouldNotBeNil)
		})

		Convey("delete user success", func() {
			result, err := api.UserDelete("zhangsan")
			So(err, ShouldBeNil)
			So(result, ShouldNotBeNil)
		})

		Convey("batch delete user success", func() {
			result, err := api.UserBatchDelete([]string{"zhangsan", "lisi"})
			So(err, ShouldBeNil)
			So(result, ShouldNotBeNil)
		})

		Convey("get department user simple list success", func() {
			result, err := api.UserSimpleList("1", "1")
			So(err, ShouldBeNil)
			So(result, ShouldNotBeNil)
		})

		Convey("get department user detail list success", func() {
			result, err := api.UserList("1", "1")
			So(err, ShouldBeNil)
			So(result, ShouldNotBeNil)
		})

		Convey("userid convert to openid success", func() {
			result, err := api.UserConvertToOpenId("zhangsan")
			So(err, ShouldBeNil)
			So(result, ShouldNotBeNil)
		})

		Convey("openid convert to userid success", func() {
			result, err := api.UserConvertToUserId("zhangsan1")
			So(err, ShouldBeNil)
			So(result, ShouldNotBeNil)
		})

		Convey("authsucc ", func() {
			result, err := api.UserAuthSucc("zhangsan1")
			So(err, ShouldBeNil)
			So(result, ShouldNotBeNil)
		})

		Convey("batch invite users success", func() {
			result, err := api.UserInvite([]string{"1"}, []string{"1"}, []int{1})
			So(err, ShouldBeNil)
			So(result, ShouldNotBeNil)
		})
	})

}
