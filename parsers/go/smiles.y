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
%{package smiles
%}

%union {
  char rune
  number int
  bond BondType
  chain Chain
  chirality Chirality
  element Element
}

%token <bond> ringbond chainbond dot
%token <chirality> chirality
%token <element> hydrogen aliphatic_organic aromatic_organic aromatic_other other wildcard
%token <char> minus error1
%token <number> digit

%type <element> ELEMENT_SYMBOLS AROMATIC_SYMBOLS
%type <number> ATOM RINGED_ATOM BRANCHED_ATOM BRACKET_ATOM BRACKET_ATOM0 BRACKET_ATOM1 BRACKET_ATOM2 BRACKET_ATOM3 BRACKET_ATOM4
%type <chain> CHAIN
%type <number> ISOTOPE HCOUNT CLASS NUMBER CHARGE

%%

top: CHAIN { FinalizeMolecule(smileslex)};

CHAIN : BRANCHED_ATOM                  { $$ = AppendToChain(smileslex, NewChain(), $1, BOND_DEFAULT)}
      | CHAIN BRANCHED_ATOM            { $$ = AppendToChain(smileslex, $1, $2, BOND_DEFAULT) }
      | CHAIN minus BRANCHED_ATOM      { $$ = AppendToChain(smileslex, $1, $3, BOND_SINGLE) }
      | CHAIN chainbond BRANCHED_ATOM  { $$ = AppendToChain(smileslex, $1, $3, $2) }
      | CHAIN dot BRANCHED_ATOM        { $$ = AppendToChain(smileslex, $1, $3, $2) }

BRACKET_ATOM0 : ELEMENT_SYMBOLS     { $$ = NewAtom(smileslex, $1, false, false) }
              | AROMATIC_SYMBOLS    { $$ = NewAtom(smileslex, $1, true, false) }
;

BRACKET_ATOM1 : "[" BRACKET_ATOM0            { $$ = $2 }
              | "[" ISOTOPE BRACKET_ATOM0    { $$ = SetIsotope(smileslex, $3, $2) }
;

BRACKET_ATOM2 : BRACKET_ATOM1
              | BRACKET_ATOM1 chirality { $$ = SetChirality(smileslex, $1, $2) }
;

BRACKET_ATOM3 : BRACKET_ATOM2
              | BRACKET_ATOM2 HCOUNT    { $$ = SetHCount(smileslex, $1, $2) }
;

BRACKET_ATOM4 : BRACKET_ATOM3
              | BRACKET_ATOM3 CHARGE    { $$ = SetCharge(smileslex, $1, $2) }
;

BRACKET_ATOM : BRACKET_ATOM4 "]"        { $$ = $1 }
             | BRACKET_ATOM4 CLASS "]"  { $$ = SetClass(smileslex, $1, $2) }
;

ATOM : BRACKET_ATOM
     | aliphatic_organic   { $$ = NewAtom(smileslex, $1, false, true) }
     | aromatic_organic    { $$ = NewAtom(smileslex, $1, true, false) }
     | wildcard            { $$ = NewAtom(smileslex, $1, false, false) }
;

RINGED_ATOM : ATOM
            | RINGED_ATOM digit                     { $$ = AddRingbond(smileslex, $1, BOND_DEFAULT, $2) }
            | RINGED_ATOM '%' digit digit           { $$ = AddRingbond(smileslex, $1, BOND_DEFAULT, $3 * 10 + $4) }
            | RINGED_ATOM ringbond  digit           { $$ = AddRingbond(smileslex, $1, $2, $3) }
            | RINGED_ATOM ringbond '%' digit digit  { $$ = AddRingbond(smileslex, $1, $2, $4 * 10 + $5) }
;


BRANCHED_ATOM : RINGED_ATOM
              | BRANCHED_ATOM "(" CHAIN ")"            { $$ = AttachBranch(smileslex, $1, $3, BOND_DEFAULT)}
              | BRANCHED_ATOM "(" minus CHAIN ")"      { $$ = AttachBranch(smileslex, $1, $4, BOND_SINGLE)}
              | BRANCHED_ATOM "(" chainbond CHAIN ")"  { $$ = AttachBranch(smileslex, $1, $4, $3)}
              | BRANCHED_ATOM "(" dot CHAIN ")"        { $$ = AttachBranch(smileslex, $1, $4, $3)}
;

ELEMENT_SYMBOLS : hydrogen | aliphatic_organic | other ;
AROMATIC_SYMBOLS : aromatic_organic | aromatic_other ;

ISOTOPE : NUMBER ;
HCOUNT : hydrogen        { $$ = 1 }
       | hydrogen NUMBER { $$ = $2 }
;

CHARGE : "+"               { $$ = 1 }
       | "+" "+"           { $$ = 2 }
       | "+" digit         { $$ = int($2) }
       | "+" digit digit   { $$ = int($2*10+$3) }
       | minus             { $$ = -1 }
       | minus minus       { $$ = -2 }
       | minus digit digit { $$ = -int($2*10+$3) }
;

CLASS : NUMBER ;

NUMBER : digit
       | NUMBER digit  { $$ = $1 * 10 + $2 }
;

%%
