package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/Siddhartha15/golang-exercise/src/tutorialpb"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func main() {

	p := getPerson()
	readWriteDemo(p)

	jsonDemo(p)

}

func jsonDemo(p proto.Message) {
	pjson := toJson(p)
	fmt.Printf("to Json %T %s \n", pjson, pjson)
	pmsg := &tutorialpb.Person{}
	fromJson(pjson, pmsg)
	fmt.Println("from Json", pmsg)
}
func toJson(pb proto.Message) string {
	// return protojson.Format(pb)
	out, err := protojson.Marshal(pb)

	if err != nil {
		log.Fatalln("toJson error: ", err)
		return ""
	}
	fmt.Println("Success to json")
	return string(out)
}

func fromJson(str string, pb proto.Message) {
	err := protojson.Unmarshal([]byte(str), pb)
	if err != nil {
		log.Fatalln("from Json error: ", err)
	}
}

func readWriteDemo(p proto.Message) {
	writeToFile("output.bin", p)
	rpmsg := &tutorialpb.Person{}
	readFromFile("output.bin", rpmsg)
	fmt.Println("Read: ", rpmsg)
}

func writeToFile(fname string, pmsg proto.Message) error {
	out, err := proto.Marshal(pmsg)

	if err != nil {
		log.Fatalln("Something wrong with serialize :", err)
		return err
	}

	err2 := ioutil.WriteFile(fname, out, 0644)

	if err2 != nil {
		log.Fatalln("Error while eriting: ", err2)
		return err2
	}
	fmt.Println("Data written")

	return nil
}

func readFromFile(fname string, rpmsg proto.Message) error {

	out, err := ioutil.ReadFile(fname)

	if err != nil {
		log.Fatalln("Error while reading: ", err)
		return err
	}

	err2 := proto.Unmarshal(out, rpmsg)
	if err2 != nil {
		log.Fatalln("error while deserializing: ", err2)
		return err2
	}

	fmt.Println("Succes reading from file")
	return nil
}

func getPerson() *tutorialpb.Person {
	person := tutorialpb.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*tutorialpb.Person_PhoneNumber{
			{Number: "555-4321", Type: tutorialpb.Person_HOME},
		},
	}

	fmt.Println(person.GetId())
	return &person
}
