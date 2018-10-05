// Code generated by adapter-generator. DO NOT EDIT.

package adapter

import (
	"github.com/gogo/protobuf/proto"
	. "github.com/ligato/vpp-agent/plugins/kvscheduler/api"
	"github.com/ligato/vpp-agent/plugins/vppv2/ifplugin/ifaceidx"
	"github.com/ligato/vpp-agent/plugins/vppv2/model/interfaces"
)

////////// type-safe key-value pair with metadata //////////

type InterfaceKVWithMetadata struct {
	Key      string
	Value    *interfaces.Interface
	Metadata *ifaceidx.IfaceMetadata
	Origin   ValueOrigin
}

////////// type-safe Descriptor structure //////////

type InterfaceDescriptor struct {
	Name               string
	KeySelector        KeySelector
	ValueTypeName      string
	KeyLabel           func(key string) string
	ValueComparator    func(key string, v1, v2 *interfaces.Interface) bool
	NBKeyPrefix        string
	WithMetadata       bool
	MetadataMapFactory MetadataMapFactory
	Add                func(key string, value *interfaces.Interface) (metadata *ifaceidx.IfaceMetadata, err error)
	Delete             func(key string, value *interfaces.Interface, metadata *ifaceidx.IfaceMetadata) error
	Modify             func(key string, oldValue, newValue *interfaces.Interface, oldMetadata *ifaceidx.IfaceMetadata) (newMetadata *ifaceidx.IfaceMetadata, err error)
	ModifyWithRecreate func(key string, oldValue, newValue *interfaces.Interface, metadata *ifaceidx.IfaceMetadata) bool
	Update             func(key string, value *interfaces.Interface, metadata *ifaceidx.IfaceMetadata) error
	IsRetriableFailure func(err error) bool
	Dependencies       func(key string, value *interfaces.Interface) []Dependency
	DerivedValues      func(key string, value *interfaces.Interface) []KeyValuePair
	Dump               func(correlate []InterfaceKVWithMetadata) ([]InterfaceKVWithMetadata, error)
	DumpDependencies   []string /* descriptor name */
}

////////// Descriptor adapter //////////

type InterfaceDescriptorAdapter struct {
	descriptor *InterfaceDescriptor
}

func NewInterfaceDescriptor(typedDescriptor *InterfaceDescriptor) *KVDescriptor {
	adapter := &InterfaceDescriptorAdapter{descriptor: typedDescriptor}
	descriptor := &KVDescriptor{
		Name:               typedDescriptor.Name,
		KeySelector:        typedDescriptor.KeySelector,
		ValueTypeName:      typedDescriptor.ValueTypeName,
		KeyLabel:           typedDescriptor.KeyLabel,
		NBKeyPrefix:        typedDescriptor.NBKeyPrefix,
		WithMetadata:       typedDescriptor.WithMetadata,
		MetadataMapFactory: typedDescriptor.MetadataMapFactory,
		IsRetriableFailure: typedDescriptor.IsRetriableFailure,
		DumpDependencies:   typedDescriptor.DumpDependencies,
	}
	if typedDescriptor.ValueComparator != nil {
		descriptor.ValueComparator = adapter.ValueComparator
	}
	if typedDescriptor.Add != nil {
		descriptor.Add = adapter.Add
	}
	if typedDescriptor.Delete != nil {
		descriptor.Delete = adapter.Delete
	}
	if typedDescriptor.Modify != nil {
		descriptor.Modify = adapter.Modify
	}
	if typedDescriptor.ModifyWithRecreate != nil {
		descriptor.ModifyWithRecreate = adapter.ModifyWithRecreate
	}
	if typedDescriptor.Update != nil {
		descriptor.Update = adapter.Update
	}
	if typedDescriptor.Dependencies != nil {
		descriptor.Dependencies = adapter.Dependencies
	}
	if typedDescriptor.DerivedValues != nil {
		descriptor.DerivedValues = adapter.DerivedValues
	}
	if typedDescriptor.Dump != nil {
		descriptor.Dump = adapter.Dump
	}
	return descriptor
}

func (da *InterfaceDescriptorAdapter) ValueComparator(key string, v1, v2 proto.Message) bool {
	typedV1, err1 := castInterfaceValue(key, v1)
	typedV2, err2 := castInterfaceValue(key, v2)
	if err1 != nil || err2 != nil {
		return false
	}
	return da.descriptor.ValueComparator(key, typedV1, typedV2)
}

func (da *InterfaceDescriptorAdapter) Add(key string, value proto.Message) (metadata Metadata, err error) {
	typedValue, err := castInterfaceValue(key, value)
	if err != nil {
		return nil, err
	}
	return da.descriptor.Add(key, typedValue)
}

func (da *InterfaceDescriptorAdapter) Modify(key string, oldValue, newValue proto.Message, oldMetadata Metadata) (newMetadata Metadata, err error) {
	oldTypedValue, err := castInterfaceValue(key, oldValue)
	if err != nil {
		return nil, err
	}
	newTypedValue, err := castInterfaceValue(key, newValue)
	if err != nil {
		return nil, err
	}
	typedOldMetadata, err := castInterfaceMetadata(key, oldMetadata)
	if err != nil {
		return nil, err
	}
	return da.descriptor.Modify(key, oldTypedValue, newTypedValue, typedOldMetadata)
}

func (da *InterfaceDescriptorAdapter) Delete(key string, value proto.Message, metadata Metadata) error {
	typedValue, err := castInterfaceValue(key, value)
	if err != nil {
		return err
	}
	typedMetadata, err := castInterfaceMetadata(key, metadata)
	if err != nil {
		return err
	}
	return da.descriptor.Delete(key, typedValue, typedMetadata)
}

func (da *InterfaceDescriptorAdapter) ModifyWithRecreate(key string, oldValue, newValue proto.Message, metadata Metadata) bool {
	oldTypedValue, err := castInterfaceValue(key, oldValue)
	if err != nil {
		return true
	}
	newTypedValue, err := castInterfaceValue(key, newValue)
	if err != nil {
		return true
	}
	typedMetadata, err := castInterfaceMetadata(key, metadata)
	if err != nil {
		return true
	}
	return da.descriptor.ModifyWithRecreate(key, oldTypedValue, newTypedValue, typedMetadata)
}

func (da *InterfaceDescriptorAdapter) Update(key string, value proto.Message, metadata Metadata) error {
	typedValue, err := castInterfaceValue(key, value)
	if err != nil {
		return err
	}
	typedMetadata, err := castInterfaceMetadata(key, metadata)
	if err != nil {
		return err
	}
	return da.descriptor.Update(key, typedValue, typedMetadata)
}

func (da *InterfaceDescriptorAdapter) Dependencies(key string, value proto.Message) []Dependency {
	typedValue, err := castInterfaceValue(key, value)
	if err != nil {
		return nil
	}
	return da.descriptor.Dependencies(key, typedValue)
}

func (da *InterfaceDescriptorAdapter) DerivedValues(key string, value proto.Message) []KeyValuePair {
	typedValue, err := castInterfaceValue(key, value)
	if err != nil {
		return nil
	}
	return da.descriptor.DerivedValues(key, typedValue)
}

func (da *InterfaceDescriptorAdapter) Dump(correlate []KVWithMetadata) ([]KVWithMetadata, error) {
	var correlateWithType []InterfaceKVWithMetadata
	for _, kvpair := range correlate {
		typedValue, err := castInterfaceValue(kvpair.Key, kvpair.Value)
		if err != nil {
			continue
		}
		typedMetadata, err := castInterfaceMetadata(kvpair.Key, kvpair.Metadata)
		if err != nil {
			continue
		}
		correlateWithType = append(correlateWithType,
			InterfaceKVWithMetadata{
				Key:      kvpair.Key,
				Value:    typedValue,
				Metadata: typedMetadata,
				Origin:   kvpair.Origin,
			})
	}

	typedDump, err := da.descriptor.Dump(correlateWithType)
	if err != nil {
		return nil, err
	}
	var dump []KVWithMetadata
	for _, typedKVWithMetadata := range typedDump {
		kvWithMetadata := KVWithMetadata{
			Key:      typedKVWithMetadata.Key,
			Metadata: typedKVWithMetadata.Metadata,
			Origin:   typedKVWithMetadata.Origin,
		}
		kvWithMetadata.Value = typedKVWithMetadata.Value
		dump = append(dump, kvWithMetadata)
	}
	return dump, err
}

////////// Helper methods //////////

func castInterfaceValue(key string, value proto.Message) (*interfaces.Interface, error) {
	typedValue, ok := value.(*interfaces.Interface)
	if !ok {
		return nil, ErrInvalidValueType(key, value)
	}
	return typedValue, nil
}

func castInterfaceMetadata(key string, metadata Metadata) (*ifaceidx.IfaceMetadata, error) {
	if metadata == nil {
		return nil, nil
	}
	typedMetadata, ok := metadata.(*ifaceidx.IfaceMetadata)
	if !ok {
		return nil, ErrInvalidMetadataType(key)
	}
	return typedMetadata, nil
}
