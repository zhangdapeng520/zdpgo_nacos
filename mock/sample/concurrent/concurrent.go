// Package concurrent demonstrates how to use gomock with goroutines.
package concurrent

//go:generate mockgen -destination mock/concurrent_mock.go github.com/zhangdapeng520/zdpgo_nacos/mock/sample/concurrent Math

type Math interface {
	Sum(a, b int) int
}
