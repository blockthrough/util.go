package victoria

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLabelsSortAndDeduplicateKeys(t *testing.T) {
	tests := []struct {
		name string
		l    labels
		want labels
	}{
		{
			name: "nil",
			l:    nil,
			want: nil,
		},
		{
			name: "empty",
			l:    labels{},
			want: labels{},
		},
		{
			name: "single",
			l:    labels{{"a", "1"}},
			want: labels{{"a", "1"}},
		},
		{
			name: "pair in order",
			l:    labels{{"a", "1"}, {"b", "2"}},
			want: labels{{"a", "1"}, {"b", "2"}},
		},
		{
			name: "pair out of order",
			l:    labels{{"b", "2"}, {"a", "1"}},
			want: labels{{"a", "1"}, {"b", "2"}},
		},
		{
			name: "pair duplicate key",
			l:    labels{{"a", "1"}, {"a", "2"}},
			want: labels{{"a", "2"}},
		},
		{
			name: "pair duplicate key, other between",
			l:    labels{{"a", "1"}, {"b", ""}, {"a", "2"}},
			want: labels{{"a", "2"}, {"b", ""}},
		},
		{
			name: "triple duplicate key",
			l:    labels{{"a", "1"}, {"a", "2"}, {"a", "3"}},
			want: labels{{"a", "3"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.l.sortAndDeduplicateKeys()
			assert.EqualValues(t, tt.want, tt.l, "labels.sortAndDeduplicateKeys() = %v, want %v", tt.l, tt.want)
		})
	}
}
