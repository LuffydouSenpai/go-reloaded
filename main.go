package main

import (
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	argument := os.Args[1:]

	//vérifier qu'il y a bien seulement 2 argument
	if len(argument) != 2 {
		return
	}

	//lire le fichier puit le mettre dans un string
	data, _ := ioutil.ReadFile(argument[0])
	stest := string(data)
	stest = strings.Replace(stest, "‘", "'", -1)

	os.Remove(argument[1]) //remove the file

	//le mettre dans un tableau de string temporaire avec un split
	stab := SplitWhiteSpaces(stest)

	//faire les vérification
	result := verification(stab)

	//cree le fichier txt et le remplir
	createFichierTxt(result, argument[1])

}

func SplitWhiteSpaces(s string) []string {
	var result []string
	var temp string
	nb := 0
	resultNB := 0
	vide := true
	parenthese := false
	for range s {
		nb++
	}
	for i := 0; i <= nb-1; i++ {
		if string(s[i]) == "(" {
			parenthese = true
		}
		if string(s[i]) == ")" {
			parenthese = false
		}
		if (string(s[i]) == "," || string(s[i]) == "." || string(s[i]) == "!" || string(s[i]) == "?" || string(s[i]) == ":" || string(s[i]) == ";") && vide == true && parenthese == false  {
			result = append(result, string(s[i]))
		} else if (string(s[i]) == " " || string(s[i]) == "," || string(s[i]) == "." || string(s[i]) == "!" || string(s[i]) == "?" || string(s[i]) == ":" || string(s[i]) == ";") && vide == false && parenthese == false {
			result = append(result, temp)
			temp = ""
			vide = true
			resultNB++
			if string(s[i]) != " " {
				result = append(result, string(s[i]))
			}
		} else if i == nb-1 {
			temp = temp + string(s[i])
			result = append(result, temp)
		} else if string(s[i]) != " " || (string(s[i]) == " " && parenthese == true) {
			temp = temp + string(s[i])
			vide = false
		}
	}
	return result
}

func hex(stab []string) {
	nblen := len(stab)-1
	
	decimal, _ := strconv.ParseInt(stab[nblen], 16, 64)


	stab[nblen] = strconv.Itoa(int(decimal))

	return
}

func bin(stab []string) {

	nblen := len(stab)-1

	nb, _ := strconv.Atoi(stab[nblen])

	nbresult := binaryToDecimal(nb)

	stab[nblen] = strconv.Itoa(nbresult)

	return
}

func binaryToDecimal(num int) int {
	var remainder int
	index := 0
	decimalNum := 0
	for num != 0{
	   remainder = num % 10
	   num = num / 10
	   decimalNum = decimalNum + remainder * int(math.Pow(2, float64(index)))
	   index++
	}
	return decimalNum
}

func up(stab []string, snb string) {

	nbdonner, _ := strconv.Atoi(snb)

	nblen := len(stab)-1
	if nblen < nbdonner {
		nbdonner = len(stab)
	}
	nb := nblen - nbdonner

	for i := nb+1; i <= nblen; i++ {
		s := stab[i]
		result := ""
		for _, v := range s {
			if v >= 'a' && v <= 'z' {
				result = result + string(v-32)
			} else {
				result = result + string(v)
			}
		}
		stab[i] = result

	}
	return 
}

func low(stab []string, snb string) {

	nbdonner, _ := strconv.Atoi(snb)

	nblen := len(stab)-1
	if nblen < nbdonner {
		nbdonner = len(stab)
	}
	nb := nblen - nbdonner

	for i := nb+1; i <= nblen; i++ {
		s := stab[i]
		result := ""
		for _, v := range s {
			if v >= 'A' && v <= 'Z' {
				result = result + string(v+32)
			} else {
				result = result + string(v)
			}
		}
		stab[i] = result

	}
	return 
}

func cap(stab []string, snb string) {
	
	nbdonner, _ := strconv.Atoi(snb)
	nblen := len(stab)-1
	if nblen < nbdonner {
		nbdonner = len(stab)
	}
	nb := nblen - nbdonner

	for i := nb+1; i <= nblen; i++ {
		s := stab[i]
		result := ""
		nb := 1
		for _, v := range s {
			if nb == 1 && v >= 'a' && v <= 'z' {
				result = result + string(v-32)
				nb++
			} else {
				result = result + string(v)
			}
		}
		stab[i] = result
	}
	return 
}

func createFichierTxt(stab []string, arg string) {

	espace := false
	s := ""

	for i := 0; i < len(stab); i++ {
		if espace == true {
			s = s + " "
		}
		s = s + stab[i]
		espace = true
	}
	file, _ := os.OpenFile(arg, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
    defer file.Close() // on ferme automatiquement à la fin de notre programme

    _, _ = file.WriteString(s) // écrire dans le fichier
}

func verification(stab []string) []string {
	var result = []string{}
	temp := ""
	tempfunc := ""
	snbr := ""
	imprim := true
	firstMark := true


	for i := 0; i < len(stab); i++ {
		temp = stab[i]	
		for j := 0; j < len(temp); j++ {
			tempfunc = ""
			snbr = ""

			//verification des ponctuation
			if temp == "," || temp == "." || temp == "!" || temp == "?" || temp == ":" || temp == ";" {
				result[len(result)-1] = result[len(result)-1] + temp
				imprim = false
			}

			//vérification des mark
			if temp == "'" && firstMark == true {
				temp = temp + stab[i+1]
				i++
				firstMark = false
			} else if temp == "'" && firstMark == false {
				result[len(result)-1] = result[len(result)-1] + temp
				imprim = false
				firstMark = true
			}

			//vérification de a
			if temp == "a" || temp == "A" {
				t := stab[i+1]
				if string(t[0]) == "a" || string(t[0]) == "A" || string(t[0]) == "e" || string(t[0]) == "E" || string(t[0]) == "i" || string(t[0]) == "I" || string(t[0]) == "o" || string(t[0]) == "O" || string(t[0]) == "u" || string(t[0]) == "U" || string(t[0]) == "h" || string(t[0]) == "H" {
					temp = temp + "n"
				}
			}

			//vérifier si c'est une fonction
			for _, v := range temp {
				if v == '(' {
					imprim = false
				}
				if imprim == false {
					if v >= 'a' && v <= 'z' {
					tempfunc = tempfunc + string(v)
					} else if v >= '0' && v <= '9'{
						snbr = snbr + string(v)
					}
				}
			}
		}	
		if imprim == false {
			if snbr == "" {
				snbr = "1"
			}

			//voir quelle function sais
			switch tempfunc {
				case "hex":
					hex(result)
					break
				case "bin":
					bin(result)
					break
				case "up":
					up(result, snbr)
					break
				case "low":
					low(result, snbr)
					break
				case "cap":
					cap(result, snbr)
					break
				default:
			}
		}
		
		if imprim == true {
			result = append(result, temp)
		}
		imprim = true
	}

	return result
}
