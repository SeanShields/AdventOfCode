package main

import (
	"fmt"
)

func main() {
	input := []string{"oiwcdpbseqgxryfmlpktnupvza", "oiwddpbsuqhxryfmlgkznujvza", "ziwcdpbsechxrvfmlgktnujvza", "oiwcgpbseqhxryfmmgktnhjvza", "owwcdpbseqhxryfmlgktnqjvze", "oiscdkbseqhxrffmlgktnujvza", "oiwcdibseqoxrnfmlgktnujvza", "oiwcdpbsejhxryfmlektnujiza", "oewcdpbsephxryfmlgkwnujvza", "riwcdpbseqhxryfmlgktnujzaa", "omwcdpbseqwxryfmlgktnujvqa", "oiwcdqqneqhxryfmlgktnujvza", "oawcdvaseqhxryfmlgktnujvza", "oiwcdabseqhxcyfmlkktnujvza", "oiwcdpbseqhxryfmlrktrufvza", "oiwcdpbseqhxdyfmlgqtnujkza", "oiwcdpbseqhxrmfolgktnujvzy", "oiwcdpeseqhxnyfmlgktnejvza", "oiwcdpbseqhxrynmlaktlujvza", "oiwcdpbseqixryfmlektncjvza", "liwtdpbseqhxryfylgktnujvza", "ouwcdpbszqhxryfmlgktnajvza", "oiwzdpbseqhxryfmngktnujvga", "wiwcfpbseqhxryfmlgktnuhvza", "oiwcdpbselhfryfmlrktnujvza", "oywcdpbveqhxryfmlgktnujdza", "oiwcdpbsiqhxryfmqiktnujvza", "obwcdhbseqhxryfmlgktnujvqa", "oitcdpbseqhfryfmlyktnujvza", "oiwcdpbseqhxryfmlnutnujqza", "oiwcdpbseqhxnyfmlhktnujtza", "oiwcdpbseqhbryfmlgkunuwvza", "oiwcopbseqhiryfmlgktnkjvza", "oiwcdpsseqhxryfklrktnujvza", "oiwcdpsrsqhxryfmlgktnujvza", "oiwcdpbsyqrxryfmlgktnujvzc", "ojwcepbseqhxryfmlgktnujvfa", "oiwcdpbseqhxrlfmlgvtnujvzr", "oiycdpbsethsryfmlgktnujvza", "eiwcdpbseqwxryfmlgktnujcza", "oiocdpbseqhxryfxlgktaujvza", "qiwydpbseqhpryfmlgktnujvza", "ziwcdpbseqhxryfmlgktsuuvza", "oiwcdpbseqheryfmygxtnujvza", "oiwidpbseqhxryfulgktnujvzm", "oiscdpdseqhxryfmlgktnujvya", "oiwmypbseqhxryfmlgktnuxvza", "oizcwbbseqhxryfmlgktnujvza", "oiwcdpbseqpxryfmlgxfnujvza", "oiwpdpbscqhxryxmlgktnujvza", "oiwcdpbseqhxrifrlgkxnujvza", "oiwcdpbsrqhxrifmlgktnzjvza", "tiwcdpbseqhxryfmegkvjujvza", "oiwcddbseqhxryfingktnujvza", "oiwcdpbsiqhiryfmlgktnucvza", "oiwcipbseqhxrkfmlgktnuvvza", "kiwcdpbseqhxryfmlbkonujvza", "qiwcdhbsedhxryfmlgktnujvza", "siwcdpbseqhxryfmsgktnujvua", "oqwcdpbseqhxryfmlyktndjvza", "oiwcqnbseehxryfmlgktnujvza", "oiwcdybseqhxryfmlgbtnujvia", "oiwcdpbsejhxryfdlgktngjvza", "oawcdpbseqhxrbfmlkktnujvza", "oilcdpbseqhhrjfmlgktnujvza", "oibcdpbseqhxryfmngktnucvza", "niwcdpbseqhxryfmlgkuaujvza", "oiwcdpbseqhxryfmqgmtnujvha", "oiwcdpbseqhcryfxlgktnzjvza", "oiwcdpaseqhxryfmqgktnujvzl", "oiwcdpbseqhxjyfmlgkonujvzx", "oiwmdzbseqlxryfmlgktnujvza", "oiwhdpbsexhxryfmlgktnujvzw", "oiwctpbseqhxryfmlgktnummza", "oiwcdpbseqoxrydmlgktnujvoa", "oiwkdpvseqhxeyfmlgktnujvza", "oixcdpbsemhxryfmlgctnujvza", "oimcdpbseqhxryfmlgktnujvlr", "oiwcdpbseehxrywmlgktnejvza", "oiwcdpbseqoxhyfmlgktnujyza", "oiwcdpbsethxryfmlgktnrjvxa", "oiwcdpbxeqhxryfmlgktnrjvzb", "ogeadpbseqhxryfmlgktnujvza", "eiwcdpbseqhxryfmlgktnvuvza", "oiwcdpbseqhxryfmlgktnujaxv", "siwcdpbsuqhxryfmlgktnuavza", "oixcdpbseqhxryfmlgatnujvzy", "oiwcdpbzeghmryfmlgktnujvza", "oiwcdpbieqhxryfmlgktyujvzr", "oiwcdpbseqhxeyfhlgktngjvza", "oiwcdpbseqhjoyrmlgktnujvza", "iiwcdpbseqhxryfmwhktnujvza", "oixcdpbsiqhxryfmagktnujvza", "oiwcdpfljqhxryfmlgktnujvza", "oivcdpbseqhxrqfmlgktnujvca", "ovwcdpbseqhxzyfmlgkenujvza", "oiwxdpgseqhxryfmlgktnhjvza", "oibcdpbbeohxryfmlgktnujvza", "oiwcrpbseqhxrygmlgwtnujvza", "jiwcdpbseqkxryfmlgntnujvza", "oiwcdbbseqhxrywmlgktnujvra", "oiwcdpbteqyxoyfmlgktnujvza", "oiwcdjbsvqvxryfmlgktnujvza", "obwcdukseqhxryfmlgktnujvza", "oifcdpdseqhxryfmlgktnujsza", "oiwcdpbseqhxryfalgktnujyda", "oiwcwpbseqhxrnfmkgktnujvza", "oswcdpbsethcryfmlgktnujvza", "oiwcdpbieqhxryfmlgktnuoiza", "oiwcdibsehhxryfmzgktnujvza", "oisjdpbseqhxryfmfgktnujvza", "oiwcjpbseqkxqyfmlgktnujvza", "obwcdpbshqhgryfmlgktnujvza", "oiwcspbseqhxryxmlgktnujvzl", "oiwcdvbswqhxryfmlgklnujvza", "oiwcdhuseqhxryfmlgdtnujvza", "oiwcdpbkeqdxryfmlgktnujvzv", "oiwcdpzseqhxcyfmlgksnujvza", "oiwcdpbseqhxryfmbkkvnujvza", "qiwcdpbseqhxrnfmlgktnujvha", "okwcdpbseqhxryfmdgktgujvza", "oiwcdpbkeqhxryfmldktnujvzu", "oiwctpxseqhxgyfmlgktnujvza", "oiwcdpbseqhxrykmlgktnujita", "oiwcdpbseqhxryfmldktyujnza", "oiwcdpbszqhxrjfmlgktnuqvza", "oiwcdpbeeqhxrykmlgktnujrza", "oiwcvpbseqhxryhmlgwtnujvza", "oiwcdpbpeehxryfmlgktnujvzz", "oiwcdbbsxqhxyyfmlgktnujvza", "oiwkdpbseqhxryfplgktnujeza", "opwcdpbseqhxdyfmlgctnujvza", "oowcdpbseqhnryfmlgktnujvga", "oawzdibseqhxryfmlgktnujvza", "oiwcdpbfeqzxrjfmlgktnujvza", "oiwcdpbseqaxryfmlgkonulvza", "oiacdpbseqvxryfmlgktvujvza", "oiwudpbseqhxryfwlgktnujvka", "oiwcdpbjeqhxryfymgktnujvza", "oiwcdpbweqhxrynmlgktnujaza", "oiwcdpbseqdxryfclgvtnujvza", "oiwcdppseqhxryfmlgmtzujvza", "oiwcdpbseqhxrhfelektnujvza", "kiwcdpbsnqhxryfmegktnujvza", "oiwcdpbseqpxryfmlgzwnujvza", "eiwcdpbseqnxryfplgktnujvza", "oiwcdbbseqhxryfmlmktnujvha", "oiwcdpbseqhxryfmlgktjhjvka", "oiwcdpbseqhxnyfylgktnujvzs", "oiwcdpbbxqhxryfmzgktnujvza", "opwcdpbseqhfryfmlgktnujzza", "oiwcdpbsjqpxryfmtgktnujvza", "oiwcdpbseqhyqbfmlgktnujvza", "oxwcdpbseqhxrffmlgktiujvza", "oiwcdpbseqhxgyfmlgytnujvzq", "oiwidpbseqhxryfmlgxtnujvzd", "oiwcdpbshqhxryzmlpktnujvza", "oifcdpbseqbxryfmlgktdujvza", "biwcdzbseqhxtyfmlgktnujvza", "oiwcdpbswqhxryfmlgutnujvca", "xiwcdpbseqhxryxmlnktnujvza", "oiwcdpzseqhxryfrlgktdujvza", "oiwudpbseqhxryfmlgkqnurvza", "oiwqdpbseihiryfmlgktnujvza", "iiwjdpbseqhxryamlgktnujvza", "oiwcdplseqhqryfmmgktnujvza", "oiwcdppseqhxrmfmlgklnujvza", "oiwcdobseqhxryfmmgkttujvza", "odwcdpbseqhxryfmlgktnujvyk", "oiwcdpkseqhxrhfmlgktntjvza", "oiocdpbseqhxryfmlgjknujvza", "oiicdpbieqhxryfmlgksnujvza", "oiwcdpbseqhxryemlgktnujdla", "oiwcdpbseqdxryfmlgktsujvzt", "oiwcdcksnqhxryfmlgktnujvza", "oowcdpbseqhxryfmlgsfnujvza", "oawcdpbseqhxryfmljktnuevza", "oiwcdpbsaqhxrffmzgktnujvza", "oiwcipbseqhcryfmlgktnujvga", "oiwcdpbsewhxrbfmlgktnuuvza", "oiwcdpbsewhxryfmlgkunujvzc", "oiwcdpbseqhxryfmlgktiujkga", "jiwcdpbseqhxrlfmlgktnurvza", "tiwcdpbseqoxryfmliktnujvza", "oiwcdpbsenhxryfmlgkskujvza", "oiwcdpbseqhxvyfmlhktoujvza", "oiwcdpbseqhxeyfmlgmtnunvza", "oiwcdpbseqhxdyfmloktnujvzu", "oiwcdpbseqgxryfmlgkynejvza", "oudcdpbseqhxryfmlgktmujvza", "oiwcdpbseqhxryfmvgktnucvzc", "oiwcdpbseqhxrysalgwtnujvza", "odwodpbseqhgryfmlgktnujvza", "oiwcdpbseqheryzmlgktnujgza", "oiwcdpbseqhxryfalgwtnujvba", "oiwcdpbseqhxryfmlgtdnuhvza", "oiocdpbseqhxrefulgktnujvza", "kiwcdpbseqhxrywolgktnujvza", "niwcdpbseksxryfmlgktnujvza", "oiwcdybseqexryfalgktnujvza", "oiwcdpbbeqhxryamlgktnujpza", "oiecdqbseqhxryfmlgktnujcza", "oiwcdpbsqqhxlyfmlpktnujvza", "oiwcdpbsaqheryfmlgktnujlza", "oiecdpbseqhxryfmlgkknujvzz", "oiwcapbsdqhxryfmlgktvujvza", "ojwcdxbseqhxryfmlgktrujvza", "oiwhdpbseqvxrzfmlgktnujvza", "oiwcdppseqhtryfmlgktnujvzs", "oikcdpbsfqhxryfmdgktnujvza", "owwczpbseqhxryfilgktnujvza", "oifwdpbseqhxryfmlgktnujfza", "oowcdpbseqhxrprmlgktnujvza", "oiwcapbseqhxryfmjgktnujvze", "oiwcdpaseqhdrybmlgktnujvza", "tiwcdpbseqhxryfmlgkvjujvza", "xiwcdpbseqhoryfmlgktnujvqa", "eiwrdpbsyqhxryfmlgktnujvza", "oiwcdpbseqhxranmlgktnujvzt", "oiwcdpbseqhxryfqlgktnudaza", "oiwcdpbsvqhxrywmlgktnsjvza", "oewcdpbseqhxryfmlgkunujvma", "oiwcdpbseqhjrywmlgktnujvzb", "omwcdpbseqhxryfmlgktwujsza", "oiwcdpbyxqhxryfmlgrtnujvza", "oiwidpbseqhxryfhlgktnunvza", "oqwcdpbweqhxrybmlgktnujvza", "oiwcdqbseqhxryfrlgktnujfza", "oiacdpbseqhdryfmlgktaujvza", "oiwcdpbstqhxmyfmlgktyujvza", "oiwcdpbseqhxeyfclgktjujvza", "wiwcdpeseqhxryfmlgktnujvzx", "viwcdpbseqhxryfmlgvtvujvza", "oircdpbseqhxcyfmlgktnujvma", "miwcdpbseqtwryfmlgktnujvza", "oiwcppbseqhxcyfmlgxtnujvza", "giwcrpbseqhxryfmlgktnudvza", "onwcvpbseqhxryfmlgktnujqza", "oiwcipbseqhxryfmlgitnuqvza", "oiwcdpbseqhxryjmlgkonufvza", "oiwnwpbseqhxtyfmlgktnujvza", "oiwcypbseqhxryfmlgetnujvzv", "oiwcdpbseqhxryqmljktnkjvza", "olwcdpbseqhxryfmlgkenujvba", "biwcdpbseqwxrywmlgktnujvza", "oiwcdpbsevhmryjmlgktnujvza", "oiwcdpbseqhxryfmlzktnkjvzv", "oiwudpbseqhxrefmlgktnujvia", "oiicdpbseqhxryfdloktnujvza", "oihcjpbsxqhxryfmlgktnujvza"}
	// test := []string{"abcdef", "bababc", "abbcde", "abcccd", "aabcdd", "abcdee", "ababab"} // expected: 12
	twos := 0
	threes := 0

	for _, boxid := range input {
		hastwo := false
		hasthree := false
		var checkedChars []rune

		for _, char := range boxid {
			if existsInArray(checkedChars, char) {
				continue
			}
			checkedChars = append(checkedChars, char)

			if !hastwo && existsInString(boxid, char, 2) {
				twos++
				hastwo = true
			}
			if !hasthree && existsInString(boxid, char, 3) {
				threes++
				hasthree = true
			}
		}
	}

	fmt.Printf("twos: %d\n", twos)
	fmt.Printf("threes: %d\n", threes)
	fmt.Printf("result: %d\n", twos*threes)
}

func existsInString(s string, e rune, count int) bool {
	for _, a := range s {
		if a == e {
			count--
		}
	}
	return count == 0
}

func existsInArray(s []rune, e rune) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
