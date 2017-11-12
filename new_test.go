package log_test

import (
	"github.com/Laily123/log"
	"testing"
)

func TestSetOutput(t *testing.T) {
	log.SetOutput("/aa.log")
	log.Info("bb")
	log.Debug("aa")
}
