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

import (
  "log"
  "unicode/utf8"
)

const eof = 0

// The parser uses the type <prefix>Lex as a lexer.  It must provide
// the methods Lex(*<prefix>SymType) int and Error(string).
type smilesLex struct {
  line     []byte      // Line that is not parsed yet
  pending  map[int]int // Maps [yet] unclaimed ringbond to index of bond in Molecule.Bonds
  Molecule Molecule    // Output
}

// The parser calls this method to get each new token.  This
// implementation returns operators and NUM.
func (x *smilesLex) Lex(yylval *smilesSymType) int {
  for {
    c := x.next()
    if c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z' {
      return x.element(c, yylval)
    }
    switch c {
    case eof:
      return eof
    case '=', '#', '$', ':', '/', '\\':
      return x.bond(c, yylval)
    case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
      yylval.number = int(c - '0')
      return digit
    case '*':
      return wildcard
    case '.':
      return dot
    case '@':
      return x.chirality(yylval)
    case '+', '(', ')', '[', ']', '%':
      return int(c)
    case '-':
      return minus
    default:
      return error1
    }
  }
}

// Lex chirality that consists of several bytes
func (x *smilesLex) chirality2(c1, c2, d2 rune, yylval *smilesSymType) int {
  var d1 rune = '0'
  var max int
  switch {
  case c1 == 'T' && c2 == 'H':
    yylval.chirality = Chirality{Type: CHIRALITY_TH}
    max = 2
  case c1 == 'A' && c2 == 'L':
    yylval.chirality = Chirality{Type: CHIRALITY_AL}
    max = 2
  case c1 == 'S' && c2 == 'P':
    yylval.chirality = Chirality{Type: CHIRALITY_SP}
    max = 3
  case c1 == 'T' && c2 == 'B':
    yylval.chirality = Chirality{Type: CHIRALITY_TB}
    d1 = d2
    d2 = x.next()
    max = 20
  case c1 == 'O' && c2 == 'H':
    yylval.chirality = Chirality{Type: CHIRALITY_OH}
    d1 = d2
    d2 = x.next()
    max = 30
  default:
    return error1
  }
  idx := int(d1-'0')*10 + int(d2-'0')
  if idx < 1 || idx > max {
    return error1
  }
  yylval.chirality.Index = idx
  return chirality
}

// Lex chirality
func (x *smilesLex) chirality(yylval *smilesSymType) int {
  switch x.peek() {
  case 'T', 'A', 'S', 'O':
    return x.chirality2(x.next(), x.next(), x.next(), yylval)
  case '@':
    x.next()
    yylval.chirality = Chirality{Type: CHIRALITY_CW}
  default:
    yylval.chirality = Chirality{Type: CHIRALITY_CCW}
  }
  return chirality
}

// Lex an element.
func (x *smilesLex) element(c rune, yylval *smilesSymType) int {
  var element Element
  var ok bool
  if element, ok = NameToElement[string([]rune{c, x.peek()})]; !ok {
    if element, ok = NameToElement[string([]rune{c})]; !ok {
      return error1
    }
  } else {
    x.next()
  }
  switch element {
  case H:
    yylval.element = element
    return hydrogen
  case B, C, N, O, S, P, F, Cl, Br, I:
    yylval.element = element
    return aliphatic_organic
  case -B, -C, -N, -O, -S, -P:
    yylval.element = -element
    return aromatic_organic
  case -Se, -As:
    yylval.element = -element
    return aromatic_other
  default:
    yylval.element = element
  }
  return other
}

// Lex a bond.
func (x *smilesLex) bond(c rune, yylval *smilesSymType) int {
  switch c {
  case '=':
    yylval.bond = BOND_DOUBLE
  case '#':
    yylval.bond = BOND_TRIPLE
  case '$':
    yylval.bond = BOND_QUADRUPLE
  case ':':
    yylval.bond = BOND_AROMATIC
  case '/':
    yylval.bond = BOND_UP
  case '\\':
    yylval.bond = BOND_DOWN
  }

  // This is a fix for syntax ambiguity - look ahead *in the lexer*
  switch x.peek() {
  case '%', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
    return ringbond
  }
  return chainbond
}

// Return the next rune but don't advance the read position.
func (x *smilesLex) peek() rune {
  if len(x.line) == 0 {
    return eof
  }
  c, _ := utf8.DecodeRune(x.line)
  return c
}

// Return the next rune for the lexer.
func (x *smilesLex) next() rune {
  if len(x.line) == 0 {
    return eof
  }
  c, size := utf8.DecodeRune(x.line)
  x.line = x.line[size:]
  if c == utf8.RuneError {
    log.Print("invalid utf8")
    return x.next()
  }
  return c
}

// The parser calls this method on a parse error.
func (x *smilesLex) Error(s string) {
  log.Printf("parse error: %s", s)
}
