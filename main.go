package main

import (
	"fmt"
)

// Shunting Yard
// Transforms regular expressions from infix to postfix
func intopost(infix string) string {

	// Maps characters into integers to keep track to special characters
	specials := map[rune]int{'*': 10, '.': 9, '|': 8, '+': 11, '?': 12}

	postfix, s := []rune{}, []rune{}

	// Loop over the infix and return index of read characters
	for _, r := range infix {
		switch {
		case r == '(':
			// Put at the end of the stack
			s = append(s, r)
		case r == ')':
			// While last character on stack isn't an open bracket append to the top of stack
			for s[len(s)-1] != '(' {
				postfix, s = append(postfix, s[len(s)-1]), s[:len(s)-1]
			}
			// Everything up to the bottom of the stack 
			s = s[:len(s)-1]
		case specials[r] > 0:
			// While character is on the stack and less than what is at the end, take element from top and put it into the end of the stack
			for len(s) > 0 && specials [r] <= specials[s[len(s)-1]] {
				// Pop the elements off top of stack and append to postfix
				postfix, s = append(postfix, s[len(s)-1]), s[:len(s)-1]
			}
			s = append(s, r)
		default:
			// Puts r to the end of the postfix
			postfix = append(postfix, r)
		}
	}

	// If anything is on the top of the stack append it to the output
	for len(s) > 0 {
		postfix, s = append(postfix, s[len(s)-1]), s[:len(s)-1]
	}

	return string(postfix)
}

// Thompsons Construction
// Stores states to other structs
type state struct {
	symbol rune 
	edge1 *state
	edge2 *state
}

type nfa struct {
	initial *state
	accept  *state
}

// Input is postfix, return pointer to nfa structs
func poRegtonfa(postfix string) *nfa {
	// Provides an array of pointers to empty nfa
	nfastack := []*nfa{}	
	
		// Loop through rune for each character
		for _, r := range postfix {
			switch r {
			case '.':
				// Takes 2 fragments off stack
				frag2 := nfastack[len(nfastack)-1]
				nfastack = nfastack[:len(nfastack)-1]
				frag1 := nfastack[len(nfastack)-1]
           		nfastack = nfastack[:len(nfastack)-1]

				// Joins the accept state of frag1 to the initial state of frag2
           		 frag1.accept.edge1 = frag2.initial

				// Push a new fragment onto the nfa stack 
				nfastack = append(nfastack, &nfa{initial: frag1.initial, accept: frag2.accept}) 
			case '|':
				// Takes 2 fragments off stack
				frag2 := nfastack[len(nfastack)-1]
				nfastack = nfastack[:len(nfastack)-1]
       		    frag1 := nfastack[len(nfastack)-1]
            	nfastack = nfastack[:len(nfastack)-1]
			
				// Join accept states to initial states
				initial := state{edge1: frag1.initial, edge2: frag2.initial}
            	accept := state{}
				frag1.accept.edge1 = &accept
            	frag2.accept.edge1 = &accept            

				// Push a new fragment onto the nfa stack 
            	nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
			case '*':
				// Takes fragment off stack
				frag := nfastack[len(nfastack)-1]
				nfastack := nfastack[:len(nfastack)-1]
				
				accept := state{}
				initial := state{edge1: frag.initial, edge2: &accept}
				frag.accept.edge1 = frag.initial
				frag.accept.edge2 = &accept
	
				// Push a new fragment to the nfa stack
				nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})	
			case '+':
				// Takes fragment off stack
				frag := nfastack[len(nfastack)-1]
				nfastack = nfastack[:len(nfastack)-1]

				accept := state{}
				initial := state{edge1: frag.initial}
				middle := state{edge1: frag.initial, edge2: &accept}
				frag.accept.edge1 = &middle

				// Push a new fragment to the nfa stack
				nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
			case '?':
				// Pop a fragment off the stack.
				frag := nfastack[len(nfastack)-1]
				nfastack = nfastack[:len(nfastack)-1]

				accept := state{}
				initial := state{edge1: frag.initial, edge2: &accept}
				frag.accept.edge1 = &accept

				// Push a new fragment to the nfa stack
				nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
			default:
				accept := state{}
				initial := state{symbol: r, edge1: &accept}
				nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
			}
		}

	// If there's more than one element on the stack
	if len(nfastack) != 1 {
		fmt.Println("\nError: ", len(nfastack), nfastack)
	}

	return nfastack[0]
}

// List of pointers to state, single pointer to state and accept state
func addState(l []*state, s *state, a *state) []*state {
	l = append(l, s)
	// If s is not an accept state and the state has no value
	if s != a && s.symbol == 0 {
		// Pass to list l
		l = addState(l, s.edge1, a)
		if s.edge2 != nil {
			// Pass to list l
			l = addState(l, s.edge2, a)
		}
	}
	return l
}

// Function to check if regexp matches string
func pomatch(userExp string, input string) bool {
	// Set boolean false 
	isMatch := false
	// New variable from poRegtonfa called
	ponfa := poRegtonfa(userExp)

	// Array of pointers to current and next states 
	current := []*state{}
	next := []*state{}

	// Passes current to addState function
	current = addState(current[:], ponfa.initial, ponfa.accept)

	// Loop through input for each character
	for _, r := range input {
		for _, s := range current {
			if s.symbol == r {
				// Add s and any state to next array
				next = addState(next[:], s.edge1, ponfa.accept)
			}
		}
		// Move from current state to the next state
		current, next = next, []*state{}
	}

	for _, s := range current {
		if s == ponfa.accept {
			// Set boolean to true to accept state
			isMatch = true
			break
		}
	}

	return isMatch
}

// UI Menu
func main() {
	var userExp, poReg, userString string

	// Header
	fmt.Print("Graph Theory Project 2018 - G00326176")
	fmt.Print("\nEnter a Regexp: ")
	fmt.Scanln(&userExp)

	// Convert expression to postfix
	poReg = intopost(userExp)
	fmt.Println("\nPostfix: ", poReg)

	fmt.Print("\nEnter String to See if it Matches Your NFA: ")
	fmt.Scan(&userString)

	// Check if string matches regexp
	fmt.Println("Does Your String Match? ",pomatch(poReg, userString))
}