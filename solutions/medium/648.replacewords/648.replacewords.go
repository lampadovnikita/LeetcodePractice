// https://leetcode.com/problems/replace-words/
// 648. Replace Words
//
// In English, we have a concept called root, which can be followed by some other word to form another longer word - let's call this word successor. For example, when the root "an" is followed by the successor word "other", we can form a new word "another".
// Given a dictionary consisting of many roots and a sentence consisting of words separated by spaces, replace all the successors in the sentence with the root forming it. If a successor can be replaced by more than one root, replace it with the root that has the shortest length.
// Return the sentence after the replacement.
//
// Example 1:
// Input: dictionary = ["cat","bat","rat"], sentence = "the cattle was rattled by the battery"
// Output: "the cat was rat by the bat"
//
// Example 2:
// Input: dictionary = ["a","b","c"], sentence = "aadsfasf absbs bbab cadsfafs"
// Output: "a a b c"

package replacewords

import "strings"

type Trie struct {
	childs map[rune]*Trie
	isWord bool
}

func Constructor() *Trie {
	return &Trie{
		childs: make(map[rune]*Trie),
		isWord: false,
	}
}

func (this *Trie) Insert(word string) {
	currTrie := this

	for _, currRune := range []rune(word) {
		_, ok := currTrie.childs[currRune]
		if ok == false {
			currTrie.childs[currRune] = Constructor()
		}

		currTrie = currTrie.childs[currRune]
	}

	currTrie.isWord = true
}

func (this *Trie) FindPrefix(word string) string {
	currTrie := this

	var ok bool

	for idx, r := range []rune(word) {
		currTrie, ok = currTrie.childs[r]
		if ok == false {
			return ""
		}

		if currTrie.isWord == true {
			return string([]rune(word)[:idx+1])
		}
	}

	return ""
}

func replaceWords(dictionary []string, sentence string) string {
	trie := Constructor()

	for i := range dictionary {
		trie.Insert(dictionary[i])
	}

	words := strings.Split(sentence, " ")
	prefixes := make([]string, len(words))

	for i := range words {
		prefix := trie.FindPrefix(words[i])
		if prefix == "" {
			prefix = words[i]
		}

		prefixes[i] = prefix
	}

	return strings.Join(prefixes, " ")
}
