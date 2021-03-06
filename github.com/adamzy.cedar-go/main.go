package main



import (
	"fmt"
	"os"

	"github.com/adamzy/cedar-go"
)

var (
	cd    *cedar.Cedar
	words = []string{
		"a", "aa", "ab", "ac", "abc", "abd",
		"abcd", "abde", "abdf", "abcdef", "abcde",
		"abcdefghijklmn", "bcd", "b", "xyz",
		"中国", "中国北京", "中国上海", "中国广州",
		"中华", "中华文明", "中华民族", "中华人民共和国",
		"this", "this is", "this is a sentence.",
	}
)

func loadTestData() {
	if cd != nil {
		return
	}
	cd = cedar.New()
	// cd.Ordered = false

	// add the keys
	for i, word := range words {
		cd.Insert([]byte(word), i)
	}

	// delete some keys
	for i := 0; i < len(words); i += 4 {
		cd.Delete([]byte(words[i]))
	}
	return
}

func TestBasic() {
	loadTestData()
	// check the consistency
	checkConsistency(cd)
}

func TestSaveAndLoad() {
	loadTestData()

	cd.SaveToFile("cedar.gob", "gob")
	defer os.Remove("cedar.gob")
	daGob := cedar.New()
	if err := daGob.LoadFromFile("cedar.gob", "gob"); err != nil {
		panic(err)
	}
	checkConsistency(daGob)

	cd.SaveToFile("cedar.json", "json")
	defer os.Remove("cedar.json")
	daJson := cedar.New()
	if err := daJson.LoadFromFile("cedar.json", "json"); err != nil {
		panic(err)
	}
	checkConsistency(daJson)
}

func TestPrefixMatch() {
	var ids []int
	var keys []string
	var values []int

	ids = cd.PrefixMatch([]byte("abcdefg"), 0)
	keys = []string{"ab", "abcd", "abcde", "abcdef"}
	values = []int{2, 6, 10, 9}
	check(cd, ids, keys, values)

	ids = cd.PrefixMatch([]byte("中华人民共和国"), 0)
	keys = []string{"中华", "中华人民共和国"}
	values = []int{19, 22}
	check(cd, ids, keys, values)

	ids = cd.PrefixMatch([]byte("this is a sentence."), 0)
	keys = []string{"this", "this is a sentence."}
	values = []int{23, 25}
	check(cd, ids, keys, values)
}

func check(cd *cedar.Cedar, ids []int, keys []string, values []int) {
	if len(ids) != len(keys) {
		panic("wrong prefix match")
	}
	for i, n := range ids {
		key, _ := cd.Key(n)
		val, _ := cd.Value(n)
		if string(key) != keys[i] || val != values[i] {
			panic("wrong prefix match")
		}
	}
}

func TestOrder() {
	c := cedar.New()
	c.Insert([]byte("a"), 1)
	c.Insert([]byte("b"), 3)
	c.Insert([]byte("d"), 6)
	c.Insert([]byte("ab"), 2)
	c.Insert([]byte("c"), 5)
	c.Insert([]byte(""), 0)
	c.Insert([]byte("bb"), 4)
	ids := c.PrefixPredict([]byte(""), 0)
	if len(ids) != 7 {
		panic("wrong order")
	}
	for i, n := range ids {
		val, _ := c.Value(n)
		if i != val {
			panic("wrong order")
		}
	}
}

func TestPrefixPredict() {
	var ids []int
	var keys []string
	var values []int
	ids = cd.PrefixPredict([]byte("中华"), 0)
	keys = []string{"中华", "中华人民共和国", "中华民族"}
	values = []int{19, 22, 21}
	check(cd, ids, keys, values)

	ids = cd.PrefixPredict([]byte("中国"), 0)
	keys = []string{"中国", "中国上海", "中国广州"}
	values = []int{15, 17, 18}
	check(cd, ids, keys, values)
}

func checkConsistency(cd *cedar.Cedar) {
	for i, word := range words {
		id, err := cd.Jump([]byte(word), 0)
		if i%4 == 0 {
			if err == cedar.ErrNoPath {
				continue
			}
			_, err := cd.Value(id)
			if err == cedar.ErrNoValue {
				continue
			}
			panic("not deleted")
		}
		key, err := cd.Key(id)
		if err != nil {
			panic(err)
		}
		if string(key) != word {
			panic("key error")
		}
		value, err := cd.Value(id)
		if err != nil || value != i {
			fmt.Println(word, i, value, err)
			panic("value error")
		}
	}
}

func main() {
	words = []string{
		"a", "aa", "ab", "ac", "abc", "abd",
		"abcd", "abde", "abdf", "abcdef", "abcde",
	}

	cd := cedar.New()
	// cd.Ordered = false

	// add the keys
	for i, word := range words {
		cd.Insert([]byte(word), i)
	}

	var ids []int
	var keys []string
	var values []int

	ids = cd.PrefixMatch([]byte("abcdefg"), 0)
	for i,id := range ids {
		key, _ := cd.Key(id)
		val, _ := cd.Value(id)
		fmt.Printf("i=%d, key=[%v] val=[%v] id=[%v]\n", i, string(key), val, id)
	}

	keys = []string{"a", "ab", "abc", "abcd", "abcde", "abcdef"}
	values = []int{0, 2, 4, 6, 10, 9}
	check(cd, ids, keys, values)
}
