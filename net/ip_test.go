package nettool

import (
	"log"
	"testing"
)

func TestIp(t *testing.T) {
	ip, err := GetLocalIp()
	if err != nil {
		t.Error(err)
	}
	log.Println(ip)
}
