package main
import "encoding/xml"

type Student struct {
	Name string
	Age int
	Number int
}

type Class struct {
	Students []Student
	Name string
	Grade int
}

func main() {
	var s Student
	s.Name = "jane"
	s.Age = 12
	s.Number = 1

	var c Class
	c.Name = "class 1-2"
	c.Grade = 5
	c.Students = append(c.Students, s)

	s.Name = "zzz"
	s.Age = 13
	s.Number = 2
	c.Students = append(c.Students, s)

	s.Name = "1231231"
	s.Age = 14
	s.Number = 3
	c.Students = append(c.Students, s)

	m, _ := xml.MarshalIndent(c, "", "    ")
	println(string(m))

	var cc Class
	xml.Unmarshal(m, &cc)

	cc.Name = "class5-6"
	cc.Grade = 6

	m, _ = xml.MarshalIndent(cc, "", "    ")
	println("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	println(string(m))
}