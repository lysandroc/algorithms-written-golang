package tree

import (
	"errors"
	"fmt"
	"sort"
)

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	sort.SliceStable(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	if records[0].ID != 0 || records[0].ID != records[0].Parent {
		return nil, errors.New("Invalid root")
	}

	parents := make(map[int]*Node)

	root := Node{ID: records[0].ID}
	parents[root.ID] = &root

	for _, record := range records[1:] {
		parent := parents[record.Parent]

		if record.ID <= record.Parent || record.ID >= len(records) {
			return nil, fmt.Errorf("outside of range")
		}

		if parents[record.ID] != nil {
			return nil, errors.New("already have a node")
		}

		if parent == nil {
			return nil, errors.New("invalid parent")
		}
		node := Node{ID: record.ID}
		parents[node.ID] = &node
		parent.Children = append(parent.Children, &node)
	}

	return &root, nil
}
