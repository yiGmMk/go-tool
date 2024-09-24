package tests

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

// goquery  https://learnku.com/articles/48022
func TestGoQuery(t *testing.T) {
	t.Run("find", func(t *testing.T) {
		html := `<html>
		<body>
			<h1 id="title">春晓</h1>
			<p class="content1">
			春眠不觉晓，
			处处闻啼鸟。
			夜来风雨声，
			花落知多少。
			</p>
		</body>
		</html>
		`
		dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
		if err != nil {
			log.Fatalln(err)
		}

		dom.Find("p").Each(func(i int, selection *goquery.Selection) {
			fmt.Println(selection.Text())
		})
	})

	t.Run("baidu", func(t *testing.T) {
		// 百度热搜
		res, err := http.Get("https://top.baidu.com/board?platform=pc&sa=pcindex_entry")
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		}

		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(doc.Html())
		fmt.Println()
		fmt.Println("-----------------------------")
		doc.Find(".list_1EDla .item-wrap_2oCLZ").Each(func(i int, s *goquery.Selection) {
			content := s.Find(".name_2Px2N .c-single-text-ellipsis").First().Text()
			fmt.Printf("%d: %s\n", i, content)
		})
		// 0:   以人口高质量发展支撑东北全面振兴
		// 1:   罗永浩：“真还传”共还了8.24亿
		// 2:   黄子韬求婚徐艺洋
		// 3:   各地采取多种措施保民生保安全
		// 4:   13岁男孩血尿医生取出20多颗磁力珠
		// 5:   男子仅退款被拒后起诉网店被驳回
		// 6:   10年来结婚登记数腰斩
		// 7:   陈梦回应佩戴定制首饰参赛
		// 8:   水果太甜是打了甜蜜素？谣言
		// 9:   男生开学打开行李箱天塌了：粉色被褥
		// 10:   女子被榴莲咬脚杀伤力满级
	})
}
