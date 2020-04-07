package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"testing"
)

func init(){
	fmt.Println("init in test")
}

func TestMain(t *testing.T) {
	level := "info"
	os.Setenv("VIPER_LOGGING_LEVEL", level)
	if viper.GetString("logging.level") != level {
		t.Error("Expected", level, "got", viper.GetString("logging.level"))
	}
}
