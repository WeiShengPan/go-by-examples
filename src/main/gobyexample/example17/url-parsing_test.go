package main

import (
	"fmt"
	"net/url"
	"strings"
	"testing"
)

// URL解析
func TestUrlParsing(t *testing.T) {

	var p = fmt.Println

	s := "postgre://user:pass@host.com:8888/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	p(u.Scheme)

	p(u.User)
	p(u.User.Username())
	pwd, _ := u.User.Password()
	p(pwd)

	p(u.Host)
	h := strings.Split(u.Host, ":")
	p(h[0])
	p(h[1])

	p(u.Path)
	p(u.Fragment)

	p(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	p(m)
	p(m["k"])
	p(m["k"][0])

}
