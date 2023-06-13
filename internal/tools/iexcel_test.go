package tools

import (
	"fmt"
	"github.com/spf13/cast"
	"sync"
	"testing"
	"time"
)

func TestExcel(t *testing.T) {

	// 读文件
	defer timeCost(time.Now())
	_, rows, err := FileToRows("gpt6.xlsx", "工作表1")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 读行
	var dbDataIds []int
	var AIpromptIds []int
	notfounf := []int{450, 452, 453, 469, 493, 494, 496, 497, 503}
	m := make(map[int]bool)
	// 将 s 中的元素添加到映射中。
	for _, v := range notfounf {
		m[v] = true
	}
	for i, row := range rows {
		// 跳过表头
		if i == 0 {
			continue
		}

		if !m[i] {
			continue
		}

		// 获取写库数据
		imageUrl := row[8]
		grade := cast.ToInt(row[1][0:1])
		question := row[7]
		id, err := getDBData(grade, question, imageUrl)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("请求与返回信息：", i, question, id)
		dbDataIds = append(dbDataIds, id)
	}

	for _, id := range dbDataIds {
		fmt.Println("id:-------", id)
		time.Sleep(time.Second)
		// 获取 AI 推荐题目
		AIPromptId, err := getAIData(id)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("写入AI数组的信息：", AIPromptId)
		AIpromptIds = append(AIpromptIds, AIPromptId)
	}

	fmt.Println("写入 excel 的数据：", dbDataIds, AIpromptIds)

	// 写文件
	//for i := 0; i < len(dbDataIds); i++ {
	//	cell := fmt.Sprintf("J%d", i+2)
	//	err := f.SetCellValue("工作表1", cell, dbDataIds[i])
	//	if err != nil {
	//		return
	//	}
	//}
	//
	//for i := 0; i < len(AIpromptIds); i++ {
	//	cell := fmt.Sprintf("K%d", i+2)
	//	if AIpromptIds[i] == 0 {
	//		err := f.SetCellValue("工作表1", cell, "无")
	//		if err != nil {
	//			return
	//		}
	//	} else {
	//		err := f.SetCellValue("工作表1", cell, AIpromptIds[i])
	//		if err != nil {
	//			return
	//		}
	//	}
	//
	//}
	//
	//// 保存文件
	//if err := f.SaveAs("gpt4-簇讲解plus测试.xlsx"); err != nil {
	//	fmt.Println(err)
	//}
}

// 并发访问接口例子
func TestGpt4Math(t *testing.T) {

	// 读文件
	defer timeCost(time.Now())
	f, rows, err := FileToRows("gptData返回数据4.xlsx", "工作表1")
	if err != nil {
		fmt.Println(err)
		return
	}

	var m sync.Map
	var wg sync.WaitGroup
	useCount := 525
	maxCon := 8
	//badCase := []int{11, 19, 31, 37, 53, 63, 69, 70, 71, 79, 80, 82, 85, 87, 107, 108, 115, 117, 123, 124, 127, 134, 143, 146, 148, 153, 154, 158, 161, 164, 165, 166, 174, 181, 183, 191, 192, 196, 199, 200, 201, 202, 212, 213, 221, 223, 232, 233, 237, 245, 251, 252, 259, 260, 263, 266, 273, 274, 277, 281, 283, 286, 289, 300, 303, 304, 305, 308, 309, 310, 314, 315, 318, 328, 329, 332, 341, 361, 364, 368, 369, 371, 376, 377, 378, 379, 380, 383, 384, 385, 386, 393, 394, 404, 405, 406, 411, 418, 420, 422, 430, 436, 439, 440, 442, 445, 446, 447, 448, 450, 467, 468, 477, 478, 481, 488, 492, 494, 495, 496, 507, 509, 515, 517, 518, 520, 521}
	//badCase := []int{70, 69, 31, 11, 53, 63, 37, 71, 79, 80, 82, 85, 87, 107, 115, 117, 123, 124, 127, 134, 146, 143, 148, 154, 161, 164, 166, 165, 181, 183, 191, 192, 196, 200, 221, 232, 223, 237, 245, 251, 252, 259, 266, 273, 274, 281, 283, 286, 289, 300, 303, 304, 305, 309, 315, 318, 328, 329, 332, 341, 361, 364, 368, 369, 371, 376, 377, 378, 379, 384, 385, 393, 394, 404, 405, 406, 411, 418, 420, 422, 430, 436, 442, 445, 446, 447, 448, 450, 467, 477, 478, 481, 488, 492, 494, 496, 507, 509, 515, 517, 518}
	//badCase := []int{82, 124, 127, 223, 303, 305, 341, 509}
	badCase := []int{509}
	c := make(chan struct{}, maxCon)
	defer close(c)

	for i, row := range rows {
		// 跳过表头
		if i == 0 {
			continue
		}
		if i > useCount {
			break
		}

		// bad_case
		bad := false
		for _, badId := range badCase {
			if badId == i {
				bad = true
				break
			}
		}

		if !bad {
			continue
		}

		// 获取写库数据
		question := row[7]
		control := row[10]
		asr := ""
		if control != "无" {
			asr = row[14]
		}

		id := row[0]
		c <- struct{}{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("取得行数据：--> ", id, asr, question)
			if asr == "" {
				m.Store(cast.ToInt(id), "无数据")
				<-c
				return
			}
			// 请求
			gptStr, err := getGpt4MathData(Gpt4Prompt, question, asr)
			if err != nil {
				fmt.Println("请求 gpt-err --> ", err)
				<-c
				return
			}
			m.Store(cast.ToInt(id), gptStr)
			fmt.Println("id", id, "gpt 返回信息：@GPT-B@[", gptStr, "]@GPT-E@")
			<-c
		}()
	}

	wg.Wait()
	// 写文件
	fmt.Println("开始写文件")

	m.Range(func(key, value interface{}) bool {
		id := key.(int)
		gptStr := value.(string)
		cell := fmt.Sprintf("Q%d", id+1)
		fmt.Println("写入数据：--> ", id, gptStr)
		err := f.SetCellValue("工作表1", cell, gptStr)
		return err == nil
	})

	//for i := 1; i <= useCount; i++ {
	//	cell := fmt.Sprintf("Q%d", i+1)
	//	if gptStr, ok := m.Load(i); ok {
	//		err := f.SetCellValue("工作表1", cell, cast.ToString(gptStr))
	//		if err != nil {
	//			return
	//		}
	//	} else {
	//		fmt.Println("有问题的数据行：", i)
	//	}
	//
	//}

	// 保存文件
	if err := f.SaveAs("gptData返回数据5.xlsx"); err != nil {
		fmt.Println(err)
	}
}

func TestText4Question(t *testing.T) {

	// 读文件
	defer timeCost(time.Now())
	f, rows, err := FileToRows("./excel/内容云试题信息0613.xlsx", "Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 读行
	var badCase []int
	for i, row := range rows {
		// 跳过表头,且跳过没有有效数据的行
		if i == 0 || len(row) < 3 {
			continue
		}

		if i > 197893 {
			break
		}

		// 获取写库数据
		questionText := row[2]
		cell := fmt.Sprintf("D%d", i+1)
		id, err := GetText4QuestionData(questionText)
		if err != nil {
			badCase = append(badCase, i)
			_ = f.SetCellValue("Sheet1", cell, err.Error())
			continue
		}

		// 写文件
		fmt.Println("请求与返回信息：", i, questionText, id)
		err = f.SetCellValue("Sheet1", cell, id)
		if err != nil {
			return
		}

		if i%1000 == 0 { // 每 1000 次保存一次数据
			// 保存文件
			if err := f.SaveAs("./excel/[结果]内容云试题信息0613.xlsx"); err != nil {
				fmt.Println(err)
			}
			fmt.Println("保存成功：", i)
		}
	}

	fmt.Println(badCase)

	// 保存文件
	if err := f.SaveAs("./excel/[结果]内容云试题信息0613.xlsx"); err != nil {
		fmt.Println(err)
	}

}
