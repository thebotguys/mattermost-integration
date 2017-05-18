package logrus

import (
	"encoding/json"
	"testing"
)

func TestAttachmentMarshalJSON(t *testing.T) {
	tests := []struct {
		desc string
		in   Attachment
		exp  string
	}{
		{
			desc: "With empty fields",
			in:   Attachment{},
			exp:  `{}`,
		},
		{
			desc: "With empty author_icon",
			in: Attachment{
				AuthorIcon: "",
				AuthorLink: "authorlink",
			},
			exp: `{"author_link":"authorlink"}`,
		},
		{
			desc: "With fields",
			in: Attachment{
				AuthorIcon: "",
				AuthorLink: "authorlink",
				Fields: Fields{
					{
						Short: false,
						Title: "t1",
						Value: "v1",
					},
					{
						Short: true,
					},
					{
						Short: true,
						Title: "t3",
						Value: "v3",
					},
				},
			},
			exp: `{"author_link":"authorlink","fields":[{"short":false,"title":"t1","value":"v1"},{"short":true,"title":"t3","value":"v3"}]}`,
		},
	}

	for _, test := range tests {
		t.Log(test.desc)

		got, err := json.Marshal(test.in)
		if err != nil {
			t.Fatal(err)
		}

		assert(t, test.exp, string(got), true)
	}
}
