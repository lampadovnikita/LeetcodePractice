// https://leetcode.com/problems/implement-trie-prefix-tree/
// 208. Implement Trie (Prefix Tree)
//
// A trie (pronounced as "try") or prefix tree is a tree data structure used to efficiently store and retrieve keys in a dataset of strings. There are various applications of this data structure, such as autocomplete and spellchecker.
// Implement the Trie class:
// Trie() Initializes the trie object.
// void insert(String word) Inserts the string word into the trie.
// boolean search(String word) Returns true if the string word is in the trie (i.e., was inserted before), and false otherwise.
// boolean startsWith(String prefix) Returns true if there is a previously inserted string word that has the prefix prefix, and false otherwise.
//
// Example 1:
// Input
// ["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
// [[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
// Output
// [null, null, true, false, true, null, true]
// Explanation
// Trie trie = new Trie();
// trie.insert("apple");
// trie.search("apple");   // return True
// trie.search("app");     // return False
// trie.startsWith("app"); // return True
// trie.insert("app");
// trie.search("app");     // return True

package implementtrie

type Trie struct {
	childs map[rune]*Trie
	isWord bool
}

func Constructor() Trie {
	return Trie{
		childs: make(map[rune]*Trie),
		isWord: false,
	}
}

func (this *Trie) Insert(word string) {
	currTrie := this

	for _, currRune := range []rune(word) {
		_, ok := currTrie.childs[currRune]
		if ok == false {
			t := Constructor()
			currTrie.childs[currRune] = &t
		}

		currTrie = currTrie.childs[currRune]
	}

	currTrie.isWord = true
}

func (this *Trie) Search(word string) bool {
	currTrie := this

	for _, currRune := range []rune(word) {
		_, ok := currTrie.childs[currRune]
		if ok == false {
			return false
		}

		currTrie = currTrie.childs[currRune]
	}

	return currTrie.isWord
}

func (this *Trie) StartsWith(prefix string) bool {
	currTrie := this

	for _, currRune := range []rune(prefix) {
		_, ok := currTrie.childs[currRune]
		if ok == false {
			return false
		}

		currTrie = currTrie.childs[currRune]
	}

	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
