package message

import "golang.org/x/text/language"

type FieldNameTag string

const (
	Bottom      FieldNameTag = "Bottom"
	Creature    FieldNameTag = "Creature"
	Plant       FieldNameTag = "Plant"
	Stone       FieldNameTag = "Stone"
	Wood        FieldNameTag = "Wood"
	Tank        FieldNameTag = "Tank"
	Equipment   FieldNameTag = "Equipment"
	Post        FieldNameTag = "Post"
	Activity    FieldNameTag = "Activity"
	Collection  FieldNameTag = "Collection"
	DeviceToken FieldNameTag = "DeviceToken"
	Record      FieldNameTag = "Record"
)

var filedNames = map[string]map[language.Tag]string{
	string(Creature): {
		language.Japanese: "生き物",
		language.English:  "creature",
	},
	string(Bottom): {
		language.Japanese: "低床",
		language.English:  "bottom",
	},
	string(Plant): {
		language.Japanese: "植物",
		language.English:  "plant",
	},
	string(Stone): {
		language.Japanese: "石",
		language.English:  "stone",
	},
	string(Wood): {
		language.Japanese: "流木",
		language.English:  "wood",
	},
	string(Tank): {
		language.Japanese: "水槽",
		language.English:  "tank",
	},
	string(Equipment): {
		language.Japanese: "機材",
		language.English:  "equipment",
	},
	string(Post): {
		language.Japanese: "投稿",
		language.English:  "post",
	},
	string(Activity): {
		language.Japanese: "アクティビティ",
		language.English:  "activity",
	},
	string(Collection): {
		language.Japanese: "コレクション",
		language.English:  "collection",
	},
	string(DeviceToken): {
		language.Japanese: "デバイストークン",
		language.English:  "device token",
	},
	string(Record): {
		language.Japanese: "記録",
		language.English:  "record",
	},
}
