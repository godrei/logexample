package log_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/godrei/logexample/log"
	"github.com/stretchr/testify/require"
)

func TestCommand(t *testing.T) {
	var buff bytes.Buffer
	log.SetOut(&buff)

	in := "xcode build --project myproject.xcproj"
	log.Command(in)

	var out map[string]string
	require.NoError(t, json.Unmarshal(buff.Bytes(), &out))
	require.NotEqual(t, "", out["@time"])
	require.Equal(t, "cmd", out["@type"])
	require.Equal(t, in, out["msg"])
}

func TestInfo(t *testing.T) {
	var buff bytes.Buffer
	log.SetOut(&buff)

	in := "Info message"
	log.Info(in)

	var out map[string]string
	require.NoError(t, json.Unmarshal(buff.Bytes(), &out))
	require.NotEqual(t, "", out["@time"])
	require.Equal(t, "inf", out["@type"])
	require.Equal(t, in, out["msg"])
}
