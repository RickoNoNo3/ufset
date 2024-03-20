package generic_disjoint_set

type setNode[K comparable] struct {
	Key    K
	rank   int64
	parent *setNode[K]
	Sets   *DisjointSets[K]
}

func (n *setNode[K]) Find() *setNode[K] {
	for n.parent != n.parent.parent {
		n.parent = n.parent.Find()
	}
	return n.parent
}

func (n *setNode[K]) Union(n2 *setNode[K]) {
	if n.Sets != n2.Sets {
		panic("can't union nodes from different disjoint sets")
	}
	p1 := n.Find()
	p2 := n2.Find()
	if p1 == p2 {
		return
	}
	if p1.rank < p2.rank {
		p1.parent = p2
	} else {
		p2.parent = p1
		if p1.rank == p2.rank {
			p1.rank++
		}
	}
}

type DisjointSets[K comparable] struct {
	nodes map[K]*setNode[K]
}

func New[K comparable]() *DisjointSets[K] {
	return &DisjointSets[K]{
		nodes: make(map[K]*setNode[K]),
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

func (s *DisjointSets[K]) Find(key K) K {
	s.add(key)
	return s.nodes[key].Find().Key
}

func (s *DisjointSets[K]) Union(k1, k2 K) {
	s.add(k1)
	s.add(k2)
	n1 := s.nodes[k1]
	n2 := s.nodes[k2]
	n1.Union(n2)
}

func (s *DisjointSets[K]) InSameSet(k1, k2 K) bool {
	return s.Find(k1) == s.Find(k2)
}
