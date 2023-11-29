// https://leetcode.com/problems/map-sum-pairs/
// 677. Map Sum Pairs
//
// Design a map that allows you to do the following:
// Maps a string key to a given value.
// Returns the sum of the values that have a key with a prefix equal to a given string.
// Implement the MapSum class:
// MapSum() Initializes the MapSum object.
// void insert(String key, int val) Inserts the key-val pair into the map. If the key already existed, the original key-value pair will be overridden to the new one.
// int sum(string prefix) Returns the sum of all the pairs' value whose key starts with the prefix.
//
// Example 1:
// Input
// ["MapSum", "insert", "sum", "insert", "sum"]
// [[], ["apple", 3], ["ap"], ["app", 2], ["ap"]]
// Output
// [null, null, 3, null, 5]

// Explanation
// MapSum mapSum = new MapSum();
// mapSum.insert("apple", 3);
// mapSum.sum("ap");           // return 3 (apple = 3)
// mapSum.insert("app", 2);
// mapSum.sum("ap");           // return 5 (apple + app = 3 + 2 = 5)

package mapsumpairs

type MapSum struct {
	childsSum int
	val       int
	childs    map[rune]*MapSum
}

func Constructor() MapSum {
	return MapSum{
		childs: make(map[rune]*MapSum),
	}
}

func (this *MapSum) Insert(key string, val int) {
	currMapSum := this

	queue := make([]*MapSum, 0)

	for _, r := range []rune(key) {
		queue = append(queue, currMapSum)

		nextMapSum, ok := currMapSum.childs[r]
		if ok == false {
			ms := Constructor()
			nextMapSum = &ms
			currMapSum.childs[r] = nextMapSum
		}

		currMapSum = nextMapSum
	}

	sumDiff := val - currMapSum.val
	currMapSum.val = val

	for len(queue) != 0 {
		currMapSum = queue[0]
		queue = queue[1:]

		currMapSum.childsSum += sumDiff
	}
}

func (this *MapSum) Sum(prefix string) int {
	currMapSum := this
	var ok bool
	for _, r := range []rune(prefix) {
		currMapSum, ok = currMapSum.childs[r]
		if ok == false {
			return 0
		}
	}

	return currMapSum.childsSum + currMapSum.val
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
