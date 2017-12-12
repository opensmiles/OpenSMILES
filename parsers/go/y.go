
//line smiles.y:16
package smiles
import __yyfmt__ "fmt"
//line smiles.y:17
		
//line smiles.y:19
type smilesSymType struct {
	yys int
  char rune
  number int
  bond BondType
  chain Chain
  chirality Chirality
  element Element
}

const ringbond = 57346
const chainbond = 57347
const dot = 57348
const chirality = 57349
const hydrogen = 57350
const aliphatic_organic = 57351
const aromatic_organic = 57352
const aromatic_other = 57353
const other = 57354
const wildcard = 57355
const minus = 57356
const error1 = 57357
const digit = 57358

var smilesToknames = []string{
	"ringbond",
	"chainbond",
	"dot",
	"chirality",
	"hydrogen",
	"aliphatic_organic",
	"aromatic_organic",
	"aromatic_other",
	"other",
	"wildcard",
	"minus",
	"error1",
	"digit",
}
var smilesStatenames = []string{}

const smilesEofCode = 1
const smilesErrCode = 2
const smilesMaxDepth = 200

//line smiles.y:117


//line yacctab:1
var smilesExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const smilesNprod = 51
const smilesPrivate = 57344

var smilesTokenNames []string
var smilesStates []string

const smilesLast = 119

var smilesAct = []int{

	3, 2, 33, 15, 17, 18, 29, 25, 7, 8,
	19, 53, 9, 16, 28, 32, 14, 43, 44, 45,
	71, 46, 37, 72, 17, 18, 17, 18, 7, 8,
	7, 8, 9, 16, 9, 16, 14, 60, 14, 59,
	70, 51, 69, 26, 52, 23, 56, 15, 54, 62,
	63, 64, 55, 38, 39, 41, 42, 40, 57, 68,
	58, 26, 67, 15, 15, 15, 17, 18, 48, 49,
	7, 8, 7, 8, 9, 16, 9, 47, 14, 31,
	14, 22, 61, 17, 18, 66, 65, 7, 8, 26,
	50, 9, 16, 20, 1, 14, 21, 7, 8, 27,
	24, 9, 30, 34, 10, 14, 38, 39, 41, 42,
	40, 11, 12, 13, 6, 4, 5, 36, 35,
}
var smilesPact = []int{

	88, -1000, 78, -10, 77, -1000, -1000, -1000, -1000, -1000,
	27, -8, 71, 8, 45, -10, 88, 88, 88, 63,
	-1000, 74, 25, -1000, -7, 32, -1000, -1000, 30, 44,
	-1000, 73, -1000, -1000, 98, -1000, -1000, 32, -1000, -1000,
	-1000, -1000, -1000, -10, -10, -10, 61, 88, 88, 88,
	70, -1000, 69, -1000, -1000, -1000, 46, -1000, 43, 32,
	-1000, -1000, 21, 19, -1, -1000, 7, -1000, -1000, -1000,
	-1000, -1000, -1000,
}
var smilesPgo = []int{

	0, 118, 117, 116, 115, 0, 114, 2, 113, 112,
	111, 104, 1, 103, 102, 100, 7, 99, 94,
}
var smilesR1 = []int{

	0, 18, 12, 12, 12, 12, 12, 7, 7, 8,
	8, 9, 9, 10, 10, 11, 11, 6, 6, 3,
	3, 3, 3, 4, 4, 4, 4, 4, 5, 5,
	5, 5, 5, 1, 1, 1, 2, 2, 13, 14,
	14, 17, 17, 17, 17, 17, 17, 17, 15, 16,
	16,
}
var smilesR2 = []int{

	0, 1, 1, 2, 3, 3, 3, 1, 1, 2,
	3, 1, 2, 1, 2, 1, 2, 2, 3, 1,
	1, 1, 1, 1, 2, 4, 3, 5, 1, 4,
	5, 5, 5, 1, 1, 1, 1, 1, 1, 1,
	2, 1, 2, 2, 3, 1, 2, 3, 1, 1,
	2,
}
var smilesChk = []int{

	-1000, -18, -12, -5, -4, -3, -6, 9, 10, 13,
	-11, -10, -9, -8, 17, -5, 14, 5, 6, 20,
	16, 19, 4, 18, -15, -16, 16, -17, 22, 14,
	-14, 8, 7, -7, -13, -1, -2, -16, 8, 9,
	12, 10, 11, -5, -5, -5, -12, 14, 5, 6,
	16, 16, 19, 18, 16, 22, 16, 14, 16, -16,
	-7, 21, -12, -12, -12, 16, 16, 16, 16, 21,
	21, 21, 16,
}
var smilesDef = []int{

	0, -2, 1, 2, 28, 23, 19, 20, 21, 22,
	0, 15, 13, 11, 0, 3, 0, 0, 0, 0,
	24, 0, 0, 17, 0, 48, 49, 16, 41, 45,
	14, 39, 12, 9, 0, 7, 8, 38, 33, 34,
	35, 36, 37, 4, 5, 6, 0, 0, 0, 0,
	0, 26, 0, 18, 50, 42, 43, 46, 0, 40,
	10, 29, 0, 0, 0, 25, 0, 44, 47, 30,
	31, 32, 27,
}
var smilesTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 19, 3, 3,
	20, 21, 3, 22, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 17, 3, 18,
}
var smilesTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16,
}
var smilesTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var smilesDebug = 0

type smilesLexer interface {
	Lex(lval *smilesSymType) int
	Error(s string)
}

const smilesFlag = -1000

func smilesTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(smilesToknames) {
		if smilesToknames[c-4] != "" {
			return smilesToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func smilesStatname(s int) string {
	if s >= 0 && s < len(smilesStatenames) {
		if smilesStatenames[s] != "" {
			return smilesStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func smileslex1(lex smilesLexer, lval *smilesSymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = smilesTok1[0]
		goto out
	}
	if char < len(smilesTok1) {
		c = smilesTok1[char]
		goto out
	}
	if char >= smilesPrivate {
		if char < smilesPrivate+len(smilesTok2) {
			c = smilesTok2[char-smilesPrivate]
			goto out
		}
	}
	for i := 0; i < len(smilesTok3); i += 2 {
		c = smilesTok3[i+0]
		if c == char {
			c = smilesTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = smilesTok2[1] /* unknown char */
	}
	if smilesDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", smilesTokname(c), uint(char))
	}
	return c
}

func smilesParse(smileslex smilesLexer) int {
	var smilesn int
	var smileslval smilesSymType
	var smilesVAL smilesSymType
	smilesS := make([]smilesSymType, smilesMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	smilesstate := 0
	smileschar := -1
	smilesp := -1
	goto smilesstack

ret0:
	return 0

ret1:
	return 1

smilesstack:
	/* put a state and value onto the stack */
	if smilesDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", smilesTokname(smileschar), smilesStatname(smilesstate))
	}

	smilesp++
	if smilesp >= len(smilesS) {
		nyys := make([]smilesSymType, len(smilesS)*2)
		copy(nyys, smilesS)
		smilesS = nyys
	}
	smilesS[smilesp] = smilesVAL
	smilesS[smilesp].yys = smilesstate

smilesnewstate:
	smilesn = smilesPact[smilesstate]
	if smilesn <= smilesFlag {
		goto smilesdefault /* simple state */
	}
	if smileschar < 0 {
		smileschar = smileslex1(smileslex, &smileslval)
	}
	smilesn += smileschar
	if smilesn < 0 || smilesn >= smilesLast {
		goto smilesdefault
	}
	smilesn = smilesAct[smilesn]
	if smilesChk[smilesn] == smileschar { /* valid shift */
		smileschar = -1
		smilesVAL = smileslval
		smilesstate = smilesn
		if Errflag > 0 {
			Errflag--
		}
		goto smilesstack
	}

smilesdefault:
	/* default state action */
	smilesn = smilesDef[smilesstate]
	if smilesn == -2 {
		if smileschar < 0 {
			smileschar = smileslex1(smileslex, &smileslval)
		}

		/* look through exception table */
		xi := 0
		for {
			if smilesExca[xi+0] == -1 && smilesExca[xi+1] == smilesstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			smilesn = smilesExca[xi+0]
			if smilesn < 0 || smilesn == smileschar {
				break
			}
		}
		smilesn = smilesExca[xi+1]
		if smilesn < 0 {
			goto ret0
		}
	}
	if smilesn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			smileslex.Error("syntax error")
			Nerrs++
			if smilesDebug >= 1 {
				__yyfmt__.Printf("%s", smilesStatname(smilesstate))
				__yyfmt__.Printf(" saw %s\n", smilesTokname(smileschar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for smilesp >= 0 {
				smilesn = smilesPact[smilesS[smilesp].yys] + smilesErrCode
				if smilesn >= 0 && smilesn < smilesLast {
					smilesstate = smilesAct[smilesn] /* simulate a shift of "error" */
					if smilesChk[smilesstate] == smilesErrCode {
						goto smilesstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if smilesDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", smilesS[smilesp].yys)
				}
				smilesp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if smilesDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", smilesTokname(smileschar))
			}
			if smileschar == smilesEofCode {
				goto ret1
			}
			smileschar = -1
			goto smilesnewstate /* try again in the same state */
		}
	}

	/* reduction by production smilesn */
	if smilesDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", smilesn, smilesStatname(smilesstate))
	}

	smilesnt := smilesn
	smilespt := smilesp
	_ = smilespt // guard against "declared and not used"

	smilesp -= smilesR2[smilesn]
	smilesVAL = smilesS[smilesp+1]

	/* consult goto table to find next state */
	smilesn = smilesR1[smilesn]
	smilesg := smilesPgo[smilesn]
	smilesj := smilesg + smilesS[smilesp].yys + 1

	if smilesj >= smilesLast {
		smilesstate = smilesAct[smilesg]
	} else {
		smilesstate = smilesAct[smilesj]
		if smilesChk[smilesstate] != -smilesn {
			smilesstate = smilesAct[smilesg]
		}
	}
	// dummy call; replaced with literal code
	switch smilesnt {

	case 1:
		//line smiles.y:41
		{ FinalizeMolecule(smileslex)}
	case 2:
		//line smiles.y:43
		{ smilesVAL.chain = AppendToChain(smileslex, NewChain(), smilesS[smilespt-0].number, BOND_DEFAULT)}
	case 3:
		//line smiles.y:44
		{ smilesVAL.chain = AppendToChain(smileslex, smilesS[smilespt-1].chain, smilesS[smilespt-0].number, BOND_DEFAULT) }
	case 4:
		//line smiles.y:45
		{ smilesVAL.chain = AppendToChain(smileslex, smilesS[smilespt-2].chain, smilesS[smilespt-0].number, BOND_SINGLE) }
	case 5:
		//line smiles.y:46
		{ smilesVAL.chain = AppendToChain(smileslex, smilesS[smilespt-2].chain, smilesS[smilespt-0].number, smilesS[smilespt-1].bond) }
	case 6:
		//line smiles.y:47
		{ smilesVAL.chain = AppendToChain(smileslex, smilesS[smilespt-2].chain, smilesS[smilespt-0].number, smilesS[smilespt-1].bond) }
	case 7:
		//line smiles.y:49
		{ smilesVAL.number = NewAtom(smileslex, smilesS[smilespt-0].element, false, false) }
	case 8:
		//line smiles.y:50
		{ smilesVAL.number = NewAtom(smileslex, smilesS[smilespt-0].element, true, false) }
	case 9:
		//line smiles.y:53
		{ smilesVAL.number = smilesS[smilespt-0].number }
	case 10:
		//line smiles.y:54
		{ smilesVAL.number = SetIsotope(smileslex, smilesS[smilespt-0].number, smilesS[smilespt-1].number) }
	case 11:
		smilesVAL.number = smilesS[smilespt-0].number
	case 12:
		//line smiles.y:58
		{ smilesVAL.number = SetChirality(smileslex, smilesS[smilespt-1].number, smilesS[smilespt-0].chirality) }
	case 13:
		smilesVAL.number = smilesS[smilespt-0].number
	case 14:
		//line smiles.y:62
		{ smilesVAL.number = SetHCount(smileslex, smilesS[smilespt-1].number, smilesS[smilespt-0].number) }
	case 15:
		smilesVAL.number = smilesS[smilespt-0].number
	case 16:
		//line smiles.y:66
		{ smilesVAL.number = SetCharge(smileslex, smilesS[smilespt-1].number, smilesS[smilespt-0].number) }
	case 17:
		//line smiles.y:69
		{ smilesVAL.number = smilesS[smilespt-1].number }
	case 18:
		//line smiles.y:70
		{ smilesVAL.number = SetClass(smileslex, smilesS[smilespt-2].number, smilesS[smilespt-1].number) }
	case 19:
		smilesVAL.number = smilesS[smilespt-0].number
	case 20:
		//line smiles.y:74
		{ smilesVAL.number = NewAtom(smileslex, smilesS[smilespt-0].element, false, true) }
	case 21:
		//line smiles.y:75
		{ smilesVAL.number = NewAtom(smileslex, smilesS[smilespt-0].element, true, false) }
	case 22:
		//line smiles.y:76
		{ smilesVAL.number = NewAtom(smileslex, smilesS[smilespt-0].element, false, false) }
	case 23:
		smilesVAL.number = smilesS[smilespt-0].number
	case 24:
		//line smiles.y:80
		{ smilesVAL.number = AddRingbond(smileslex, smilesS[smilespt-1].number, BOND_DEFAULT, smilesS[smilespt-0].number) }
	case 25:
		//line smiles.y:81
		{ smilesVAL.number = AddRingbond(smileslex, smilesS[smilespt-3].number, BOND_DEFAULT, smilesS[smilespt-1].number * 10 + smilesS[smilespt-0].number) }
	case 26:
		//line smiles.y:82
		{ smilesVAL.number = AddRingbond(smileslex, smilesS[smilespt-2].number, smilesS[smilespt-1].bond, smilesS[smilespt-0].number) }
	case 27:
		//line smiles.y:83
		{ smilesVAL.number = AddRingbond(smileslex, smilesS[smilespt-4].number, smilesS[smilespt-3].bond, smilesS[smilespt-1].number * 10 + smilesS[smilespt-0].number) }
	case 28:
		smilesVAL.number = smilesS[smilespt-0].number
	case 29:
		//line smiles.y:88
		{ smilesVAL.number = AttachBranch(smileslex, smilesS[smilespt-3].number, smilesS[smilespt-1].chain, BOND_DEFAULT)}
	case 30:
		//line smiles.y:89
		{ smilesVAL.number = AttachBranch(smileslex, smilesS[smilespt-4].number, smilesS[smilespt-1].chain, BOND_SINGLE)}
	case 31:
		//line smiles.y:90
		{ smilesVAL.number = AttachBranch(smileslex, smilesS[smilespt-4].number, smilesS[smilespt-1].chain, smilesS[smilespt-2].bond)}
	case 32:
		//line smiles.y:91
		{ smilesVAL.number = AttachBranch(smileslex, smilesS[smilespt-4].number, smilesS[smilespt-1].chain, smilesS[smilespt-2].bond)}
	case 33:
		smilesVAL.element = smilesS[smilespt-0].element
	case 34:
		smilesVAL.element = smilesS[smilespt-0].element
	case 35:
		smilesVAL.element = smilesS[smilespt-0].element
	case 36:
		smilesVAL.element = smilesS[smilespt-0].element
	case 37:
		smilesVAL.element = smilesS[smilespt-0].element
	case 38:
		smilesVAL.number = smilesS[smilespt-0].number
	case 39:
		//line smiles.y:98
		{ smilesVAL.number = 1 }
	case 40:
		//line smiles.y:99
		{ smilesVAL.number = smilesS[smilespt-0].number }
	case 41:
		//line smiles.y:102
		{ smilesVAL.number = 1 }
	case 42:
		//line smiles.y:103
		{ smilesVAL.number = 2 }
	case 43:
		//line smiles.y:104
		{ smilesVAL.number = int(smilesS[smilespt-0].number) }
	case 44:
		//line smiles.y:105
		{ smilesVAL.number = int(smilesS[smilespt-1].number*10+smilesS[smilespt-0].number) }
	case 45:
		//line smiles.y:106
		{ smilesVAL.number = -1 }
	case 46:
		//line smiles.y:107
		{ smilesVAL.number = -2 }
	case 47:
		//line smiles.y:108
		{ smilesVAL.number = -int(smilesS[smilespt-1].number*10+smilesS[smilespt-0].number) }
	case 48:
		smilesVAL.number = smilesS[smilespt-0].number
	case 49:
		smilesVAL.number = smilesS[smilespt-0].number
	case 50:
		//line smiles.y:114
		{ smilesVAL.number = smilesS[smilespt-1].number * 10 + smilesS[smilespt-0].number }
	}
	goto smilesstack /* stack new state and value */
}
