package expect

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"
)

func TestAwaitDetectsSubstring(t *testing.T) {
	e := New(t, strings.NewReader("ready set go"), io.Discard)
	e.Await("set")
}

func TestRecvMatchesExactString(t *testing.T) {
	e := New(t, strings.NewReader("pong"), io.Discard)
	e.Recv("pong")
}

func TestEchoRoundTrip(t *testing.T) {
	var sent bytes.Buffer
	e := New(t, strings.NewReader("ping"), &sent)
	e.Echo("ping")

	if got := sent.String(); got != "ping" {
		t.Fatalf("expected writer to capture %q, got %q", "ping", got)
	}
}

func TestSendWritesAllBytes(t *testing.T) {
	var sent bytes.Buffer
	e := New(t, strings.NewReader(""), &sent)
	e.Send("ok")

	if got := sent.String(); got != "ok" {
		t.Fatalf("expected writer to capture %q, got %q", "ok", got)
	}
}

func ExampleExpect_Await() {
	t := &testing.T{}
	e := New(t, strings.NewReader("ready set go"), io.Discard)
	e.Await("set")
	fmt.Println("matched set")
	// Output:
	// matched set
}

func ExampleExpect_Recv() {
	t := &testing.T{}
	e := New(t, strings.NewReader("pong"), io.Discard)
	e.Recv("pong")
	fmt.Println("received pong")
	// Output:
	// received pong
}

func ExampleExpect_Echo() {
	t := &testing.T{}
	var sent bytes.Buffer
	e := New(t, strings.NewReader("ping"), &sent)
	e.Echo("ping")
	fmt.Println(sent.String())
	// Output:
	// ping
}

func ExampleExpect_Send() {
	t := &testing.T{}
	var sent bytes.Buffer
	e := New(t, strings.NewReader(""), &sent)
	e.Send("ok")
	fmt.Println(sent.String())
	// Output:
	// ok
}
