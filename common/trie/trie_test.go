package trie

import (
	"testing"
)

func TestTrie_AddWord(t *testing.T) {
	errMsg := "cannot add empty word to trie"
	testCases := []struct {
		input        string
		expectedWord string
		shouldError  bool
	}{
		{"added", "added", false},
		{"games night", "games night", false},
		{"0813 022 6951", "0813 022 6951", false},
		{"", "", true},
	}

	tr := NewTrie()
	for _, testCase := range testCases {
		got, err := tr.AddWord(testCase.input)
		if testCase.shouldError {
			if err != nil && err.Error() != errMsg {
				t.Errorf("expected error message to be %q", errMsg)
			}
			continue
		}
		if testCase.expectedWord != got.Word {
			t.Errorf("expected: %s, got %s", testCase.expectedWord, got.Word)
		}
		if !got.IsEnd {
			t.Errorf("word is supposed to have IsEnd marked")
		}
	}
}

func TestTrie_Find(t *testing.T) {
	toAdd := []string{"games night", "0813 022 6951"}

	testCases := []struct {
		toFind      string
		isEnd       bool
		shouldError bool
		errMsg      string
	}{
		{"games night", true, false, ""},
		{"0813 022 6951", true, false, ""},
		{"", false, true, "cannot search empty word"},
		{"0813", false, false, ""},
		{"08130", false, true, "word not in trie"},
	}

	tr := NewTrie()
	for _, val := range toAdd {
		_, _ = tr.AddWord(val)
	}

	for _, testCase := range testCases {
		got, err := tr.Find(testCase.toFind)
		if testCase.shouldError {
			if err == nil {
				t.Errorf("expected err, got nil")
			} else if testCase.errMsg != err.Error() {
				t.Errorf("expected err message: %v, got %v", testCase.errMsg, err)
			}
			continue
		}

		if got.IsEnd != testCase.isEnd {
			t.Errorf("expected isEnd to be %v, got %v", testCase.isEnd, got.IsEnd)
		}
	}
}

func TestTrie_FindWord(t *testing.T) {
	toAdd := []string{"games night", "0813 022 6951"}

	testCases := []struct {
		toFind      string
		isEnd       bool
		shouldError bool
		errMsg      string
	}{
		{"games night", true, false, ""},
		{"0813 022 6951", true, false, ""},
		{"", false, true, "cannot search empty word"},
		{"0813", false, true, "word not in trie"},
		{"08130", false, true, "word not in trie"},
		{"0813 0", false, true, "word not in trie"},
	}

	tr := NewTrie()
	for _, val := range toAdd {
		_, _ = tr.AddWord(val)
	}

	for _, testCase := range testCases {
		got, err := tr.FindWord(testCase.toFind)
		if testCase.shouldError {
			if err == nil {
				t.Errorf("expected err, got nil")
			} else if testCase.errMsg != err.Error() {
				t.Errorf("expected err message: %v, got %v", testCase.errMsg, err)
			}
			continue
		}

		if got.IsEnd != testCase.isEnd {
			t.Errorf("expected isEnd to be %v, got %v", testCase.isEnd, got.IsEnd)
		}
	}
}

func TestTrie_DeleteWord(t *testing.T) {
	toAdd := []string{"games night", "0813 022 6951"}
	errMsg := "word not in trie"

	testCases := []struct {
		toDel string
	}{
		{"games night"},
		{"0813 022 6951"},
	}

	tr := NewTrie()
	for _, val := range toAdd {
		_, _ = tr.AddWord(val)
	}

	for _, testCase := range testCases {
		// find the node representing the last character of the word
		node, err := tr.FindWord(testCase.toDel)
		if err != nil {
			t.Errorf("expected error to be nil, got %v", err)
		}

		// delete the node from the trie
		if err = tr.DeleteWord(node); err != nil {
			t.Errorf("expected error to be nil, got %v", err)
		}

		// check that the word is not a complete word in the trie
		node, err = tr.FindWord(testCase.toDel)
		if err == nil {
			t.Errorf("expected error to be %v, got %v", errMsg, err)
		}
		if err.Error() != errMsg {
			t.Errorf("expected error message to be %s, got %s", errMsg, err.Error())
		}
		if node != nil {
			t.Errorf("expected node to be nil, got %v", node)
		}

		// check that the word is still present as a substring
		node, err = tr.Find(testCase.toDel)
		if err != nil {
			t.Errorf("expected error to be nil, got %v", err)
		}
	}
}

func TestTrie_UpdateWord(t *testing.T) {
	toAdd := []string{"games night", "0813 022 6951"}
	errMsg := "word not in trie"

	testCases := []struct {
		toUpdate string
		newValue string
	}{
		{"games night", "gamer night"},
		{"0813 022 6951", "0803 951 8636"},
	}

	tr := NewTrie()
	for _, val := range toAdd {
		_, _ = tr.AddWord(val)
	}

	for _, testCase := range testCases {
		// get the node for the word
		node, err := tr.FindWord(testCase.toUpdate)
		if err != nil {
			t.Errorf("expected error to be nil, got %v", err)
		}

		// update the node with the new value
		node, err = tr.UpdateWord(node, testCase.newValue)
		if err != nil {
			t.Errorf("expected error to be nil, got %v", err)
		}

		// check that the previous word cannot be found anymore
		node, err = tr.FindWord(testCase.toUpdate)
		if err == nil {
			t.Errorf("expected error to be %v, got %v", errMsg, err)
		}
		if err.Error() != errMsg {
			t.Errorf("expected error message to be %s, got %s", errMsg, err.Error())
		}
		if node != nil {
			t.Errorf("expected node to be nil, got %v", node)
		}

		// check that the new word can be retrieved
		node, err = tr.FindWord(testCase.newValue)
		if err != nil {
			t.Errorf("expected error to be nil, got %v", err)
		}

		if testCase.newValue != node.Word {
			t.Errorf("expected node word to be %v, got %v", testCase.newValue, node.Word)
		}
	}
}
