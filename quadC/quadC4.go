// package main

package piscine

import "github.com/01-edu/z01"

func QuadC4(x int, y int) {
	// je verifie ici si x et y sont positive
    if x > 0 && y > 0 {
		// si ils sont positive on boucle sur le y
        for i := 0; i < y; i++ {
			// on verifie si on est sur la ligne 
            if i == 0  {
				// on affiche le A  
                z01.PrintRune('A')
                z01.PrintRune('\n')
				// ici on fait le boucle pour pouvoir afficher les 3  B 
                  for j := 0; j < y-2; j++ {
                    z01.PrintRune('B')
                    z01.PrintRune('\n')
                }
				// on affiche le C
                z01.PrintRune('C')
                z01.PrintRune('\n')
				
            } 
    }
}
}


// func main() {
// 	QuadC4(1,5)
// }