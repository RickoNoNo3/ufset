// Package ufset implements a disjoint set (union-find) data structure.
// It provides efficient union and find operations with optional path compression.
package ufset

// setNode represents a node in the disjoint set forest.
type setNode[K comparable] struct {
	Key    K
	rank   int64
	parent *setNode[K]
	Sets   *DisjointSets[K]
}

// Find finds the root of the set containing the node.
// If pathCompress is true, it applies path compression to flatten the structure.
func (n *setNode[K]) Find(pathCompress bool) *setNode[K] {
	root := n
	for root != root.parent {
		root = root.parent
	}
	if pathCompress {
		t := n
		for t != root {
			pa := t.parent
			t.parent = root
			t = pa
		}
	}
	return root
}

// Union merges the sets containing the two nodes.
// If pathCompress is true, it applies path compression.
// If unionByRank is true, it uses union by rank to keep the tree flat, only available if pathCompress is true.
func (n *setNode[K]) Union(n2 *setNode[K], pathCompress bool, unionByRank bool) {
	if n.Sets != n2.Sets {
		panic("can't union nodes from different disjoint sets")
	}
	p1 := n.Find(pathCompress)
	p2 := n2.Find(pathCompress)
	if p1 == p2 {
		return
	}
	if pathCompress && unionByRank {
		if p1.rank < p2.rank {
			p1.parent = p2
		} else {
			p2.parent = p1
			if p1.rank == p2.rank {
				p1.rank++
			}
		}
	} else if pathCompress {
		p2.parent = p1
	} else {
		p2.parent = n
	}
}

// DisjointSets implements a union-find data structure with optional path compression and union by rank.
type DisjointSets[K comparable] struct {
	nodes    map[K]*setNode[K]
	compress bool
}

// New creates a new DisjointSets instance which uses path compression and union by rank.
// It allows for efficient union and find operations on disjoint sets, with time complexity of O(alpha(n)).
func New[K comparable]() *DisjointSets[K] {
	return &DisjointSets[K]{
		nodes:    make(map[K]*setNode[K]),
		compress: true,
	}
}

// NewRigid creates a new DisjointSets instance which does not use path compression or union by rank.
// It allows for union and find operations on disjoint sets, but does not alter the depth of the trees.
// The union operation does not merge the source tree to the root of the destination tree but just connects the source tree to the targeted node.
// The ufset.GetTree function receives this kind of DisjointSets instance and returns the tree structure of merging history.
// It has a time complexity of O(n) for every union and find operation.
func NewRigid[K comparable]() *DisjointSets[K] {
	return &DisjointSets[K]{
		nodes:    make(map[K]*setNode[K]),
		compress: false,
	}
}

func (s *DisjointSets[K]) newNode(key K) *setNode[K] {
	node := &setNode[K]{
		Key:  key,
		Sets: s,
	}
	node.parent = node
	s.nodes[key] = node
	return node
}

func (s *DisjointSets[K]) add(key K) {
	if !s.has(key) {
		s.newNode(key)
	}
}

func (s *DisjointSets[K]) has(key K) bool {
	return s.nodes[key] != nil
}

// Find finds the root of the set containing the key.
func (s *DisjointSets[K]) Find(key K) K {
	s.add(key)
	return s.nodes[key].Find(s.compress).Key
}

// Union merges the sets containing the two keys.
func (s *DisjointSets[K]) Union(k1, k2 K) {
	s.add(k1)
	s.add(k2)
	n1 := s.nodes[k1]
	n2 := s.nodes[k2]
	n1.Union(n2, s.compress, s.compress)
}

// InSameSet checks if two keys are in the same set.
func (s *DisjointSets[K]) InSameSet(k1, k2 K) bool {
	return s.Find(k1) == s.Find(k2)
}

// IsCompressed checks if the disjoint sets use path compression.
func (s *DisjointSets[K]) IsCompressed() bool {
	return s.compress
}

type TreeNode[K comparable] struct {
	Key      K
	Children []*TreeNode[K]
}

// GetTree returns a forest of tree nodes representing the structure of the disjoint sets.
func GetTree[K comparable](s *DisjointSets[K]) (forest []*TreeNode[K]) {
	if s.IsCompressed() {
		panic("GetTree is not supported for DisjointSets with path compression enabled")
	}
	allNodes := make(map[K]*TreeNode[K])
	for key := range s.nodes {
		allNodes[key] = &TreeNode[K]{Key: key, Children: []*TreeNode[K]{}}
	}
	for key, node := range s.nodes {
		if node.parent != node {
			allNodes[node.parent.Key].Children = append(allNodes[node.parent.Key].Children, allNodes[key])
		}
	}
	forest = []*TreeNode[K]{}
	for key, treeNode := range allNodes {
		if s.nodes[key].parent.Key == key {
			forest = append(forest, treeNode)
		}
	}
	return forest
}
