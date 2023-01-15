// package main
package piscine

import "github.com/01-edu/z01"

func QuadC3(x int, y int) {
	// je verifie ici si x et y sont positive
    if x > 0 && y > 0 {
		// si ils sont positive on boucle sur le y
        for i := 0; i < y; i++ {
			// on verifie si on est sur la ligne 1 colonne 1
            if i == 0  {
				// on affiche le A
                z01.PrintRune('A')
                z01.PrintRune('\n')
				
            } 
    }
}
}


// func main() {
// 	QuadC3(1,1)
// }