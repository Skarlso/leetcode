package main

import (
	"fmt"
)

func longestPalindrome(s string) string {
	longest := ""
	l := len(s)
	for i := 0; i < l; i++ {
		pal := ""
		runes := []rune(s[i:])
		rl := len(runes)
		for j := 0; j < rl; j++ {
			pal += string(runes[j])
			if isPalindrome(pal) {
				palLength := len(pal)
				if palLength > len(longest) {
					longest = pal
				}
			}
		}
	}
	return longest
}

func isPalindrome(s string) bool {
	return reverse(s) == s
}

func reverse(s string) string {
	runes := []rune(s)
	l := len(runes) - 1
	for i, j := 0, l; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	fmt.Println("Longest: ", longestPalindrome("civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth"))
}
