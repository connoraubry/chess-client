package client

import "testing"

func TestBaseAddress(t *testing.T) {
	cfg := ClientConfig{Server: "http://test.com", Port: 3030}
	c := Client{cfg: cfg}

	got := c.getBaseAddress()
	exp := "http://test.com:3030/api/v1/"

	if got != exp {
		t.Fatalf(`getBaseAddress() failed. Got %v. Expected %v`, got, exp)
	}

	c.cfg.Server = "http://test.com/"
	got = c.getBaseAddress()

	if got != exp {
		t.Fatalf(`getBaseAddress() failed. Got %v. Expected %v`, got, exp)
	}
}
