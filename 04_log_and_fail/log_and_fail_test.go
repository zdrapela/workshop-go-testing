package log_and_fail

import "testing"

func TestLogAndFail(t *testing.T) {
	t.Log("Testing Foo")

	// t.Fail()
	// t.Logf("Log number %d from Foo", 2)

	// t.FailNow()
	// t.Log("Log after FailNow")

	// t.Error("Log, signal error and continue")

	// t.Fatalf("Log, signal %s and stop", "fatal")
	// t.Log("Log after Fatalf")
}

// go test -v

// Error()  = Log()		+ Fail()
// Errorf() = Logf() 	+ Fail()
// Fatal()  = Log() 	+ FailNow()
// Fatalf() = Logf()	+ FailNow()
