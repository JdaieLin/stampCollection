package main

import (
	"cici/andromeda/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"strconv"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

var serials, atlases, stamps []interface{}

func main() {
	var insts []map[string]string
	data, err := ioutil.ReadFile("../../conf/demo2.json")
	if err != nil {
		log.Fatalf("load data faild:%s", err)
	}
	if err = json.Unmarshal(data, &insts); err != nil {
		log.Fatalf("unmarshal faild:%s", err)
	}
	var sm = map[string]int64{}
	var sm_n int64
	for i, inst := range insts {
		// log.Printf("atlas:%+v", inst)
		var atlas models.Atlas
		atlas.ID, _ = strconv.ParseInt(inst["stamp_id"], 10, 64)
		atlas.Name = inst["name"]
		atlas.SetName = inst["set_name"]
		atlas.SetName = strings.TrimPrefix(atlas.SetName, "《")
		atlas.SetName = strings.TrimSuffix(atlas.SetName, "》")
		if serial_id, ok := sm[atlas.SetName]; !ok {
			sm_n++
			sm[atlas.SetName] = sm_n
			serial := models.Serial{sm_n, atlas.SetName}
			serials = append(serials, serial)
			atlas.SerialID = sm_n
			updateSerial(serial)
		} else {
			atlas.SerialID = serial_id
		}
		atlas.Year, _ = strconv.Atoi(inst["year"])
		atlas.Date = atlas.Date * 10000
		atlas.Set = inst["set"]
		atlas.Order, _ = strconv.ParseInt(inst["order"], 10, 64)
		atlas.RealVolume, _ = strconv.ParseInt(inst["real_volume"], 10, 64)
		volume, _ := strconv.ParseFloat(inst["volume"], 64)
		atlas.Volume = int64(volume)
		atlas.Remain = atlas.Volume
		atlas.Amount = atlas.Volume
		atlas.Floor = rand.Intn(10)
		atlas.Image = "http://xxx.aipeli.com/data/image/" + inst["year"] + "/" + inst["stamp_id"] + ".jpg"
		atlases = append(atlases, atlas)
		updateAtlas(atlas)
		for j := int64(0); j < atlas.Amount; j++ {
			var stamp models.Stamp
			stamp.ID = int64(i)<<32 + j
			stamp.AtlasId = atlas.ID
			stamp.Atlas = atlas
			stamp.Brightness = float64(rand.Intn(100)) / 100
			stamp.Stain = float64(rand.Intn(100)/100) / 100
			stamp.Mark = float64(rand.Intn(100)) / 100
			stamp.Crack = float64(rand.Intn(100)) / 100
			if stamp.Stain < 0.1 {
				stamp.Rarity++
			}
			if stamp.Mark > 0.18 {
				stamp.Rarity++
			}
			if stamp.Crack > 0.18 {
				stamp.Rarity++
			}
			if stamp.Brightness > 0.18 {
				stamp.Rarity++
			}
			stamp.Score = float64(90-stamp.Rarity*10+rand.Intn(10)) + float64(rand.Intn(1000))/1000.0
			stamp.OwnerID = 100
			stamp.Floor = 330000.0 / float64(atlas.Year-1959) / float64(atlas.Volume)
			switch stamp.Rarity {
			case 4:
				stamp.Floor = stamp.Floor
			case 3:
				stamp.Floor = stamp.Floor * 2.5
			case 2:
				stamp.Floor = stamp.Floor * 6
			case 1:
				stamp.Floor = stamp.Floor * 10.16
			case 0:
				stamp.Floor = stamp.Floor * 600
			}
			// log.Printf("%+v", stamp)
			// return
			stamps = append(stamps, stamp)
			// Flush(false)
			updateStamp(stamp)
		}
	}
	// Flush(true)
	return
}

func updateSerial(serial models.Serial) {
	var changes = bson.M{
		"$set": bson.M{
			"Name": serial.Name,
		},
	}
	if err := models.DB().Update(models.COL_SERIAL, bson.M{"_id": serial.ID}, changes); err != nil {
		log.Fatalf("inset faild:%s", err)
	}
}

func updateAtlas(atlas models.Atlas) {
	var changes = bson.M{
		"$set": bson.M{
			"SetName": atlas.SetName,
		},
	}
	if err := models.DB().Update(models.COL_ATLAS, bson.M{"_id": atlas.ID}, changes); err != nil {
		log.Fatalf("inset faild:%s", err)
	}
}

func updateStamp(stamp models.Stamp) {
	var changes = bson.M{
		"$set": bson.M{
			"Rarity":        stamp.Rarity,
			"Score":         stamp.Score,
			"atlas.SetName": stamp.SetName,
		},
	}
	if err := models.DB().Update(models.COL_STAMP, bson.M{"_id": stamp.ID}, changes); err != nil {
		log.Fatalf("inset faild:%s", err)
	}
}

func Flush(force bool) {
	var err error
	if len(atlases) > 1000 || force {
		if err = models.DB().Insert(models.COL_ATLAS, atlases...); err != nil {
			log.Fatalf("inset faild:%s", err)
		}
		atlases = atlases[:0]
	}
	if len(stamps) > 1000 || force {
		if err = models.DB().Insert(models.COL_STAMP, stamps...); err != nil {
			log.Fatalf("inset faild:%s", err)
		}
		stamps = stamps[:0]
	}
	if len(serials) > 1000 || force {
		if err = models.DB().Insert(models.COL_SERIAL, serials...); err != nil {
			log.Fatalf("inset faild:%s", err)
		}
		serials = serials[:0]
	}
}
