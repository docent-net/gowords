package words

// type wordInfo struct {
// 	word         string
// 	partOfSpeech string
// 	synonyms     []string
// 	antonyms     []string
// 	derivation   []string
// }

type wordResults struct {
	word string
	// results []wordInfo
}

func FetchWordInfo(word string) wordResults {
	var ww = wordResults{
		word: word,
		// results:
	}

	return ww
}
