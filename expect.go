package expect

import (
	"bufio"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type Expect struct {
	r bufio.Reader
	w io.Writer
	t testing.TB
}

func New(t testing.TB, r io.Reader, w io.Writer) *Expect {
	t.Helper()
	return &Expect{r: *bufio.NewReader(r), w: w, t: t}
}

// Await consumes input until it detects s.
func (e *Expect) Await(s string) {
	e.t.Helper()
	var sb strings.Builder
	require := require.New(e.t)
	for {
		b, err := e.r.ReadByte()
		require.NoError(err)
		sb.WriteByte(b)
		if strings.HasSuffix(sb.String(), s) {
			break
		}
	}
	e.t.Logf("received %q", s)
}

func (e *Expect) Recv(s string) {
	e.t.Helper()
	var sb strings.Builder
	require := require.New(e.t)
	for i, c := range s {
		b, err := e.r.ReadByte()
		require.NoError(err)
		sb.WriteByte(b)
		require.EqualValues(c, b, "at position %d, expected %q, got %q; %q", i, c, b, sb.String())
	}
	e.t.Logf("received %q", s)
}

func (e *Expect) Echo(s string) {
	e.t.Helper()
	e.t.Logf("sending %q", s)
	require := require.New(e.t)
	var sb strings.Builder
	for i, c := range s {
		_, err := e.w.Write([]byte{byte(c)})
		require.NoError(err)
		b, err := e.r.ReadByte()
		require.NoError(err)
		sb.WriteByte(b)
		require.EqualValues(c, b, "at position %d, expected %q, got %q; %q", i, c, b, sb.String())
	}
	e.t.Logf("received %q", s)
}

func (e *Expect) Send(s string) {
	e.t.Helper()
	e.t.Logf("sending %q", s)
	for _, c := range s {
		_, err := e.w.Write([]byte{byte(c)})
		require.NoError(e.t, err)
		time.Sleep(10 * time.Millisecond)
	}
}
