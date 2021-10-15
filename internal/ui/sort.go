package ui

import (
	"sort"

	"github.com/MarinX/era/rpc"
)

func sortKeysAsc(keys []*rpc.Key) {
	sort.Slice(keys, func(i, j int) bool {
		return keys[i].CreatedAt.After(keys[j].CreatedAt)
	})
}

func sortContactsAsc(cnts []*rpc.Contact) {
	sort.Slice(cnts, func(i, j int) bool {
		return cnts[i].CreatedAt.After(cnts[j].CreatedAt)
	})
}
