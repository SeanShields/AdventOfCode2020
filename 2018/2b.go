package main

import (
	"fmt"
)

func main() {
	input := []string{"oiwcdpbseqgxryfmlpktnupvza", "oiwddpbsuqhxryfmlgkznujvza", "ziwcdpbsechxrvfmlgktnujvza", "oiwcgpbseqhxryfmmgktnhjvza", "owwcdpbseqhxryfmlgktnqjvze", "oiscdkbseqhxrffmlgktnujvza", "oiwcdibseqoxrnfmlgktnujvza", "oiwcdpbsejhxryfmlektnujiza", "oewcdpbsephxryfmlgkwnujvza", "riwcdpbseqhxryfmlgktnujzaa", "omwcdpbseqwxryfmlgktnujvqa", "oiwcdqqneqhxryfmlgktnujvza", "oawcdvaseqhxryfmlgktnujvza", "oiwcdabseqhxcyfmlkktnujvza", "oiwcdpbseqhxryfmlrktrufvza", "oiwcdpbseqhxdyfmlgqtnujkza", "oiwcdpbseqhxrmfolgktnujvzy", "oiwcdpeseqhxnyfmlgktnejvza", "oiwcdpbseqhxrynmlaktlujvza", "oiwcdpbseqixryfmlektncjvza", "liwtdpbseqhxryfylgktnujvza", "ouwcdpbszqhxryfmlgktnajvza", "oiwzdpbseqhxryfmngktnujvga", "wiwcfpbseqhxryfmlgktnuhvza", "oiwcdpbselhfryfmlrktnujvza", "oywcdpbveqhxryfmlgktnujdza", "oiwcdpbsiqhxryfmqiktnujvza", "obwcdhbseqhxryfmlgktnujvqa", "oitcdpbseqhfryfmlyktnujvza", "oiwcdpbseqhxryfmlnutnujqza", "oiwcdpbseqhxnyfmlhktnujtza", "oiwcdpbseqhbryfmlgkunuwvza", "oiwcopbseqhiryfmlgktnkjvza", "oiwcdpsseqhxryfklrktnujvza", "oiwcdpsrsqhxryfmlgktnujvza", "oiwcdpbsyqrxryfmlgktnujvzc", "ojwcepbseqhxryfmlgktnujvfa", "oiwcdpbseqhxrlfmlgvtnujvzr", "oiycdpbsethsryfmlgktnujvza", "eiwcdpbseqwxryfmlgktnujcza", "oiocdpbseqhxryfxlgktaujvza", "qiwydpbseqhpryfmlgktnujvza", "ziwcdpbseqhxryfmlgktsuuvza", "oiwcdpbseqheryfmygxtnujvza", "oiwidpbseqhxryfulgktnujvzm", "oiscdpdseqhxryfmlgktnujvya", "oiwmypbseqhxryfmlgktnuxvza", "oizcwbbseqhxryfmlgktnujvza", "oiwcdpbseqpxryfmlgxfnujvza", "oiwpdpbscqhxryxmlgktnujvza", "oiwcdpbseqhxrifrlgkxnujvza", "oiwcdpbsrqhxrifmlgktnzjvza", "tiwcdpbseqhxryfmegkvjujvza", "oiwcddbseqhxryfingktnujvza", "oiwcdpbsiqhiryfmlgktnucvza", "oiwcipbseqhxrkfmlgktnuvvza", "kiwcdpbseqhxryfmlbkonujvza", "qiwcdhbsedhxryfmlgktnujvza", "siwcdpbseqhxryfmsgktnujvua", "oqwcdpbseqhxryfmlyktndjvza", "oiwcqnbseehxryfmlgktnujvza", "oiwcdybseqhxryfmlgbtnujvia", "oiwcdpbsejhxryfdlgktngjvza", "oawcdpbseqhxrbfmlkktnujvza", "oilcdpbseqhhrjfmlgktnujvza", "oibcdpbseqhxryfmngktnucvza", "niwcdpbseqhxryfmlgkuaujvza", "oiwcdpbseqhxryfmqgmtnujvha", "oiwcdpbseqhcryfxlgktnzjvza", "oiwcdpaseqhxryfmqgktnujvzl", "oiwcdpbseqhxjyfmlgkonujvzx", "oiwmdzbseqlxryfmlgktnujvza", "oiwhdpbsexhxryfmlgktnujvzw", "oiwctpbseqhxryfmlgktnummza", "oiwcdpbseqoxrydmlgktnujvoa", "oiwkdpvseqhxeyfmlgktnujvza", "oixcdpbsemhxryfmlgctnujvza", "oimcdpbseqhxryfmlgktnujvlr", "oiwcdpbseehxrywmlgktnejvza", "oiwcdpbseqoxhyfmlgktnujyza", "oiwcdpbsethxryfmlgktnrjvxa", "oiwcdpbxeqhxryfmlgktnrjvzb", "ogeadpbseqhxryfmlgktnujvza", "eiwcdpbseqhxryfmlgktnvuvza", "oiwcdpbseqhxryfmlgktnujaxv", "siwcdpbsuqhxryfmlgktnuavza", "oixcdpbseqhxryfmlgatnujvzy", "oiwcdpbzeghmryfmlgktnujvza", "oiwcdpbieqhxryfmlgktyujvzr", "oiwcdpbseqhxeyfhlgktngjvza", "oiwcdpbseqhjoyrmlgktnujvza", "iiwcdpbseqhxryfmwhktnujvza", "oixcdpbsiqhxryfmagktnujvza", "oiwcdpfljqhxryfmlgktnujvza", "oivcdpbseqhxrqfmlgktnujvca", "ovwcdpbseqhxzyfmlgkenujvza", "oiwxdpgseqhxryfmlgktnhjvza", "oibcdpbbeohxryfmlgktnujvza", "oiwcrpbseqhxrygmlgwtnujvza", "jiwcdpbseqkxryfmlgntnujvza", "oiwcdbbseqhxrywmlgktnujvra", "oiwcdpbteqyxoyfmlgktnujvza", "oiwcdjbsvqvxryfmlgktnujvza", "obwcdukseqhxryfmlgktnujvza", "oifcdpdseqhxryfmlgktnujsza", "oiwcdpbseqhxryfalgktnujyda", "oiwcwpbseqhxrnfmkgktnujvza", "oswcdpbsethcryfmlgktnujvza", "oiwcdpbieqhxryfmlgktnuoiza", "oiwcdibsehhxryfmzgktnujvza", "oisjdpbseqhxryfmfgktnujvza", "oiwcjpbseqkxqyfmlgktnujvza", "obwcdpbshqhgryfmlgktnujvza", "oiwcspbseqhxryxmlgktnujvzl", "oiwcdvbswqhxryfmlgklnujvza", "oiwcdhuseqhxryfmlgdtnujvza", "oiwcdpbkeqdxryfmlgktnujvzv", "oiwcdpzseqhxcyfmlgksnujvza", "oiwcdpbseqhxryfmbkkvnujvza", "qiwcdpbseqhxrnfmlgktnujvha", "okwcdpbseqhxryfmdgktgujvza", "oiwcdpbkeqhxryfmldktnujvzu", "oiwctpxseqhxgyfmlgktnujvza", "oiwcdpbseqhxrykmlgktnujita", "oiwcdpbseqhxryfmldktyujnza", "oiwcdpbszqhxrjfmlgktnuqvza", "oiwcdpbeeqhxrykmlgktnujrza", "oiwcvpbseqhxryhmlgwtnujvza", "oiwcdpbpeehxryfmlgktnujvzz", "oiwcdbbsxqhxyyfmlgktnujvza", "oiwkdpbseqhxryfplgktnujeza", "opwcdpbseqhxdyfmlgctnujvza", "oowcdpbseqhnryfmlgktnujvga", "oawzdibseqhxryfmlgktnujvza", "oiwcdpbfeqzxrjfmlgktnujvza", "oiwcdpbseqaxryfmlgkonulvza", "oiacdpbseqvxryfmlgktvujvza", "oiwudpbseqhxryfwlgktnujvka", "oiwcdpbjeqhxryfymgktnujvza", "oiwcdpbweqhxrynmlgktnujaza", "oiwcdpbseqdxryfclgvtnujvza", "oiwcdppseqhxryfmlgmtzujvza", "oiwcdpbseqhxrhfelektnujvza", "kiwcdpbsnqhxryfmegktnujvza", "oiwcdpbseqpxryfmlgzwnujvza", "eiwcdpbseqnxryfplgktnujvza", "oiwcdbbseqhxryfmlmktnujvha", "oiwcdpbseqhxryfmlgktjhjvka", "oiwcdpbseqhxnyfylgktnujvzs", "oiwcdpbbxqhxryfmzgktnujvza", "opwcdpbseqhfryfmlgktnujzza", "oiwcdpbsjqpxryfmtgktnujvza", "oiwcdpbseqhyqbfmlgktnujvza", "oxwcdpbseqhxrffmlgktiujvza", "oiwcdpbseqhxgyfmlgytnujvzq", "oiwidpbseqhxryfmlgxtnujvzd", "oiwcdpbshqhxryzmlpktnujvza", "oifcdpbseqbxryfmlgktdujvza", "biwcdzbseqhxtyfmlgktnujvza", "oiwcdpbswqhxryfmlgutnujvca", "xiwcdpbseqhxryxmlnktnujvza", "oiwcdpzseqhxryfrlgktdujvza", "oiwudpbseqhxryfmlgkqnurvza", "oiwqdpbseihiryfmlgktnujvza", "iiwjdpbseqhxryamlgktnujvza", "oiwcdplseqhqryfmmgktnujvza", "oiwcdppseqhxrmfmlgklnujvza", "oiwcdobseqhxryfmmgkttujvza", "odwcdpbseqhxryfmlgktnujvyk", "oiwcdpkseqhxrhfmlgktntjvza", "oiocdpbseqhxryfmlgjknujvza", "oiicdpbieqhxryfmlgksnujvza", "oiwcdpbseqhxryemlgktnujdla", "oiwcdpbseqdxryfmlgktsujvzt", "oiwcdcksnqhxryfmlgktnujvza", "oowcdpbseqhxryfmlgsfnujvza", "oawcdpbseqhxryfmljktnuevza", "oiwcdpbsaqhxrffmzgktnujvza", "oiwcipbseqhcryfmlgktnujvga", "oiwcdpbsewhxrbfmlgktnuuvza", "oiwcdpbsewhxryfmlgkunujvzc", "oiwcdpbseqhxryfmlgktiujkga", "jiwcdpbseqhxrlfmlgktnurvza", "tiwcdpbseqoxryfmliktnujvza", "oiwcdpbsenhxryfmlgkskujvza", "oiwcdpbseqhxvyfmlhktoujvza", "oiwcdpbseqhxeyfmlgmtnunvza", "oiwcdpbseqhxdyfmloktnujvzu", "oiwcdpbseqgxryfmlgkynejvza", "oudcdpbseqhxryfmlgktmujvza", "oiwcdpbseqhxryfmvgktnucvzc", "oiwcdpbseqhxrysalgwtnujvza", "odwodpbseqhgryfmlgktnujvza", "oiwcdpbseqheryzmlgktnujgza", "oiwcdpbseqhxryfalgwtnujvba", "oiwcdpbseqhxryfmlgtdnuhvza", "oiocdpbseqhxrefulgktnujvza", "kiwcdpbseqhxrywolgktnujvza", "niwcdpbseksxryfmlgktnujvza", "oiwcdybseqexryfalgktnujvza", "oiwcdpbbeqhxryamlgktnujpza", "oiecdqbseqhxryfmlgktnujcza", "oiwcdpbsqqhxlyfmlpktnujvza", "oiwcdpbsaqheryfmlgktnujlza", "oiecdpbseqhxryfmlgkknujvzz", "oiwcapbsdqhxryfmlgktvujvza", "ojwcdxbseqhxryfmlgktrujvza", "oiwhdpbseqvxrzfmlgktnujvza", "oiwcdppseqhtryfmlgktnujvzs", "oikcdpbsfqhxryfmdgktnujvza", "owwczpbseqhxryfilgktnujvza", "oifwdpbseqhxryfmlgktnujfza", "oowcdpbseqhxrprmlgktnujvza", "oiwcapbseqhxryfmjgktnujvze", "oiwcdpaseqhdrybmlgktnujvza", "tiwcdpbseqhxryfmlgkvjujvza", "xiwcdpbseqhoryfmlgktnujvqa", "eiwrdpbsyqhxryfmlgktnujvza", "oiwcdpbseqhxranmlgktnujvzt", "oiwcdpbseqhxryfqlgktnudaza", "oiwcdpbsvqhxrywmlgktnsjvza", "oewcdpbseqhxryfmlgkunujvma", "oiwcdpbseqhjrywmlgktnujvzb", "omwcdpbseqhxryfmlgktwujsza", "oiwcdpbyxqhxryfmlgrtnujvza", "oiwidpbseqhxryfhlgktnunvza", "oqwcdpbweqhxrybmlgktnujvza", "oiwcdqbseqhxryfrlgktnujfza", "oiacdpbseqhdryfmlgktaujvza", "oiwcdpbstqhxmyfmlgktyujvza", "oiwcdpbseqhxeyfclgktjujvza", "wiwcdpeseqhxryfmlgktnujvzx", "viwcdpbseqhxryfmlgvtvujvza", "oircdpbseqhxcyfmlgktnujvma", "miwcdpbseqtwryfmlgktnujvza", "oiwcppbseqhxcyfmlgxtnujvza", "giwcrpbseqhxryfmlgktnudvza", "onwcvpbseqhxryfmlgktnujqza", "oiwcipbseqhxryfmlgitnuqvza", "oiwcdpbseqhxryjmlgkonufvza", "oiwnwpbseqhxtyfmlgktnujvza", "oiwcypbseqhxryfmlgetnujvzv", "oiwcdpbseqhxryqmljktnkjvza", "olwcdpbseqhxryfmlgkenujvba", "biwcdpbseqwxrywmlgktnujvza", "oiwcdpbsevhmryjmlgktnujvza", "oiwcdpbseqhxryfmlzktnkjvzv", "oiwudpbseqhxrefmlgktnujvia", "oiicdpbseqhxryfdloktnujvza", "oihcjpbsxqhxryfmlgktnujvza"}
	// test := []string{"abcde", "fghij", "klmno", "pqrst", "fguij", "axcye", "wvxyz"} // fghij and fguij are shared. Result = fgij

	for _, boxid1 := range input {
		for _, boxid2 := range input {
			if boxid1 == boxid2 {
				continue
			}

			for i := 0; i < len(boxid2); i++ {
				updatedboxid1 := string(remove(stringToSlice(boxid1), i))
				updatedboxid2 := string(remove(stringToSlice(boxid2), i))
				if updatedboxid1 == updatedboxid2 {
					fmt.Printf("common letters: %s", updatedboxid1)
					return
				}
			}
		}
	}
}

func remove(slice []rune, s int) []rune {
	runes := slice
	return append(runes[:s], runes[s+1:]...)
}

func stringToSlice(s string) []rune {
	slice := []rune{}
	for _, char := range s {
		slice = append(slice, char)
	}
	return slice
}
