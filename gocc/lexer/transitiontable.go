// Code generated by gocc; DO NOT EDIT.

package lexer

/*
Let s be the current state
Let r be the current input rune
transitionTable[s](r) returns the next state.
*/
type TransitionTable [NumStates]func(rune) int

var TransTab = TransitionTable{
	// S0
	func(r rune) int {
		switch {
		case r == 9: // ['\t','\t']
			return 1
		case r == 10: // ['\n','\n']
			return 1
		case r == 13: // ['\r','\r']
			return 1
		case r == 32: // [' ',' ']
			return 1
		case r == 34: // ['"','"']
			return 2
		case r == 38: // ['&','&']
			return 3
		case r == 39: // [''',''']
			return 4
		case r == 40: // ['(','(']
			return 5
		case r == 41: // [')',')']
			return 6
		case r == 42: // ['*','*']
			return 7
		case r == 43: // ['+','+']
			return 8
		case r == 44: // [',',',']
			return 9
		case r == 45: // ['-','-']
			return 10
		case r == 46: // ['.','.']
			return 11
		case r == 47: // ['/','/']
			return 12
		case 48 <= r && r <= 57: // ['0','9']
			return 13
		case r == 58: // [':',':']
			return 14
		case r == 59: // [';',';']
			return 15
		case r == 60: // ['<','<']
			return 16
		case r == 61: // ['=','=']
			return 17
		case r == 62: // ['>','>']
			return 18
		case r == 65: // ['A','A']
			return 19
		case r == 66: // ['B','B']
			return 20
		case r == 67: // ['C','C']
			return 21
		case 68 <= r && r <= 72: // ['D','H']
			return 19
		case r == 73: // ['I','I']
			return 22
		case 74 <= r && r <= 82: // ['J','R']
			return 19
		case r == 83: // ['S','S']
			return 23
		case r == 84: // ['T','T']
			return 24
		case 85 <= r && r <= 90: // ['U','Z']
			return 19
		case r == 91: // ['[','[']
			return 25
		case r == 93: // [']',']']
			return 26
		case r == 97: // ['a','a']
			return 19
		case r == 98: // ['b','b']
			return 27
		case 99 <= r && r <= 100: // ['c','d']
			return 19
		case r == 101: // ['e','e']
			return 28
		case r == 102: // ['f','f']
			return 29
		case 103 <= r && r <= 104: // ['g','h']
			return 19
		case r == 105: // ['i','i']
			return 30
		case 106 <= r && r <= 107: // ['j','k']
			return 19
		case r == 108: // ['l','l']
			return 31
		case 109 <= r && r <= 111: // ['m','o']
			return 19
		case r == 112: // ['p','p']
			return 32
		case r == 113: // ['q','q']
			return 19
		case r == 114: // ['r','r']
			return 33
		case r == 115: // ['s','s']
			return 34
		case r == 116: // ['t','t']
			return 35
		case r == 117: // ['u','u']
			return 19
		case r == 118: // ['v','v']
			return 36
		case r == 119: // ['w','w']
			return 37
		case 120 <= r && r <= 122: // ['x','z']
			return 19
		case r == 123: // ['{','{']
			return 38
		case r == 124: // ['|','|']
			return 39
		case r == 125: // ['}','}']
			return 40
		}
		return NoState
	},
	// S1
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S2
	func(r rune) int {
		switch {
		case r == 32: // [' ',' ']
			return 2
		case r == 33: // ['!','!']
			return 2
		case r == 34: // ['"','"']
			return 41
		case r == 35: // ['#','#']
			return 2
		case 48 <= r && r <= 57: // ['0','9']
			return 42
		case r == 63: // ['?','?']
			return 2
		case 65 <= r && r <= 90: // ['A','Z']
			return 43
		case 97 <= r && r <= 122: // ['a','z']
			return 43
		}
		return NoState
	},
	// S3
	func(r rune) int {
		switch {
		case r == 38: // ['&','&']
			return 44
		}
		return NoState
	},
	// S4
	func(r rune) int {
		switch {
		case r == 32: // [' ',' ']
			return 45
		case 48 <= r && r <= 57: // ['0','9']
			return 46
		case 65 <= r && r <= 90: // ['A','Z']
			return 47
		case 97 <= r && r <= 122: // ['a','z']
			return 47
		}
		return NoState
	},
	// S5
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S6
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S7
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S8
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S9
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S10
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S11
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S12
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S13
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 13
		}
		return NoState
	},
	// S14
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S15
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S16
	func(r rune) int {
		switch {
		case r == 61: // ['=','=']
			return 49
		case r == 62: // ['>','>']
			return 49
		}
		return NoState
	},
	// S17
	func(r rune) int {
		switch {
		case r == 61: // ['=','=']
			return 49
		}
		return NoState
	},
	// S18
	func(r rune) int {
		switch {
		case r == 61: // ['=','=']
			return 49
		}
		return NoState
	},
	// S19
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 122: // ['a','z']
			return 19
		}
		return NoState
	},
	// S20
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case r == 97: // ['a','a']
			return 51
		case 98 <= r && r <= 122: // ['b','z']
			return 19
		}
		return NoState
	},
	// S21
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 104: // ['a','h']
			return 19
		case r == 105: // ['i','i']
			return 52
		case 106 <= r && r <= 122: // ['j','z']
			return 19
		}
		return NoState
	},
	// S22
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 108: // ['a','l']
			return 19
		case r == 109: // ['m','m']
			return 53
		case 110 <= r && r <= 122: // ['n','z']
			return 19
		}
		return NoState
	},
	// S23
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 112: // ['a','p']
			return 19
		case r == 113: // ['q','q']
			return 54
		case 114 <= r && r <= 122: // ['r','z']
			return 19
		}
		return NoState
	},
	// S24
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 100: // ['a','d']
			return 19
		case r == 101: // ['e','e']
			return 55
		case 102 <= r && r <= 122: // ['f','z']
			return 19
		}
		return NoState
	},
	// S25
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S26
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S27
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 110: // ['a','n']
			return 19
		case r == 111: // ['o','o']
			return 56
		case 112 <= r && r <= 122: // ['p','z']
			return 19
		}
		return NoState
	},
	// S28
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 107: // ['a','k']
			return 19
		case r == 108: // ['l','l']
			return 57
		case 109 <= r && r <= 122: // ['m','z']
			return 19
		}
		return NoState
	},
	// S29
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case r == 97: // ['a','a']
			return 58
		case 98 <= r && r <= 107: // ['b','k']
			return 19
		case r == 108: // ['l','l']
			return 59
		case 109 <= r && r <= 110: // ['m','n']
			return 19
		case r == 111: // ['o','o']
			return 60
		case 112 <= r && r <= 122: // ['p','z']
			return 19
		}
		return NoState
	},
	// S30
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 101: // ['a','e']
			return 19
		case r == 102: // ['f','f']
			return 61
		case 103 <= r && r <= 109: // ['g','m']
			return 19
		case r == 110: // ['n','n']
			return 62
		case 111 <= r && r <= 122: // ['o','z']
			return 19
		}
		return NoState
	},
	// S31
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 104: // ['a','h']
			return 19
		case r == 105: // ['i','i']
			return 63
		case 106 <= r && r <= 122: // ['j','z']
			return 19
		}
		return NoState
	},
	// S32
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 113: // ['a','q']
			return 19
		case r == 114: // ['r','r']
			return 64
		case 115 <= r && r <= 122: // ['s','z']
			return 19
		}
		return NoState
	},
	// S33
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 100: // ['a','d']
			return 19
		case r == 101: // ['e','e']
			return 65
		case 102 <= r && r <= 122: // ['f','z']
			return 19
		}
		return NoState
	},
	// S34
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 115: // ['a','s']
			return 19
		case r == 116: // ['t','t']
			return 66
		case 117 <= r && r <= 122: // ['u','z']
			return 19
		}
		return NoState
	},
	// S35
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 113: // ['a','q']
			return 19
		case r == 114: // ['r','r']
			return 67
		case 115 <= r && r <= 122: // ['s','z']
			return 19
		}
		return NoState
	},
	// S36
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 110: // ['a','n']
			return 19
		case r == 111: // ['o','o']
			return 68
		case 112 <= r && r <= 122: // ['p','z']
			return 19
		}
		return NoState
	},
	// S37
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 103: // ['a','g']
			return 19
		case r == 104: // ['h','h']
			return 69
		case 105 <= r && r <= 122: // ['i','z']
			return 19
		}
		return NoState
	},
	// S38
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S39
	func(r rune) int {
		switch {
		case r == 124: // ['|','|']
			return 44
		}
		return NoState
	},
	// S40
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S41
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S42
	func(r rune) int {
		switch {
		case r == 32: // [' ',' ']
			return 2
		case r == 33: // ['!','!']
			return 2
		case r == 34: // ['"','"']
			return 41
		case r == 35: // ['#','#']
			return 2
		case 48 <= r && r <= 57: // ['0','9']
			return 42
		case r == 63: // ['?','?']
			return 2
		case 65 <= r && r <= 90: // ['A','Z']
			return 43
		case 97 <= r && r <= 122: // ['a','z']
			return 43
		}
		return NoState
	},
	// S43
	func(r rune) int {
		switch {
		case r == 32: // [' ',' ']
			return 2
		case r == 33: // ['!','!']
			return 2
		case r == 34: // ['"','"']
			return 41
		case r == 35: // ['#','#']
			return 2
		case 48 <= r && r <= 57: // ['0','9']
			return 42
		case r == 63: // ['?','?']
			return 2
		case 65 <= r && r <= 90: // ['A','Z']
			return 43
		case 97 <= r && r <= 122: // ['a','z']
			return 43
		}
		return NoState
	},
	// S44
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S45
	func(r rune) int {
		switch {
		case r == 39: // [''',''']
			return 70
		}
		return NoState
	},
	// S46
	func(r rune) int {
		switch {
		case r == 39: // [''',''']
			return 70
		}
		return NoState
	},
	// S47
	func(r rune) int {
		switch {
		case r == 39: // [''',''']
			return 70
		}
		return NoState
	},
	// S48
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 71
		}
		return NoState
	},
	// S49
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S50
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 122: // ['a','z']
			return 19
		}
		return NoState
	},
	// S51
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 98: // ['a','b']
			return 19
		case r == 99: // ['c','c']
			return 72
		case 100 <= r && r <= 122: // ['d','z']
			return 19
		}
		return NoState
	},
	// S52
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 113: // ['a','q']
			return 19
		case r == 114: // ['r','r']
			return 73
		case 115 <= r && r <= 122: // ['s','z']
			return 19
		}
		return NoState
	},
	// S53
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case r == 97: // ['a','a']
			return 74
		case 98 <= r && r <= 122: // ['b','z']
			return 19
		}
		return NoState
	},
	// S54
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 116: // ['a','t']
			return 19
		case r == 117: // ['u','u']
			return 75
		case 118 <= r && r <= 122: // ['v','z']
			return 19
		}
		return NoState
	},
	// S55
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 119: // ['a','w']
			return 19
		case r == 120: // ['x','x']
			return 76
		case 121 <= r && r <= 122: // ['y','z']
			return 19
		}
		return NoState
	},
	// S56
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 110: // ['a','n']
			return 19
		case r == 111: // ['o','o']
			return 77
		case 112 <= r && r <= 122: // ['p','z']
			return 19
		}
		return NoState
	},
	// S57
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 114: // ['a','r']
			return 19
		case r == 115: // ['s','s']
			return 78
		case 116 <= r && r <= 122: // ['t','z']
			return 19
		}
		return NoState
	},
	// S58
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 107: // ['a','k']
			return 19
		case r == 108: // ['l','l']
			return 79
		case 109 <= r && r <= 122: // ['m','z']
			return 19
		}
		return NoState
	},
	// S59
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 110: // ['a','n']
			return 19
		case r == 111: // ['o','o']
			return 80
		case 112 <= r && r <= 122: // ['p','z']
			return 19
		}
		return NoState
	},
	// S60
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 113: // ['a','q']
			return 19
		case r == 114: // ['r','r']
			return 81
		case 115 <= r && r <= 122: // ['s','z']
			return 19
		}
		return NoState
	},
	// S61
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 122: // ['a','z']
			return 19
		}
		return NoState
	},
	// S62
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 115: // ['a','s']
			return 19
		case r == 116: // ['t','t']
			return 82
		case 117 <= r && r <= 122: // ['u','z']
			return 19
		}
		return NoState
	},
	// S63
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 114: // ['a','r']
			return 19
		case r == 115: // ['s','s']
			return 83
		case 116 <= r && r <= 122: // ['t','z']
			return 19
		}
		return NoState
	},
	// S64
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 110: // ['a','n']
			return 19
		case r == 111: // ['o','o']
			return 84
		case 112 <= r && r <= 122: // ['p','z']
			return 19
		}
		return NoState
	},
	// S65
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 115: // ['a','s']
			return 19
		case r == 116: // ['t','t']
			return 85
		case 117 <= r && r <= 122: // ['u','z']
			return 19
		}
		return NoState
	},
	// S66
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 113: // ['a','q']
			return 19
		case r == 114: // ['r','r']
			return 86
		case 115 <= r && r <= 122: // ['s','z']
			return 19
		}
		return NoState
	},
	// S67
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 116: // ['a','t']
			return 19
		case r == 117: // ['u','u']
			return 87
		case 118 <= r && r <= 122: // ['v','z']
			return 19
		}
		return NoState
	},
	// S68
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 104: // ['a','h']
			return 19
		case r == 105: // ['i','i']
			return 88
		case 106 <= r && r <= 122: // ['j','z']
			return 19
		}
		return NoState
	},
	// S69
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 104: // ['a','h']
			return 19
		case r == 105: // ['i','i']
			return 89
		case 106 <= r && r <= 122: // ['j','z']
			return 19
		}
		return NoState
	},
	// S70
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S71
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 71
		}
		return NoState
	},
	// S72
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 106: // ['a','j']
			return 19
		case r == 107: // ['k','k']
			return 90
		case 108 <= r && r <= 122: // ['l','z']
			return 19
		}
		return NoState
	},
	// S73
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 98: // ['a','b']
			return 19
		case r == 99: // ['c','c']
			return 91
		case 100 <= r && r <= 122: // ['d','z']
			return 19
		}
		return NoState
	},
	// S74
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 102: // ['a','f']
			return 19
		case r == 103: // ['g','g']
			return 92
		case 104 <= r && r <= 122: // ['h','z']
			return 19
		}
		return NoState
	},
	// S75
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case r == 97: // ['a','a']
			return 93
		case 98 <= r && r <= 122: // ['b','z']
			return 19
		}
		return NoState
	},
	// S76
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 115: // ['a','s']
			return 19
		case r == 116: // ['t','t']
			return 94
		case 117 <= r && r <= 122: // ['u','z']
			return 19
		}
		return NoState
	},
	// S77
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 107: // ['a','k']
			return 19
		case r == 108: // ['l','l']
			return 95
		case 109 <= r && r <= 122: // ['m','z']
			return 19
		}
		return NoState
	},
	// S78
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 100: // ['a','d']
			return 19
		case r == 101: // ['e','e']
			return 96
		case 102 <= r && r <= 122: // ['f','z']
			return 19
		}
		return NoState
	},
	// S79
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 114: // ['a','r']
			return 19
		case r == 115: // ['s','s']
			return 97
		case 116 <= r && r <= 122: // ['t','z']
			return 19
		}
		return NoState
	},
	// S80
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case r == 97: // ['a','a']
			return 98
		case 98 <= r && r <= 122: // ['b','z']
			return 19
		}
		return NoState
	},
	// S81
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 122: // ['a','z']
			return 19
		}
		return NoState
	},
	// S82
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 122: // ['a','z']
			return 19
		}
		return NoState
	},
	// S83
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 115: // ['a','s']
			return 19
		case r == 116: // ['t','t']
			return 99
		case 117 <= r && r <= 122: // ['u','z']
			return 19
		}
		return NoState
	},
	// S84
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 102: // ['a','f']
			return 19
		case r == 103: // ['g','g']
			return 100
		case 104 <= r && r <= 122: // ['h','z']
			return 19
		}
		return NoState
	},
	// S85
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 116: // ['a','t']
			return 19
		case r == 117: // ['u','u']
			return 101
		case 118 <= r && r <= 122: // ['v','z']
			return 19
		}
		return NoState
	},
	// S86
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 104: // ['a','h']
			return 19
		case r == 105: // ['i','i']
			return 102
		case 106 <= r && r <= 122: // ['j','z']
			return 19
		}
		return NoState
	},
	// S87
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 100: // ['a','d']
			return 19
		case r == 101: // ['e','e']
			return 103
		case 102 <= r && r <= 122: // ['f','z']
			return 19
		}
		return NoState
	},
	// S88
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 99: // ['a','c']
			return 19
		case r == 100: // ['d','d']
			return 104
		case 101 <= r && r <= 122: // ['e','z']
			return 19
		}
		return NoState
	},
	// S89
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 107: // ['a','k']
			return 19
		case r == 108: // ['l','l']
			return 105
		case 109 <= r && r <= 122: // ['m','z']
			return 19
		}
		return NoState
	},
	// S90
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 102: // ['a','f']
			return 19
		case r == 103: // ['g','g']
			return 106
		case 104 <= r && r <= 122: // ['h','z']
			return 19
		}
		return NoState
	},
	// S91
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 107: // ['a','k']
			return 19
		case r == 108: // ['l','l']
			return 107
		case 109 <= r && r <= 122: // ['m','z']
			return 19
		}
		return NoState
	},
	// S92
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 100: // ['a','d']
			return 19
		case r == 101: // ['e','e']
			return 108
		case 102 <= r && r <= 122: // ['f','z']
			return 19
		}
		return NoState
	},
	// S93
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 113: // ['a','q']
			return 19
		case r == 114: // ['r','r']
			return 109
		case 115 <= r && r <= 122: // ['s','z']
			return 19
		}
		return NoState
	},
	// S94
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 122: // ['a','z']
			return 19
		}
		return NoState
	},
	// S95
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 122: // ['a','z']
			return 19
		}
		return NoState
	},
	// S96
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 122: // ['a','z']
			return 19
		}
		return NoState
	},
	// S97
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 100: // ['a','d']
			return 19
		case r == 101: // ['e','e']
			return 110
		case 102 <= r && r <= 122: // ['f','z']
			return 19
		}
		return NoState
	},
	// S98
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 115: // ['a','s']
			return 19
		case r == 116: // ['t','t']
			return 111
		case 117 <= r && r <= 122: // ['u','z']
			return 19
		}
		return NoState
	},
	// S99
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 122: // ['a','z']
			return 19
		}
		return NoState
	},
	// S100
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 113: // ['a','q']
			return 19
		case r == 114: // ['r','r']
			return 112
		case 115 <= r && r <= 122: // ['s','z']
			return 19
		}
		return NoState
	},
	// S101
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 113: // ['a','q']
			return 19
		case r == 114: // ['r','r']
			return 113
		case 115 <= r && r <= 122: // ['s','z']
			return 19
		}
		return NoState
	},
	// S102
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 109: // ['a','m']
			return 19
		case r == 110: // ['n','n']
			return 114
		case 111 <= r && r <= 122: // ['o','z']
			return 19
		}
		return NoState
	},
	// S103
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 122: // ['a','z']
			return 19
		}
		return NoState
	},
	// S104
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 122: // ['a','z']
			return 19
		}
		return NoState
	},
	// S105
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 100: // ['a','d']
			return 19
		case r == 101: // ['e','e']
			return 115
		case 102 <= r && r <= 122: // ['f','z']
			return 19
		}
		return NoState
	},
	// S106
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 113: // ['a','q']
			return 19
		case r == 114: // ['r','r']
			return 116
		case 115 <= r && r <= 122: // ['s','z']
			return 19
		}
		return NoState
	},
	// S107
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 100: // ['a','d']
			return 19
		case r == 101: // ['e','e']
			return 117
		case 102 <= r && r <= 122: // ['f','z']
			return 19
		}
		return NoState
	},
	// S108
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 122: // ['a','z']
			return 19
		}
		return NoState
	},
	// S109
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 100: // ['a','d']
			return 19
		case r == 101: // ['e','e']
			return 118
		case 102 <= r && r <= 122: // ['f','z']
			return 19
		}
		return NoState
	},
	// S110
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 122: // ['a','z']
			return 19
		}
		return NoState
	},
	// S111
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 122: // ['a','z']
			return 19
		}
		return NoState
	},
	// S112
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case r == 97: // ['a','a']
			return 119
		case 98 <= r && r <= 122: // ['b','z']
			return 19
		}
		return NoState
	},
	// S113
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 109: // ['a','m']
			return 19
		case r == 110: // ['n','n']
			return 120
		case 111 <= r && r <= 122: // ['o','z']
			return 19
		}
		return NoState
	},
	// S114
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 102: // ['a','f']
			return 19
		case r == 103: // ['g','g']
			return 121
		case 104 <= r && r <= 122: // ['h','z']
			return 19
		}
		return NoState
	},
	// S115
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 122: // ['a','z']
			return 19
		}
		return NoState
	},
	// S116
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 110: // ['a','n']
			return 19
		case r == 111: // ['o','o']
			return 122
		case 112 <= r && r <= 122: // ['p','z']
			return 19
		}
		return NoState
	},
	// S117
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 122: // ['a','z']
			return 19
		}
		return NoState
	},
	// S118
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 122: // ['a','z']
			return 19
		}
		return NoState
	},
	// S119
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 108: // ['a','l']
			return 19
		case r == 109: // ['m','m']
			return 123
		case 110 <= r && r <= 122: // ['n','z']
			return 19
		}
		return NoState
	},
	// S120
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 122: // ['a','z']
			return 19
		}
		return NoState
	},
	// S121
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 122: // ['a','z']
			return 19
		}
		return NoState
	},
	// S122
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 116: // ['a','t']
			return 19
		case r == 117: // ['u','u']
			return 124
		case 118 <= r && r <= 122: // ['v','z']
			return 19
		}
		return NoState
	},
	// S123
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 122: // ['a','z']
			return 19
		}
		return NoState
	},
	// S124
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 109: // ['a','m']
			return 19
		case r == 110: // ['n','n']
			return 125
		case 111 <= r && r <= 122: // ['o','z']
			return 19
		}
		return NoState
	},
	// S125
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 99: // ['a','c']
			return 19
		case r == 100: // ['d','d']
			return 126
		case 101 <= r && r <= 122: // ['e','z']
			return 19
		}
		return NoState
	},
	// S126
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 50
		case 65 <= r && r <= 90: // ['A','Z']
			return 19
		case 97 <= r && r <= 122: // ['a','z']
			return 19
		}
		return NoState
	},
}
