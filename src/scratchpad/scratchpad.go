package scratchpad

import (
	"sort"
	"time"
)

type Data struct {
	ID    string
	State string
}

type Pointer struct {
	createdAt *time.Time
}

func FindMatchesAndCombine(success, requested []Data) []Data {
	r := make([]Data, len(requested))
	for i := range requested {
		r[i] = requested[i]
		for j := range success {
			if requested[i].ID == success[j].ID {
				r[i] = Data{
					ID:    success[j].ID,
					State: success[j].State,
				}
			}
		}
	}
	return r
}

func FindMatchesAndCombineOp(success, requested []Data) []Data {
	sort.SliceStable(requested, func(i, j int) bool {
		return requested[i].ID > requested[j].ID
	})
	for _, s := range success {
		index := sort.Search(len(requested), func(i int) bool {
			return requested[i].ID >= s.ID
		})
		if index < len(requested) && requested[index].ID == s.ID {
			requested[index] = s
		}
	}
	return requested
}
