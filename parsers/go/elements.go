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

type Element int

const (
  H = Element(iota + 1)
  He
  Li
  Be
  B
  C
  N
  O
  F
  Ne
  Na
  Mg
  Al
  Si
  P
  S
  Cl
  Ar
  K
  Ca
  Sc
  Ti
  V
  Cr
  Mn
  Fe
  Co
  Ni
  Cu
  Zn
  Ga
  Ge
  As
  Se
  Br
  Kr
  Rb
  Sr
  Y
  Zr
  Nb
  Mo
  Tc
  Ru
  Rh
  Pd
  Ag
  Cd
  In
  Sn
  Sb
  Te
  I
  Xe
  Cs
  Ba
  La
  Ce
  Pr
  Nd
  Pm
  Sm
  Eu
  Gd
  Tb
  Dy
  Ho
  Er
  Tm
  Yb
  Lu
  Hf
  Ta
  W
  Re
  Os
  Ir
  Pt
  Au
  Hg
  Tl
  Pb
  Bi
  Po
  At
  Rn
  Fr
  Ra
  Ac
  Th
  Pa
  U
  Np
  Pu
  Am
  Cm
  Bk
  Cf
  Es
  Fm
  Md
  No
  Lr
  Rf
  Db
  Sg
  Bh
  Hs
  Mt
  Ds
  Rg
  Cn
  Fl = Element(114)
  Lv = Element(116)
)

var NameToElement = map[string]Element{
  "H":  H,
  "He": He,
  "Li": Li,
  "Be": Be,
  "B":  B,
  "C":  C,
  "N":  N,
  "O":  O,
  "F":  F,
  "Ne": Ne,
  "Na": Na,
  "Mg": Mg,
  "Al": Al,
  "Si": Si,
  "P":  P,
  "S":  S,
  "Cl": Cl,
  "Ar": Ar,
  "K":  K,
  "Ca": Ca,
  "Sc": Sc,
  "Ti": Ti,
  "V":  V,
  "Cr": Cr,
  "Mn": Mn,
  "Fe": Fe,
  "Co": Co,
  "Ni": Ni,
  "Cu": Cu,
  "Zn": Zn,
  "Ga": Ga,
  "Ge": Ge,
  "As": As,
  "Se": Se,
  "Br": Br,
  "Kr": Kr,
  "Rb": Rb,
  "Sr": Sr,
  "Y":  Y,
  "Zr": Zr,
  "Nb": Nb,
  "Mo": Mo,
  "Tc": Tc,
  "Ru": Ru,
  "Rh": Rh,
  "Pd": Pd,
  "Ag": Ag,
  "Cd": Cd,
  "In": In,
  "Sn": Sn,
  "Sb": Sb,
  "Te": Te,
  "I":  I,
  "Xe": Xe,
  "Cs": Cs,
  "Ba": Ba,
  "La": La,
  "Ce": Ce,
  "Pr": Pr,
  "Nd": Nd,
  "Pm": Pm,
  "Sm": Sm,
  "Eu": Eu,
  "Gd": Gd,
  "Tb": Tb,
  "Dy": Dy,
  "Ho": Ho,
  "Er": Er,
  "Tm": Tm,
  "Yb": Yb,
  "Lu": Lu,
  "Hf": Hf,
  "Ta": Ta,
  "W":  W,
  "Re": Re,
  "Os": Os,
  "Ir": Ir,
  "Pt": Pt,
  "Au": Au,
  "Hg": Hg,
  "Tl": Tl,
  "Pb": Pb,
  "Bi": Bi,
  "Po": Po,
  "At": At,
  "Rn": Rn,
  "Fr": Fr,
  "Ra": Ra,
  "Ac": Ac,
  "Th": Th,
  "Pa": Pa,
  "U":  U,
  "Np": Np,
  "Pu": Pu,
  "Am": Am,
  "Cm": Cm,
  "Bk": Bk,
  "Cf": Cf,
  "Es": Es,
  "Fm": Fm,
  "Md": Md,
  "No": No,
  "Lr": Lr,
  "Rf": Rf,
  "Db": Db,
  "Sg": Sg,
  "Bh": Bh,
  "Hs": Hs,
  "Mt": Mt,
  "Ds": Ds,
  "Rg": Rg,
  "Cn": Cn,
  "Fl": Fl,
  "Lv": Lv,

  // aromatic
  "b":  -B,
  "c":  -C,
  "n":  -N,
  "o":  -O,
  "p":  -P,
  "s":  -S,
  "as": -As,
  "se": -Se,
}

var Weights = []float32{0, 1.007948, 4.002602, 6.941222222, 9.012182333333, 10.811777778, 12.0107888889, 14.0067222222, 15.9994333333, 18.9984032555556, 20.1797666667, 22.98976928222222, 24.3050666667, 26.9815386888889, 28.0855333333, 30.973762222222, 32.065555556, 35.453222222, 39.948111111, 39.0983111111, 40.078444444, 44.955912666667, 47.867111111, 50.9415111111, 51.9961666667, 54.938045555556, 55.845222222, 58.933195555556, 58.6934444444, 63.546333333, 65.38222222, 69.723111111, 72.63111111, 74.92160222222, 78.96333333, 79.904111111, 83.798222222, 85.4678333333, 87.62111111, 88.90585222222, 91.224222222, 92.90638222222, 95.96222222, 98.0, 101.07222222, 102.90550222222, 106.42111111, 107.8682222222, 112.411888889, 114.818333333, 118.710777778, 121.760111111, 127.60333333, 126.90447333333, 131.293666667, 132.9054519222222, 137.327777778, 138.90547777778, 140.116111111, 140.90765222222, 144.242333333, 145.0, 150.36222222, 151.964111111, 157.25333333, 158.92535222222, 162.500111111, 164.93032222222, 167.259333333, 168.93421222222, 173.054555556, 174.9668111111, 178.49222222, 180.94788222222, 183.84111111, 186.207111111, 190.23333333, 192.217333333, 195.085, 196.966569444444, 200.59222222, 204.3833222222, 207.2111111, 208.98040111111, 209.0, 210.0, 222.0, 223.0, 226.0, 227.0, 232.03806222222, 231.03588222222, 238.02891333333, 237.0, 244.0, 243.0, 247.0, 247.0, 251.0, 252.0, 257.0, 258.0, 259.0, 262.0, 267.0, 268.0, 269.0, 270.0, 269.0, 278.0, 281.0, 281.0, 285.0, 286.0, 289.0, 288.0, 293.0, 294.0, 294.0}
