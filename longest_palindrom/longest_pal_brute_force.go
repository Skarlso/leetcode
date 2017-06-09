package main

import (
	"fmt"
)

func longestPalindromeBrute(s string) string {
	longest := ""
	l := len(s)
	for i := 0; i < l; i++ {
		pal := ""
		runes := []rune(s[i:])
		rl := len(runes)
		for j := 0; j < rl; j++ {
			pal += string(runes[j])
			if isPalindromeBrute(pal) {
				palLength := len(pal)
				if palLength > len(longest) {
					longest = pal
				}
			}
		}
	}
	return longest
}

func isPalindromeBrute(s string) bool {
	return reverseBrute(s) == s
}

func reverseBrute(s string) string {
	runes := []rune(s)
	l := len(runes) - 1
	for i, j := 0, l; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func xmain() {
	fmt.Println("Longest: ", longestPalindromeBrute("civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth"))
}
