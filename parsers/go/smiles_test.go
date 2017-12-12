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
package smiles_test

import (
  "fmt"
  "code.google.com/p/opensmiles"
  "sort"
)

type Bond struct {
  a, b int
}

type BondSlice []Bond

func (b BondSlice) Len() int { return len(b) }
func (b BondSlice) Less(i, j int) bool {
  if b[i].a == b[j].a {
    return b[i].b < b[j].b
  }
  return b[i].a < b[j].a
}
func (b BondSlice) Swap(i, j int) { b[i], b[j] = b[j], b[i] }

func ParseAndPrint(src string) {
  molecule, err := smiles.Parse(src)
  if err != nil {
    fmt.Println("Error happened", err)
    return
  }
  var weights []int
  var bonds BondSlice
  for _, atom := range molecule.Atoms {
    weights = append(weights, atom.Isotope)
  }
  for _, b := range molecule.Bonds {
    if b.Atom1 > b.Atom2 {
      bonds = append(bonds, Bond{b.Atom2, b.Atom1})
    } else {
      bonds = append(bonds, Bond{b.Atom1, b.Atom2})
    }
  }
  sort.Sort(bonds)
  fmt.Println("Weights:", weights)
  fmt.Println("Bonds:", bonds)
}

func ExampleParse() {
  src := "C1=CC=CC=C1"
  ParseAndPrint(src)
  // Output:
  // Weights: [12 12 12 12 12 12]
  // Bonds: [{0 1} {0 5} {1 2} {2 3} {3 4} {4 5}]
}
