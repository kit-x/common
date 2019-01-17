package stringx

import "testing"

func TestListContain(t *testing.T) {
	cases := []struct {
		list   []string
		s      string
		expect bool
	}{
		{
			list:   []string{"1", "2", "3"},
			s:      "1",
			expect: true,
		},
		{
			list:   []string{"1", "2", "3"},
			s:      "x",
			expect: false,
		},
		{
			list:   []string{},
			s:      "1",
			expect: false,
		},
	}

	for _, c := range cases {
		if ListContain(c.list, c.s) != c.expect {
			t.Errorf("list %v contain %s, should return %v", c.list, c.s, c.expect)
		}
	}
}
