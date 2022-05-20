package strutil

import (
	"regexp"
	"strings"
	"testing"

	"github.com/duke-git/lancet/v2/strutil"
	"github.com/smartystreets/goconvey/convey"
)

//
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

func TestLancetString(t *testing.T) {
	res := strutil.CamelCase("hello world")
	if res != "helloWorld" {
		t.Errorf("CamelCase error, expect helloWorld, get %s", res)
	}
}

// 驼峰
func CamelCase(s string) string {
	if len(s) == 0 {
		return ""
	}

	res := strings.Builder{}
	blankSpace := " "
	regex, _ := regexp.Compile("[-_&]+")
	ss := regex.ReplaceAllString(s, blankSpace)
	for i, v := range strings.Split(ss, blankSpace) {
		vv := []rune(v)
		if i == 0 {
			if vv[i] >= 65 && vv[i] <= 96 {
				vv[0] += 32
			}
			res.WriteString(string(vv))
		} else {
			res.WriteString(strutil.Capitalize(v))
		}
	}
	return res.String()
}

// 驼峰 benchmark
func BenchmarkCamelCase(b *testing.B) {
	ts := []string{"hello world", "foo_bar", "Foo-Bar", "Foo&bar", "foo bar"}
	n := len(ts)
	for i := 0; i < b.N; i++ {
		CamelCase(ts[i%n])
	}
}

// 柳叶刀 驼峰 benchmark
func BenchmarkLancetCamelCase(b *testing.B) {
	ts := []string{"hello world", "foo_bar", "Foo-Bar", "Foo&bar", "foo bar"}
	n := len(ts)
	for i := 0; i < b.N; i++ {
		strutil.CamelCase(ts[i%n])
	}
}
