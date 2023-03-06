package data_structure

type trieNode struct {
	Pass  int
	End   int
	Nexts [26]*trieNode
}

type Trie struct {
	BeginChar rune
	Root      *trieNode
}

func NewTrie(beginChar rune) *Trie {
	return &Trie{BeginChar: beginChar, Root: &trieNode{}}
}

func (t Trie) Insert(word string) {
	cur := t.Root
	for _, v := range word {
		path := int(v) - int(t.BeginChar)
		if cur.Nexts[path] == nil {
			cur.Nexts[path] = &trieNode{}
		}
		cur = cur.Nexts[path]
		cur.Pass++
	}
	cur.End++
}

// Search 返回字符串出现了多少次
func (t Trie) Search(word string) int {
	cur := t.Root
	for _, v := range word {
		path := int(v) - int(t.BeginChar)
		if cur.Nexts[path] == nil {
			return 0
		}
		cur = cur.Nexts[path]
	}
	return cur.End
}

func (t Trie) Delete(word string) {
	if t.Search(word) <= 0 {
		return
	}
	cur := t.Root
	for _, v := range word {
		path := int(v) - int(t.BeginChar)
		if cur.Nexts[path].Pass == 1 {
			cur.Nexts[path] = nil
			return
		}
	}
	cur.End--
}

// PrefixNumber 返回字符串出现了多少次
func (t Trie) PrefixNumber(word string) int {
	cur := t.Root
	for _, v := range word {
		path := int(v) - int(t.BeginChar)
		if cur.Nexts[path] == nil {
			return 0
		}
		cur = cur.Nexts[path]
	}
	return cur.Pass
}
