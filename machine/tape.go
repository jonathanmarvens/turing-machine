/**
 * The MIT License (MIT).
 *
 * https://github.com/jonathanmarvens/turing-machine
 *
 * Copyright (c) 2014 Jonathan Barronville (jonathan@belairlabs.com)
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package machine

import (
	"fmt"
)

type Tape struct {
	syms []Sym
}

func NewTape(syms []Sym) *Tape {
	return &Tape{
		syms: syms,
	}
}

func (self *Tape) ReadSym(i uint) Sym {
	if int(i) >= cap(self.syms) {
		return Sym("")
	}
	return self.syms[i]
}

func (self *Tape) String() string {
	var val string
	for i := range self.syms {
		switch i {
		case 0:
			val += fmt.Sprintf("[ %s | ", self.syms[i])
		case (len(self.syms) - 1):
			val += fmt.Sprintf("%s ]\n", self.syms[i])
		default:
			val += fmt.Sprintf("%s | ", self.syms[i])
		}
	}
	return val
}

func (self *Tape) WriteSym(i uint, sym Sym) {
	if int(i) >= cap(self.syms) {
		syms := make([]Sym, len(self.syms), (int(i) + 1))
		copy(syms, self.syms)
		self.syms = syms
	}
	self.syms[int(i)] = sym
}
