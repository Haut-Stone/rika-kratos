package tools

const (
	DbOnlineUrl            = "https://api-internal.tipaipai.com/bx-question-service/njf/question/create"
	AnswerOnlineUrl        = "https://internal-baodian.tal.com/base-search/api/v1/recommend/question/medium"
	Text4QuestionOnlineUrl = "http://10.156.66.74/search/jf/question/content"
)

const (
	DbTestUrl            = "https://qz.chengjiukehu.com/test/bx-question-service/njf/question/create"
	AnswerTestUrl        = "https://internal-test-baodian.tal.com/base-search/api/v1/recommend/question/medium"
	Text4QuestionTestUrl = "https://qz.chengjiukehu.com/test/qingzhou-search-api/search/jf/question/content"
)

const (
	Gpt4Url    = "https://hmi.chengjiukehu.com/gpt-service-develop/v1/condition/chat/completions"
	Gpt4Prompt = "根据给的示范例题的解题方法逐步思考给定的数学题目，确保各步骤准确无误，讲解详细。\n以下是示范例题的解题方法：{{asr}}\n以下是给定的数学题目：{{question}}"
	Template   = "<div class=\"question-item-container\"><img class=\"question-item-pic\" src=\"{{IMAGE_URL}}\"><p class=\"ocr_text_invisible\">{{OCR_TXT}}</p></div>"
)

type DbReq struct {
	SubjectId int    `json:"subject_id"`
	GradeId   int    `json:"grade_id"`
	Origin    int    `json:"origin"`
	SourceId  int    `json:"source_id"`
	OriginId  string `json:"origin_id"`
	Question  string `json:"question"`
	BookId    int    `json:"book_id"`
}

type DbResp struct {
	ErrorCode  int    `json:"error_code"`
	ErrorMsg   string `json:"error_msg"`
	ServerTime int64  `json:"server_time"`
	TraceId    string `json:"trace_id"`
	Data       DbData `json:"data"`
}

type Gpt4Req struct {
	Message         string  `form:"message" json:"message"`
	Model           string  `form:"model" json:"model"`
	PresencePenalty float32 `form:"presence_penalty" json:"presence_penalty"`
	Temperature     float32 `form:"temperature" json:"temperature"`
}

type Gpt4Resp struct {
	ErrorCode  int      `json:"error_code"`
	ErrorMsg   string   `json:"error_msg"`
	ServerTime int64    `json:"server_time"`
	Data       Gpt4Data `json:"data"`
}

type Gpt4Data struct {
	ConversationId string `json:"conversation_id"`
	Result         string `json:"result"`
}

type DbData struct {
	Id       string `json:"id"`
	IsRepeat bool   `json:"is_repeat"`
}

type AnswerReq struct {
	QuestionId string `json:"question_id"`
}

type AIResp struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	Data    AIData `json:"data"`
	Trace   string `json:"trace"`
}

type AIData struct {
	MediumRecommend []InnerData `json:"medium_recommend"`
}

type InnerData struct {
	QuestionId string `json:"question_id"`
}

type Text4QuestionResp struct {
	ErrorCode  int               `json:"error_code"`
	ErrorMsg   string            `json:"error_msg"`
	ServerTime int64             `json:"server_time"`
	TraceId    string            `json:"trace_id"`
	Data       Text4QuestionData `json:"data"`
}

type Text4QuestionData struct {
	Total       int        `json:"total"`
	QuestionArr []Question `json:"questionArr"`
}

type Question struct {
	Question     string  `json:"question"`
	QuestionId   string  `json:"question_id"`
	MatchPercent float32 `json:"match_percent"`
}
