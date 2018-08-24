package main

//go:generate esc -o generator/data/tpl/tpl.go -pkg=tpl generator/template

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/golang/protobuf/proto"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"

	"version.uuzu.com/Merlion/protoapi/generator"
)

func main() {
	args := os.Args
	var input []byte
	// if command has parameter then call protoc, if no parameter meaning calling from protoc
	if len(args) > 2 {
		var err error
		protoFile := args[2]
		input, err = ioutil.ReadFile(protoFile)
		if err != nil {
			log.Fatalf("reading file %s error: %s\n", protoFile, err.Error())
		}
		// protoapi binary
		executable, _ := os.Executable()

		langFlag := flag.String("lang", "echo:.", "output language and output directory")
		flag.Parse()

		flags := []string{}
		flags = append(flags, "--plugin=protoc-gen-custom="+executable)
		flags = append(flags, "--custom_out=lang="+*langFlag)
		flags = append(flags, protoFile)

		// Run protoc command
		cmd := exec.Command("protoc", flags...)
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}

	} else {
		var err error
		input, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatalf("reading stdin error: %s\n", err.Error())
		}

		request := new(plugin.CodeGeneratorRequest)

		proto.Unmarshal(input, request)
		if len(request.FileToGenerate) != 1 {
			log.Fatalf("input files are： %v\nwe only support one proto file\n", request.FileToGenerate)
		}

		response, err := generator.Generate(request)
		if err != nil {
			log.Fatalf("generate response failed: %s\n", err.Error())
		}

		output, err := proto.Marshal(response)
		if err != nil {
			log.Fatalf("create response failed: %s\n", err.Error())
		}
		os.Stdout.Write(output)
	}
}
