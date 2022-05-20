package strutil

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestRemoveExceptDigit(t *testing.T) {
	ts := []struct {
		in  string
		out string
	}{
		{
			in:  " 123",
			out: "123",
		},
		{
			in:  "123456789",
			out: "123456789",
		},
		{
			in:  "‭15583887773‬",
			out: "15583887773",
		},
	}
	convey.Convey("test", t, func() {
		for _, tt := range ts {
			ret := RemoveExceptDigit(tt.in)
			convey.So(ret, convey.ShouldEqual, tt.out)
		}
	})
}

func BenchmarkRemoveExceptDigit(b *testing.B) {
	ts := []struct {
		in  string
		out string
	}{
		{
			in:  " 123",
			out: "123",
		},
		{
			in:  "123456789",
			out: "123456789",
		},
		{
			in:  "‭15583887773‬",
			out: "15583887773",
		},
	}
	for i := 0; i < b.N; i++ {
		RemoveExceptDigit(ts[i%3].in)
	}
}
