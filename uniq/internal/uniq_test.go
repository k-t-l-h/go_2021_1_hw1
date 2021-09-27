package internal

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUniq(t *testing.T) {
	var cases = []struct {
		input    []string
		params   Params
		expected []string
	}{
		{
			input: []string{
				"Test", "test", "Test",
			},
			params:   Params{},
			expected: []string{"Test", "test", "Test"},
		},
		{
			input: []string{
				"Test", "test",
			},
			params:   Params{Case: true, Double: true},
			expected: []string{"Test"},
		},
		{
			input: []string{
				"Test", "test",
			},
			params:   Params{Case: true, Uniq: true},
			expected: []string{},
		},
		{
			input: []string{
				"I love music.", "I love music.", "I love music.", "I love music of Kartik.",
				"I love music of Kartik.", "Thanks.", "I love music of Kartik.", "I love music of Kartik.",
			},
			params:   Params{},
			expected: []string{"I love music.", "I love music of Kartik.", "Thanks.", "I love music of Kartik."},
		},
		{
			input: []string{
				"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.",
				"I love music of Kartik.", "Thanks.", "I love music of Kartik.", "I love music of Kartik.",
			},
			params:   Params{Count: true},
			expected: []string{"3 I love music.", "1 ", "2 I love music of Kartik.", "1 Thanks.", "2 I love music of Kartik."},
		},
		{
			input: []string{
				"I love music.", "I love music.", "I love music.", "I love music of Kartik.",
				"I love music of Kartik.", "Thanks.", "I love music of Kartik.", "I love music of Kartik.",
			},
			params:   Params{Double: true},
			expected: []string{"I love music.", "I love music of Kartik.", "I love music of Kartik."},
		},
		{
			input: []string{
				"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.",
				"I love music of Kartik.", "Thanks.", "I love music of Kartik.", "I love music of Kartik.",
			},
			params:   Params{Uniq: true},
			expected: []string{"", "Thanks."},
		},
		{
			input: []string{
				"I LOVE MUSIC.", "I love music.", "I LoVe music.", "", "I love MuSIC of Kartik.",
				"I love music of Kartik.", "Thanks.", "I love music of Kartik.", "I love MuSIC of Kartik.",
			},
			params:   Params{Case: true},
			expected: []string{"I LOVE MUSIC.", "", "I love MuSIC of Kartik.", "Thanks.", "I love music of Kartik."},
		},
		{
			input: []string{
				"We love music.", "I love music.", "I love music.", "I love music of Kartik.",
				"We love music of Kartik.", "Thanks.",
			},
			params:   Params{FieldsOffset: 1},
			expected: []string{"We love music.", "I love music of Kartik.", "Thanks."},
		},
		{
			input: []string{
				"We love music.", "I love music.", "I love music.", "I love music of Kartik.",
				"We love music of Kartik.", "Thanks.",
			},
			params:   Params{FieldsOffset: 1},
			expected: []string{"We love music.", "I love music of Kartik.", "Thanks."},
		},
		{
			input: []string{
				"We love music.", "I love music.", "I love music.", "I love music of Kartik.",
				"We love music of Kartik.", "Thanks.",
			},
			params: Params{CharOffset: 1},
			expected: []string{"We love music.", "I love music.", "I love music of Kartik.",
				"We love music of Kartik.", "Thanks."},
		},
		{
			input: []string{
				"A love.", "B love.", "C love.",
			},
			params:   Params{CharOffset: 1},
			expected: []string{"A love."},
		},
		{
			input: []string{
				"A dove.", "B love.", "C love.",
			},
			params:   Params{CharOffset: 1, FieldsOffset: 1},
			expected: []string{"A dove."},
		},
		{
			input: []string{
				"A dove.", "B love.", "C love.", "D LOVE.",
			},
			params:   Params{CharOffset: 1, FieldsOffset: 1, Case: true},
			expected: []string{"A dove."},
		},
	}
	for _, item := range cases {
		ok := Uniq(item.input, item.params)
		require.Equal(t, item.expected, ok)
	}
}
