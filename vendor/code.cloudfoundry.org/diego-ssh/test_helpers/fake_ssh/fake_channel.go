// This file was generated by counterfeiter
package fake_ssh

import (
	"io"
	"sync"

	"golang.org/x/crypto/ssh"
)

type FakeChannel struct {
	ReadStub        func(data []byte) (int, error)
	readMutex       sync.RWMutex
	readArgsForCall []struct {
		data []byte
	}
	readReturns struct {
		result1 int
		result2 error
	}
	WriteStub        func(data []byte) (int, error)
	writeMutex       sync.RWMutex
	writeArgsForCall []struct {
		data []byte
	}
	writeReturns struct {
		result1 int
		result2 error
	}
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct{}
	closeReturns     struct {
		result1 error
	}
	CloseWriteStub        func() error
	closeWriteMutex       sync.RWMutex
	closeWriteArgsForCall []struct{}
	closeWriteReturns     struct {
		result1 error
	}
	SendRequestStub        func(name string, wantReply bool, payload []byte) (bool, error)
	sendRequestMutex       sync.RWMutex
	sendRequestArgsForCall []struct {
		name      string
		wantReply bool
		payload   []byte
	}
	sendRequestReturns struct {
		result1 bool
		result2 error
	}
	StderrStub        func() io.ReadWriter
	stderrMutex       sync.RWMutex
	stderrArgsForCall []struct{}
	stderrReturns     struct {
		result1 io.ReadWriter
	}
}

func (fake *FakeChannel) Read(data []byte) (int, error) {
	fake.readMutex.Lock()
	fake.readArgsForCall = append(fake.readArgsForCall, struct {
		data []byte
	}{data})
	fake.readMutex.Unlock()
	if fake.ReadStub != nil {
		return fake.ReadStub(data)
	} else {
		return fake.readReturns.result1, fake.readReturns.result2
	}
}

func (fake *FakeChannel) ReadCallCount() int {
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	return len(fake.readArgsForCall)
}

func (fake *FakeChannel) ReadArgsForCall(i int) []byte {
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	return fake.readArgsForCall[i].data
}

func (fake *FakeChannel) ReadReturns(result1 int, result2 error) {
	fake.ReadStub = nil
	fake.readReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeChannel) Write(data []byte) (int, error) {
	fake.writeMutex.Lock()
	fake.writeArgsForCall = append(fake.writeArgsForCall, struct {
		data []byte
	}{data})
	fake.writeMutex.Unlock()
	if fake.WriteStub != nil {
		return fake.WriteStub(data)
	} else {
		return fake.writeReturns.result1, fake.writeReturns.result2
	}
}

func (fake *FakeChannel) WriteCallCount() int {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	return len(fake.writeArgsForCall)
}

func (fake *FakeChannel) WriteArgsForCall(i int) []byte {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	return fake.writeArgsForCall[i].data
}

func (fake *FakeChannel) WriteReturns(result1 int, result2 error) {
	fake.WriteStub = nil
	fake.writeReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeChannel) Close() error {
	fake.closeMutex.Lock()
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		return fake.CloseStub()
	} else {
		return fake.closeReturns.result1
	}
}

func (fake *FakeChannel) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeChannel) CloseReturns(result1 error) {
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeChannel) CloseWrite() error {
	fake.closeWriteMutex.Lock()
	fake.closeWriteArgsForCall = append(fake.closeWriteArgsForCall, struct{}{})
	fake.closeWriteMutex.Unlock()
	if fake.CloseWriteStub != nil {
		return fake.CloseWriteStub()
	} else {
		return fake.closeWriteReturns.result1
	}
}

func (fake *FakeChannel) CloseWriteCallCount() int {
	fake.closeWriteMutex.RLock()
	defer fake.closeWriteMutex.RUnlock()
	return len(fake.closeWriteArgsForCall)
}

func (fake *FakeChannel) CloseWriteReturns(result1 error) {
	fake.CloseWriteStub = nil
	fake.closeWriteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeChannel) SendRequest(name string, wantReply bool, payload []byte) (bool, error) {
	fake.sendRequestMutex.Lock()
	fake.sendRequestArgsForCall = append(fake.sendRequestArgsForCall, struct {
		name      string
		wantReply bool
		payload   []byte
	}{name, wantReply, payload})
	fake.sendRequestMutex.Unlock()
	if fake.SendRequestStub != nil {
		return fake.SendRequestStub(name, wantReply, payload)
	} else {
		return fake.sendRequestReturns.result1, fake.sendRequestReturns.result2
	}
}

func (fake *FakeChannel) SendRequestCallCount() int {
	fake.sendRequestMutex.RLock()
	defer fake.sendRequestMutex.RUnlock()
	return len(fake.sendRequestArgsForCall)
}

func (fake *FakeChannel) SendRequestArgsForCall(i int) (string, bool, []byte) {
	fake.sendRequestMutex.RLock()
	defer fake.sendRequestMutex.RUnlock()
	return fake.sendRequestArgsForCall[i].name, fake.sendRequestArgsForCall[i].wantReply, fake.sendRequestArgsForCall[i].payload
}

func (fake *FakeChannel) SendRequestReturns(result1 bool, result2 error) {
	fake.SendRequestStub = nil
	fake.sendRequestReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeChannel) Stderr() io.ReadWriter {
	fake.stderrMutex.Lock()
	fake.stderrArgsForCall = append(fake.stderrArgsForCall, struct{}{})
	fake.stderrMutex.Unlock()
	if fake.StderrStub != nil {
		return fake.StderrStub()
	} else {
		return fake.stderrReturns.result1
	}
}

func (fake *FakeChannel) StderrCallCount() int {
	fake.stderrMutex.RLock()
	defer fake.stderrMutex.RUnlock()
	return len(fake.stderrArgsForCall)
}

func (fake *FakeChannel) StderrReturns(result1 io.ReadWriter) {
	fake.StderrStub = nil
	fake.stderrReturns = struct {
		result1 io.ReadWriter
	}{result1}
}

var _ ssh.Channel = new(FakeChannel)
