package gen

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

type Generator struct {
	rawBytes    []byte
	title       string
	description string
}

func NewGenerator(files []*protogen.File, opts ...GeneratorOption) (*Generator, error) {
	g := &Generator{}

	for _, file := range files {
		if !file.Generate {
			continue
		}

		filteredFileDescriptorProto := filterFileDescriptorProto(file.Proto)
		req := &pluginpb.CodeGeneratorRequest{
			ProtoFile: []*descriptorpb.FileDescriptorProto{filteredFileDescriptorProto},
		}
		marshaledBytes := prototext.Format(req)

		g.rawBytes = append(g.rawBytes, []byte(marshaledBytes)...)
	}

	return g, nil
}

// TextProto returns protobuf with textproto format data in bytes.
func (g *Generator) TextProto() ([]byte, error) {
	return g.rawBytes, nil
}

func filterFileDescriptorProto(src *descriptorpb.FileDescriptorProto) *descriptorpb.FileDescriptorProto {
	dst := &descriptorpb.FileDescriptorProto{
		Name:             src.Name,
		Package:          src.Package,
		Dependency:       nil,
		MessageType:      src.MessageType,
		EnumType:         src.EnumType,
		Service:          src.Service,
		Extension:        src.Extension,
		Options:          nil, // remove
		SourceCodeInfo:   nil, // remove
		Syntax:           nil, // remove
		PublicDependency: nil, // remove
		WeakDependency:   nil, // remove
	}
	return dst
}

func (g *Generator) SetTitle(title string) {
	g.title = title
}

func (g *Generator) SetDescription(description string) {
	g.description = description
}
