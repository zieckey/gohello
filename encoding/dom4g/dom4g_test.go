package dom4g

import (
    "testing"
    "strings"
    "github.com/bmizerany/assert"
)

//func TestXYZ(t *testing.T) {
//    input := `<!DOCTYPE html>
//<html>
//	<head>
//		<title>
//		the title of the page
//		</title>
//	</head>
//	<body>
//		<div class="hey" custom_attr="wow"><h2>Title here</h2></div>
//		<span><h2>Yoyoyo</h2></span>
//		<div id="x">
//			<span>
//				span content<a href="xxx"><div><li>1st div content</li></div></a>
//			</span>
//		</div>
//		<div class="yo hey">
//			<a href="xyz"><div class="cow sheep bunny"><h8>h8 content</h8></div></a>
//		</div>
//	</body>
//</html>
//`
//}

func TestParse(t *testing.T) {
    input := `
    <Person>xxx
    	<FirstName>Xu</FirstName>
    	<LastName>Xinhua</LastName>
    </Person>`
    doc, err := ParseString(input)
    println(doc.ToString())
    assert.Equal(t, err, nil)
    assert.NotEqual(t, doc, nil)
    assert.Equal(t, doc.Name, "Person")
    assert.Equal(t, strings.TrimSpace(doc.Value), "xxx")
    assert.Equal(t, len(doc.Childs), 2)
    assert.Equal(t, len(doc.Attrs), 0)
}