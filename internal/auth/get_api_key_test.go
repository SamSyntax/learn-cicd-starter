package auth

import (
	"testing"
)

type Case struct {
	name       string
	input      map[string][]string
	expected   string
	shouldFail bool
}

var Cases []Case = []Case{
	{
		name:       "Case 1",
		input:      map[string][]string{"Authorization": []string{"ApiKey asfkj1kjr133u1h13ohfe1"}},
		expected:   "asfkj1kjr133u1h13ohfe1",
		shouldFail: false,
	},
	{
		name:       "Case 2",
		input:      map[string][]string{"": []string{"Bearer asfkj1kjr133u1h13ohfe1"}},
		expected:   "",
		shouldFail: true,
	},
	{
		name:       "Case 3",
		input:      map[string][]string{"Authorizat": []string{"ApiKey asfkj1kjr133u1h13ohfe1"}},
		expected:   "",
		shouldFail: true,
	},
}

func TestGetApiKey(t *testing.T) {
	for _, c := range Cases {
		key, err := GetAPIKey(c.input)
		if c.shouldFail {
			if err == nil {
				t.Errorf("[%s] expected an error, but got nil (key: %s)", c.name, key)
			}
		} else {
			if err != nil {
				t.Errorf("[%s] unexpected an error: %v", c.name, err)
			} else if key != c.expected {

				t.Errorf("[%s] expected key %q, got %q", c.name, c.expected, key)
			}
		}
	}
}
