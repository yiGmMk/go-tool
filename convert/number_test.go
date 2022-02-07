package convert

import (
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/smartystreets/goconvey/convey"
)

func TestConverter(t *testing.T) {
	type data struct {
		in  int64
		out string
	}
	ts := []data{
		{
			in:  1,
			out: "壹",
		},
		{
			in:  123,
			out: "壹佰贰拾叁",
		},
		{
			in:  9999,
			out: "玖仟玖佰玖拾玖",
		},
	}
	convey.Convey("chinese", t, func() {
		for _, v := range ts {
			out := toChinese(v.in, true)
			convey.So(v.out, convey.ShouldEqual, out)
		}
	})
}

func TestConverterCur(t *testing.T) {
	big, _ := decimal.NewFromString("9999999999999999.99")
	fmt.Println(decimal.NewFromFloat(345_9999_7890_1234.56).String(), decimal.NewFromFloat(345_9999_7890_1234.56).Mul(decimal.NewFromFloat(10000)).String(), big.String())

	type data struct {
		in  decimal.Decimal
		out string
	}
	v1, _ := decimal.NewFromString("2345999999991232.23")
	ts := []data{
		{
			in:  decimal.NewFromFloat(1_0000_0012.12),
			out: "壹亿零壹拾贰元壹角贰分",
		},
		{
			in:  decimal.NewFromFloat(0.123),
			out: "零元壹角贰分",
		},
		{
			in:  decimal.NewFromFloat(1_0012.123),
			out: "壹万零壹拾贰元壹角贰分",
		},
		{
			in:  v1,
			out: "贰仟叁佰肆拾伍万玖仟玖佰玖拾玖亿玖仟玖佰玖拾玖万壹仟贰佰叁拾贰元贰角叁分",
		},
		{
			in:  decimal.NewFromInt(0),
			out: "零元整",
		},
		{
			in:  decimal.NewFromFloat(1.12),
			out: "壹元壹角贰分",
		},
		{
			in:  decimal.NewFromFloat(1.1),
			out: "壹元壹角整",
		},
		{
			in:  decimal.NewFromFloat(123.23),
			out: "壹佰贰拾叁元贰角叁分",
		},
	}
	convey.Convey("chinese", t, func() {
		for _, v := range ts {
			out, err := Number2ChineseYUAN(v.in, true)
			fmt.Println("in:", v.in, "out:", out)
			convey.So(err, convey.ShouldBeNil)
			convey.So(v.out, convey.ShouldEqual, out)
		}
	})
}

func BenchmarkConvert(b *testing.B) {
	type data struct {
		in  decimal.Decimal
		out string
	}
	ts := []data{
		{
			in:  decimal.NewFromInt(0),
			out: "零元整",
		},
		{
			in:  decimal.NewFromFloat(1.12),
			out: "壹元壹角贰分",
		},
		{
			in:  decimal.NewFromFloat(1_0000_0012.12),
			out: "壹亿零壹拾贰元壹角贰分",
		},
		{
			in:  decimal.NewFromFloat(1.1),
			out: "壹元壹角整",
		},
		{
			in:  decimal.NewFromFloat(123.23),
			out: "壹佰贰拾叁元贰角叁分",
		},
		{
			in:  decimal.NewFromFloat(9999_9999_1232.23),
			out: "玖仟玖佰玖拾玖亿玖仟玖佰玖拾玖万壹仟贰佰叁拾贰元贰角叁分",
		},
	}
	for i := 0; i < b.N; i++ {
		for _, test := range ts {
			Number2ChineseYUAN(test.in, true)
		}
	}
}
