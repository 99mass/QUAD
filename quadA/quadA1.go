// package main

package piscine

import "github.com/01-edu/z01"

func QuadA1(x,y int) {
	// je verifie ici si x et y sont positive
    if x > 0 && y > 0 {
		// si ils sont positive on boucle sur le y
        for i := 0; i < y; i++ {
			// on verifie si on est sur la ligne 1 ou 3 
            if i == 0 || i == y-1 {
				// on affiche le premiere o  de la ligne 1
                z01.PrintRune('o')
				// ici on fait le boucle pour pouvoir afficher les 3  --- qui suit  le premiere o
                  for j := 0; j < x-2; j++ {
                    z01.PrintRune('-')
                }
				// on affiche le derniere o de la ligne 
                z01.PrintRune('o')

            } else {
                // ici on return a la ligne avant d'afficher la deuxieme ligne
                z01.PrintRune('\n')
				// ici on est sur la 2émé ligne
				// on affiche le | 
                z01.PrintRune('|')
				// ici on fait le boucle pour pouvoir afficher le space 3 fois
                for j := 0; j < x-2; j++ {
                    z01.PrintRune(' ')
                }
				// on affiche le derniere | de la ligne 
                z01.PrintRune('|')
				// ici on return a la ligne
                z01.PrintRune('\n')
            }
        }
        z01.PrintRune('\n')
    }
}


// func main() {
// 	QuadA1(5,3)
// }