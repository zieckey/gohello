package main

import (
    "encoding/xml"
    "strings"
    "fmt"
)

func main() {
    var t xml.Token
    var err error

    input := `<!DOCTYPE html>
<html>
	<head>
		<title>
		the title of the page
		</title>
	</head>
	<body>
		<div class="hey" custom_attr="wow"><h2>Title here</h2></div>
		<span><h2>Yoyoyo</h2></span>
		<div id="x">
			<span>
				span content<a href="xxx"><div><li>1st div content</li></div></a>
			</span>
		</div>
		<div class="yo hey">
			<a href="xyz"><div class="cow sheep bunny"><h8>h8 content</h8></div></a>
		</div>
	</body>
</html>
`
    inputReader := strings.NewReader(input)

    decoder := xml.NewDecoder(inputReader)
    indent := 0
    for {
    	t, err = decoder.Token()
    	if err != nil {
    		fmt.Printf("!!!!!!! ERROR %v", err.Error())
    		break
    	}
    	
        switch token := t.(type) {
        // 处理元素开始（标签）
        case xml.StartElement:
            name := token.Name.Local
            fmt.Printf("%vStartElement ====> Token name: %s\n", strings.Repeat(" ", indent), name)
        	indent += 4
            for _, attr := range token.Attr {
                attrName := attr.Name.Local
                attrValue := attr.Value
                fmt.Printf("%vAn attribute is: %s %s\n", strings.Repeat(" ", indent), attrName, attrValue)
            }
        // 处理元素结束（标签）
        case xml.EndElement:
        	indent -= 4
            fmt.Printf("%vEndElement ===> Token of '%s' end\n", strings.Repeat(" ", indent), token.Name.Local)
        // 处理字符数据（这里就是元素的文本）
        case xml.CharData:
            content := string([]byte(token))
            fmt.Printf("%v    CharData ===> This is the content: %v\n", strings.Repeat(" ", indent), content)
        case xml.Directive:
            content := string([]byte(token))
            fmt.Printf("%vDirective ===> This is the content: %v\n", strings.Repeat(" ", indent), content)
        case xml.ProcInst:
            fmt.Printf("%vProcInst ===> Inst:[%v] Target:[%v]\n", strings.Repeat(" ", indent), string(token.Inst), token.Target)
        case xml.Comment:
            fmt.Printf("%vComment ===> [%v]\n", strings.Repeat(" ", indent), string([]byte(token)))            
        default:
            // ...
        }
        
        //fmt.Printf("%v\n", decoder.Entity)
    }
}

/*

*/