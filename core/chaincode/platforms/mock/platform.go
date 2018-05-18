// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"archive/tar"
	"sync"

	"github.com/hyperledger/fabric/core/chaincode/platforms/ccmetadata"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type Platform struct {
	NameStub        func() string
	nameMutex       sync.RWMutex
	nameArgsForCall []struct{}
	nameReturns     struct {
		result1 string
	}
	nameReturnsOnCall map[int]struct {
		result1 string
	}
	ValidatePathStub        func(path string) error
	validatePathMutex       sync.RWMutex
	validatePathArgsForCall []struct {
		path string
	}
	validatePathReturns struct {
		result1 error
	}
	validatePathReturnsOnCall map[int]struct {
		result1 error
	}
	ValidateDeploymentSpecStub        func(spec *pb.ChaincodeDeploymentSpec) error
	validateDeploymentSpecMutex       sync.RWMutex
	validateDeploymentSpecArgsForCall []struct {
		spec *pb.ChaincodeDeploymentSpec
	}
	validateDeploymentSpecReturns struct {
		result1 error
	}
	validateDeploymentSpecReturnsOnCall map[int]struct {
		result1 error
	}
	GetDeploymentPayloadStub        func(spec *pb.ChaincodeSpec) ([]byte, error)
	getDeploymentPayloadMutex       sync.RWMutex
	getDeploymentPayloadArgsForCall []struct {
		spec *pb.ChaincodeSpec
	}
	getDeploymentPayloadReturns struct {
		result1 []byte
		result2 error
	}
	getDeploymentPayloadReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	GenerateDockerfileStub        func(spec *pb.ChaincodeDeploymentSpec) (string, error)
	generateDockerfileMutex       sync.RWMutex
	generateDockerfileArgsForCall []struct {
		spec *pb.ChaincodeDeploymentSpec
	}
	generateDockerfileReturns struct {
		result1 string
		result2 error
	}
	generateDockerfileReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	GenerateDockerBuildStub        func(spec *pb.ChaincodeDeploymentSpec, tw *tar.Writer) error
	generateDockerBuildMutex       sync.RWMutex
	generateDockerBuildArgsForCall []struct {
		spec *pb.ChaincodeDeploymentSpec
		tw   *tar.Writer
	}
	generateDockerBuildReturns struct {
		result1 error
	}
	generateDockerBuildReturnsOnCall map[int]struct {
		result1 error
	}
	GetMetadataProviderStub        func(spec *pb.ChaincodeDeploymentSpec) ccmetadata.MetadataProvider
	getMetadataProviderMutex       sync.RWMutex
	getMetadataProviderArgsForCall []struct {
		spec *pb.ChaincodeDeploymentSpec
	}
	getMetadataProviderReturns struct {
		result1 ccmetadata.MetadataProvider
	}
	getMetadataProviderReturnsOnCall map[int]struct {
		result1 ccmetadata.MetadataProvider
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Platform) Name() string {
	fake.nameMutex.Lock()
	ret, specificReturn := fake.nameReturnsOnCall[len(fake.nameArgsForCall)]
	fake.nameArgsForCall = append(fake.nameArgsForCall, struct{}{})
	fake.recordInvocation("Name", []interface{}{})
	fake.nameMutex.Unlock()
	if fake.NameStub != nil {
		return fake.NameStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.nameReturns.result1
}

func (fake *Platform) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *Platform) NameReturns(result1 string) {
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 string
	}{result1}
}

func (fake *Platform) NameReturnsOnCall(i int, result1 string) {
	fake.NameStub = nil
	if fake.nameReturnsOnCall == nil {
		fake.nameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.nameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *Platform) ValidatePath(path string) error {
	fake.validatePathMutex.Lock()
	ret, specificReturn := fake.validatePathReturnsOnCall[len(fake.validatePathArgsForCall)]
	fake.validatePathArgsForCall = append(fake.validatePathArgsForCall, struct {
		path string
	}{path})
	fake.recordInvocation("ValidatePath", []interface{}{path})
	fake.validatePathMutex.Unlock()
	if fake.ValidatePathStub != nil {
		return fake.ValidatePathStub(path)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.validatePathReturns.result1
}

func (fake *Platform) ValidatePathCallCount() int {
	fake.validatePathMutex.RLock()
	defer fake.validatePathMutex.RUnlock()
	return len(fake.validatePathArgsForCall)
}

func (fake *Platform) ValidatePathArgsForCall(i int) string {
	fake.validatePathMutex.RLock()
	defer fake.validatePathMutex.RUnlock()
	return fake.validatePathArgsForCall[i].path
}

func (fake *Platform) ValidatePathReturns(result1 error) {
	fake.ValidatePathStub = nil
	fake.validatePathReturns = struct {
		result1 error
	}{result1}
}

func (fake *Platform) ValidatePathReturnsOnCall(i int, result1 error) {
	fake.ValidatePathStub = nil
	if fake.validatePathReturnsOnCall == nil {
		fake.validatePathReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.validatePathReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Platform) ValidateDeploymentSpec(spec *pb.ChaincodeDeploymentSpec) error {
	fake.validateDeploymentSpecMutex.Lock()
	ret, specificReturn := fake.validateDeploymentSpecReturnsOnCall[len(fake.validateDeploymentSpecArgsForCall)]
	fake.validateDeploymentSpecArgsForCall = append(fake.validateDeploymentSpecArgsForCall, struct {
		spec *pb.ChaincodeDeploymentSpec
	}{spec})
	fake.recordInvocation("ValidateDeploymentSpec", []interface{}{spec})
	fake.validateDeploymentSpecMutex.Unlock()
	if fake.ValidateDeploymentSpecStub != nil {
		return fake.ValidateDeploymentSpecStub(spec)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.validateDeploymentSpecReturns.result1
}

func (fake *Platform) ValidateDeploymentSpecCallCount() int {
	fake.validateDeploymentSpecMutex.RLock()
	defer fake.validateDeploymentSpecMutex.RUnlock()
	return len(fake.validateDeploymentSpecArgsForCall)
}

func (fake *Platform) ValidateDeploymentSpecArgsForCall(i int) *pb.ChaincodeDeploymentSpec {
	fake.validateDeploymentSpecMutex.RLock()
	defer fake.validateDeploymentSpecMutex.RUnlock()
	return fake.validateDeploymentSpecArgsForCall[i].spec
}

func (fake *Platform) ValidateDeploymentSpecReturns(result1 error) {
	fake.ValidateDeploymentSpecStub = nil
	fake.validateDeploymentSpecReturns = struct {
		result1 error
	}{result1}
}

func (fake *Platform) ValidateDeploymentSpecReturnsOnCall(i int, result1 error) {
	fake.ValidateDeploymentSpecStub = nil
	if fake.validateDeploymentSpecReturnsOnCall == nil {
		fake.validateDeploymentSpecReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.validateDeploymentSpecReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Platform) GetDeploymentPayload(spec *pb.ChaincodeSpec) ([]byte, error) {
	fake.getDeploymentPayloadMutex.Lock()
	ret, specificReturn := fake.getDeploymentPayloadReturnsOnCall[len(fake.getDeploymentPayloadArgsForCall)]
	fake.getDeploymentPayloadArgsForCall = append(fake.getDeploymentPayloadArgsForCall, struct {
		spec *pb.ChaincodeSpec
	}{spec})
	fake.recordInvocation("GetDeploymentPayload", []interface{}{spec})
	fake.getDeploymentPayloadMutex.Unlock()
	if fake.GetDeploymentPayloadStub != nil {
		return fake.GetDeploymentPayloadStub(spec)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getDeploymentPayloadReturns.result1, fake.getDeploymentPayloadReturns.result2
}

func (fake *Platform) GetDeploymentPayloadCallCount() int {
	fake.getDeploymentPayloadMutex.RLock()
	defer fake.getDeploymentPayloadMutex.RUnlock()
	return len(fake.getDeploymentPayloadArgsForCall)
}

func (fake *Platform) GetDeploymentPayloadArgsForCall(i int) *pb.ChaincodeSpec {
	fake.getDeploymentPayloadMutex.RLock()
	defer fake.getDeploymentPayloadMutex.RUnlock()
	return fake.getDeploymentPayloadArgsForCall[i].spec
}

func (fake *Platform) GetDeploymentPayloadReturns(result1 []byte, result2 error) {
	fake.GetDeploymentPayloadStub = nil
	fake.getDeploymentPayloadReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *Platform) GetDeploymentPayloadReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.GetDeploymentPayloadStub = nil
	if fake.getDeploymentPayloadReturnsOnCall == nil {
		fake.getDeploymentPayloadReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.getDeploymentPayloadReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *Platform) GenerateDockerfile(spec *pb.ChaincodeDeploymentSpec) (string, error) {
	fake.generateDockerfileMutex.Lock()
	ret, specificReturn := fake.generateDockerfileReturnsOnCall[len(fake.generateDockerfileArgsForCall)]
	fake.generateDockerfileArgsForCall = append(fake.generateDockerfileArgsForCall, struct {
		spec *pb.ChaincodeDeploymentSpec
	}{spec})
	fake.recordInvocation("GenerateDockerfile", []interface{}{spec})
	fake.generateDockerfileMutex.Unlock()
	if fake.GenerateDockerfileStub != nil {
		return fake.GenerateDockerfileStub(spec)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.generateDockerfileReturns.result1, fake.generateDockerfileReturns.result2
}

func (fake *Platform) GenerateDockerfileCallCount() int {
	fake.generateDockerfileMutex.RLock()
	defer fake.generateDockerfileMutex.RUnlock()
	return len(fake.generateDockerfileArgsForCall)
}

func (fake *Platform) GenerateDockerfileArgsForCall(i int) *pb.ChaincodeDeploymentSpec {
	fake.generateDockerfileMutex.RLock()
	defer fake.generateDockerfileMutex.RUnlock()
	return fake.generateDockerfileArgsForCall[i].spec
}

func (fake *Platform) GenerateDockerfileReturns(result1 string, result2 error) {
	fake.GenerateDockerfileStub = nil
	fake.generateDockerfileReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *Platform) GenerateDockerfileReturnsOnCall(i int, result1 string, result2 error) {
	fake.GenerateDockerfileStub = nil
	if fake.generateDockerfileReturnsOnCall == nil {
		fake.generateDockerfileReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.generateDockerfileReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *Platform) GenerateDockerBuild(spec *pb.ChaincodeDeploymentSpec, tw *tar.Writer) error {
	fake.generateDockerBuildMutex.Lock()
	ret, specificReturn := fake.generateDockerBuildReturnsOnCall[len(fake.generateDockerBuildArgsForCall)]
	fake.generateDockerBuildArgsForCall = append(fake.generateDockerBuildArgsForCall, struct {
		spec *pb.ChaincodeDeploymentSpec
		tw   *tar.Writer
	}{spec, tw})
	fake.recordInvocation("GenerateDockerBuild", []interface{}{spec, tw})
	fake.generateDockerBuildMutex.Unlock()
	if fake.GenerateDockerBuildStub != nil {
		return fake.GenerateDockerBuildStub(spec, tw)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.generateDockerBuildReturns.result1
}

func (fake *Platform) GenerateDockerBuildCallCount() int {
	fake.generateDockerBuildMutex.RLock()
	defer fake.generateDockerBuildMutex.RUnlock()
	return len(fake.generateDockerBuildArgsForCall)
}

func (fake *Platform) GenerateDockerBuildArgsForCall(i int) (*pb.ChaincodeDeploymentSpec, *tar.Writer) {
	fake.generateDockerBuildMutex.RLock()
	defer fake.generateDockerBuildMutex.RUnlock()
	return fake.generateDockerBuildArgsForCall[i].spec, fake.generateDockerBuildArgsForCall[i].tw
}

func (fake *Platform) GenerateDockerBuildReturns(result1 error) {
	fake.GenerateDockerBuildStub = nil
	fake.generateDockerBuildReturns = struct {
		result1 error
	}{result1}
}

func (fake *Platform) GenerateDockerBuildReturnsOnCall(i int, result1 error) {
	fake.GenerateDockerBuildStub = nil
	if fake.generateDockerBuildReturnsOnCall == nil {
		fake.generateDockerBuildReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.generateDockerBuildReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Platform) GetMetadataProvider(spec *pb.ChaincodeDeploymentSpec) ccmetadata.MetadataProvider {
	fake.getMetadataProviderMutex.Lock()
	ret, specificReturn := fake.getMetadataProviderReturnsOnCall[len(fake.getMetadataProviderArgsForCall)]
	fake.getMetadataProviderArgsForCall = append(fake.getMetadataProviderArgsForCall, struct {
		spec *pb.ChaincodeDeploymentSpec
	}{spec})
	fake.recordInvocation("GetMetadataProvider", []interface{}{spec})
	fake.getMetadataProviderMutex.Unlock()
	if fake.GetMetadataProviderStub != nil {
		return fake.GetMetadataProviderStub(spec)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getMetadataProviderReturns.result1
}

func (fake *Platform) GetMetadataProviderCallCount() int {
	fake.getMetadataProviderMutex.RLock()
	defer fake.getMetadataProviderMutex.RUnlock()
	return len(fake.getMetadataProviderArgsForCall)
}

func (fake *Platform) GetMetadataProviderArgsForCall(i int) *pb.ChaincodeDeploymentSpec {
	fake.getMetadataProviderMutex.RLock()
	defer fake.getMetadataProviderMutex.RUnlock()
	return fake.getMetadataProviderArgsForCall[i].spec
}

func (fake *Platform) GetMetadataProviderReturns(result1 ccmetadata.MetadataProvider) {
	fake.GetMetadataProviderStub = nil
	fake.getMetadataProviderReturns = struct {
		result1 ccmetadata.MetadataProvider
	}{result1}
}

func (fake *Platform) GetMetadataProviderReturnsOnCall(i int, result1 ccmetadata.MetadataProvider) {
	fake.GetMetadataProviderStub = nil
	if fake.getMetadataProviderReturnsOnCall == nil {
		fake.getMetadataProviderReturnsOnCall = make(map[int]struct {
			result1 ccmetadata.MetadataProvider
		})
	}
	fake.getMetadataProviderReturnsOnCall[i] = struct {
		result1 ccmetadata.MetadataProvider
	}{result1}
}

func (fake *Platform) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	fake.validatePathMutex.RLock()
	defer fake.validatePathMutex.RUnlock()
	fake.validateDeploymentSpecMutex.RLock()
	defer fake.validateDeploymentSpecMutex.RUnlock()
	fake.getDeploymentPayloadMutex.RLock()
	defer fake.getDeploymentPayloadMutex.RUnlock()
	fake.generateDockerfileMutex.RLock()
	defer fake.generateDockerfileMutex.RUnlock()
	fake.generateDockerBuildMutex.RLock()
	defer fake.generateDockerBuildMutex.RUnlock()
	fake.getMetadataProviderMutex.RLock()
	defer fake.getMetadataProviderMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Platform) recordInvocation(key string, args []interface{}) {
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
