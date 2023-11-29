// https://leetcode.com/problems/design-add-and-search-words-data-structure/
// 211. Design Add and Search Words Data Structure
//
// Design a data structure that supports adding new words and finding if a string matches any previously added string.
// Implement the WordDictionary class:
// WordDictionary() Initializes the object.
// void addWord(word) Adds word to the data structure, it can be matched later.
// bool search(word) Returns true if there is any string in the data structure that matches word or false otherwise. word may contain dots '.' where dots can be matched with any letter.
//
// Example:
// Input
// ["WordDictionary","addWord","addWord","addWord","search","search","search","search"]
// [[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]
// Output
// [null,null,null,null,false,true,true,true]
// Explanation
// WordDictionary wordDictionary = new WordDictionary();
// wordDictionary.addWord("bad");
// wordDictionary.addWord("dad");
// wordDictionary.addWord("mad");
// wordDictionary.search("pad"); // return False
// wordDictionary.search("bad"); // return True
// wordDictionary.search(".ad"); // return True
// wordDictionary.search("b.."); // return True

package designaddandsearchwordsdatastructure

type WordDictionary struct {
	childs map[rune]*WordDictionary
	isWord bool
}

func Constructor() WordDictionary {
	return WordDictionary{
		childs: make(map[rune]*WordDictionary),
		isWord: false,
	}
}

func (this *WordDictionary) AddWord(word string) {
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

func (this *WordDictionary) Search(word string) bool {
	wordRunes := []rune(word)

	wdStack := make([]*WordDictionary, 1)
	wdStack[0] = this

	ridxStack := make([]int, 1)
	ridxStack[0] = 0

	for len(wdStack) != 0 {
		currTrie := wdStack[len(wdStack)-1]
		wdStack = wdStack[:len(wdStack)-1]

		currRidx := ridxStack[len(ridxStack)-1]
		ridxStack = ridxStack[:len(ridxStack)-1]

		if currRidx == len(word) {
			if currTrie.isWord == true {
				return true
			} else {
				continue
			}
		}

		if wordRunes[currRidx] == rune('.') {
			for _, val := range currTrie.childs {
				wdStack = append(wdStack, val)
				ridxStack = append(ridxStack, currRidx+1)
			}
		} else {
			val, ok := currTrie.childs[wordRunes[currRidx]]
			if ok == true {
				wdStack = append(wdStack, val)
				ridxStack = append(ridxStack, currRidx+1)
			}
		}
	}

	return false
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
