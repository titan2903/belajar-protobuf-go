package main

import (
	"bytes"
	"fmt"
	"latihan-protobuf-go/model"
	"os"

	"github.com/golang/protobuf/jsonpb"
)

var user1 = &model.User{
	Id: "u001",
	Name: "Jane Doe",
	Password: "password",
	Gender: model.UserGender_FEMALE,
}

var userList = &model.UserList{
	List: []*model.User{
		user1,
	},
}

var garage1 = &model.Garage{
	Id: "g001",
	Name: "kalimdor",
	Coordinate: &model.GarageCoordinate{
		Latitude: 23.2345657,
		Longitude: 53.5645454,
	},
}

var garageList = &model.GarageList{
	List: []*model.Garage{
		garage1,
	},
}

var garageListBuyer = &model.GarageListByUser{
	List: map[string]*model.GarageList{
		user1.Id: garageList,
	},
}

func main() {

	// =========== original
	fmt.Printf("# ==== Original\n       %#v \n", user1)

	// =========== as string
	fmt.Printf("# ==== As String\n       %v \n", user1.String())

	var buf bytes.Buffer
	err1 := (&jsonpb.Marshaler{}).Marshal(&buf, garageList)
	if err1 != nil {
		fmt.Println(err1.Error())
		os.Exit(0)
	}

	jsonString := buf.String()
	fmt.Printf("# ==== As JSON String\n       %v \n", jsonString)

	// ====================== Proses unmarshal dari json string ke objek proto, bisa dilakukan lewat dua cara: ==========================

	/*
	Cara pertama : Membuat objek pointer dari struct jsonpb.Unmarshaler, lalu mengakses method .Unmarshal(). Parameter pertama diisi dengan objek io.Reader yang isinya json string, dan parameter kedua adalah receiver.
	*/

	// buf2 := strings.NewReader(jsonString)
	// protoObject := new(model.GarageList)
   
	// err2 := (&jsonpb.Unmarshaler{}).Unmarshal(buf2, protoObject)
	// if err2 != nil {
	// 	fmt.Println(err2.Error())
	// 	os.Exit(0)
	// }
   
	// fmt.Printf("# ==== As String\n       %v \n", protoObject.String())


	// =========== from json string to protobuf object =================
	// Cara ke dua : Menggunakan jsonpb.UnmarshalString, dengan parameter pertama disisipi data json string.
	protoObject := new(model.GarageList)
	err2 := jsonpb.UnmarshalString(jsonString, protoObject)
	if err2 != nil {
		fmt.Println(err2.Error())
		os.Exit(0)
	}
	fmt.Printf("# ==== As String\n       %v \n", protoObject.String())
}

