package main

import (
	"fmt"
	complexpb_ "github.com/IkezawaYuki/protobuf-lesson-go/lesson-pb/src/complex"
	enumpb "github.com/IkezawaYuki/protobuf-lesson-go/lesson-pb/src/enum_example"
	tutorial "github.com/IkezawaYuki/protobuf-lesson-go/lesson-pb/src/examples"
	simplepb "github.com/IkezawaYuki/protobuf-lesson-go/lesson-pb/src/simple"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/timestamp"
	"io/ioutil"
	"log"
)

func main(){
	//sm := doSimple()
	// doEnum()
	// sm := doComplex()
	ab := doExample()
	readAndWriteDemo(ab)
	jsonDemo(ab)
}

func doExample()*tutorial.AddressBook {
	ab := tutorial.AddressBook{
		People:               []*tutorial.Person{
			{
				Name:                 "test0",
				Id:                   0,
				Email:                "test0@gmail.com",
				Phones:               []*tutorial.Person_PhoneNumber{
					{
						Number: "08013009ho",
						Type:   tutorial.Person_HOME,
					},
					{
						Number: "08013009mo",
						Type:   tutorial.Person_MOBILE,
					},
				},
				LastUpdated:          &timestamp.Timestamp{
					Seconds:              0,
					Nanos:                0,
				},
			},
			{
				Name:                 "test1",
				Id:                   1,
				Email:                "test1@hotmail.com",
				Phones:               []*tutorial.Person_PhoneNumber{
					{
						Number: "0801309wo",
						Type:   tutorial.Person_WORK,
					},
				},
				LastUpdated:       &timestamp.Timestamp{
					Seconds:              0,
					Nanos:                0,
				},
			},
			{
				Name:                 "test2",
				Id:                   2,
				Email:                "test2@docomo.ne.jp",
				Phones:               []*tutorial.Person_PhoneNumber{
					{
						Number: "08013009wo2",
						Type:   tutorial.Person_WORK,
					},
				},
				LastUpdated:          &timestamp.Timestamp{
					Seconds:              0,
					Nanos:                0,
				},
			},
		},
	}
	return &ab
}

func doComplex()*complexpb_.ComplexMessage {
	cm := complexpb_.ComplexMessage{
		OneDummy:             &complexpb_.DummyMessage{
			Id:                   1,
			Name:                 "First Message",
		},
		MultipleDummy:        []*complexpb_.DummyMessage{
			{
				Id:   2,
				Name: "Second Message",
			},
			{
				Id:   3,
				Name: "Third Message",
			},
		},
	}
	return &cm
}

func doEnum(){
	em := enumpb.EnumMessage{
		Id:           42,
		DayOfTheWeek: enumpb.DayOfTheWeek_SATURDAY,
	}
	fmt.Println(em)
}

func jsonDemo(sm proto.Message){
	smAsString := toJSON(sm)
	fmt.Println(smAsString)
	sm2 := &tutorial.AddressBook{}
	fromJSON(smAsString, sm2)
	fmt.Println(sm2)
}

func fromJSON(in string, pb proto.Message){
	err := jsonpb.UnmarshalString(in, pb)
	if err != nil{
		log.Fatalf("couldn't unmarshal the JSON int the pb :%v", err)
	}
}

func toJSON(pb proto.Message)string{
	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(pb)
	if err != nil{
		log.Fatalf("can't convert to JSON :%v", err)
	}
	return out
}

func readAndWriteDemo(sm proto.Message){
	_ = writeToFile("addressbook.bin", sm)
	sm2 := &tutorial.AddressBook{}
	_ = readFromFile("addressbook.bin", sm2)
	fmt.Println(sm2)
}

func readFromFile(fname string, pb proto.Message)error{
	in, err := ioutil.ReadFile(fname)
	if err != nil{
		log.Fatalf("something went wrong when reading the file : %v", err)
	}

	err = proto.Unmarshal(in, pb)
	if err != nil{
		log.Fatalf("couldn't put the bytes into the protocol buffers struct : %v", err)
	}
	return nil
}

func writeToFile(fname string, pb proto.Message)error{
	out, err := proto.Marshal(pb)
	if err != nil{
		log.Fatalf("can't serialise to bytes: %v", err)
	}
	if err := ioutil.WriteFile(fname, out, 0644); err != nil{
		log.Fatalf("can't write to file: %v", err)
	}
	return nil
}

func doSimple() *simplepb.SimpleMessage {
	 sm := simplepb.SimpleMessage{
		 Id:                   12345,
		 IsSimple:             true,
		 Name:                 "my simple message",
		 SampleList:           []int32{1,2,4,8},
	 }
	 fmt.Println(sm)

	 sm.Name = "I renamed you"
	 fmt.Println("This ID is :", sm.GetId())
	 return &sm
}