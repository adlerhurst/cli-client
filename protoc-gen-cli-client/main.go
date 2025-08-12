package main

import (
	"context"
	"log"

	v7 "github.com/adlerhurst/cli-client/protoc-gen-cli-client/v7"
	"github.com/bufbuild/protoplugin"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/descriptorpb"
)

const version = "0.0.1"

func main() {
	protoplugin.Main(protoplugin.HandlerFunc(handle), protoplugin.WithVersion(version))
}

func handle(
	_ context.Context,
	_ protoplugin.PluginEnv,
	responseWriter protoplugin.ResponseWriter,
	request protoplugin.Request,
) error {
	plugin, err := protogen.Options{}.New(request.CodeGeneratorRequest())
	if err != nil {
		return err
	}
	if err := handleProtogenPlugin(plugin); err != nil {
		plugin.Error(err)
	}
	response := plugin.Response()
	responseWriter.AddCodeGeneratorResponseFiles(response.GetFile()...)
	responseWriter.AddError(response.GetError())
	responseWriter.SetFeatureProto3Optional()
	responseWriter.SetFeatureSupportsEditions(descriptorpb.Edition_EDITION_2023, descriptorpb.Edition_EDITION_2024)
	return nil
}

func handleProtogenPlugin(plugin *protogen.Plugin) error {
	for _, file := range plugin.Files {
		if !file.Generate {
			continue
		}
		err := v7.NewFile(plugin.NewGeneratedFile(string(file.Desc.FullName())+".cli.pb.go", file.GoImportPath), file).Execute()
		log.Println(string(file.Desc.FullName()) + ".cli.pb.go")
		if err != nil {
			return err
		}
	}
	return nil
}

// package main

// import (
// 	"flag"
// 	"log"

// 	// "github.com/adlerhurst/cli-client/protoc-gen-cli-clienttypes"
// 	// v4_types "github.com/adlerhurst/cli-client/protoc-gen-cli-client/types/v4"
// 	// "github.com/adlerhurst/cli-client/protoc-gen-cli-client/v5"
// 	v6 "github.com/adlerhurst/cli-client/protoc-gen-cli-client/v6"
// 	"google.golang.org/protobuf/compiler/protogen"
// 	"google.golang.org/protobuf/reflect/protoregistry"
// 	"google.golang.org/protobuf/types/descriptorpb"
// 	"google.golang.org/protobuf/types/pluginpb"
// )

// var registry protoregistry.Files

// func main() {
// 	// typesV4()
// 	// callV5()
// 	callV6()
// 	// protogen.Options{
// 	// 	ParamFunc: flag.CommandLine.Set,
// 	// }.Run(func(plugin *protogen.Plugin) error {
// 	// 	plugin.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

// 	// 	for _, file := range plugin.Files {
// 	// 		if err := registry.RegisterFile(file.Desc); err != nil {
// 	// 			log.Println("register failed", err)
// 	// 		}

// 	// 		if !file.Generate {
// 	// 			continue
// 	// 		}

// 	// 		for _, svc := range file.Services {
// 	// 			service := types.NewService(svc)
// 	// 			if err := service.Generate(plugin, file); err != nil {
// 	// 				return err
// 	// 			}
// 	// 		}

// 	// 		messages := types.MessagesFromFile(file)
// 	// 		err := messages.GenerateMessages(plugin, file)
// 	// 		if err != nil {
// 	// 			return err
// 	// 		}
// 	// 	}

// 	// 	return nil
// 	// })
// }

// // func typesV4() {
// // 	protogen.Options{
// // 		ParamFunc: flag.CommandLine.Set,
// // 	}.Run(func(plugin *protogen.Plugin) error {
// // 		plugin.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

// // 		for _, file := range plugin.Files {
// // 			if err := registry.RegisterFile(file.Desc); err != nil {
// // 				log.Println("register failed", err)
// // 			}

// // 			if !file.Generate {
// // 				continue
// // 			}

// // 			for _, service := range file.Services {
// // 				if err := v4_types.NewService(service).Generate(plugin, file); err != nil {
// // 					return err
// // 				}
// // 			}

// // 			if err := v4_types.GenerateMessages(plugin, file); err != nil {
// // 				return err
// // 			}
// // 		}

// // 		return nil
// // 	})
// // }

// // func callV5() {
// // 	protogen.Options{
// // 		ParamFunc: flag.CommandLine.Set,
// // 	}.Run(func(plugin *protogen.Plugin) error {
// // 		plugin.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

// // 		for _, file := range plugin.Files {
// // 			if err := registry.RegisterFile(file.Desc); err != nil {
// // 				log.Println("register failed", err)
// // 			}

// // 			if !file.Generate {
// // 				continue
// // 			}

// // 			if err := v5.GenerateMessages(plugin, file); err != nil {
// // 				return err
// // 			}

// // 			if err := v5.GenerateServices(plugin, file); err != nil {
// // 				return err
// // 			}

// // 			// for _, message := range file.Messages {
// // 			// 	if err := v5.GenerateMessage(message); err != nil {
// // 			// 		return err
// // 			// 	}
// // 			// }
// // 		}

// // 		return nil
// // 	})
// // }

// func callV6() {
// 	protogen.Options{
// 		ParamFunc: flag.CommandLine.Set,
// 	}.Run(func(plugin *protogen.Plugin) error {
// 		plugin.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL | pluginpb.CodeGeneratorResponse_FEATURE_SUPPORTS_EDITIONS)
// 		plugin.SupportedEditionsMinimum = descriptorpb.Edition_EDITION_2023
// 		plugin.SupportedEditionsMaximum = descriptorpb.Edition_EDITION_2024

// 		for _, file := range plugin.Files {
// 			if err := registry.RegisterFile(file.Desc); err != nil {
// 				log.Println("register failed", err)
// 			}

// 			if !file.Generate {
// 				continue
// 			}

// 			if err := v6.GenerateMessages(plugin, file); err != nil {
// 				return err
// 			}
// 		}

// 		return nil
// 	})
// }
