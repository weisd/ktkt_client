package quan

import (
	"testing"
)

func TestClient(t *testing.T) {
	InitClient("http://127.0.0.1:3033")

	order, err := OrdersClient.OrdersGetInfoById(3)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(order)

	comment, err := KtktQuanCommentClient.QuanGetCommentInfoById(1)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(comment)

	thread, err := KtktQuanThreadClient.QuanGetThreadInfoById(1)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(thread)

}
