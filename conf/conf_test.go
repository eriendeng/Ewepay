package conf_test

import (
	"testing"
	"../conf"
	"fmt"
)

func TestC_Use(t *testing.T) {
	conf.LoadConfig("/Users/erien/go/myprojects/Ewepay/.properties").Use("DEVELOPMENT")
	fmt.Println()
	fmt.Printf("%v",conf.GetConfig())
	fmt.Println("\n")
}
