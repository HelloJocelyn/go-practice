package test

import "testing"
func TestReverse(t *testing.T){
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _,c := range cases {
		if len(c.in) <= 0 {
			t.Errorf("length of string less than 0")
		}
	}
}
