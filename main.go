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

package main

import (
	"encoding/json"
	"github.com/codegangsta/cli"
	error_ "github.com/jonathanmarvens/turing-machine/error"
	"github.com/jonathanmarvens/turing-machine/machine"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Action = func(ctx *cli.Context) {
		prog := ctx.GlobalString("prog")
		if prog == "" {
			log.Fatal(
				error_.New("--prog (or -p) required!"),
			)
		}
		file, err := os.Open(prog)
		if err != nil {
			log.Fatal(err)
		}
		jsonDec := json.NewDecoder(file)
		var machineDesc machine.MachineDesc
		err = jsonDec.Decode(&machineDesc)
		if err != nil {
			log.Fatal(err)
		}
		machine, err := machine.NewMachine(&machineDesc)
		if err != nil {
			log.Fatal(err)
		}
		machine.Run()
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			"prog,p",
			"",
			"Path to the turing machine \"program\" to run.",
		},
	}
	app.Name = "turing-machine"
	app.Usage = "Turing machinez tho."
	app.Version = "0.0.1"
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
