package earthquake

import (
	"github.com/kazuhito-m/earthquake-twitter-bot/earthquake"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func Test地震発生時の地震データJSONテキストを構造体に変換出来る(t *testing.T) {
	jsonText := loadTestJson("testEarthquakeReportRise.json")

	actual := earthquake.ParseJsonOf(jsonText)

	result := actual.Result
	assertJsonValue(result.Status, "success", t)
	assertJsonValue(result.Message, "", t)
	assertJsonValue(actual.Report_Id, "20180916115045", t)

	assertJsonValue(actual.Region_Name, "十勝地方南部", t)
	assertJsonValue(actual.Report_Time, "2018/09/16 11:51:37", t)
	assertJsonFloat(actual.LatitudeF64(), 42.6, t)
	assertJsonFloat(actual.LongitudeF64(), 143.5, t)
	assertJsonFloat(actual.MagunitudeF64(), 3.6, t)
	assertJsonInt(actual.Sindo(), 2, t)

	assertJsonValue(actual.Request_Time, "20180916115437", t)
}

func Test空データ時の地震データJSONテキストを構造体に変換出来る(t *testing.T) {
	jsonText := loadTestJson("testEarthquakeReportEmpty.json")

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

func Test地震データが同じか否かをは判定出来る(t *testing.T) {
	base := earthquake.EarthquakeReport{Report_Id: "同一レポート番号"}
	same := earthquake.EarthquakeReport{Report_Id: "同一レポート番号"}
	if !base.Same(same) {
		t.Errorf("「地震データが同一か」判定が「同一でない」結果を返した。")
	}

	other := earthquake.EarthquakeReport{Report_Id: "全く違うレポート番号"}
	if base.Same(other) {
		t.Errorf("「地震データが同一じゃない」判定が「同一」結果を返した。")
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

func assertJsonFloat(actual float64, expect float64, t *testing.T) {
	if actual != expect {
		t.Errorf("JSON読み取り結果の値が異なります。結果:'%v',期待:'%v'", actual, expect)
	}
}

func assertJsonInt(actual int, expect int, t *testing.T) {
	if actual != expect {
		t.Errorf("JSON読み取り結果の値が異なります。結果:'%v',期待:'%v'", actual, expect)
	}
}
