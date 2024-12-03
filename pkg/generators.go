package pkg

import (
	"fmt"
	"strings"
	"unicode"

	"albatross.pe/hexagonal/internal/constants"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"golang.org/x/exp/rand"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func GenerateID() string {
	id, _ := gonanoid.New()
	return id
}

func RemoveStrips(name string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	slug, _, err := transform.String(t, name)
	if err != nil {
		panic(err)
	}

	return slug
}

func FindRandomColor() string {
	index := rand.Intn(len(constants.COLORS))
	return constants.COLORS[index]
}

func FindAbbreviation(name string) string {

	name = RemoveStrips(name)
	abbr := name[0:2]

	if splits := strings.Split(name, " "); len(splits) > 1 {
		abbr = fmt.Sprintf("%s%s", splits[0][0:1], splits[1][0:1])
	}

	return strings.ToUpper(abbr)
}
