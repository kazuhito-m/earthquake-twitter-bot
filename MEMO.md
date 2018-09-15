基本的に、 golang には、

- 構造体
- インターフェイス

しか「カタチを表すもの」はない。

例えば、ClassBasedな擬似言語において、Classをつくるような行為、例えば↓

```
class Site {
    // private field
	private client web.Client

    // constructor
    Site(outClient web.Client) {
        client = outClient
    }

    // private method
    private func createUrl() string {
        return strings.Replace(SITE_URL_TEMPLATE, "{time_id}", CreateTimeId(), 1)
    }

    // public method
    public func GetNowEarthquakeAnalyze() EarthquakeInformation {
        url := createUrl()
        jsonText := this.client.Get(url)
        return ParseJsonOf(jsonText)
    }

    // factory method
    public static func CreateDefaultSite() Site {
        return new Site(new WebClient())
    }
}

var site = Site.CreateDefaultSite()
var analyze = site.GetNowEarthquakeAnalyze()
Println(analyze.status)
```

というような「クラス」みたいなものを作りたい場合、golang的には

```golang
package earthquake

type Site struct {
	client web.Client
}

func CreateTimeId() string {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	return time.Now().In(loc).Format("20060102150405")
}

func createUrl() string {
	return strings.Replace(SITE_URL_TEMPLATE, "{time_id}", CreateTimeId(), 1)
}

func (site Site) GetNowEarthquakeAnalyze() EarthquakeInformation {
	url := createUrl()
	jsonText := site.client.Get(url)
	return ParseJsonOf(jsonText)
}

func createSite(client web.Client) Site {
	return Site{client}
}

func CreateDefaultSite() {
    return createSite(WebClient{})
}

var site Site = earthquake.CreateDefaultSite()
var analyze = site.GetNowEarthquakeAnalyze()
fmt.Println(analyze.status)
```
