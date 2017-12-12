package smiles

type BondType int
type ChiralityType int

const (
  BOND_DEFAULT = BondType(iota)
  BOND_SINGLE
  BOND_DOUBLE
  BOND_TRIPLE
  BOND_QUADRUPLE
  BOND_UP
  BOND_DOWN
  BOND_AROMATIC
)

const (
  CHIRALITY_DEFAULT = ChiralityType(iota)
  CHIRALITY_CCW
  CHIRALITY_CW
  CHIRALITY_TH // Valid indices are 1, 2
  CHIRALITY_AL // Valid indices are 1, 2
  CHIRALITY_SP // Valid indices are 1, 2, 3
  CHIRALITY_TB // Valid indices are 1...20
  CHIRALITY_OH // Valid indices are 1...30
)

type Atom struct {
  Bonds     []Bond
  Branches  []Chain
  Charge    int
  Chirality Chirality
  Class     int
  Element   Element
  HCount    int
  Isotope   int

  // used only while parsing SMILES
  aliphatic bool
  aromatic  bool
  index     int
}

type Bond struct {
  Type         BondType
  Atom1, Atom2 int // <0 are the ringbonds that are still unresolved
}

type Chain []int

type Chirality struct {
  Type  ChiralityType
  Index int
}

type Molecule struct {
  Atoms []*Atom
  Bonds []Bond
}
