// Code generated by counterfeiter. DO NOT EDIT.
package curlfakes

import (
	"io"
	"net/http"
	"sync"

	"github.com/pivotal-cf/pivnet-cli/v3/commands/curl"
)

type FakePivnetClient struct {
	MakeRequestStub        func(string, string, int, io.Reader) (*http.Response, error)
	makeRequestMutex       sync.RWMutex
	makeRequestArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 int
		arg4 io.Reader
	}
	makeRequestReturns struct {
		result1 *http.Response
		result2 error
	}
	makeRequestReturnsOnCall map[int]struct {
		result1 *http.Response
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePivnetClient) MakeRequest(arg1 string, arg2 string, arg3 int, arg4 io.Reader) (*http.Response, error) {
	fake.makeRequestMutex.Lock()
	ret, specificReturn := fake.makeRequestReturnsOnCall[len(fake.makeRequestArgsForCall)]
	fake.makeRequestArgsForCall = append(fake.makeRequestArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 int
		arg4 io.Reader
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("MakeRequest", []interface{}{arg1, arg2, arg3, arg4})
	fake.makeRequestMutex.Unlock()
	if fake.MakeRequestStub != nil {
		return fake.MakeRequestStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.makeRequestReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePivnetClient) MakeRequestCallCount() int {
	fake.makeRequestMutex.RLock()
	defer fake.makeRequestMutex.RUnlock()
	return len(fake.makeRequestArgsForCall)
}

func (fake *FakePivnetClient) MakeRequestCalls(stub func(string, string, int, io.Reader) (*http.Response, error)) {
	fake.makeRequestMutex.Lock()
	defer fake.makeRequestMutex.Unlock()
	fake.MakeRequestStub = stub
}

func (fake *FakePivnetClient) MakeRequestArgsForCall(i int) (string, string, int, io.Reader) {
	fake.makeRequestMutex.RLock()
	defer fake.makeRequestMutex.RUnlock()
	argsForCall := fake.makeRequestArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakePivnetClient) MakeRequestReturns(result1 *http.Response, result2 error) {
	fake.makeRequestMutex.Lock()
	defer fake.makeRequestMutex.Unlock()
	fake.MakeRequestStub = nil
	fake.makeRequestReturns = struct {
		result1 *http.Response
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) MakeRequestReturnsOnCall(i int, result1 *http.Response, result2 error) {
	fake.makeRequestMutex.Lock()
	defer fake.makeRequestMutex.Unlock()
	fake.MakeRequestStub = nil
	if fake.makeRequestReturnsOnCall == nil {
		fake.makeRequestReturnsOnCall = make(map[int]struct {
			result1 *http.Response
			result2 error
		})
	}
	fake.makeRequestReturnsOnCall[i] = struct {
		result1 *http.Response
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.makeRequestMutex.RLock()
	defer fake.makeRequestMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePivnetClient) recordInvocation(key string, args []interface{}) {
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

var _ curl.PivnetClient = new(FakePivnetClient)
