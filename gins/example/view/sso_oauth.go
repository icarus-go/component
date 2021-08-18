package view

import "pmo-test4.yz-intelligence.com/kit/component/gins"

func init() {
	gins.AddTemplate("sso_oauth", ssoOAuth)
}

const ssoOAuth = `<!DOCTYPE html>
<html>
<head>
    <title>Gins SSO</title>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=0">
</head>
<body style="height:100%;width:100%;">
    <div id="app">Hello Gins SSO OAuth</div>
</body>
</html>`
