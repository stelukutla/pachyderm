// Code generated by protoc-gen-protolog
// source: pfs/fuse/fuse.proto
// DO NOT EDIT!

package fuse

import "go.pedge.io/protolog"

func init() {
	protolog.Register("fuse.Root", protolog.MessageType_MESSAGE_TYPE_EVENT, func() protolog.Message { return &Root{} })
	protolog.Register("fuse.DirectoryAttr", protolog.MessageType_MESSAGE_TYPE_EVENT, func() protolog.Message { return &DirectoryAttr{} })
	protolog.Register("fuse.DirectoryLookup", protolog.MessageType_MESSAGE_TYPE_EVENT, func() protolog.Message { return &DirectoryLookup{} })
	protolog.Register("fuse.DirectoryReadDirAll", protolog.MessageType_MESSAGE_TYPE_EVENT, func() protolog.Message { return &DirectoryReadDirAll{} })
	protolog.Register("fuse.DirectoryCreate", protolog.MessageType_MESSAGE_TYPE_EVENT, func() protolog.Message { return &DirectoryCreate{} })
	protolog.Register("fuse.DirectoryMkdir", protolog.MessageType_MESSAGE_TYPE_EVENT, func() protolog.Message { return &DirectoryMkdir{} })
	protolog.Register("fuse.FileAttr", protolog.MessageType_MESSAGE_TYPE_EVENT, func() protolog.Message { return &FileAttr{} })
	protolog.Register("fuse.FileRead", protolog.MessageType_MESSAGE_TYPE_EVENT, func() protolog.Message { return &FileRead{} })
	protolog.Register("fuse.FileOpen", protolog.MessageType_MESSAGE_TYPE_EVENT, func() protolog.Message { return &FileOpen{} })
	protolog.Register("fuse.FileWrite", protolog.MessageType_MESSAGE_TYPE_EVENT, func() protolog.Message { return &FileWrite{} })
}

func (m *Root) ProtologName() string {
	return "fuse.Root"
}
func (m *DirectoryAttr) ProtologName() string {
	return "fuse.DirectoryAttr"
}
func (m *DirectoryLookup) ProtologName() string {
	return "fuse.DirectoryLookup"
}
func (m *DirectoryReadDirAll) ProtologName() string {
	return "fuse.DirectoryReadDirAll"
}
func (m *DirectoryCreate) ProtologName() string {
	return "fuse.DirectoryCreate"
}
func (m *DirectoryMkdir) ProtologName() string {
	return "fuse.DirectoryMkdir"
}
func (m *FileAttr) ProtologName() string {
	return "fuse.FileAttr"
}
func (m *FileRead) ProtologName() string {
	return "fuse.FileRead"
}
func (m *FileOpen) ProtologName() string {
	return "fuse.FileOpen"
}
func (m *FileWrite) ProtologName() string {
	return "fuse.FileWrite"
}