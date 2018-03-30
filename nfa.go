/*package main

import (
	"fmt"
)// imports

type state struct{
	symbol rune
	edge1 *state
	edge2 *state
}// state struct

type nfa struct{
	initial *state
	accept  *state
}// nfa struct



func poregtonfa(pofix string) *nfa{
	nfastack := []*nfa{}

	for _, r :=  range pofix{
		switch r {
		case '.':
			frag2 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			frag1 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			frag1.accept.edge1 = frag2.initial

			nfastack = append(nfastack, &nfa{initial: frag1.initial, accept: frag2.accept})
			// case .
		case '|':
			frag2 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			frag1 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			accept := state{}
			initial := state{edge1: frag1.initial, edge2: frag2.initial}
			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept

			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})

			// case |
		case '*':
			frag := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			accept := state{}
			initial := state{edge1: frag.initial, edge2: &accept}
			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept

			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
			// case *
		default:
			accept := state{}
			initial := state{symbol: r ,edge1: &accept}

			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
		}/// switch
	}// for 

	

	return nfastack[0]
}// poregtonfa

func main(){
	nfa := poregtonfa("ab.c*|")
	fmt.Println(nfa)
}*/