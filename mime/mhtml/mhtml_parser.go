package mhtml

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/zieckey/goini"
	"io/ioutil"
	"net/textproto"
	"strings"
)

type MHtml struct {
	Title  string
	Iframe string
}

func New() *MHtml {
	return &MHtml{}
}

func (m *MHtml) GetMHtmlFromLog(logxml string) ([]byte, error) {
	ini := goini.New()
	logxml = strings.Trim(logxml, "<>")
	err := ini.Parse([]byte(logxml), "><", ":")
	if err != nil {
		fmt.Printf("parse xml log error : %v\n", err.Error())
		return nil, err
	}

	req, ok := ini.Get("req")
	if !ok {
		fmt.Printf("cannot found 'req'\n")
		return nil, errors.New("cannot found 'req'")
	}

	ini.Reset()
	err = ini.Parse([]byte(req), "&", "=")
	if err != nil {
		fmt.Printf("parse mhtml error : %v\n", err.Error())
		return nil, err
	}
	mhtmlenc, ok := ini.Get("mhtml")
	if !ok {
		fmt.Printf("cannot found 'mhtml'\n")
		return nil, errors.New("cannot found 'mhtml'")
	}

	mhtml, err := base64.StdEncoding.DecodeString(mhtmlenc)
	if err != nil {
		fmt.Printf("'mhtml' base64 decode error : %v\n", err.Error())
		return nil, err
	}
	ioutil.WriteFile("mime.txt", mhtml, 0644)

	//	mediatype , params , err  := mime.ParseMediaType(string(mhtml))
	//	fmt.Printf("mediatype=%v, params=%v, err=%v\n", mediatype , params , err)
	return mhtml, nil
}

func (m *MHtml) Parse(mht []byte) error {
	r := textproto.NewReader(bufio.NewReader(bytes.NewReader(mht)))
	mimeHeader, err := r.ReadMIMEHeader()
	if err != nil {
		return err
	}
	fmt.Printf("%v %v\n", mimeHeader, err)
	contentType := mimeHeader.Get("Content-Type")
	fmt.Printf("Content-Type = %v %v\n", contentType)
	
		//	mediatype , params , err  := mime.ParseMediaType(string(mhtml))
	//	fmt.Printf("mediatype=%v, params=%v, err=%v\n", mediatype , params , err)
	
	//	boundary := m.GetBoundary(mht)
	//	mr := multipart.NewReader(bytes.NewReader(mht), boundary)
	//	form, _ := mr.ReadForm(int64(len(mht)))
	//	fmt.Printf("%v\n", form)
	//	return nil
	//
	//	var index = 0
	//	for {
	//		part, err := mr.NextPart()
	//		if err != nil {
	//			break
	//		}
	//
	//		fmt.Println("\n\n================================================================================================================================================================\n\n")
	//		d := make([]byte, len(mht))
	//		n, err := part.Read(d)
	//		d = d[:n]
	//		//TODO check err
	//		fmt.Printf("filename=%v formname=%v n=%v err=%v content=\n%v", part.FileName(), part.FormName(), n, err, string(d))
	//		ioutil.WriteFile(
	//			fmt.Sprintf("part-%v.txt", index),
	//			[]byte(fmt.Sprintf("filename=%v formname=%v n=%v err=%v Header=%v content=\n%v", part.FileName(), part.FormName(), n, err, part.Header, string(d))),
	//			0644)
	//
	//		index++
	//	}
	return nil
}

func (m *MHtml) GetBoundary(mht []byte) string {
	//TODO how to parse an boundary from mht content
	boundary := "----=_NextPart_000_3E0F_9DFE9458.D3FADCE4"
	return boundary
}
