package mockHelper

var (
	HtmlStr = `<html>
					<body>
						<p>Hello World!</p>
						<a href="https://deeptest.com">Deeptest WebSite</a>

						<form id="form1" name="f" action="/s" class="fm">
							<input id="kw" name="wd" class="small" value="" maxlength="255" autocomplete="off">
						</form>

						<h1 id="h1">
							This is a H1
						</h1>
						<ul>
							<li><a id="1" href="/">Home</a></li>
							<li><a id="2" href="/about">about</a></li>
							<li><a id="3" href="/account">login</a></li>
							<li></li>
						</ul>
						<p>
							Hello,This is an example for gxpath.
						</p>
						<footer>footer script</footer>

					</body>
				<html>`
)

func GetHtmlData() (result string) {
	result = HtmlStr
	return
}
