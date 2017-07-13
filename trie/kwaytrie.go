package trie

type Alphabet interface {
	Len() int
	ToIndex(c uint8) int
	ToChar(index int) uint8
}

func NewNode(r int) *Node {
	node := &Node{}
	node.next = make([]*Node, r)
	return node
}

type Node struct {
	value interface{}
	next  []*Node
}

func NewKWayTrie(al Alphabet) KWayTrie {
	return KWayTrie{nil, al}
}

type KWayTrie struct {
	root  *Node
	alpha Alphabet
}

// Get gets the value paird with key, nil if key is absent
func (t *KWayTrie) Get(key string) interface{} {
	node := t.get(t.root, key, 0)
	if node == nil {
		return nil
	}
	return node.value
}
func (t *KWayTrie) get(node *Node, key string, d int) *Node {
	if node == nil {
		return nil
	}
	if len(key) == d {
		return node
	}
	return t.get(node.next[t.alpha.ToIndex(key[d])], key, d+1)
}

// Put puts key-value pair into the trie, remove key if value is nil
func (t *KWayTrie) Put(key string, value interface{}) {
	t.root = t.put(t.root, key, value, 0)
}

func (t *KWayTrie) put(node *Node, key string, val interface{}, d int) *Node {
	if node == nil {
		node = NewNode(t.alpha.Len())
	}
	if len(key) == d {
		node.value = val
		return node
	}
	index := t.alpha.ToIndex(key[d])
	if node.next[index] == nil {
		node.next[index] = NewNode(t.alpha.Len())
	}
	t.put(node.next[index], key, val, d+1)
	return node
}

// Delete removes key and it's value
func (t *KWayTrie) Delete(key string) {
}

// Contains checks if there is a value with key
func (t *KWayTrie) Contains(key string) bool {
	return t.Get(key) != nil
}

// IsEmpty checks if the trie is empty
func (t *KWayTrie) IsEmpty() bool {
	return t.root == nil
}

// Size returns the number of key-value pairs
func (t *KWayTrie) Size() int {
	return t.size(t.root)
}

func (t *KWayTrie) size(node *Node) int {
	if node == nil {
		return 0
	}
	var size int
	if node.value != nil {
		size = 1
	}

	for _, n := range node.next {
		size += t.size(n)
	}
	return size
}

// Keys returns all keys in the trie
func (t *KWayTrie) Keys() []string {
	return t.KeysWithPrefix("")
}

// LongestPrefixOf returns the longest key that is prefix of s
func (t *KWayTrie) LongestPrefixOf(s string) string {
	return ""
}

// KeysWithPrefix returns all keys having s as prefix
func (t *KWayTrie) KeysWithPrefix(s string) []string {
	node := t.get(t.root, s, 0)
	var keys []string
	if node == nil {
		return keys
	}
	t.collect(node, s, &keys)
	return keys
}
func (t *KWayTrie) collect(node *Node, prefix string, keys *[]string) {
	if node.value != nil {
		*keys = append(*keys, prefix)
	}

	for i, n := range node.next {
		if n != nil {
			t.collect(n, prefix+string(t.alpha.ToChar(i)), keys)
		}
	}
}

// KeysThatMatch returns all keys that matches s (where . matches any character)
func (t *KWayTrie) KeysThatMatch(s string) []string {
	return []string{}
}
