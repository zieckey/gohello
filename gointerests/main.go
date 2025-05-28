package main

import (
	"fmt"
	"html/template"
	"math"
	"net/http"
	"strconv"
)

// 定义 HTML 模板
const htmlTemplate = `
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <title>复利年化收益率计算器</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: #f4f4f4;
            display: flex;
            flex-direction: column;
            align-items: center;
            justify-content: center;
            min-height: 100vh;
            margin: 0;
        }
        h1 {
            color: #333;
            margin-bottom: 20px;
        }
        form {
            background-color: #fff;
            padding: 20px;
            border-radius: 5px;
            box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
            width: 300px;
        }
        label {
            display: block;
            margin-bottom: 5px;
            color: #666;
        }
        input[type="number"] {
            width: 100%;
            padding: 8px;
            margin-bottom: 15px;
            border: 1px solid #ccc;
            border-radius: 3px;
            box-sizing: border-box;
        }
        input[type="submit"] {
            width: 100%;
            padding: 10px;
            background-color: #007BFF;
            color: #fff;
            border: none;
            border-radius: 3px;
            cursor: pointer;
            transition: background-color 0.3s;
        }
        input[type="submit"]:hover {
            background-color: #0056b3;
        }
        p {
            color: #333;
            font-weight: bold;
            margin-top: 20px;
        }
    </style>
</head>
<body>
    <h1>复利年化收益率计算器</h1>
    <form method="post">
        <label for="initialAmount">初始金额:</label>
        <input type="number" id="initialAmount" name="initialAmount" value="{{.InitialAmount}}" required><br>
        <label for="finalAmount">终期金额:</label>
        <input type="number" id="finalAmount" name="finalAmount" value="{{.FinalAmount}}" required><br>
        <label for="years">年限:</label>
        <input type="number" id="years" name="years" value="{{.Years}}" required><br>
        <input type="submit" value="计算">
    </form>
    {{if .Result}}
    <p>每年复利年化收益率: {{.Result}}%</p>
    {{end}}
</body>
</html>
`

// 定义数据结构，添加输入值字段
type CalculatorData struct {
	Result         float64
	InitialAmount  string
	FinalAmount    string
	Years          string
}

// 处理根路径请求
func handler(w http.ResponseWriter, r *http.Request) {
	var data CalculatorData
	if r.Method == http.MethodPost {
		// 解析表单数据
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "表单解析失败", http.StatusBadRequest)
			return
		}

		// 获取输入值
		initialAmountStr := r.FormValue("initialAmount")
		finalAmountStr := r.FormValue("finalAmount")
		yearsStr := r.FormValue("years")

		// 保存输入值到数据结构
		data.InitialAmount = initialAmountStr
		data.FinalAmount = finalAmountStr
		data.Years = yearsStr

		// 转换为浮点数
		initialAmount, err := strconv.ParseFloat(initialAmountStr, 64)
		if err != nil {
			http.Error(w, "初始金额格式错误", http.StatusBadRequest)
			return
		}

		finalAmount, err := strconv.ParseFloat(finalAmountStr, 64)
		if err != nil {
			http.Error(w, "终期金额格式错误", http.StatusBadRequest)
			return
		}

		years, err := strconv.ParseFloat(yearsStr, 64)
		if err != nil {
			http.Error(w, "年限格式错误", http.StatusBadRequest)
			return
		}

		// 计算复利年化收益率
		rate := (math.Pow(finalAmount/initialAmount, 1/years) - 1) * 100
		// 格式化结果，保留 4 位小数
		formattedRate, err := strconv.ParseFloat(fmt.Sprintf("%.4f", rate), 64)
		if err != nil {
			http.Error(w, "结果格式化失败", http.StatusInternalServerError)
			return
		}
		data.Result = formattedRate
	}

	// 渲染模板
	tmpl := template.Must(template.New("calculator").Parse(htmlTemplate))
	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "模板渲染失败", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}