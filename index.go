package main

func Index(s string, toFind string) int {
	for i := 0; i < len(s)-len(toFind); i++ {
		if toFind == s[i:i+len(toFind)] {
			return i
		}
	}
	return -1
}

/* À l'intérieur de la boucle, elle vérifie si la sous-chaîne toFind est égale 
à la tranche de s commençant à l'indice i et ayant une longueur de toFind.
 Si c'est le cas, cela signifie que la sous-chaîne a été trouvée et la fonction renvoie l'indice i.