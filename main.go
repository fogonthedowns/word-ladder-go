package main

// Time Complexity: O(LENGTH OF WORD^2 * Total number of words) O(M^2 N)
// Word Ladder
// J* = ["JZ", "JY"]

func ladderLength(beginWord string, endWord string, wordList []string) int {
	if beginWord == endWord {
		return 1
	}

	endWordExists := false

	// do pre-processing of wordList
	// e.g. map[*og:[cog dog hog]
	transformations := map[string][]string{}
	for _, w := range wordList {
		if w == endWord {
			endWordExists = true
		}
		for i := range w {
			t := replaceAtIndex(w, i)
			if _, ok := transformations[t]; !ok {
				transformations[t] = []string{w}
			} else {
				transformations[t] = append(transformations[t], w)
			}
		}
	}

	// speed up if impossible
	if !endWordExists {
		return 0
	}

	type item struct {
		w string
		l int
	}

	// create queue
	// push in the first word
	// The 1 represents the level number of a node
	queue := []item{
		{w: beginWord, l: 1},
	}

	for q := 0; q < len(queue); q++ {
		for i := range queue[q].w {
			// replaces char with * at index
			t := replaceAtIndex(queue[q].w, i)

			// if the transformation exists
			if _, ok := transformations[t]; ok {

				// range over the list of words matching transformation
				for _, w := range transformations[t] {
					// reached the end
					if w == endWord {
						return queue[q].l + 1
					}

					// append item to the queue
					queue = append(queue, item{w: w, l: queue[q].l + 1})

					// clean up transformations map
					delete(transformations, t)
				}
			}
		}
	}

	return 0
}

// replaces char with * at an index
// eg "cog", 1 -> "c*g"
func replaceAtIndex(in string, i int) string {
	out := []rune(in)
	out[i] = '*'
	return string(out)
}
