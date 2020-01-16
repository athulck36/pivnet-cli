// Code generated by counterfeiter. DO NOT EDIT.
package commandsfakes

import (
	"sync"

	pivnet "github.com/pivotal-cf/go-pivnet/v3"
	"github.com/pivotal-cf/pivnet-cli/commands"
)

type FakeHelmChartReferenceClient struct {
	ListStub        func(productSlug string, releaseVersion string) error
	listMutex       sync.RWMutex
	listArgsForCall []struct {
		productSlug    string
		releaseVersion string
	}
	listReturns struct {
		result1 error
	}
	listReturnsOnCall map[int]struct {
		result1 error
	}
	GetStub        func(productSlug string, releaseVersion string, helmChartReferenceID int) error
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		productSlug          string
		releaseVersion       string
		helmChartReferenceID int
	}
	getReturns struct {
		result1 error
	}
	getReturnsOnCall map[int]struct {
		result1 error
	}
	CreateStub        func(config pivnet.CreateHelmChartReferenceConfig) error
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		config pivnet.CreateHelmChartReferenceConfig
	}
	createReturns struct {
		result1 error
	}
	createReturnsOnCall map[int]struct {
		result1 error
	}
	UpdateStub        func(productSlug string, helmChartReferenceID int, description *string, docsURL *string, systemRequirements *[]string) error
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		productSlug          string
		helmChartReferenceID int
		description          *string
		docsURL              *string
		systemRequirements   *[]string
	}
	updateReturns struct {
		result1 error
	}
	updateReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteStub        func(productSlug string, helmChartReferenceID int) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		productSlug          string
		helmChartReferenceID int
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	AddToReleaseStub        func(productSlug string, helmChartReferenceID int, releaseVersion string) error
	addToReleaseMutex       sync.RWMutex
	addToReleaseArgsForCall []struct {
		productSlug          string
		helmChartReferenceID int
		releaseVersion       string
	}
	addToReleaseReturns struct {
		result1 error
	}
	addToReleaseReturnsOnCall map[int]struct {
		result1 error
	}
	RemoveFromReleaseStub        func(productSlug string, helmChartReferenceID int, releaseVersion string) error
	removeFromReleaseMutex       sync.RWMutex
	removeFromReleaseArgsForCall []struct {
		productSlug          string
		helmChartReferenceID int
		releaseVersion       string
	}
	removeFromReleaseReturns struct {
		result1 error
	}
	removeFromReleaseReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeHelmChartReferenceClient) List(productSlug string, releaseVersion string) error {
	fake.listMutex.Lock()
	ret, specificReturn := fake.listReturnsOnCall[len(fake.listArgsForCall)]
	fake.listArgsForCall = append(fake.listArgsForCall, struct {
		productSlug    string
		releaseVersion string
	}{productSlug, releaseVersion})
	fake.recordInvocation("List", []interface{}{productSlug, releaseVersion})
	fake.listMutex.Unlock()
	if fake.ListStub != nil {
		return fake.ListStub(productSlug, releaseVersion)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.listReturns.result1
}

func (fake *FakeHelmChartReferenceClient) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *FakeHelmChartReferenceClient) ListArgsForCall(i int) (string, string) {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return fake.listArgsForCall[i].productSlug, fake.listArgsForCall[i].releaseVersion
}

func (fake *FakeHelmChartReferenceClient) ListReturns(result1 error) {
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeHelmChartReferenceClient) ListReturnsOnCall(i int, result1 error) {
	fake.ListStub = nil
	if fake.listReturnsOnCall == nil {
		fake.listReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.listReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeHelmChartReferenceClient) Get(productSlug string, releaseVersion string, helmChartReferenceID int) error {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		productSlug          string
		releaseVersion       string
		helmChartReferenceID int
	}{productSlug, releaseVersion, helmChartReferenceID})
	fake.recordInvocation("Get", []interface{}{productSlug, releaseVersion, helmChartReferenceID})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(productSlug, releaseVersion, helmChartReferenceID)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getReturns.result1
}

func (fake *FakeHelmChartReferenceClient) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeHelmChartReferenceClient) GetArgsForCall(i int) (string, string, int) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return fake.getArgsForCall[i].productSlug, fake.getArgsForCall[i].releaseVersion, fake.getArgsForCall[i].helmChartReferenceID
}

func (fake *FakeHelmChartReferenceClient) GetReturns(result1 error) {
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeHelmChartReferenceClient) GetReturnsOnCall(i int, result1 error) {
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeHelmChartReferenceClient) Create(config pivnet.CreateHelmChartReferenceConfig) error {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		config pivnet.CreateHelmChartReferenceConfig
	}{config})
	fake.recordInvocation("Create", []interface{}{config})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(config)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.createReturns.result1
}

func (fake *FakeHelmChartReferenceClient) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeHelmChartReferenceClient) CreateArgsForCall(i int) pivnet.CreateHelmChartReferenceConfig {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].config
}

func (fake *FakeHelmChartReferenceClient) CreateReturns(result1 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeHelmChartReferenceClient) CreateReturnsOnCall(i int, result1 error) {
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeHelmChartReferenceClient) Update(productSlug string, helmChartReferenceID int, description *string, docsURL *string, systemRequirements *[]string) error {
	fake.updateMutex.Lock()
	ret, specificReturn := fake.updateReturnsOnCall[len(fake.updateArgsForCall)]
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		productSlug          string
		helmChartReferenceID int
		description          *string
		docsURL              *string
		systemRequirements   *[]string
	}{productSlug, helmChartReferenceID, description, docsURL, systemRequirements})
	fake.recordInvocation("Update", []interface{}{productSlug, helmChartReferenceID, description, docsURL, systemRequirements})
	fake.updateMutex.Unlock()
	if fake.UpdateStub != nil {
		return fake.UpdateStub(productSlug, helmChartReferenceID, description, docsURL, systemRequirements)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.updateReturns.result1
}

func (fake *FakeHelmChartReferenceClient) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeHelmChartReferenceClient) UpdateArgsForCall(i int) (string, int, *string, *string, *[]string) {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return fake.updateArgsForCall[i].productSlug, fake.updateArgsForCall[i].helmChartReferenceID, fake.updateArgsForCall[i].description, fake.updateArgsForCall[i].docsURL, fake.updateArgsForCall[i].systemRequirements
}

func (fake *FakeHelmChartReferenceClient) UpdateReturns(result1 error) {
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeHelmChartReferenceClient) UpdateReturnsOnCall(i int, result1 error) {
	fake.UpdateStub = nil
	if fake.updateReturnsOnCall == nil {
		fake.updateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeHelmChartReferenceClient) Delete(productSlug string, helmChartReferenceID int) error {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		productSlug          string
		helmChartReferenceID int
	}{productSlug, helmChartReferenceID})
	fake.recordInvocation("Delete", []interface{}{productSlug, helmChartReferenceID})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(productSlug, helmChartReferenceID)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteReturns.result1
}

func (fake *FakeHelmChartReferenceClient) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeHelmChartReferenceClient) DeleteArgsForCall(i int) (string, int) {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].productSlug, fake.deleteArgsForCall[i].helmChartReferenceID
}

func (fake *FakeHelmChartReferenceClient) DeleteReturns(result1 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeHelmChartReferenceClient) DeleteReturnsOnCall(i int, result1 error) {
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeHelmChartReferenceClient) AddToRelease(productSlug string, helmChartReferenceID int, releaseVersion string) error {
	fake.addToReleaseMutex.Lock()
	ret, specificReturn := fake.addToReleaseReturnsOnCall[len(fake.addToReleaseArgsForCall)]
	fake.addToReleaseArgsForCall = append(fake.addToReleaseArgsForCall, struct {
		productSlug          string
		helmChartReferenceID int
		releaseVersion       string
	}{productSlug, helmChartReferenceID, releaseVersion})
	fake.recordInvocation("AddToRelease", []interface{}{productSlug, helmChartReferenceID, releaseVersion})
	fake.addToReleaseMutex.Unlock()
	if fake.AddToReleaseStub != nil {
		return fake.AddToReleaseStub(productSlug, helmChartReferenceID, releaseVersion)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.addToReleaseReturns.result1
}

func (fake *FakeHelmChartReferenceClient) AddToReleaseCallCount() int {
	fake.addToReleaseMutex.RLock()
	defer fake.addToReleaseMutex.RUnlock()
	return len(fake.addToReleaseArgsForCall)
}

func (fake *FakeHelmChartReferenceClient) AddToReleaseArgsForCall(i int) (string, int, string) {
	fake.addToReleaseMutex.RLock()
	defer fake.addToReleaseMutex.RUnlock()
	return fake.addToReleaseArgsForCall[i].productSlug, fake.addToReleaseArgsForCall[i].helmChartReferenceID, fake.addToReleaseArgsForCall[i].releaseVersion
}

func (fake *FakeHelmChartReferenceClient) AddToReleaseReturns(result1 error) {
	fake.AddToReleaseStub = nil
	fake.addToReleaseReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeHelmChartReferenceClient) AddToReleaseReturnsOnCall(i int, result1 error) {
	fake.AddToReleaseStub = nil
	if fake.addToReleaseReturnsOnCall == nil {
		fake.addToReleaseReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addToReleaseReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeHelmChartReferenceClient) RemoveFromRelease(productSlug string, helmChartReferenceID int, releaseVersion string) error {
	fake.removeFromReleaseMutex.Lock()
	ret, specificReturn := fake.removeFromReleaseReturnsOnCall[len(fake.removeFromReleaseArgsForCall)]
	fake.removeFromReleaseArgsForCall = append(fake.removeFromReleaseArgsForCall, struct {
		productSlug          string
		helmChartReferenceID int
		releaseVersion       string
	}{productSlug, helmChartReferenceID, releaseVersion})
	fake.recordInvocation("RemoveFromRelease", []interface{}{productSlug, helmChartReferenceID, releaseVersion})
	fake.removeFromReleaseMutex.Unlock()
	if fake.RemoveFromReleaseStub != nil {
		return fake.RemoveFromReleaseStub(productSlug, helmChartReferenceID, releaseVersion)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.removeFromReleaseReturns.result1
}

func (fake *FakeHelmChartReferenceClient) RemoveFromReleaseCallCount() int {
	fake.removeFromReleaseMutex.RLock()
	defer fake.removeFromReleaseMutex.RUnlock()
	return len(fake.removeFromReleaseArgsForCall)
}

func (fake *FakeHelmChartReferenceClient) RemoveFromReleaseArgsForCall(i int) (string, int, string) {
	fake.removeFromReleaseMutex.RLock()
	defer fake.removeFromReleaseMutex.RUnlock()
	return fake.removeFromReleaseArgsForCall[i].productSlug, fake.removeFromReleaseArgsForCall[i].helmChartReferenceID, fake.removeFromReleaseArgsForCall[i].releaseVersion
}

func (fake *FakeHelmChartReferenceClient) RemoveFromReleaseReturns(result1 error) {
	fake.RemoveFromReleaseStub = nil
	fake.removeFromReleaseReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeHelmChartReferenceClient) RemoveFromReleaseReturnsOnCall(i int, result1 error) {
	fake.RemoveFromReleaseStub = nil
	if fake.removeFromReleaseReturnsOnCall == nil {
		fake.removeFromReleaseReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.removeFromReleaseReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeHelmChartReferenceClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.addToReleaseMutex.RLock()
	defer fake.addToReleaseMutex.RUnlock()
	fake.removeFromReleaseMutex.RLock()
	defer fake.removeFromReleaseMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeHelmChartReferenceClient) recordInvocation(key string, args []interface{}) {
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

var _ commands.HelmChartReferenceClient = new(FakeHelmChartReferenceClient)
