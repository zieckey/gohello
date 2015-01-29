package dom4g

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"strings"
)

type ElementArray []*Element

type Element struct {
	Name   string
	Value  string
	Attrs  map[string]string
	Childs map[string]ElementArray
	Parent *Element
	//Root *Element
}

func New() *Element {
	el := &Element{
		Attrs:     make(map[string]string),
		Childs:    make(map[string]ElementArray),
	}
	return el
}

func Parse(r io.Reader) (current *Element, err error) {
	var root *Element
	decoder := xml.NewDecoder(r)
	for {
		t, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		switch token := t.(type) {
		case xml.StartElement:
			el := New()
			el.Name = token.Name.Local
			for _, a := range token.Attr {
				el.Attrs[a.Name.Local] = a.Value
			}

			if root == nil {
				root = el
			} else {
				current.Childs[el.Name] = append(current.Childs[el.Name], el)
				el.Parent = current
			}
			current = el
		case xml.EndElement:
			current = current.Parent
		case xml.CharData:
            /*
            <Person>xxx
                <FirstName>Xu</FirstName>
                <LastName>Xinhua</LastName>
            </Person>

            The Person element's CharData will return 3 times.
            */
			if current != nil {
				if len(current.Value) == 0 {
					current.Value = strings.TrimSpace(string(token))
				} else {
					//TODO how to process this case :
                    /*
                    <Person>
                        xxx
                        <FirstName>Xu</FirstName>
                        yyy
                        <LastName>Xinhua</LastName>
                        zzz
                    </Person>
                    */
				}
			}
		default:
			return nil, fmt.Errorf("parse xml fail!")
		}
	}
	return root, nil
}

func ParseString(xmlstr string) (current *Element, err error) {
	r := strings.NewReader(xmlstr)
	return Parse(r)
}

func (e *Element) ToString() string {
	var buf bytes.Buffer
	e.write(&buf)
	return buf.String()
}

func (e *Element) write(w io.Writer) {
	w.Write([]byte("<"))
	w.Write([]byte(e.Name))
	if len(e.Attrs) > 0 {
		for n, v := range e.Attrs {
			w.Write([]byte(" "))
			w.Write([]byte(n))
			w.Write([]byte("=\""))
			w.Write([]byte(v))
			w.Write([]byte("\""))
		}
	}
	w.Write([]byte(">"))

	w.Write([]byte(e.Value))
	if len(e.Childs) > 0 {
		for _, cl := range e.Childs {
			for _, c := range cl {
				c.write(w)
			}
		}
	}

	w.Write([]byte("</"))
	w.Write([]byte(e.Name))
	w.Write([]byte(">"))
}
