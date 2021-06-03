package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/kaitolucifer/protobuf-example/proto/complex"
	"github.com/kaitolucifer/protobuf-example/proto/enum"
	"github.com/kaitolucifer/protobuf-example/proto/simple"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func main() {
	// simple example
	sm := NewSimple()
	// Protobuf to/from file
	WriteToFile("./simple.bin", sm)
	var newSM simple.SimpleMessage
	ReadFromFile("./simple.bin", &newSM)
	fmt.Println(&newSM)
	// Protobuf to/from JSON
	smJSON := ToJSON(sm)
	fmt.Println(smJSON)
	var anotherSM simple.SimpleMessage
	FromJSON(smJSON, &anotherSM)
	fmt.Println(&anotherSM)

	// enum example
	em := NewEnum()
	fmt.Println(em)

	// complex
	cm := NewComplex()
	fmt.Println(cm)
}

func NewSimple() *simple.SimpleMessage {
	sm := simple.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My Simple Message",
		SampleList: []int32{1, 4, 7, 8},
	}

	return &sm
}

func NewEnum() *enum.EnumMessage {
	em := enum.EnumMessage{
		Id:           42,
		DayOfTheWeek: enum.DayOfTheWeek_THURSDAY,
	}

	return &em
}

func NewComplex() *complex.ComplexMessage {
	cm := complex.ComplexMessage{
		Dummy: &complex.DummyMessage{Id: 1, Name: "first message"},
		MultipleDummpy: []*complex.DummyMessage{
			{Id: 2, Name: "second message"},
			{Id: 3, Name: "third message"},
		},
	}

	return &cm
}

func WriteToFile(fname string, pb proto.Message) {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("can't serailise to bytes", err)
	}
	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("can't write to file", err)
	}
	fmt.Println("data has been written")
}

func ReadFromFile(fname string, pb proto.Message) {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("can't read file", err)
	}

	err = proto.Unmarshal(in, pb)
	if err != nil {
		log.Fatalln("couldn't put the bytes into the protobuf struct")
	}

	fmt.Println("data has been readed")
}

func ToJSON(pb proto.Message) string {
	marshaler := protojson.MarshalOptions{}
	out, err := marshaler.Marshal(pb)
	if err != nil {
		log.Fatalln("can't convert to JSON", err)
	}

	return string(out)
}

func FromJSON(in string, pb proto.Message) {
	marshaler := protojson.UnmarshalOptions{}
	err := marshaler.Unmarshal([]byte(in), pb)
	if err != nil {
		log.Fatalln("can't convert from JSON", err)
	}
}
