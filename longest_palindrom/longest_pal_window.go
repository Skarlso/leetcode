package main

import (
	"fmt"
	"sync"
)

type pals struct {
	m map[string]int
	sync.RWMutex
}

func (p *pals) get(s string) int {
	p.RLock()
	defer p.RUnlock()
	return p.m[s]
}

func (p *pals) set(s string, v int) {
	p.Lock()
	defer p.Unlock()
	p.m[s] = v
}

func longestPalindrome(s string) string {
	pal := pals{m: make(map[string]int)}
	l := len(s)
	var wg sync.WaitGroup
	for i := 0; i < l; i++ {
		wg.Add(1)
		go func(c int) {
			defer wg.Done()
			for j := c; j < l; j++ {
				sub := s[c : j+1]
				if isPalindrome(sub) {
					pal.set(sub, len(sub))
				}
			}
		}(i)
	}
	wg.Wait()
	longest := ""
	for k, v := range pal.m {
		if v > len(longest) {
			longest = k
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
	//fmt.Println("Longest: ", longestPalindrome("cec"))
	//fmt.Println("Longest:", longestPalindrome("aaaabaaa"))
}
