package main

import "testing"

func TestIsEven(t *testing.T) {
	// t.Parallel() // параллельно, но всё равно один тест тормозиться, чтобы выполнить другой

	// Table-driver test
	var testcases = []struct {
		text string
		got  int
		want string
	}{
		{"-5 - отрицательное", -5, "no"},
		{"0 - zero", 0, "yes"},
		{"2 - Even", 2, "yes"},
		{"1 - Not Even", 99999999999991, "no"},
	}
	for _, tc := range testcases {
		t.Run(tc.text, func(t *testing.T) {
			result := IsEven(tc.got)
			if result != tc.want {
				t.Errorf("Error in test %s: got: %s want: %s",
					tc.text, result, tc.want)
			}
		})
	}
}
