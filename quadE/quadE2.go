// package main

package piscine

import "github.com/01-edu/z01"

func QuadE2(x int, y int) {
	// je verifie ici si x et y sont positive
    if x > 0 && y > 0 {
		// si ils sont positive on boucle sur le y
        for i := 0; i < y; i++ {
			// on verifie si on est sur la ligne 1 
            if i == 0 || i == y-1 {
				// on affiche le premiere A  de la ligne 
                z01.PrintRune('A')
				// ici on fait le boucle pour pouvoir afficher les 3 B 
                  for j := 0; j < x-2; j++ {
                    z01.PrintRune('B')
                }
				// on affiche le C
                z01.PrintRune('C')
                z01.PrintRune('\n')
				
            } 
    }
}
}

// func main() {
// 	QuadE2(5,1)
// }
