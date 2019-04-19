package lib

// import (
// 	"fmt"
// 	"github.com/PuerkitoBio/goquery"
// 	"log"
// 	"net/http"
// 	"strings"
// )

// // ExampleScrape ...
// func ExampleScrape() {
// 	// Request the HTML page.
// 	res, err := http.Get("http://metalsucks.net")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer res.Body.Close()
// 	if res.StatusCode != 200 {
// 		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
// 	}

// 	strHTML := `
// <!----- Conversion time: 0.885 seconds.

// Using this HTML file:

// 1. Cut and paste this output into your source file.
// 2. See the notes and action items below regarding this conversion run.
// 3. Check the rendered output (headings, lists, code blocks, tables) for proper
//    formatting and use a linkchecker before you publish this page.

// Conversion notes:

// * Docs to Markdown version 1.0β14
// * Wed Feb 06 2019 00:17:35 GMT-0800 (PST)
// * Source doc: https://docs.google.com/a/michael-inc.com/open?id=1Ma_5cZy7D9LXdNu2T6-Dxe2PZh3JLgdD8Vk7jOiv630
// * This is a partial selection. Check to make sure intra-doc links work.
// ----->

// <h1>タイトル</h1>

// <p>サブタイトル
// </p>
// <h1>見出し1</h1>

// <h2>見出し2</h2>

// <h3>見出し3</h3>

// <p>
// <strong>これは太字です</strong>
// </p>
// <p>
// <em>これはイタリックです</em>
// </p>
// <p>
// <span style="text-decoration:underline;">これはアンダーバーです</span>
// </p>
// <p>
// これは赤字です
// </p>
// <p>
// これは背景黄色です
// </p>
// <p>
// <a href="https://godoc.org/github.com/PuerkitoBio/goquery#Selection.Children">これはりんくです。</a>
// </p>
// <p>
// <a href="https://godoc.org/github.com/PuerkitoBio/goquery#Selection.Children">https://godoc.org/github.com/PuerkitoBio/goquery#Selection.Children</a>
// </p>
// <p>
// これはプレーンテキストです。
// </p>

// <table>
//   <tr>
//    <td>ヘッダー1
//    </td>
//    <td>ヘッダー2
//    </td>
//    <td>ヘッダー3
//    </td>
//   </tr>
//   <tr>
//    <td>1,1
//    </td>
//    <td>1,2
//    </td>
//    <td>1,3
//    </td>
//   </tr>
//   <tr>
//    <td>2,1
//    </td>
//    <td>2,2
//    </td>
//    <td>2,3
//    </td>
//   </tr>
// </table>

// <ul>

// <li>箇条書きです
// <li>箇条書きです
// <li>箇条書きです<ol>

// <li>箇条書き1です
// <li>箇条書き2です
// <li>箇条書き3です
// `

// 	// Load the HTML document
// 	// doc, err := goquery.NewDocumentFromReader(res.Body)
// 	stringReader := strings.NewReader(strHTML)
// 	doc, err := goquery.NewDocumentFromReader(stringReader)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Find the review items
// 	// doc.Find(".sidebar-reviews article .content-block").Each(func(i int, s *goquery.Selection) {
// 	// 	// For each item found, get the band and title
// 	// 	band := s.Find("a").Text()
// 	// 	title := s.Find("i").Text()
// 	// 	fmt.Printf("Review %d: %s - %s\n", i, band, title)
// 	// })
// 	// html, _ := doc.Html()
// 	// doc2 := goquery.NewDocumentFromNode(doc.Nodes[0])
// 	// doc2.Find("body").Children().Each(func(index int, selector *goquery.Selection) {
// 	// 	ret, _ := selector.Html()
// 	// 	// ret := selector.Text()
// 	// 	fmt.Printf("%d: %s\n", index, ret)
// 	// })

// 	sel := doc.Find("body").Children()
// 	for i := range sel.Nodes {
// 		s := sel.Eq(i)
// 		fmt.Println(i)
// 		// fmt.Println(s)
// 		// html, _ := s.Html()
// 		// fmt.Println(html)
// 		// fmt.Println(s.Text())
// 		// fmt.Println(s.Parent())
// 		// fmt.Println(s.Parent().Text())

// 		ret, _ := goquery.OuterHtml(s)
// 		fmt.Println(ret)
// 	}
// 	// fmt.Println(html2)
// 	// doc.Selection.Children().Each(func(i int, s *goquery.Selection) {
// 	// 	// For each item found, get the band and title
// 	// 	content, _ := s.Html()
// 	// 	fmt.Printf("%d: %s", i, content)
// 	// })
// }

// // PointerExtract ...
// func PointerExtract(data interface{}) {
// 	valType := fmt.Sprintf("%T", data)

// 	fmt.Println(valType)
// 	fmt.Println(*data.(*goquery.Document))
// }
