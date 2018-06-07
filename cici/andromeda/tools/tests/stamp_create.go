package main

import (
	"cici/andromeda/libs/log"
	"cici/andromeda/models"
	"math/rand"
)

func main() {
	var stamps []interface{}
	var names = []string{"天山童姥", "丁春秋", "段正淳", "阮星竹", "萧敬腾", "陈独秀", "逍遥子", "李秋水", "无崖子"}
	for i := 0; i < 100; i++ {
		stamp := &models.Stamp{
			Name:      names[rand.Intn(len(names))],
			Expain:    "天下白玉京，五楼十二城，仙人抚我顶，结发授长生",
			Rarity:    1,
			Crack:     1,
			Mark:      1,
			Stain:     1,
			Score:     100,
			SerialID:  rand.Int63n(3),
			SerialNum: 1,
			OwnerID:   100,
			Floor:     10,
			Image:     "https://img.zcool.cn/community/0133ee59a973a8a8012028a9985b34.png",
		}
		stamp.ID, _ = models.DB().AutoIncId(models.COL_STAMP)
		stamps = append(stamps, stamp)
	}
	if err := models.DB().Insert(models.COL_STAMP, stamps...); err != nil {
		log.Errorf("insert stamps faild : %d %+v", len(stamps), stamps[0])
	}
	return
}
