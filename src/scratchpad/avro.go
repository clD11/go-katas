package scratchpad

import (
	json2 "encoding/json"
	"github.com/linkedin/goavro/v2"
	"log"
)

var avro_schema_txt = `{
        "type": "record",
        "name": "AvroData",
        "namespace": "data.avro",
        "doc": "docstring",
        "fields": [
            {
                "name": "f1",
                "type": "int"
            },
            {
                "name": "f2",
                "type": "string"
            },
            {
                "name": "f3",
                "type": "string"
            },
            {
                "name": "dependencies",
                "type": {
                    "type": "array",
                    "items": "string"
                }
            }
        ]
    }`

// added json to match field names in avro
type SimpleRecord struct {
	F1           int      `avro:"f1" json:"f1"`
	F2           string   `avro:"f2" json:"f2"`
	F3           string   `avro:"f3" json:"f3"`
	Dependencies []string `avro:"dependencies" json:"dependencies"`
}

func encodeDecode(data SimpleRecord) SimpleRecord {

	codec, err := goavro.NewCodec(avro_schema_txt)
	if err != nil {
		log.Fatal(err.Error())
	}

	// encode

	textualIn, err := json2.Marshal(data)
	if err != nil {
		log.Fatal(err.Error())
	}

	nativeIn, _, err := codec.NativeFromTextual(textualIn)
	if err != nil {
		log.Fatal(err.Error())
	}

	binaryIn, err := codec.BinaryFromNative(nil, nativeIn)
	if err != nil {
		log.Fatal(err.Error())
	}

	// decode

	nativeOut, _, err := codec.NativeFromBinary(binaryIn)
	if err != nil {
		log.Fatal(err.Error())
	}

	textualOut, err := codec.TextualFromNative(nil, nativeOut)
	if err != nil {
		log.Fatal(err.Error())
	}

	var out = SimpleRecord{}
	err = json2.Unmarshal(textualOut, &out)
	if err != nil {
		log.Fatal(err.Error())
	}

	return out
}
