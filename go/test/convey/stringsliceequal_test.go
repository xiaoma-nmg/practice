package convey

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStringSliceEqual1(t *testing.T) {
	Convey("TestStringSliceEqual should return true when a!=nil && b!=nil", t, func() {
		a := []string{"hello", "goConvey"}
		b := []string{"hello", "goConvey"}
		So(StringSliceEqual(a, b), ShouldBeTrue)
	})

	Convey("TestStringSliceEqual should return true when a==nil && b==nil", t, func() {
		So(StringSliceEqual(nil, nil), ShouldBeTrue)
	})

	Convey("TestStingSliceEqual should return false when a==nil && b!=nil", t, func() {
		a := []string(nil)
		b := []string{}
		So(StringSliceEqual(a, b), ShouldBeFalse)
	})

	Convey("TestStringSliceEqual should return false when a!=nil && b!=nil", t, func() {
		a := []string{"hello", "test"}
		b := []string{"hello", "goConvey"}
		So(StringSliceEqual(a, b), ShouldBeFalse)
	})
}

func TestStringSliceEqual2(t *testing.T) {
	Convey("TestStringSliceEqual", t, func() {
		SkipConvey("should return true when a!=nil && b!=nil", func() {
			a := []string{"hello", "goConvey"}
			b := []string{"hello", "goConvey"}
			So(StringSliceEqual(a, b), ShouldBeTrue)
		})

		Convey("should return true when a==nil && b==nil", func() {
			SkipSo(StringSliceEqual(nil, nil), ShouldBeTrue)
		})

		Convey("should return false when a==nil && b!=nil", func() {
			a := []string(nil)
			b := []string{}
			So(StringSliceEqual(a, b), ShouldBeFalse)
		})

		Convey("should return false when a!=nil && b!=nil", func() {
			a := []string{"hello", "test"}
			b := []string{"hello", "goConvey"}
			So(StringSliceEqual(a, b), ShouldBeFalse)
		})
	})
}
