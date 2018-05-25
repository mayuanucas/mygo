package problem696

//preRun count the same item happend before (let say you have 0011, preRun = 2 when you hit the first 1, means there are two zeros before first '1')
//curRun count the current number of items (let say you have 0011, curRun = 2 when you hit the second 1, means there are two 1s so far)
//Whenever item change (from 0 to 1 or from 1 to 0), preRun change to curRun, reset curRun to 1 (store the curRun number into PreRun, reset curRun)
//Every time preRun >= curRun means there are more 0s before 1s, so could do count++ .
// (This was the tricky one, ex. 0011 when you hit the first '1', curRun = 1, preRun = 2, means 0s number is larger than 1s number,
// so we could form "01" at this time, count++ . When you hit the second '1', curRun = 2, preRun = 2, means 0s' number equals to 1s' number,
// so we could form "0011" at this time, that is why count++)
func countBinarySubstrings(s string) int {
	prevRun, curRun, ans := 0, 1, 0
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			curRun++
		} else {
			prevRun = curRun
			curRun = 1
		}
		if prevRun >= curRun {
			ans++
		}
	}
	return ans
}
