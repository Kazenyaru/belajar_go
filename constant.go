package main

import "fmt"

const (
	a = iota
	b = iota
	c = iota
)

const (
	isAdmin = 1 << iota
	isHeadquarters
	canSeeFinancials

	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

func runConstant() {
	var myRoles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%v\b", myRoles)
}
