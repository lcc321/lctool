package generater

import (
	"bufio"
	"encoding/json"
	"fmt"
	md "github.com/JohannesKaufmann/html-to-markdown"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var leetcodePayload string = `{
    "operationName": "questionData",
    "variables": {
        "titleSlug": "%s"
    },
    "query": "query questionData($titleSlug: String!) {\n  question(titleSlug: $titleSlug) {\n    questionId\n    questionFrontendId\n    boundTopicId\n    title\n    titleSlug\n    content\n    translatedTitle\n    translatedContent\n    isPaidOnly\n    difficulty\n    likes\n    dislikes\n    isLiked\n    similarQuestions\n    contributors {\n      username\n      profileUrl\n      avatarUrl\n      __typename\n    }\n    langToValidPlayground\n    topicTags {\n      name\n      slug\n      translatedName\n      __typename\n    }\n    companyTagStats\n    codeSnippets {\n      lang\n      langSlug\n      code\n      __typename\n    }\n    stats\n    hints\n    solution {\n      id\n      canSeeDetail\n      __typename\n    }\n    status\n    sampleTestCase\n    metaData\n    judgerAvailable\n    judgeType\n    mysqlSchemas\n    enableRunCode\n    envInfo\n    book {\n      id\n      bookName\n      pressName\n      source\n      shortDescription\n      fullDescription\n      bookImgUrl\n      pressImgUrl\n      productUrl\n      __typename\n    }\n    isSubscribed\n    isDailyQuestion\n    dailyRecordStatus\n    editorType\n    ugcQuestionId\n    style\n    exampleTestcases\n    __typename\n  }\n}\n"
}`

func QuestionDesc(q string, path string) error {
	res, err := RequestLeetcode(q)
	if err != nil {
		return err
	}
	markdown, err := FormatResponse(res)
	if err != nil {
		return err
	}

	return WriteStringToFile(markdown, path)
}

func RequestLeetcode(q string) (*http.Response, error) {
	url := "https://leetcode-cn.com/graphql/"
	method := "POST"

	s := fmt.Sprintf(leetcodePayload, q)
	payload := strings.NewReader(s)

	client := &http.Client {}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cookie", "csrftoken=UOmYariLbvZizoOWGAlvqSKTmmOTvA57P6r9Tws9iURnZgZ6PUV4EeCYiAaCE2gd")

	return client.Do(req)
}

func FormatResponse(res *http.Response) (string, error) {
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	questionInfo := NewQuestionResponse()
	err = json.Unmarshal(body, questionInfo)
	if err != nil {
		return "", err
	}

	converter := md.NewConverter("", true, nil)
	markdown, err := converter.ConvertString(questionInfo.GetQuestion())
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return markdown, nil
}

func WriteStringToFile(content, path string) error {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	write := bufio.NewWriter(file)
	write.WriteString(content)
	write.Flush()

	return nil
}
