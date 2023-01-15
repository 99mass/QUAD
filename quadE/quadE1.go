// package main

package piscine

import "github.com/01-edu/z01"


func QuadE1(x,y int) {
	// je verifie ici si x et y sont positive
    if x > 0 && y > 0 {
		// si ils sont positive on boucle sur le y
        for i := 0; i < y; i++ {
			// on verifie si on est sur la ligne 1 ou 3 
            if i == 0  {
				// on affiche le premiere A  de la ligne 1
                z01.PrintRune('A')
				// ici on fait le boucle pour pouvoir afficher les 3 B qui suit  le premiere A
                  for j := 0; j < x-2; j++ {
                    z01.PrintRune('B')
                }
				// on affiche le derniere C de la ligne 
                z01.PrintRune('C')
                z01.PrintRune('\n')
				
            } else if i == y-2 {
                // ici on return a la ligne avant d'afficher la deuxieme ligne
				// ici on est sur la 2émé ligne
				// on affiche le B
                z01.PrintRune('B')
				// ici on fait le boucle pour pouvoir afficher le space 3 fois
                for j := 0; j < x-2; j++ {
                    z01.PrintRune(' ')
                }
				// on affiche le derniere B de la ligne 
                z01.PrintRune('B')
				// ici on return a la ligne
                z01.PrintRune('\n')                

            } else {
                // on affiche le premiere C  de la ligne 1
                z01.PrintRune('C')
				// ici on fait le boucle pour pouvoir afficher les 3 *** qui suit  le premiere B
                  for j := 0; j < x-2; j++ {
                    z01.PrintRune('B')
                }
				// on affiche le derniere a de la ligne 
                z01.PrintRune('A')    
                z01.PrintRune('\n')             

            }
        }
    }
}


// func main() {
// 	QuadE1(5,3)
// }