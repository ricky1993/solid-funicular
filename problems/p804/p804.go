package p804

var (
	morse = []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.",
		"--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
)

func uniqueMorseRepresentations(words []string) int {
	set := map[string]struct{}{}
	for _, word := range words {
		s := ""
		for j := range word {
			s += morse[word[j]-'a']
		}
		set[s] = struct{}{}
	}
	return len(set)
}
