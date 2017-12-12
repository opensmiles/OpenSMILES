/*
Copyright 2014 Google Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package smiles

import "fmt"

func NewAtom(lexer smilesLexer, e Element, aromacity, aliphacity bool) int {
  molecule := &lexer.(*smilesLex).Molecule
  isotope := int(Weights[e] + 0.5)
  a := &Atom{Element: e, index: len(molecule.Atoms), aromatic: aromacity, aliphatic: aliphacity, Isotope: isotope}
  molecule.Atoms = append(molecule.Atoms, a)
  return a.index
}

func NewChain() Chain {
  return Chain{}
}

func SetChirality(lexer smilesLexer, a int, c Chirality) int {
  lexer.(*smilesLex).Molecule.Atoms[a].Chirality = c
  return a
}

func SetHCount(lexer smilesLexer, a int, c int) int {
  lexer.(*smilesLex).Molecule.Atoms[a].HCount = c
  return a
}

func SetCharge(lexer smilesLexer, a int, c int) int {
  lexer.(*smilesLex).Molecule.Atoms[a].Charge = c
  return a
}

func SetClass(lexer smilesLexer, a int, c int) int {
  lexer.(*smilesLex).Molecule.Atoms[a].Class = c
  return a
}

func SetIsotope(lexer smilesLexer, a int, i int) int {
  lexer.(*smilesLex).Molecule.Atoms[a].Isotope = i
  return a
}

func GetBondType(a Atom, b BondType) BondType {
  if b != BOND_DEFAULT {
    return b
  }
  if a.aromatic {
    return BOND_AROMATIC
  }
  return BOND_SINGLE
}

func NewBond(m *Molecule, b BondType, a1, a2 int) {
  m.Bonds = append(m.Bonds, Bond{GetBondType(*m.Atoms[a1], b), a1, a2})
}

func AddRingbond(_lex smilesLexer, a int, b BondType, i int) int {
  lexer := _lex.(*smilesLex)
  molecule := &lexer.Molecule

  if b1, ok := lexer.pending[i]; ok {
    delete(lexer.pending, i)
    b = GetBondType(*molecule.Atoms[a], b)
    if molecule.Bonds[b1].Type != b {
      panic(fmt.Sprintf("Ringbond %v has differnt types: %v vs %v.", i, molecule.Bonds[b1].Type, b))
    }
    molecule.Bonds[b1].Atom2 = a
  } else {
    lexer.pending[i] = len(molecule.Bonds)
    NewBond(molecule, b, a, -1)
  }
  return a
}

func AppendToChain(lexer smilesLexer, c Chain, a int, b BondType) Chain {
  if len(c) > 0 {
    NewBond(&lexer.(*smilesLex).Molecule, b, c[len(c)-1], a)
  }
  c = append(c, a)
  return c
}

func AttachBranch(lexer smilesLexer, a int, c Chain, b BondType) int {
  NewBond(&lexer.(*smilesLex).Molecule, b, c[0], a)
  return a
}

func FinalizeMolecule(lexer smilesLexer) {
  if len(lexer.(*smilesLex).pending) > 0 {
    panic("Some ringbonds are still not claimed.")
  }
  molecule := &lexer.(*smilesLex).Molecule
  perAtom := make(map[int][]BondType)
  for _, B := range molecule.Bonds {
    perAtom[B.Atom1] = append(perAtom[B.Atom1], B.Type)
    perAtom[B.Atom2] = append(perAtom[B.Atom2], B.Type)
    if B.Type == BOND_DEFAULT {
      panic("Default bond made it through to Finalize(). Not good.")
    }
  }
  for a, A := range molecule.Atoms {
    if !A.aliphatic {
      continue
    }
    v := 0
    for _, bt := range perAtom[a] {
      switch bt {
      case BOND_QUADRUPLE:
        v += 4
      case BOND_TRIPLE:
        v += 3
      case BOND_DOUBLE:
        v += 2
      case BOND_SINGLE, BOND_UP, BOND_DOWN:
        v += 1
      default: //  BOND_AROMATIC
        panic("Aromatic bonds are not supported yet")
      }
      h := 0
      switch A.Element {
      case B:
        h = 3 - v
      case C:
        h = 4 - v
      case O:
        h = 2 - v
      case F, Cl, Br, I:
        h = 1 - v
      case N, P:
        if v > 3 {
          h = 5 - v
        } else {
          h = 3 - v
        }
      case S:
        if v > 4 {
          h = 6 - v
        } else if v > 2 {
          h = 4 - v
        } else {
          h = 2 - v
        }
      }
      if h < 0 {
        panic("Hydrogen count for aliphatic become <0")
      }
      A.HCount = h
    }
  }
}
