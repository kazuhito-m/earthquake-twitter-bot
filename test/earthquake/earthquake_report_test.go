package earthquake

import (
	"github.com/kazuhito-m/earthquake-twitter-bot/earthquake"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func Test地震データのJSONテキストを構造体に変換出来る(t *testing.T) {
	jsonText := loadTestJson("testEarthquakeReport.json")

	actual := earthquake.ParseJsonOf(jsonText)

	result := actual.Result
	assertJsonValue(result.Status, "success", t)
	assertJsonValue(result.Message, "データがありません", t)
	assertJsonValue(actual.Report_Id, "れぽーとあいでぃ", t)
}

func Testカタチが合わないテキストを送り込んだ場合にはStatusが空文字になる(t *testing.T) {
	jsonText := "構造も持たなければ、分解もできない無構造テキスト"

	actual := earthquake.ParseJsonOf(jsonText)

	result := actual.Result
	assertJsonValue(result.Status, "", t)
}

func Test地震データが存在する場合を判定出来る(t *testing.T) {
	sut := earthquake.EarthquakeReport{}
	sut.Result.Message = ""
	if sut.Exists() == false {
		t.Errorf("「地震データが存在する」判定が「存在しない」結果を返した。")
	}
}

func Test地震データが存在しない場合を判定出来る(t *testing.T) {
	sut := earthquake.EarthquakeReport{}
	sut.Result.Message = "データがありません"
	if sut.Exists() == true {
		t.Errorf("「地震データが存在しない」判定が「存在する」結果を返した。")
	}
}


// utility functions

func createTestFilePath(fileName string) string {
	testDirPath, _ := os.Getwd()
	return path.Join(testDirPath, fileName)
}

func loadTestJson(fileName string) string {
	filePath := createTestFilePath(fileName)
	contents, _ := ioutil.ReadFile(filePath)
	return string(contents);
}

func assertJsonValue(actual string, expect string, t *testing.T) {
	if actual != expect {
		t.Errorf("JSON読み取り結果の値が異なります。結果:'%v',期待:'%v'", actual, expect)
	}
}
