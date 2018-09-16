package japanesemap

type RegionalClassifications struct {
	Classifications []RegionalClassification
}

type RegionalClassification struct {
	PrefectureName string
	RegionName     string
}

func (r RegionalClassifications) GetRegionName(prefectureName string) string {
	for _, v := range r.Classifications {
		if v.PrefectureName == prefectureName {
			return v.RegionName
		}
	}
	return ""
}

func CreateRegionalClassifications() RegionalClassifications {
	return RegionalClassifications{defineClassificationsList()}
}

func defineClassificationsList() []RegionalClassification {
	classfications := []RegionalClassification{}
	regions := defineRegionMap()
	for rName, value := range regions {
		for _, pName := range value {
			classfications = append(classfications, RegionalClassification{pName, rName})
		}
	}
	return classfications
}

func defineRegionMap() map[string][]string {
	return map[string][]string{
		"北海道": []string{"北海道"},
		"東北":  []string{"青森県", "岩手県", "秋田県", "宮城県", "山形県", "福島県"},
		"関東":  []string{"茨城県", "栃木県", "群馬県", "埼玉県", "千葉県", "東京都", "神奈川県"},
		"中部":  []string{"新潟県", "富山県", "石川県", "福井県", "山梨県", "長野県", "岐阜県", "静岡県", "愛知県"},
		"近畿":  []string{"三重県", "滋賀県", "奈良県", "和歌山県", "京都府", "大阪府", "兵庫県"},
		"中国":  []string{"岡山県", "広島県", "鳥取県", "島根県", "山口県"},
		"四国":  []string{"香川県", "徳島県", "愛媛県", "高知県"},
		"九州":  []string{"福岡県", "佐賀県", "長崎県", "大分県", "熊本県", "宮崎県", "鹿児島県", "沖縄県"},
	}
}
