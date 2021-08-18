package api

import (
	"testing"

	"pmo-test4.yz-intelligence.com/kit/component/apiconstant"

	"pmo-test4.yz-intelligence.com/kit/component/gins/test"
)

func TestDemoPostByJSON(t *testing.T) {
	req := test.NewTestRequest()
	req.SetJSON()
	req.AddJSON("id", "123")

	res, err := req.Call("post", "/api/demo/")
	if err != nil {
		t.Errorf("\n请求失败：%s", err)
		return
	}

	t.Logf("\n响应内容：%s", res.Content())

	if res.Code != apiconstant.RESPONSE_OK {
		t.Errorf("\n响应错误：代码应该为 [%d]，现在得到 [%d]", apiconstant.RESPONSE_OK, res.Code)
	}
}

func TestDemoPostByForm(t *testing.T) {
	req := test.NewTestRequest()
	req.AddParam("id", "123")

	res, err := req.Call("post", "/api/demo/")
	if err != nil {
		t.Errorf("\n请求失败：%s", err)
		return
	}

	t.Logf("\n响应内容：%s", res.Content())

	if res.Code != apiconstant.RESPONSE_OK {
		t.Errorf("\n响应错误：代码应该为 [%d]，现在得到 [%d]", apiconstant.RESPONSE_OK, res.Code)
	}
}

func TestDemoPanic(t *testing.T) {
	req := test.NewTestRequest()

	res, err := req.Call("get", "/api/demo/panic")
	if err != nil {
		t.Errorf("\n请求失败：%s", err)
		return
	}

	t.Logf("\n响应内容：%s", res.Content())

	if res.Code != apiconstant.RESPONSE_OK {
		t.Errorf("\n响应错误：代码应该为 [%d]，现在得到 [%d]", apiconstant.RESPONSE_OK, res.Code)
	}
}
