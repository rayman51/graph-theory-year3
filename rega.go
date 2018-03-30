package main

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
	nfastack := []*nfa{}// array of pointers

	for _, r :=  range pofix {
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

	if len(nfastack) != 1 {
		fmt.Println("Uh oh!", len(nfastack), nfastack)
	}// if

	return nfastack[0]
}// poregtonfa

func addstate(l []*state, s *state, a *state) []*state{
	l = append(l, s)

	if s != a && s.symbol == 0 {
		l = addstate(l, s.edge1, a)
		if s.edge2 != nil{
			l = addstate(l, s.edge2, a)
		}
	}// if
	return l
}

func pomatch(po string, s string) bool{
	ismatch := false
	ponfa := poregtonfa(po)

	current := []*state{}
	next := []*state{}

	current = addstate(current[:], ponfa.initial, ponfa.accept)


	for _, r := range s{
		for _, c := range current{
			if c.symbol == r {
				next = addstate(next[:], c.edge1, ponfa.accept)
			}// if
		}// inner for
		current, next = next, []*state{}
	}// for
	
	for _, c := range current {
		if c == ponfa.accept{
			ismatch = true
			break
		}// if
	}// for


	return ismatch
}// pomatch
func main(){
	fmt.Println(pomatch("ab.c*|", "ccccc"))
	fmt.Println(pomatch("ab.c*|", "abc"))
	fmt.Println(pomatch("ab.c*|", ""))
	fmt.Println(pomatch("ab.c*|", "ccccccccccccc"))
	fmt.Println(pomatch("ab.c*|", "def"))


}// main