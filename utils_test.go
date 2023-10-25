package dextk

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

type ulebCase struct {
	bytes []byte
	size  uint32
	value uint32
}

var ulebCases = []ulebCase{
	{
		bytes: []byte{0x0},
		size:  1,
		value: 0,
	},
	{
		bytes: []byte{0x1},
		size:  1,
		value: 1,
	},
	{
		bytes: []byte{0x7f},
		size:  1,
		value: 127,
	},
	{
		bytes: []byte{0x80, 0x7f},
		size:  2,
		value: 16256,
	},
}

func TestReadUleb128(t *testing.T) {
	for _, c := range ulebCases {
		bytes := bytes.NewReader(c.bytes)
		reader := &Reader{
			file: bytes,
		}

		gotVal, gotSize, gotErr := reader.readUleb128(0)

		require.NoError(t, gotErr, "read error")
		require.Equal(t, c.size, gotSize, "read size")
		require.Equal(t, c.value, gotVal, "read value")
	}
}

type slebCase struct {
	bytes []byte
	size  uint32
	value int32
}

var slebCases = []slebCase{
	{
		bytes: []byte{0x0},
		size:  1,
		value: 0,
	},
	{
		bytes: []byte{0x1},
		size:  1,
		value: 1,
	},
	{
		bytes: []byte{0x7f},
		size:  1,
		value: -1,
	},
	{
		bytes: []byte{0x80, 0x7f},
		size:  2,
		value: -128,
	},
}

func TestReadSleb128(t *testing.T) {
	for _, c := range slebCases {
		bytes := bytes.NewReader(c.bytes)
		reader := &Reader{
			file: bytes,
		}

		gotVal, gotSize, gotErr := reader.readSleb128(0)

		require.NoError(t, gotErr, "read error")
		require.Equal(t, c.size, gotSize, "read size")
		require.Equal(t, c.value, gotVal, "read value")
	}
}
