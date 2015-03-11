package funcs

import (
	"fmt"
	"io/ioutil"

	"github.com/ikeikeikeike/shuffler"
)

func RandomImgname(ftype string) string {
	files, _ := ioutil.ReadDir(fmt.Sprintf("./static/img/%s/", ftype))

	var names []string
	for _, f := range files {
		names = append(names, f.Name())
	}

	shuffler.Shuffle(names)
	return fmt.Sprintf("/static/img/%s/%s", ftype, names[0])
}

func IsImgFallback(url string) bool {
	switch url {
	case "http://erolog.info":
		return true
	case "http://sexbox.sexy":
		return true
	case "http://eroconnect.net":
		return true
	case "http://localhost:3000":
		return true
	default:
		return false
	}
}
