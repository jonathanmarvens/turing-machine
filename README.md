# Turing machinez tho.

This is a simple turing machine I created for an article I’ve been writing titled, “*Neurons, turing machines, & lambda abstractions: an interesting combo?*”

## Setup.

There are 3 ways you can install this on your system in order to play with it.

-----

__(1) Download a pre-built binary.__

You should be able to grab a pre-built binary on the [__Releases__ page](https://github.com/jonathanmarvens/turing-machine/releases) (oh, the power of Go cross-compiling __:)__ … I’m looking at you, GCC).

-----

Assuming you have your [Go workspace](http://golang.org/doc/code.html) all set up, the following are your two other options.

-----

__(2) Using `go get`.__

```sh
go get -t -u github.com/jonathanmarvens/turing-machine
```

-----

__(3) Building from source.__

```sh
go get -t -u github.com/kisielk/errcheck github.com/mitchellh/gox github.com/tools/godep
git clone https://github.com/jonathanmarvens/turing-machine.git $GOPATH/src/github.com/jonathanmarvens/turing-machine
cd $GOPATH/src/github.com/jonathanmarvens/turing-machine
make 
```

Now look in the __.bin/__ directory.

-----

## Usage.

Running `turing-machine help` should be helpful enough about the usage.

-----

Feel free to use the example *programs* in [__examples/__](https://github.com/jonathanmarvens/turing-machine/tree/master/examples).

```sh
# ./.bin/{OS}-{ARCH}/turing-machine --prog="./examples/n-plus-1.btm.json"
```

## Author.

__Jonathan Barronville__ < [__http://乔纳森.com__](http://乔纳森.com) > ( *jonathan@belairlabs.com* )

## Acknowledgements.

Here’s a great document that includes a formal (mathematical) definition of a turing machine (from Cornell University’s CS department): [http://www.cs.cornell.edu/courses/cs4820/2012su/handouts/turingm.pdf](http://www.cs.cornell.edu/courses/cs4820/2012su/handouts/turingm.pdf) … I liked this document more compared to like 6 others I read (the math is clearer and pretty straightforward to follow, there are real examples, the author doesn’t make stupid assumptions about things __-.-__, *et cetera*), so I (loosely) modeled the turing machine around it. You should read it if you want a better mathematical understanding of turing machines! Please don’t rely on Wikipedia for a good understanding of turing machines.

Here’s a good document, which also includes a formal definition of a turing machine: [http://plato.stanford.edu/entries/turing-machine](http://plato.stanford.edu/entries/turing-machine) … although I didn’t particularly love this document, it’s pretty informative (you probably wanna read it first if you’re new to turing machines).

## License.

See [__LICENSE__](https://github.com/jonathanmarvens/turing-machine/blob/master/LICENSE).
