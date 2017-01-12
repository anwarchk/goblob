// This file was generated by counterfeiter
package blobstorefakes

import (
	"io"
	"sync"

	"github.com/c0-ops/goblob/blobstore"
)

type FakeBlobstore struct {
	NameStub        func() string
	nameMutex       sync.RWMutex
	nameArgsForCall []struct{}
	nameReturns     struct {
		result1 string
	}
	ListStub        func() ([]*blobstore.Blob, error)
	listMutex       sync.RWMutex
	listArgsForCall []struct{}
	listReturns     struct {
		result1 []*blobstore.Blob
		result2 error
	}
	ReadStub        func(src *blobstore.Blob) (io.ReadCloser, error)
	readMutex       sync.RWMutex
	readArgsForCall []struct {
		src *blobstore.Blob
	}
	readReturns struct {
		result1 io.ReadCloser
		result2 error
	}
	ChecksumStub        func(src *blobstore.Blob) (string, error)
	checksumMutex       sync.RWMutex
	checksumArgsForCall []struct {
		src *blobstore.Blob
	}
	checksumReturns struct {
		result1 string
		result2 error
	}
	WriteStub        func(dst *blobstore.Blob, src io.Reader) error
	writeMutex       sync.RWMutex
	writeArgsForCall []struct {
		dst *blobstore.Blob
		src io.Reader
	}
	writeReturns struct {
		result1 error
	}
	ExistsStub        func(*blobstore.Blob) bool
	existsMutex       sync.RWMutex
	existsArgsForCall []struct {
		arg1 *blobstore.Blob
	}
	existsReturns struct {
		result1 bool
	}
	NewBucketIteratorStub        func(string) (blobstore.BucketIterator, error)
	newBucketIteratorMutex       sync.RWMutex
	newBucketIteratorArgsForCall []struct {
		arg1 string
	}
	newBucketIteratorReturns struct {
		result1 blobstore.BucketIterator
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBlobstore) Name() string {
	fake.nameMutex.Lock()
	fake.nameArgsForCall = append(fake.nameArgsForCall, struct{}{})
	fake.recordInvocation("Name", []interface{}{})
	fake.nameMutex.Unlock()
	if fake.NameStub != nil {
		return fake.NameStub()
	} else {
		return fake.nameReturns.result1
	}
}

func (fake *FakeBlobstore) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *FakeBlobstore) NameReturns(result1 string) {
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeBlobstore) List() ([]*blobstore.Blob, error) {
	fake.listMutex.Lock()
	fake.listArgsForCall = append(fake.listArgsForCall, struct{}{})
	fake.recordInvocation("List", []interface{}{})
	fake.listMutex.Unlock()
	if fake.ListStub != nil {
		return fake.ListStub()
	} else {
		return fake.listReturns.result1, fake.listReturns.result2
	}
}

func (fake *FakeBlobstore) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *FakeBlobstore) ListReturns(result1 []*blobstore.Blob, result2 error) {
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 []*blobstore.Blob
		result2 error
	}{result1, result2}
}

func (fake *FakeBlobstore) Read(src *blobstore.Blob) (io.ReadCloser, error) {
	fake.readMutex.Lock()
	fake.readArgsForCall = append(fake.readArgsForCall, struct {
		src *blobstore.Blob
	}{src})
	fake.recordInvocation("Read", []interface{}{src})
	fake.readMutex.Unlock()
	if fake.ReadStub != nil {
		return fake.ReadStub(src)
	} else {
		return fake.readReturns.result1, fake.readReturns.result2
	}
}

func (fake *FakeBlobstore) ReadCallCount() int {
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	return len(fake.readArgsForCall)
}

func (fake *FakeBlobstore) ReadArgsForCall(i int) *blobstore.Blob {
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	return fake.readArgsForCall[i].src
}

func (fake *FakeBlobstore) ReadReturns(result1 io.ReadCloser, result2 error) {
	fake.ReadStub = nil
	fake.readReturns = struct {
		result1 io.ReadCloser
		result2 error
	}{result1, result2}
}

func (fake *FakeBlobstore) Checksum(src *blobstore.Blob) (string, error) {
	fake.checksumMutex.Lock()
	fake.checksumArgsForCall = append(fake.checksumArgsForCall, struct {
		src *blobstore.Blob
	}{src})
	fake.recordInvocation("Checksum", []interface{}{src})
	fake.checksumMutex.Unlock()
	if fake.ChecksumStub != nil {
		return fake.ChecksumStub(src)
	} else {
		return fake.checksumReturns.result1, fake.checksumReturns.result2
	}
}

func (fake *FakeBlobstore) ChecksumCallCount() int {
	fake.checksumMutex.RLock()
	defer fake.checksumMutex.RUnlock()
	return len(fake.checksumArgsForCall)
}

func (fake *FakeBlobstore) ChecksumArgsForCall(i int) *blobstore.Blob {
	fake.checksumMutex.RLock()
	defer fake.checksumMutex.RUnlock()
	return fake.checksumArgsForCall[i].src
}

func (fake *FakeBlobstore) ChecksumReturns(result1 string, result2 error) {
	fake.ChecksumStub = nil
	fake.checksumReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeBlobstore) Write(dst *blobstore.Blob, src io.Reader) error {
	fake.writeMutex.Lock()
	fake.writeArgsForCall = append(fake.writeArgsForCall, struct {
		dst *blobstore.Blob
		src io.Reader
	}{dst, src})
	fake.recordInvocation("Write", []interface{}{dst, src})
	fake.writeMutex.Unlock()
	if fake.WriteStub != nil {
		return fake.WriteStub(dst, src)
	} else {
		return fake.writeReturns.result1
	}
}

func (fake *FakeBlobstore) WriteCallCount() int {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	return len(fake.writeArgsForCall)
}

func (fake *FakeBlobstore) WriteArgsForCall(i int) (*blobstore.Blob, io.Reader) {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	return fake.writeArgsForCall[i].dst, fake.writeArgsForCall[i].src
}

func (fake *FakeBlobstore) WriteReturns(result1 error) {
	fake.WriteStub = nil
	fake.writeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBlobstore) Exists(arg1 *blobstore.Blob) bool {
	fake.existsMutex.Lock()
	fake.existsArgsForCall = append(fake.existsArgsForCall, struct {
		arg1 *blobstore.Blob
	}{arg1})
	fake.recordInvocation("Exists", []interface{}{arg1})
	fake.existsMutex.Unlock()
	if fake.ExistsStub != nil {
		return fake.ExistsStub(arg1)
	} else {
		return fake.existsReturns.result1
	}
}

func (fake *FakeBlobstore) ExistsCallCount() int {
	fake.existsMutex.RLock()
	defer fake.existsMutex.RUnlock()
	return len(fake.existsArgsForCall)
}

func (fake *FakeBlobstore) ExistsArgsForCall(i int) *blobstore.Blob {
	fake.existsMutex.RLock()
	defer fake.existsMutex.RUnlock()
	return fake.existsArgsForCall[i].arg1
}

func (fake *FakeBlobstore) ExistsReturns(result1 bool) {
	fake.ExistsStub = nil
	fake.existsReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeBlobstore) NewBucketIterator(arg1 string) (blobstore.BucketIterator, error) {
	fake.newBucketIteratorMutex.Lock()
	fake.newBucketIteratorArgsForCall = append(fake.newBucketIteratorArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("NewBucketIterator", []interface{}{arg1})
	fake.newBucketIteratorMutex.Unlock()
	if fake.NewBucketIteratorStub != nil {
		return fake.NewBucketIteratorStub(arg1)
	} else {
		return fake.newBucketIteratorReturns.result1, fake.newBucketIteratorReturns.result2
	}
}

func (fake *FakeBlobstore) NewBucketIteratorCallCount() int {
	fake.newBucketIteratorMutex.RLock()
	defer fake.newBucketIteratorMutex.RUnlock()
	return len(fake.newBucketIteratorArgsForCall)
}

func (fake *FakeBlobstore) NewBucketIteratorArgsForCall(i int) string {
	fake.newBucketIteratorMutex.RLock()
	defer fake.newBucketIteratorMutex.RUnlock()
	return fake.newBucketIteratorArgsForCall[i].arg1
}

func (fake *FakeBlobstore) NewBucketIteratorReturns(result1 blobstore.BucketIterator, result2 error) {
	fake.NewBucketIteratorStub = nil
	fake.newBucketIteratorReturns = struct {
		result1 blobstore.BucketIterator
		result2 error
	}{result1, result2}
}

func (fake *FakeBlobstore) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	fake.checksumMutex.RLock()
	defer fake.checksumMutex.RUnlock()
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	fake.existsMutex.RLock()
	defer fake.existsMutex.RUnlock()
	fake.newBucketIteratorMutex.RLock()
	defer fake.newBucketIteratorMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeBlobstore) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ blobstore.Blobstore = new(FakeBlobstore)
