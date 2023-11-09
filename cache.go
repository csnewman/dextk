package dextk

import (
	"io"
)

type cachedReader struct {
	slots   int
	store   []byte
	mapping []int64
	lengths []int
	r       io.ReaderAt
}

func newCachedReader(r io.ReaderAt, slots int) *cachedReader {
	store := make([]byte, 4096*slots)
	mappings := make([]int64, slots)
	lengths := make([]int, slots)

	for i := 0; i < slots; i++ {
		mappings[i] = -1
	}

	return &cachedReader{
		slots,
		store,
		mappings,
		lengths,
		r,
	}
}

func (r *cachedReader) ReadAt(tgt []byte, pos int64) (int, error) {
	remaining := len(tgt)
	cur := 0

	for remaining > 0 {
		chunk := pos >> 12
		inner := int(pos & 0xFFF)

		data, err := r.getChunk(chunk)
		if err != nil {
			return len(tgt) - remaining, err
		}

		rsize := len(data) - inner
		if rsize > remaining {
			rsize = remaining
		}

		copy(tgt[cur:cur+rsize], data[inner:inner+rsize])

		remaining -= rsize
		cur += rsize
		pos += int64(rsize)

		if remaining > 0 && rsize != 4096-inner {
			return len(tgt) - remaining, io.EOF
		}
	}

	if remaining > 0 {
		panic("remaining != 0")
	}

	return len(tgt), nil
}

func (r *cachedReader) getChunk(chunk int64) ([]byte, error) {
	slot := chunk % int64(r.slots)

	slotStart := int(slot << 12)
	slotEnd := slotStart + 4096

	if r.mapping[slot] == chunk {
		l := r.lengths[slot]

		return r.store[slotStart : slotStart+l], nil
	}

	data := r.store[slotStart:slotEnd]

	size, err := r.r.ReadAt(data, chunk<<12)
	if err != nil && err != io.EOF {
		return nil, err
	}

	r.mapping[slot] = chunk
	r.lengths[slot] = size

	block := data[:size]
	return block, nil
}
