package exporter

import (
	"slices"
	"strings"
)

// labelKV is a key-value pair for a label
type labelKV struct {
	key   string
	value string
}

// labels is used to build a metric name with labels.
// It is ideally created in alphabetical order by key with no duplicate keys.
// The final value for a duplicated key will be the only one reported.
type labels []labelKV

// sortAndDeduplicateKeys sorts by key in place and removes duplicate keys from labels, keeping the last value for each key
func (l *labels) sortAndDeduplicateKeys() {
	if l == nil || len(*l) <= 1 {
		return
	}

	slices.SortStableFunc(*l, func(i, j labelKV) int {
		return strings.Compare(i.key, j.key)
	})

	// remove duplicates by key, retain the last value for each key
	for i := 0; i < len(*l)-1; i++ {
		if (*l)[i].key == (*l)[i+1].key {
			*l = append((*l)[:i], (*l)[i+1:]...)
			i-- // stay at the same index to check the next element
		}
	}
}
