package builder

import (
	"io/ioutil"
	"strings"

	"github.com/ikeikeikeike/shuffler"

	// "github.com/k0kubun/pp"

	"gopkg.in/yaml.v2"
)

type Fixable struct {
	Tags      map[string]string
	Sentences map[string]string
}

func NewFixable() *Fixable {
	return &Fixable{
		Tags:      map[string]string{},
		Sentences: map[string]string{},
	}
}

func (f *Fixable) SetFilterByPath(pth string) (err error) {
	buf, err := ioutil.ReadFile(pth)
	if err != nil {
		return
	}

	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal(buf, &m)
	if err != nil {
		return
	}

	for _, val := range m["tags"].([]interface{}) {
		v := val.(map[interface{}]interface{})
		f.Tags[v["from_word"].(string)] = v["to_word"].(string)
	}

	for _, val := range m["sentences"].([]interface{}) {
		v := val.(map[interface{}]interface{})
		f.Sentences[v["from_word"].(string)] = v["to_word"].(string)
	}

	// pp.Println(f)
	return
}

// We expect args data format below.
//
//  - from_word: 'from word'
//  - to_word: 'to_word'
//  or
//  - from_word: 'from word'
//  - to_word: 'word1,word2,ward3'
//
func (f *Fixable) Sentence(w string) string {
	if w == "" {
		return w
	}

	for k, v := range f.Tags {
		w = strings.Replace(w, k, v, 1)
	}

	for k, v := range f.Sentences {
		words := strings.Split(v, ",")
		shuffler.Shuffle(words)
		w = strings.Replace(w, k, words[0], 1)
	}

	return w
}

func (f *Fixable) Tag(w string) string {
	if w == "" {
		return w
	}

	if v, ok := f.Tags[w]; ok {
		w = v
	}

	return w
}
