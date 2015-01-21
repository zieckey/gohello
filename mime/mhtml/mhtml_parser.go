package mhtml

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/zieckey/goini"
	"strings"
	"bytes"
	//"mime"
	"mime/multipart"
	"io/ioutil"
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
	mhtml, ok := ini.Get("mhtml")
	if !ok {
		fmt.Printf("cannot found 'mhtml'\n")
		return nil, errors.New("cannot found 'mhtml'")
	}

	mineenc, err := base64.StdEncoding.DecodeString(mhtml)
	if err != nil {
		fmt.Printf("'mhtml' base64 decode error : %v\n", err.Error())
		return nil, err
	}

	return mineenc, nil
}



func (m *MHtml) Parse(mht []byte) error {
	boundary := m.GetBoundary(mht)
	mr := multipart.NewReader(bytes.NewReader(mht), boundary)
	var index = 0
	for {
		part, err := mr.NextPart()
		if err != nil {
			break
		}
		fmt.Println("\n\n================================================================================================================================================================\n\n")
		d := make([]byte, 1024*1024)
		n, err := part.Read(d)
		fmt.Printf("filename=%v formname=%v n=%v err=%v content=\n%v", part.FileName(), part.FormName(), n, err, string(d))
		ioutil.WriteFile(
			fmt.Sprintf("part-%v.txt", index),
			[]byte(fmt.Sprintf("filename=%v formname=%v n=%v err=%v Header=%v content=\n%v", part.FileName(), part.FormName(), n, err, part.Header, string(d))),
			0755)
		
		index ++ 
	}
	return nil
}

func (m *MHtml) GetBoundary(mht []byte) string {
	//TODO
	boundary := "----=_NextPart_000_3E0F_9DFE9458.D3FADCE4"
	return boundary
}