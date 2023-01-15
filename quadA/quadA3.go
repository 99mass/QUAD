// package main
package piscine

import "github.com/01-edu/z01"

func QuadA3(x int, y int) {
	// je verifie ici si x et y sont positive
    if x > 0 && y > 0 {
		// si ils sont positive on boucle sur le y
        for i := 0; i < y; i++ {
			// on verifie si on est sur la ligne 1 ou 3 
            if i == 0  {
				// on affiche le premiere o  
                z01.PrintRune('o')
                z01.PrintRune('\n')
				// ici on fait le boucle pour pouvoir afficher les 3  | 
                  for j := 0; j < y-2; j++ {
                    z01.PrintRune('|')
                    z01.PrintRune('\n')
                }
				// on affiche le derniere o 
                z01.PrintRune('o')
                z01.PrintRune('\n')
				
            } 
    }
}
}


// func main() {
// 	QuadA3(1,5)
// }