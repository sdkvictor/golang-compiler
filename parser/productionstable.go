// Code generated by gocc; DO NOT EDIT.

package parser



type (
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index      int
		NumSymbols int
		ReduceFunc func([]Attrib) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab{
	ProdTabEntry{
		String: `S' : Programa	<<  >>`,
		Id:         "S'",
		NTType:     0,
		Index:      0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Programa : program id semicolon Vars Functions Tick Main	<<  >>`,
		Id:         "Programa",
		NTType:     1,
		Index:      1,
		NumSymbols: 7,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Programa : program id semicolon Tick Main	<<  >>`,
		Id:         "Programa",
		NTType:     1,
		Index:      2,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Programa : program id semicolon Functions Tick Main	<<  >>`,
		Id:         "Programa",
		NTType:     1,
		Index:      3,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Vars : Type VarsAux	<<  >>`,
		Id:         "Vars",
		NTType:     2,
		Index:      4,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `VarsAux : Ids semicolon Vars	<<  >>`,
		Id:         "VarsAux",
		NTType:     3,
		Index:      5,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `VarsAux : Ids semicolon	<<  >>`,
		Id:         "VarsAux",
		NTType:     3,
		Index:      6,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Ids : id comma Ids	<<  >>`,
		Id:         "Ids",
		NTType:     4,
		Index:      7,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Ids : id	<<  >>`,
		Id:         "Ids",
		NTType:     4,
		Index:      8,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Functions : FunctionsAux id leftparenthesis Vars rightparenthesis Block Functions	<<  >>`,
		Id:         "Functions",
		NTType:     5,
		Index:      9,
		NumSymbols: 7,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Functions : FunctionsAux id leftparenthesis Vars rightparenthesis Block	<<  >>`,
		Id:         "Functions",
		NTType:     5,
		Index:      10,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `FunctionsAux : Type	<<  >>`,
		Id:         "FunctionsAux",
		NTType:     6,
		Index:      11,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `FunctionsAux : voidtype	<<  >>`,
		Id:         "FunctionsAux",
		NTType:     6,
		Index:      12,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Block : leftbracket BlockAux	<<  >>`,
		Id:         "Block",
		NTType:     7,
		Index:      13,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `BlockAux : Statement rightbracket	<<  >>`,
		Id:         "BlockAux",
		NTType:     8,
		Index:      14,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `BlockAux : Statement BlockAux	<<  >>`,
		Id:         "BlockAux",
		NTType:     8,
		Index:      15,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Statement : Assign	<<  >>`,
		Id:         "Statement",
		NTType:     9,
		Index:      16,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Statement : Condition	<<  >>`,
		Id:         "Statement",
		NTType:     9,
		Index:      17,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Statement : Write	<<  >>`,
		Id:         "Statement",
		NTType:     9,
		Index:      18,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Statement : Return	<<  >>`,
		Id:         "Statement",
		NTType:     9,
		Index:      19,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Statement : For	<<  >>`,
		Id:         "Statement",
		NTType:     9,
		Index:      20,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Statement : While	<<  >>`,
		Id:         "Statement",
		NTType:     9,
		Index:      21,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Statement : CallFunction	<<  >>`,
		Id:         "Statement",
		NTType:     9,
		Index:      22,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Statement : PredefinedFunction	<<  >>`,
		Id:         "Statement",
		NTType:     9,
		Index:      23,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Expression : Exp	<<  >>`,
		Id:         "Expression",
		NTType:     10,
		Index:      24,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Expression : Exp Operations Exp	<<  >>`,
		Id:         "Expression",
		NTType:     10,
		Index:      25,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Operations : relop	<<  >>`,
		Id:         "Operations",
		NTType:     11,
		Index:      26,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Operations : logicalop	<<  >>`,
		Id:         "Operations",
		NTType:     11,
		Index:      27,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `BasicType : inttype	<<  >>`,
		Id:         "BasicType",
		NTType:     12,
		Index:      28,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `BasicType : floattype	<<  >>`,
		Id:         "BasicType",
		NTType:     12,
		Index:      29,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `BasicType : booltype	<<  >>`,
		Id:         "BasicType",
		NTType:     12,
		Index:      30,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `BasicType : stringtype	<<  >>`,
		Id:         "BasicType",
		NTType:     12,
		Index:      31,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `BasicType : Object	<<  >>`,
		Id:         "BasicType",
		NTType:     12,
		Index:      32,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Object : squaretype	<<  >>`,
		Id:         "Object",
		NTType:     13,
		Index:      33,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Object : circletype	<<  >>`,
		Id:         "Object",
		NTType:     13,
		Index:      34,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Object : imagetype	<<  >>`,
		Id:         "Object",
		NTType:     13,
		Index:      35,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Object : texttype	<<  >>`,
		Id:         "Object",
		NTType:     13,
		Index:      36,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Object : backgroundtype	<<  >>`,
		Id:         "Object",
		NTType:     13,
		Index:      37,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Term : Factor	<<  >>`,
		Id:         "Term",
		NTType:     14,
		Index:      38,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Term : Factor mult Factor	<<  >>`,
		Id:         "Term",
		NTType:     14,
		Index:      39,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Term : Factor div Factor	<<  >>`,
		Id:         "Term",
		NTType:     14,
		Index:      40,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Exp : Term	<<  >>`,
		Id:         "Exp",
		NTType:     15,
		Index:      41,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Exp : Term plus Term	<<  >>`,
		Id:         "Exp",
		NTType:     15,
		Index:      42,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Exp : Term minus Term	<<  >>`,
		Id:         "Exp",
		NTType:     15,
		Index:      43,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Factor : leftparenthesis Expression rightparenthesis	<<  >>`,
		Id:         "Factor",
		NTType:     16,
		Index:      44,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Factor : plus Varcte	<<  >>`,
		Id:         "Factor",
		NTType:     16,
		Index:      45,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Factor : minus Varcte	<<  >>`,
		Id:         "Factor",
		NTType:     16,
		Index:      46,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Factor : Varcte	<<  >>`,
		Id:         "Factor",
		NTType:     16,
		Index:      47,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Assign : id equals Expression semicolon	<<  >>`,
		Id:         "Assign",
		NTType:     17,
		Index:      48,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Assign : Attribute equals Expression semicolon	<<  >>`,
		Id:         "Assign",
		NTType:     17,
		Index:      49,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Condition : if leftparenthesis Expression rightparenthesis Block semicolon	<<  >>`,
		Id:         "Condition",
		NTType:     18,
		Index:      50,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Condition : if leftparenthesis Expression rightparenthesis else Block semicolon	<<  >>`,
		Id:         "Condition",
		NTType:     18,
		Index:      51,
		NumSymbols: 7,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Write : print leftparenthesis WriteAux rightparenthesis semicolon	<<  >>`,
		Id:         "Write",
		NTType:     19,
		Index:      52,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `WriteAux : Expression comma WriteAux	<<  >>`,
		Id:         "WriteAux",
		NTType:     20,
		Index:      53,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `WriteAux : ctestring comma WriteAux	<<  >>`,
		Id:         "WriteAux",
		NTType:     20,
		Index:      54,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `WriteAux : Expression	<<  >>`,
		Id:         "WriteAux",
		NTType:     20,
		Index:      55,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `WriteAux : ctestring	<<  >>`,
		Id:         "WriteAux",
		NTType:     20,
		Index:      56,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Return : return Expression semicolon	<<  >>`,
		Id:         "Return",
		NTType:     21,
		Index:      57,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `For : for leftparenthesis Assign semicolon Expression semicolon Expression semicolon rightparenthesis Block	<<  >>`,
		Id:         "For",
		NTType:     22,
		Index:      58,
		NumSymbols: 10,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `While : while leftparenthesis Expression rightparenthesis Block	<<  >>`,
		Id:         "While",
		NTType:     23,
		Index:      59,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `CallFunction : id CallFunctionAux	<<  >>`,
		Id:         "CallFunction",
		NTType:     24,
		Index:      60,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `CallFunctionAux : leftparenthesis Expression rightparenthesis	<<  >>`,
		Id:         "CallFunctionAux",
		NTType:     25,
		Index:      61,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `CallFunctionAux : leftparenthesis Expression comma CallFunctionAux	<<  >>`,
		Id:         "CallFunctionAux",
		NTType:     25,
		Index:      62,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `CallFunctionAux : leftparenthesis rightparenthesis	<<  >>`,
		Id:         "CallFunctionAux",
		NTType:     25,
		Index:      63,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Varcte : id	<<  >>`,
		Id:         "Varcte",
		NTType:     26,
		Index:      64,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Varcte : cteint	<<  >>`,
		Id:         "Varcte",
		NTType:     26,
		Index:      65,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Varcte : ctefloat	<<  >>`,
		Id:         "Varcte",
		NTType:     26,
		Index:      66,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Varcte : ctestring	<<  >>`,
		Id:         "Varcte",
		NTType:     26,
		Index:      67,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Varcte : ListElem	<<  >>`,
		Id:         "Varcte",
		NTType:     26,
		Index:      68,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Varcte : ctebool	<<  >>`,
		Id:         "Varcte",
		NTType:     26,
		Index:      69,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Varcte : Attribute	<<  >>`,
		Id:         "Varcte",
		NTType:     26,
		Index:      70,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Main : voidtype main leftparenthesis rightparenthesis Block	<<  >>`,
		Id:         "Main",
		NTType:     27,
		Index:      71,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Tick : voidtype tick leftparenthesis rightparenthesis Block	<<  >>`,
		Id:         "Tick",
		NTType:     28,
		Index:      72,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `SetColor : setColor leftparenthesis ctestring rightparenthesis	<<  >>`,
		Id:         "SetColor",
		NTType:     29,
		Index:      73,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ListElem : id leftsqrbracket cteint rightsqrbracket	<<  >>`,
		Id:         "ListElem",
		NTType:     30,
		Index:      74,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `SetSize : setSize leftparenthesis cteint comma cteint rightparenthesis	<<  >>`,
		Id:         "SetSize",
		NTType:     31,
		Index:      75,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `SetPosition : setPosition leftparenthesis cteint comma cteint rightparenthesis	<<  >>`,
		Id:         "SetPosition",
		NTType:     32,
		Index:      76,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Move : move leftparenthesis cteint comma cteint rightparenthesis	<<  >>`,
		Id:         "Move",
		NTType:     33,
		Index:      77,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `SetText : setText leftparenthesis ctestring rightparenthesis	<<  >>`,
		Id:         "SetText",
		NTType:     34,
		Index:      78,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Render : render leftparenthesis rightparenthesis	<<  >>`,
		Id:         "Render",
		NTType:     35,
		Index:      79,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `KeyPressed : keyPressed leftparenthesis ctestring rightparenthesis	<<  >>`,
		Id:         "KeyPressed",
		NTType:     36,
		Index:      80,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `SetFontSize : setFontSize leftparenthesis cteint rightparenthesis	<<  >>`,
		Id:         "SetFontSize",
		NTType:     37,
		Index:      81,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Type : BasicType	<<  >>`,
		Id:         "Type",
		NTType:     38,
		Index:      82,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Type : list greaterthan BasicType lessthan	<<  >>`,
		Id:         "Type",
		NTType:     38,
		Index:      83,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Attribute : id dot id	<<  >>`,
		Id:         "Attribute",
		NTType:     39,
		Index:      84,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `LoadImage : loadImage leftparenthesis ctestring rightparenthesis	<<  >>`,
		Id:         "LoadImage",
		NTType:     40,
		Index:      85,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `CheckCollision : checkCollision leftparenthesis Object dot Object rightparenthesis	<<  >>`,
		Id:         "CheckCollision",
		NTType:     41,
		Index:      86,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `SetImage : setImage leftparenthesis ctestring rightparenthesis	<<  >>`,
		Id:         "SetImage",
		NTType:     42,
		Index:      87,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Pow : pow leftparenthesis ctefloat comma ctefloat rightparenthesis	<<  >>`,
		Id:         "Pow",
		NTType:     43,
		Index:      88,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `SquareRoot : squareRoot leftparenthesis ctefloat rightparenthesis	<<  >>`,
		Id:         "SquareRoot",
		NTType:     44,
		Index:      89,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `PredefinedFunction : SetColor	<<  >>`,
		Id:         "PredefinedFunction",
		NTType:     45,
		Index:      90,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `PredefinedFunction : SetSize	<<  >>`,
		Id:         "PredefinedFunction",
		NTType:     45,
		Index:      91,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `PredefinedFunction : SetPosition	<<  >>`,
		Id:         "PredefinedFunction",
		NTType:     45,
		Index:      92,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `PredefinedFunction : Move	<<  >>`,
		Id:         "PredefinedFunction",
		NTType:     45,
		Index:      93,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `PredefinedFunction : SetText	<<  >>`,
		Id:         "PredefinedFunction",
		NTType:     45,
		Index:      94,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `PredefinedFunction : KeyPressed	<<  >>`,
		Id:         "PredefinedFunction",
		NTType:     45,
		Index:      95,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `PredefinedFunction : SetFontSize	<<  >>`,
		Id:         "PredefinedFunction",
		NTType:     45,
		Index:      96,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `PredefinedFunction : Render	<<  >>`,
		Id:         "PredefinedFunction",
		NTType:     45,
		Index:      97,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `PredefinedFunction : LoadImage	<<  >>`,
		Id:         "PredefinedFunction",
		NTType:     45,
		Index:      98,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `PredefinedFunction : CheckCollision	<<  >>`,
		Id:         "PredefinedFunction",
		NTType:     45,
		Index:      99,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `PredefinedFunction : SetImage	<<  >>`,
		Id:         "PredefinedFunction",
		NTType:     45,
		Index:      100,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `PredefinedFunction : Pow	<<  >>`,
		Id:         "PredefinedFunction",
		NTType:     45,
		Index:      101,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `PredefinedFunction : SquareRoot	<<  >>`,
		Id:         "PredefinedFunction",
		NTType:     45,
		Index:      102,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
}
