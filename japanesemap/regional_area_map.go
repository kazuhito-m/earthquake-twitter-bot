package japanesemap

import "github.com/kazuhito-m/earthquake-twitter-bot/coordinate2d"

type RegionalAreaMap struct {
	areas []RegionalArea
}

type RegionalArea struct {
	RegionName string
	Vertexes   coordinate2d.Polygon
}

func (ram RegionalAreaMap) GetRegionName(latitude float64, longitude float64) string {
	nowGeo := coordinate2d.Point{latitude, longitude}
	for _, area := range ram.areas {
		if area.Vertexes.InsideOf(nowGeo) {
			return area.RegionName
		}
	}
	return ""
}

func createAreas() []RegionalArea {
	return []RegionalArea{
		RegionalArea{
			"北海道",
			coordinate2d.CreatePolygon([][2]float64{
				{43.14004, 134.73596},
				{38.1805, 151.78732},
				{44.2949, 165.48915},
				{59.21354, 143.19398},
			}),
		},
		RegionalArea{
			"東北",
			coordinate2d.CreatePolygon([][2]float64{
				{33.18017, 149.07629},
				{37.01685, 139.02926},
				{41.49734, 130.04797},
				{43.14004, 134.73596},
				{38.1805, 151.78732},
			}),
		},
		RegionalArea{
			"関東",
			coordinate2d.CreatePolygon([][2]float64{
				{36.40032, 138.28219},
				{29.14861, 145.57711},
				{33.18017, 149.07629},
				{37.01685, 139.02926},
			}),
		},
		RegionalArea{
			"中部",
			coordinate2d.CreatePolygon([][2]float64{
				{26.29323, 142.82949},
				{41.49734, 130.04797},
				{29.14861, 145.57711},
			}),
		},
		RegionalArea{
			"関西",
			coordinate2d.CreatePolygon([][2]float64{
				{26.29323, 142.82949},
				{25.0322, 140.72012},
				{26.29323, 142.82949},
				{41.49734, 130.04797},
			}),
		},
		RegionalArea{
			"四国",
			coordinate2d.CreatePolygon([][2]float64{
				{34.47608, 134.26016},
				{33.41892, 131.93106},
				{23.77915, 137.2759},
				{25.0322, 140.72012},
			}),
		},
		RegionalArea{
			"中国",
			coordinate2d.CreatePolygon([][2]float64{
				{36.66025, 122.58408},
				{33.41892, 131.93106},
				{34.47608, 134.26016},
				{40.56919, 129.25039},
			}),
		},
		RegionalArea{
			"九州",
			coordinate2d.CreatePolygon([][2]float64{
				{34.07985, 119.44924},
				{23.77915, 137.2759},
				{33.41892, 131.93106},
				{36.66025, 122.58408},
			}),
		},
		RegionalArea{
			"九州",
			coordinate2d.CreatePolygon([][2]float64{
				{34.07985, 119.44924},
				{26.44178, 120.0425},
				{15.8036, 130.65529},
				{23.77915, 137.2759},
			}),
		},
	}
}

func CreateRegionalAreaMap() RegionalAreaMap {
	return RegionalAreaMap{createAreas()}
}
