package mockHelper

var (
	HtmlStr = `<html>
					<body>
						<p>Hello World!</p>
						<a href="https://deeptest.com">Deeptest WebSite</a>
					</body>
				<html>`
)

func GetHtmlData() (result string) {
	result = HtmlStr
	return
}
