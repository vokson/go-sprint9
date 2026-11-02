package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{name: "zero size", size: 0},
		{name: "one element", size: 1},
		{name: "several elements", size: 100},
	}

	for _, tt := range tests {
		got := generateRandomElements(tt.size)
		assert.Len(t, got, tt.size, tt.name)
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name    string
		data    []int
		want    int
		wantErr bool
	}{
		{name: "nil slice", data: nil, wantErr: true},
		{name: "empty slice", data: []int{}, wantErr: true},
		{name: "single element", data: []int{42}, want: 42},
		{name: "max at start", data: []int{9, 1, 2}, want: 9},
		{name: "max in middle", data: []int{1, 9, 2}, want: 9},
		{name: "max at end", data: []int{1, 2, 9}, want: 9},
		{name: "all equal", data: []int{5, 5, 5}, want: 5},
		{name: "negative values", data: []int{-5, -1, -9}, want: -1},
		{name: "with zero", data: []int{0, 3, 1}, want: 3},
	}

	for _, tt := range tests {
		got, err := maximum(tt.data)

		if tt.wantErr {
			assert.Error(t, err, tt.name)
			continue
		}

		assert.NoError(t, err, tt.name)
		assert.Equal(t, tt.want, got, tt.name)
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name    string
		data    []int
		want    int
		wantErr bool
	}{
		{name: "nil slice", data: nil, wantErr: true},
		{name: "empty slice", data: []int{}, wantErr: true},
		{name: "single element", data: []int{7}, want: 7},
		{name: "less than chunks", data: []int{1, 2, 3}, want: 3},
		{name: "exactly chunks", data: []int{1, 2, 3, 4, 5, 6, 7, 8}, want: 8},
		{
			name: "not divisible by chunks, max in tail",
			data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want: 10,
		},
		{
			name: "max beyond last full chunk",
			data: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1000},
			want: 1000,
		},
	}

	for _, tt := range tests {
		got, err := maxChunks(tt.data)

		if tt.wantErr {
			assert.Error(t, err, tt.name)
			continue
		}

		assert.NoError(t, err, tt.name)
		assert.Equal(t, tt.want, got, tt.name)
	}
}

func TestMaxChunksMatchesMaximum(t *testing.T) {
	sizes := []int{1, 2, 7, 8, 9, 15, 16, 17, 100, 1001}

	for _, size := range sizes {
		data := generateRandomElements(size)

		want, err := maximum(data)
		assert.NoError(t, err, "size=%d", size)

		got, err := maxChunks(data)
		assert.NoError(t, err, "size=%d", size)

		assert.Equal(t, want, got, "size=%d", size)
	}
}
