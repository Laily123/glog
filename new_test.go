package log_test

import (
	"github.com/Laily123/log"
	"testing"
)

func TestSetOutput(t *testing.T) {
	log.SetOutput("/testlog/")
	log.SetLevel(log.LEVEL_DEBUG)
	log.Info("bb")
	log.Debug("aa")
	log.Flush()
}
