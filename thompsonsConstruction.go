package main 

import (
	"fmt"
)

type state struct {
	symbol rune
	edge1 *state
	edge2 *state
}

type nfa struct {
	initial *state
	accept *state
}

func poregtonfa(pofix string) *nfa{
	nfastack := []*nfa{}

	for _, r := range pofix {
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
		case'|':
			// Takes 2 fragments off stack
			frag2 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			frag1 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			// Join accept states to initial states
			accept := state{}
			initial := state{edge1: frag1.initial, edge2: frag2.initial}
			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept

			// Push a new fragment onto the nfa stack
			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
		case'*':
			// Takes fragment off stack
			frag := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			accept := state{}
			initial := state{edge1: frag.initial, edge2: &accept}
			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept

			// Push a new fragment to the nfa stack
			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
		default:
			accept := state{}
			initial := state{symbol: r, edge1: &accept}

			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
		}
	}

	if len(nfastack) != 1 {
		fmt.Println("Nope:", len(nfastack), nfastack)
	}
	return nfastack[0]
}

func addState(l []*state, s *state, a *state ) []*state {
	l = append(l, s)

	if s != a && s.symbol == 0 {
		l = addState(l, s.edge1, a)
		if s.edge2 != nil {
			l = addState(l, s.edge2, a)
		}
	}
	return l;
}

func pomatch(po string, s string) bool{
	ismatch := false
	ponfa := poregtonfa(po)

	current := []*state{}
	next := []*state{}

	current = addState(current[:], ponfa.initial, ponfa.accept)

	for _, r:= range s {
		for _, c := range current {
			if c.symbol == r {
				next = addState(next[:], c.edge1, ponfa.accept)
			}

		}
		current, next = next, []*state{}
	}

	for _, c := range current {
		if c == ponfa.accept {
			ismatch = true
			break
		}
	}

	return ismatch
}

func main() {
	fmt.Println(pomatch("ab.c*|", "cccc"))
}