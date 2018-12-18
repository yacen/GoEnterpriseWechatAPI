package GoEnterpriseWechatAPI

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDepartment(t *testing.T) {

	Convey("init api should success", t, func() {

		api := NewApi("ww7138020fc27495ae", "NEg9fKNZcV4NJhSmqfHAdf_PBOyDsNXwwFHNGMkVacM")
		So(api, ShouldNotBeNil)

		Convey("create department success", func() {
			result, err := api.DepartmentCreate(&Department{Name: "广州研发中心", ParentId: 1, Order: 1, Id: 1})
			So(err, ShouldBeNil)
			So(result, ShouldNotBeNil)
		})

		Convey("update department success", func() {
			result, err := api.DepartmentUpdate(&Department{Name: "上海研发中心", ParentId: 1, Order: 1, Id: 1})
			So(err, ShouldBeNil)
			So(result, ShouldNotBeNil)
		})

		Convey("delete department success", func() {
			result, err := api.DepartmentDelete(1)
			So(err, ShouldBeNil)
			So(result, ShouldNotBeNil)
		})

		Convey("get department list success", func() {
			result, err := api.DepartmentList(1)
			So(err, ShouldBeNil)
			So(result, ShouldNotBeNil)
		})
	})

}
