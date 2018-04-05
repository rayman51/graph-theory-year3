package main

import (
	"fmt" // allows user input
	"os" // exits programm
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
			frag2 := nfastack[len(nfastack)-1] //pops the last nfa off the stack
			nfastack = nfastack[:len(nfastack)-1] // sets stack to accept all but last element
			
			frag1 := nfastack[len(nfastack)-1] //pops the last nfa off the stack
			nfastack = nfastack[:len(nfastack)-1] // sets stack to accept all but last element

			frag1.accept.edge1 = frag2.initial // frag1 points to frag2

			nfastack = append(nfastack, &nfa{initial: frag1.initial, accept: frag2.accept}) //sets initial state to that of frag1
																						   //and the accept state of frag2
			// case .
		case '|':
			frag2 := nfastack[len(nfastack)-1] //pops the last nfa off the stack
			nfastack = nfastack[:len(nfastack)-1]  // sets stack to accept all but last element
			
			frag1 := nfastack[len(nfastack)-1] //pops the last nfa off the stack
			nfastack = nfastack[:len(nfastack)-1] // sets stack to accept all but last element

			accept := state{} // becomes an empty state
			initial := state{edge1: frag1.initial, edge2: frag2.initial} //initial state points to frag1 & frag2 initial states
			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept

			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept}) // set the initial and accept state
												 								 // to the new states
			// case |
		case '*':
			frag := nfastack[len(nfastack)-1] //pops the last nfa off the stack
			nfastack = nfastack[:len(nfastack)-1] // sets stack to accept all but last element

			accept := state{} // becomes an empty state
			initial := state{edge1: frag.initial, edge2: &accept}
			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept

			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept}) // pushs to the nfa stack
			// case *
			case '+':
			frag := nfastack[len(nfastack)-1] //pops the last nfa off the stack
			
			accept := state{}// becomes an empty state
			initial := state{edge1: frag.initial, edge2: &accept}
			frag.accept.edge1 = &initial

			nfastack = append(nfastack, &nfa{initial: frag.initial, accept: &accept}) // pushs to the nfa stack
			// case +
		default:
			accept := state{} //empty state for accept
			initial := state{symbol: r ,edge1: &accept}

			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept}) // pushs to the nfa stack
		}/// switch
	}// for 

	/*if len(nfastack) != 1 {
		fmt.Println("Uh oh!", len(nfastack), nfastack)
	}// if*/

	return nfastack[0]
}// poregtonfa

func addstate(l []*state, s *state, a *state) []*state{
	l = append(l, s)

	if s != a && s.symbol == 0 {
		l = addstate(l, s.edge1, a)// keep running method

		if s.edge2 != nil{ // if edge2 is not nil run again
			l = addstate(l, s.edge2, a)// keep running method
		}
	}// if
	return l
}// addstate

func pomatch(po string, s string) bool{
	ismatch := false // value for return

	ponfa := poregtonfa(po)// pofix to nfa reg expression

	current := []*state{}
	next := []*state{}

	current = addstate(current[:], ponfa.initial, ponfa.accept)// adds all initial states to the array

	for _, r := range s{// loop to run through the string
		for _, c := range current{
			if c.symbol == r {
				next = addstate(next[:], c.edge1, ponfa.accept)
			}// if
		}// inner for
		current, next = next, []*state{}// replaces current array with next array
	}// for
	
	for _, c := range current {// loop to run through the state
		if c == ponfa.accept{
			ismatch = true// if current state matches accept state return true
			break
		}// if
	}// for


	return ismatch// return statement
}// pomatch
// intopost method for shunting taken from shunt.go
func intopost(infix string) string{
	specials := map[rune]int{'*': 10,'.': 9,'|': 8}

	pofix , s:= []rune{}, []rune{}

	for _, r := range infix {
		switch {
		case r == '(':
		s = append(s, r)// add r to stack
		//case (
		case r == ')':
			for s[len(s)-1] != '(' {// pop off the stack till ( is found
				pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
			}// for
		s = s[:len(s)-1]// add all elements but last to stack
		//case )
		case specials[r] > 0:
			for len(s) > 0 && specials[r] <= specials[s[len(s)-1]] {
				pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]// add top of stack to pofix
			}// for
			s = append(s,r)
		//case specials[r]
		default:
			pofix = append(pofix,r)// add other characters to pofix rune 
		}// switch

	}// for
	for len(s) > 0 {
		pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]// add everything except the last element to the stack
	}// for

	return string(pofix)// return statement
}// intopost


func runProg(){
	var condition string
	var test string
	
	fmt.Print("Please enter condition: ")// user prompt for condition
	fmt.Scanln(&condition) 

	for _, r := range condition {// loops through condition
		if r == '(' || r == ')'{// if string has brackets 
			condition = intopost(condition)// change infix to pofix
			break
		}// if

	}// for
	fmt.Print("Please enter a string: ")// user prompt for condition
	fmt.Scanln(&test)
	fmt.Print("=========================\n")
	if pomatch(condition, test) == false {
		fmt.Printf("String is not accepted\n")// if test is false
	} else {
		fmt.Printf("String is accepted\n")// if test is true
	}
	
}// runProg
func main(){
	
	for {
		var option string // var for switch

		fmt.Println("Enter an Option")// user menu
		fmt.Println("================")
		fmt.Println("Press 1 to run a test")
		fmt.Println("Press 2 to exit ")
		fmt.Println("================")
		fmt.Print("option : ")
		fmt.Scanln(&option)

		switch option {
		case "1":
			runProg()//runs the test
		case "2":
			os.Exit(3)//Use os.Exit to immediately exit with a given status.
		default:
			fmt.Println("Invalid input press 1 to test or 2 to exit")
		}// switch

	}// for
}// main