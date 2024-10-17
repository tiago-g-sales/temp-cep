package pkg

import (
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)



func RemoveAcentos(s string) string  {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	output, _, e := transform.String(t, s)
	if e != nil {
		panic(e)
	}
	return output
}


func ConvertTemp( temp float64 ) (float64){

	const K = 273
	return temp + K

}

func Replace(s string) (string){
	
	return strings.Replace(RemoveAcentos(s), " ", "%20", -1)

}

