package waitfor_test

import (
	"os"
	"testing"
	"testing_tasks/task_3/waitfor"
	"testing_tasks/task_3/waitfor/fake_server"
)

type testServerSettings struct {
	Host      string
	Port      int
	Delay     int
	DelayOK   int
	DelayFail int
}

var settings testServerSettings

func startTestServer() {
	srv := fake_server.New(settings.Host, settings.Port)
	srv.StartWithDelay(settings.Delay)
}

func init() {
	settings = testServerSettings{
		Host:      "localhost",
		Port:      12345,
		Delay:     3,
		DelayOK:   4,
		DelayFail: 2,
	}
}

func TestMain(m *testing.M) {
	startTestServer()
	os.Exit(m.Run())
}

func TestIt_Ok(t *testing.T) {
	t.Parallel()
	err := waitfor.It(settings.Host, settings.Port, settings.DelayOK)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
}

func TestIt_Fail(t *testing.T) {
	t.Parallel()
	err := waitfor.It(settings.Host, settings.Port, settings.DelayFail)
	if err == nil {
		t.Errorf("Expected an error, got: %s", err)
	}
}
