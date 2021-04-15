package mbdpay

import (
	"bytes"
	"fmt"
	"sort"

	"github.com/datewu/security"
)

const (
	apiAddress = "https://api.mianbaoduo.com"
)

type client struct {
	id, key string
}

func newClient(id, key string) *client {
	c := new(client)
	c.id = id
	c.key = key
	return c
}

func (c client) sign(params map[string]string) string {
	keys := make([]string, len(params))
	i := 0
	for k := range params {
		keys[i] = k
		i++
	}
	sort.Strings(keys)

	var plain bytes.Buffer
	for _, k := range keys {
		kv := fmt.Sprintf("%s=%s&", k, params[k])
		plain.WriteString(kv)
	}
	plain.WriteString("key=" + c.key)
	hashBytes := security.SimpleMd5(plain.Bytes())
	return security.ToHexString(hashBytes)
}
