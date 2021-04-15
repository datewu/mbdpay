package mbdpay

import "fmt"

func ExampleHash() {
	cli := New("abc-id", "xyz-secret-key")
	params := map[string]string{
		"key1":  "value",
		"abc":   "xyz",
		"12df":  "0834",
		"hello": "world",
		"nice":  "shot",
	}
	sign := cli.sign(params)
	fmt.Println(sign)
	// Output:
	// 28ea05ae777e1dbad038c7d58e5a3167
}
