package jmespath

import "github.com/zhangdapeng520/zdpgo_nacos/jmespath"

// Fuzz will fuzz test the JMESPath parser.
func Fuzz(data []byte) int {
	p := jmespath.NewParser()
	_, err := p.Parse(string(data))
	if err != nil {
		return 1
	}
	return 0
}
