// Code generated by protoc-gen-protolog
// source: protolog.proto
// DO NOT EDIT!

package dockervolume

import "go.pedge.io/protolog"

func init() {
	protolog.Register("dockervolume.MethodInvocation", func() protolog.Message { return &MethodInvocation{} })
}

func (m *MethodInvocation) ProtologName() string {
	return "dockervolume.MethodInvocation"
}