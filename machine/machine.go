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
	error_ "github.com/jonathanmarvens/turing-machine/error"
	"log"
)

type Machine struct {
	conf     *Conf
	tape     *Tape
	transMap map[MachineTransMapKey]MachineTransMapVal
}

type MachineDesc struct {
	Alphabet  []Sym   `json:"Σ"`
	InitState State   `json:"s"`
	InputSyms []Sym   `json:"x"`
	States    []State `json:"K"`
	TransMap  []struct {
		Map struct {
			Dir   Dir   `json:"dir"`
			State State `json:"state"`
			Sym   Sym   `json:"sym"`
		}
		State State `json:"state"`
		Sym   Sym   `json:"sym"`
	} `json:"δ"`
}

type MachineTransMapKey struct {
	state State
	sym   Sym
}

type MachineTransMapVal struct {
	dir   Dir
	state State
	sym   Sym
}

func NewMachine(machineDesc *MachineDesc) (*Machine, error) {
	alphabet := make(map[Sym]bool)
	for i := range machineDesc.Alphabet {
		alphabet[machineDesc.Alphabet[i]] = true
	}
	states := make(map[State]bool)
	for i := range machineDesc.States {
		states[machineDesc.States[i]] = true
	}
	states[StateValHalt] = true
	for i := range machineDesc.InputSyms {
		if !alphabet[machineDesc.InputSyms[i]] {
			return nil, error_.New(
				fmt.Sprintf("<machine> Input symbol \"%s\" isn't in the alphabet!", machineDesc.InputSyms[i]),
			)
		}
	}
	machine := Machine{
		conf: &Conf{
			currState:   machineDesc.InitState,
			currTapePos: uint(0),
		},
		tape:     NewTape(machineDesc.InputSyms),
		transMap: make(map[MachineTransMapKey]MachineTransMapVal),
	}
	for i := range machineDesc.TransMap {
		if !states[machineDesc.TransMap[i].Map.State] {
			return nil, error_.New(
				fmt.Sprintf("<machine> Transition map output state \"%s\" isn't in the set of finite states!", machineDesc.TransMap[i].Map.State),
			)
		}
		if !alphabet[machineDesc.TransMap[i].Map.Sym] {
			return nil, error_.New(
				fmt.Sprintf("<machine> Transition map output symbol \"%s\" isn't in the alphabet!", machineDesc.TransMap[i].Map.Sym),
			)
		}
		if !states[machineDesc.TransMap[i].State] {
			return nil, error_.New(
				fmt.Sprintf("<machine> Transition map input state \"%s\" isn't in the set of finite states!", machineDesc.TransMap[i].State),
			)
		}
		if !alphabet[machineDesc.TransMap[i].Sym] {
			return nil, error_.New(
				fmt.Sprintf("<machine> Transition map input symbol \"%s\" isn't in the alphabet!", machineDesc.TransMap[i].Sym),
			)
		}
		switch machineDesc.TransMap[i].Map.Dir {
		case DirMovLeft, DirMovRight, DirStay:
			break
		default:
			return nil, error_.New(
				fmt.Sprintf("<machine> \"%s\" isn't a valid direction!", machineDesc.TransMap[i].Map.Dir),
			)
		}
		machine.transMap[MachineTransMapKey{
			state: machineDesc.TransMap[i].State,
			sym:   machineDesc.TransMap[i].Sym,
		}] = MachineTransMapVal{
			dir:   machineDesc.TransMap[i].Map.Dir,
			state: machineDesc.TransMap[i].Map.State,
			sym:   machineDesc.TransMap[i].Map.Sym,
		}
	}
	return &machine, nil
}

func (self *Machine) Run() {
	self.nextTrans()
}

func (self *Machine) nextTrans() {
	currSym := self.tape.ReadSym(self.conf.currTapePos)
	nextTrans := self.transMap[MachineTransMapKey{
		state: self.conf.currState,
		sym:   currSym,
	}]
	if nextTrans.state == StateValHalt {
		_, err := fmt.Println("Machine halted!")
		if err != nil {
			log.Print(err)
		}
		_, err = fmt.Print("Tape: ", self.tape)
		if err != nil {
			log.Print(err)
		}
		return
	}
	self.conf.SetCurrState(nextTrans.state)
	if currSym != nextTrans.sym {
		self.tape.WriteSym(self.conf.currTapePos, nextTrans.sym)
	}
	switch nextTrans.dir {
	case DirMovLeft:
		self.conf.MovTapePosLeft()
	case DirMovRight:
		self.conf.MovTapePosRight()
	case DirStay:
		break
	}
	self.nextTrans()
}
