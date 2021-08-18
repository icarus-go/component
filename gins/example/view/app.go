package view

import "pmo-test4.yz-intelligence.com/kit/component/gins"

func init() {
	gins.AddTemplate("app", app)
}

const app = `<!DOCTYPE html>
<html>
<head>
    <title>Gins Example</title>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=0">
</head>
<body style="height:100%;width:100%;">
    <div id="app">Hello Gins Example</div>
</body>
</html>`
