package piscine




func PrintComb2(){

	for i:= 0; i < 100 ; i++{
		for j:= i + 1; j < 100; i++{
			PrintRune(rune(i) + 48)
			PrintRune(rune(j) + 48)
		}

	}
	


}

