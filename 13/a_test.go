package day13

import (
	"encoding/json"
	"strings"
	"testing"

	"aoc-2022/pkg/utils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadRawPacketLine(t *testing.T) {
	type testData struct {
		input    string
		expected packetData
	}
	tts := []testData{
		{
			input: "[1,1,3,1,1]",
			expected: packetData{
				l: []packetData{
					{v: utils.AsPointer(1)},
					{v: utils.AsPointer(1)},
					{v: utils.AsPointer(3)},
					{v: utils.AsPointer(1)},
					{v: utils.AsPointer(1)},
				},
			},
		},
		{
			input: "[1,[2,[3,[4,[5,6,7]]]],8,9]",
			expected: packetData{
				l: []packetData{
					{v: utils.AsPointer(1)},
					{l: []packetData{
						{v: utils.AsPointer(2)},
						{l: []packetData{
							{v: utils.AsPointer(3)},
							{l: []packetData{
								{v: utils.AsPointer(4)},
								{l: []packetData{
									{v: utils.AsPointer(5)},
									{v: utils.AsPointer(6)},
									{v: utils.AsPointer(7)},
								}},
							}},
						}},
					}},
					{v: utils.AsPointer(8)},
					{v: utils.AsPointer(9)},
				},
			},
		},
	}

	for _, tt := range tts {
		t.Run(tt.input, func(t *testing.T) {
			jsonStr := []byte(tt.input)
			expected := tt.expected
			var rawData rawPacketData
			require.NoError(t, json.Unmarshal(jsonStr, &rawData))

			actual := NewPacketDataFromRaw(rawData)

			assert.Equal(t, expected, actual)
		})
	}

}

func TestCheckOrder(t *testing.T) {
	type testData struct {
		pairLines []string
		expected  order
	}
	tts := []testData{
		{
			pairLines: []string{"[1,1,3,1,1]", "[1,1,5,1,1]"},
			expected:  orderCorrect,
		},
		{
			pairLines: []string{"[[1],[2,3,4]]", "[[1],4]"},
			expected:  orderCorrect,
		},
		{
			pairLines: []string{"[9]", "[[8,7,6]]"},
			expected:  orderIncorrect,
		},
		{
			pairLines: []string{"[]", "[3]"},
			expected:  orderCorrect,
		},
		{
			pairLines: []string{"[[[]]]", "[[]]"},
			expected:  orderIncorrect,
		},
		{
			pairLines: []string{"[1,[2,[3,[4,[5,6,7]]]],8,9]", "[1,[2,[3,[4,[5,6,0]]]],8,9]"},
			expected:  orderIncorrect,
		},
		{
			pairLines: []string{"[4, 4]", "[4, 4]"},
			expected:  orderContinue,
		},
	}

	for _, tt := range tts {
		t.Run(strings.Join(tt.pairLines, "\n"), func(t *testing.T) {
			pair := pairFromLines(tt.pairLines[0], tt.pairLines[1])

			actual := pair.checkOrder()

			assert.Equal(t, tt.expected, actual)
		})
	}

}
