/*
<https://leetcode.com/problems/student-attendance-record-ii/>
//constraints
//1 <= n <= 105

'A': Absent.
'L': Late.
'P': Present.

Given an integer n, return the number of possible attendance records of
length n that make a student eligible for an attendance award.
The answer may be very large, so return it modulo 10^9 + 7.

notice that we can:
- have  <= 1 day absent and no more than <= 3 consecutive late days

n = 0; res = 1
empty set {}

n = 1; res = 3
note: nump opts w/o "A" = 2
{P, L, A}

n = 2; res = 8  
note: num opts w/o "A" = 4
{PP, LL, AP, PA, PL, LP, AL, LA}

n = 3 ; res = 19
note: num opts w/o "A" = 7
if not absent any day opts are: PPP, PLL, PPL, LPL, LLL, PLP, LLP 
if absent only day1 opts are : APP, ALL, APL, ALP
if absent only day2 opts are : PAP, LAP,LAL, PAL,  
if absent only day3 opts are : PLA, LPA, PPA, LLA

for n = 4 things get interesting
- let's enumerate all the possibilities w/o "A"
- if not absent any day opts are: PPPP, PPLL, PPPL, PLPL, PLLL, PPLP, PLLP 
                                  LPPP, LPLL, LPPL, LLPL, LPLP, LLLP 
- let's maintain a cache of subproblems for non "A" days
- therefore, for n = 1 we had 2 opts, for n=2 we had 4 opts, and for n =3, 7 opts
-  dp[n-1] + dp[n-2] + dp[n-3] = 7 + 4 + 2 = 13
- next we need to find the number of opts to be absent for 1 day out of these n=4 days 
  - if we're absent on day1, how many opts do we have?
    - A$$$  
       ^^^- dp[0]*dp[3] = 1*7
  - if we're absent on day2, how many opts do we have?
    - $A$$  
      ^-^^-- dp[1]*dp[2] = 2*4
    - aka num opts we have for day1 * num opts we have for day2
    
  - if we're absent on day3, how many opts do we have?
    - $$A$  
      ^^-^ dp[2]*dp[1] = 4*2
  - if we're absent on day4, how many opts do we have?
    - $$$A  
      ^^^-- dp[3]*dp[0] = 7*1
  - therefore, num ways to can be absent 1 of these 4 day opts is = 7 + 8 + 8 + 7 = 30
  - the last thing we have to do is employ the sum rule and add these two disjoint sets
  - the num ways to do A or B is A+B;  where A is the set of nonAbsent days and B is the set of 1day Absent opts 
  - res = 13 + 30 = 43


*/

type attendance byte

const (
    Absent attendance = 'A'
    Late attendance = 'L'
    Present attendance = 'P'

)

const mod = 1e9 + 7

//time: O(2n) ~ O(n)
//space: O(n)

func checkRecord(n int) int {
	switch n {
	case 1:
		return 3
	case 2:
		return 8
	case 3:
		return 19
	}
    //dp table will capture num ways w/o absent days at numDays=n
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 2
	dp[2] = 4
	dp[3] = 7
	for day := 4; day <= n; day++ {
		dp[day] = (dp[day-1] + dp[day-2] + dp[day-3]) % mod
	}
    //num records with 1 out of n days absent
    var optsOneAbs int
	for day := 0; day <= n-1; day++ {
		optsOneAbs += (dp[day] * dp[n-1-day])
        optsOneAbs = optsOneAbs % mod
	}
	return (dp[n] + optsOneAbs) % mod
}
    

