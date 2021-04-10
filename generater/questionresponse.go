package generater

type QuestionResponse struct {
	Data Data `json:"data"`
}
type TopicTags struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
	TranslatedName string `json:"translatedName"`
	Typename string `json:"__typename"`
}
type CodeSnippets struct {
	Lang string `json:"lang"`
	LangSlug string `json:"langSlug"`
	Code string `json:"code"`
	Typename string `json:"__typename"`
}
type Question struct {
	QuestionID string `json:"questionId"`
	QuestionFrontendID string `json:"questionFrontendId"`
	BoundTopicID int `json:"boundTopicId"`
	Title string `json:"title"`
	TitleSlug string `json:"titleSlug"`
	Content string `json:"content"`
	TranslatedTitle string `json:"translatedTitle"`
	TranslatedContent string `json:"translatedContent"`
	IsPaidOnly bool `json:"isPaidOnly"`
	Difficulty string `json:"difficulty"`
	Likes int `json:"likes"`
	Dislikes int `json:"dislikes"`
	IsLiked interface{} `json:"isLiked"`
	SimilarQuestions string `json:"similarQuestions"`
	Contributors []interface{} `json:"contributors"`
	LangToValidPlayground string `json:"langToValidPlayground"`
	TopicTags []TopicTags `json:"topicTags"`
	CompanyTagStats string `json:"companyTagStats"`
	CodeSnippets []CodeSnippets `json:"codeSnippets"`
	Stats string `json:"stats"`
	Hints []interface{} `json:"hints"`
	Solution interface{} `json:"solution"`
	Status string `json:"status"`
	SampleTestCase string `json:"sampleTestCase"`
	MetaData string `json:"metaData"`
	JudgerAvailable bool `json:"judgerAvailable"`
	JudgeType string `json:"judgeType"`
	MysqlSchemas []interface{} `json:"mysqlSchemas"`
	EnableRunCode bool `json:"enableRunCode"`
	EnvInfo string `json:"envInfo"`
	Book interface{} `json:"book"`
	IsSubscribed bool `json:"isSubscribed"`
	IsDailyQuestion bool `json:"isDailyQuestion"`
	DailyRecordStatus string `json:"dailyRecordStatus"`
	EditorType string `json:"editorType"`
	UgcQuestionID interface{} `json:"ugcQuestionId"`
	Style string `json:"style"`
	ExampleTestcases string `json:"exampleTestcases"`
	Typename string `json:"__typename"`
}
type Data struct {
	Question Question `json:"question"`
}

func NewQuestionResponse() *QuestionResponse {
	return &QuestionResponse{}
}

func (q *QuestionResponse) GetQuestion() string {
	return q.Data.Question.TranslatedContent
}