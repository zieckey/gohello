package mhtml

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"net/textproto"
	"strings"
	"encoding/xml"
)

var _ = ioutil.ReadAll

type MHtml struct {
	ContentLocation string // The url of this page
	Title           string
	Iframe          string
	Html            string // The original html content
}

func New() *MHtml {
	return &MHtml{}
}

func (m *MHtml) Parse(mht []byte) error {
	if err := m.ParseHTML(mht); err != nil {
		return err
	}
	
	r := strings.NewReader(m.Html)
	xmldec := xml.NewDecoder(r)
	xmldec.Token()
	//TODO
	return nil
}

func (m *MHtml) ParseHTML(mht []byte) error {	
	br := bufio.NewReader(bytes.NewReader(mht)) // The buffer reader
	tr := textproto.NewReader(br)
	boundary := m.GetBoundary(tr)
	if len(boundary) == 0 {
		return fmt.Errorf("Cannot get boundary")
	}

	mr := multipart.NewReader(br, boundary)
	var index = 0
	for {
		part, err := mr.NextPart()
		if err != nil {
			break
		}

		fmt.Println("\n\n================================================================================================================================================================\n\n")
		d := make([]byte, len(mht))
		n, err := part.Read(d)
		if err != nil && err != io.EOF {
			return err
		}
		d = d[:n]
		fmt.Printf("filename=%v formname=%v n=%v err=%v content=\n", part.FileName(), part.FormName(), n, err)
		ioutil.WriteFile(
			fmt.Sprintf("part-%v.txt", index),
			[]byte(fmt.Sprintf("filename=%v formname=%v n=%v err=%v Header=%v content=\n%v", part.FileName(), part.FormName(), n, err, part.Header, string(d))),
			0644)
		index++

		contentType := part.Header["Content-Type"]
		if len(contentType) == 0 {
			continue
		}
		fmt.Printf("Content-Type=%v\n", contentType[0])
		if contentType[0] == "text/html" {
			m.Html = string(d)
			break
		}
	}
	return nil
}

func (m *MHtml) GetBoundary(r *textproto.Reader) string {
	mimeHeader, err := r.ReadMIMEHeader()
	if err != nil {
		return ""
	}
	fmt.Printf("%v %v\n", mimeHeader, err)
	contentType := mimeHeader.Get("Content-Type")
	fmt.Printf("Content-Type = %v %v\n", contentType)

	mediatype, params, err := mime.ParseMediaType(contentType)
	fmt.Printf("mediatype=%v,  params=%v %v, err=%v\n", mediatype, len(params), params, err)
	boundary := params["boundary"]
	fmt.Printf("boundary=%v\n", boundary)
	return boundary
}
