package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var RouteMap = `
<html>
	<head>
		<title>Tricking API Routes</title>
	</head>
	<body>
		<h1>Welcome to api.trickingapi.dev</h1>
		<p>These are the routes available in the Tricking API</p>
		<ul>
			<li><a href="/tricks">/tricks</a></li>
			<li><a href="/tricks/ids">/tricks/:id</a></li>
			<li><a href="/tricks/names">/tricks/names</a></li>
			<li><a href="/tricks/:name">/tricks/:name</a></li>
			<li><a href="/categories">/categories</a></li>
			<li><a href="/categories/tricks">/categories/tricks</a></li>
			<li><a href="/categories/:name">/categories/:name</a></li>
			<li><a href="/transitions">/transitions</a></li>
			<li><a href="/transitions/ids">/transitions/ids</a></li>
			<li><a href="/transitions/:id">/transitions/:id</a></li>
			<li><a href="/landingstances">/landingstances</a></li>
			<li><a href="/landingstances/ids">/landingstances/ids</a></li>
			<li><a href="/landingstances/:id">/landingstances/:id</a></li>
		</ul>
	</body>
</html>
`

func GetRootHandler() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(RouteMap))
	}

	return gin.HandlerFunc(fn)
}
