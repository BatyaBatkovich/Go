package main

import ("fmt"; "math/rand";"time")

func main() {
	leght := 8
	rand.Seed(time.Now().UnixNano())
	res := ""
	chars := []rune{

    'a','b','c','d','e','f','g','h','i','j','k','l','m',
    'n','o','p','q','r','s','t','u','v','w','x','y','z',

    
    'A','B','C','D','E','F','G','H','I','J','K','L','M',
    'N','O','P','Q','R','S','T','U','V','W','X','Y','Z',

    '0','1','2','3','4','5','6','7','8','9',

    '!','"','#','$','%','&','\'','(',')','*','+',
    ',','-','.','/',':',';','<','=','>','?','@',
    '[','\\',']','^','_','`','{','|','}','~',
		}
	for i := 0; i < leght; i++ {
				random_append := rand.Intn(len(chars))
		res += string(chars[random_append])
	}
	fmt.Println(res)
}
