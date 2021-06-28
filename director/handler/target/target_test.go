package target

import (
	"testing"

	"github.com/greenbone/eulabeia/internal/test"
	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/messages/handler"
)

func TestSuccessResponse(t *testing.T) {
	h := handler.New(handler.FromAggregate(New(NoopStorage{})))
	tests := []test.HandleTests{
		{
			Input: messages.Create{
				Message: messages.NewMessage("create.target", "1", "1"),
			},
			Handler:         h,
			ExpectedMessage: messages.NewMessage("created.target", "1", "1"),
		},
		{
			Input: messages.Get{
				Message: messages.NewMessage("get.target", "1", "1"),
				ID:      "someid",
			},
			Handler:         h,
			ExpectedMessage: messages.NewMessage("got.target", "1", "1"),
		},
		{
			Input: messages.Modify{
				Message: messages.NewMessage("modify.target", "1", "1"),
				ID:      "1",
				Values: map[string]interface{}{
					"scanner":  "openvas",
					"hosts":    []string{"a", "b"},
					"plugins":  []string{"a", "b"},
					"alive":    true,
					"parallel": false,
				},
			},
			Handler:         h,
			ExpectedMessage: messages.NewMessage("modified.target", "1", "1"),
		},
	}
	for _, test := range tests {
		test.Verify(t)
	}
}
