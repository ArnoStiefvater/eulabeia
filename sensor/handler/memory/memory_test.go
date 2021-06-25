package memory

import (
	"errors"
	"testing"

	"github.com/greenbone/eulabeia/messages"
	mem "github.com/mackerelio/go-osstat/memory"
)

func TestMemoryHandlerErrors(t *testing.T) {
	const reactOn = "get.memory"
	var tests = []struct {
		mt  string
		f   func() (*mem.Stats, error)
		err bool
	}{
		{reactOn, func() (*mem.Stats, error) { return nil, nil }, true},
		{reactOn, func() (*mem.Stats, error) { return nil, errors.New("someerror") }, true},
		{reactOn, func() (*mem.Stats, error) { return &mem.Stats{}, nil }, false},
		{"nope", func() (*mem.Stats, error) { panic("do not call me") }, false},
	}
	for i, test := range tests {
		h := getMemory{
			stats: test.f,
		}
		_, _, err := h.Get(messages.Get{
			Message: messages.NewMessage(test.mt, "", ""),
			ID:      "it's ignored here",
		})
		if (err != nil) != test.err {
			t.Errorf("[%d] expected error == %t; error: %v", i, test.err, err)
		}
	}
}
