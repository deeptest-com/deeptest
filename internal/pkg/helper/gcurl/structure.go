package curlHelper

import (
	"container/heap"
)

// trieWord Trie 需要的Word接口
type trieWord interface {
	GetWord() string
}

// TrieStrWord 最简单的TrieWord 结构
type trieStrWord string

// GetWord 获取单词
func (tsw *trieStrWord) GetWord() string {
	return (string)(*tsw)
}

// Trie 前缀树
type hTrie struct {
	isWord bool
	value  interface{}
	char   byte
	prev   *hTrie
	next   map[byte]*hTrie
}

// newTrie Initialize your data structure here.
func newTrie() *hTrie {
	return &hTrie{next: make(map[byte]*hTrie)}
}

// Insert a word into the trie.
func (trie *hTrie) Insert(iword trieWord) {
	cur := trie
	word := iword.GetWord()
	l := len(word)

	for i := 0; i < l; i++ {
		c := word[i]
		if next, ok := cur.next[c]; ok {
			cur = next
		} else {
			create := newTrie()
			cur.next[c] = create
			create.char = c
			create.prev = cur
			cur = create
		}
	}

	cur.isWord = true
	cur.value = iword
}

// AllWords 所有单词
func (trie *hTrie) AllWords() []string {
	var result []string
	for _, v := range trie.next {
		look(v, "", &result)
	}
	return result
}

func look(cur *hTrie, content string, result *[]string) {
	content += string(cur.char)
	if cur.isWord {
		*result = append(*result, content)
	}
	for _, v := range cur.next {
		look(v, content, result)
	}
}

// Remove 移除单词
func (trie *hTrie) Remove(word string) {
	cur := trie
	l := len(word)
	for i := 0; i < l; i++ {
		c := word[i]
		if next, ok := cur.next[c]; ok {
			cur = next
		} else {
			return
		}
	}

	if cur != nil {
		cur.isWord = false
		cur.value = nil

		lastchar := cur.char

		if len(cur.next) == 0 {
			for !cur.isWord && cur.prev != nil {
				lastchar = cur.char
				cur = cur.prev
				if len(cur.next) > 1 {
					return
				}
			}
			delete(cur.next, lastchar)
		}
	}
}

// SearchMostPrefix Returns if the word is in the trie.
func (trie *hTrie) SearchDepth(iword trieWord) interface{} {
	cur := trie
	word := iword.GetWord()

	l := len(word)

	var result interface{}
	for i := 0; i < l; i++ {
		c := word[i]
		if next, ok := cur.next[c]; ok {
			cur = next
			if cur.isWord {
				result = cur.value
			} else {
				result = nil
			}
		} else {
			return result
		}
	}
	return result
}

// Match Returns  the word is in the trie.
func (trie *hTrie) Match(iword trieWord) interface{} {
	cur := trie
	word := iword.GetWord()

	l := len(word)
	for i := 0; i < l; i++ {
		c := word[i]
		if next, ok := cur.next[c]; ok {
			cur = next
		} else {
			return nil
		}
	}

	return cur.value
}

// StartsWith Returns if there is any word in the trie that starts with the given prefix. */
func (trie *hTrie) StartsWith(prefix string) bool {
	cur := trie
	l := len(prefix)
	for i := 0; i < l; i++ {
		c := prefix[i]
		if next, ok := cur.next[c]; ok {
			cur = next
		} else {
			return false
		}
	}
	return true
}

// 优先队列 所在的域

// parseQueue for Heap, Container GetScript
type parseQueue []*parseFunction

// parseFunction 优先执行参数
type parseFunction struct {
	ExecuteFunction func(u *CURL, soption string)
	ParamCURL       *CURL
	ParamData       string
	Priority        int
}

// Execute 执行 函数
func (pf *parseFunction) Execute() {
	pf.ExecuteFunction(pf.ParamCURL, pf.ParamData)
}

// Swap 实现sort.Iterface
func (nodes *parseQueue) Swap(i, j int) {
	ns := *nodes
	ns[i], ns[j] = ns[j], ns[i]
}

// Less Priority Want Less
func (nodes *parseQueue) Less(i, j int) bool {
	ns := *nodes
	return ns[i].Priority < ns[j].Priority
}

// Push 实现heap.Interface接口定义的额外方法
func (nodes *parseQueue) Push(exec interface{}) {
	*nodes = append(*nodes, exec.(*parseFunction))
}

// Pop 堆顶
func (nodes *parseQueue) Pop() (exec interface{}) {
	nlen := nodes.Len()
	exec = (*nodes)[nlen-1]    // 返回删除的元素
	*nodes = (*nodes)[:nlen-1] // [n:m]不包括下标为m的元素
	return exec
}

// Len len(nodes)
func (nodes *parseQueue) Len() int {
	return len(*nodes)
}

// pQueueExecute 优先函数队列
type pQueueExecute struct {
	nodes parseQueue
}

// newPQueueExecute CreateExpression A pQueueExecute
func newPQueueExecute() *pQueueExecute {
	pe := &pQueueExecute{}
	pe.nodes = make(parseQueue, 0)
	heap.Init(&pe.nodes)
	return pe
}

// Push CreateExpression A pQueueExecute
func (pqe *pQueueExecute) Push(exec *parseFunction) {
	heap.Push(&pqe.nodes, exec)
}

// Pop CreateExpression A pQueueExecute
func (pqe *pQueueExecute) Pop() *parseFunction {
	return heap.Pop(&pqe.nodes).(*parseFunction)
}

// Len CreateExpression A pQueueExecute
func (pqe *pQueueExecute) Len() int {
	return pqe.nodes.Len()
}

// func (pqe *pQueueExecute) String() string {
// 	content := ""
// 	for _, node := range pqe.nodes {
// 		content += strconv.Itoa(node.Prioty)
// 		content += " "
// 	}
// 	return content
// }
