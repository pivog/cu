package cudnn

/* Generated by gencudnn. DO NOT EDIT */

// #include <cudnn.h>
import "C"
import "runtime"

// Backend is a representation of cudnnBackendDescriptor_t.
type Backend struct {
	internal C.cudnnBackendDescriptor_t

	attributeName   BackendAttributeName
	attributeType   BackendAttributeType
	elementCount    TODO
	arrayOfElements Memory
}

// NewBackend creates a new Backend.
func NewBackend(attributeName BackendAttributeName, attributeType BackendAttributeType, elementCount TODO, arrayOfElements Memory) (retVal *Backend, err error) {
	var internal C.cudnnBackendDescriptor_t
	if err := result(C.cudnnBackendCreateDescriptor(&internal)); err != nil {
		return nil, err
	}

	if err := result(C.cudnnBackendSetAttribute(internal, attributeName.C(), attributeType.C(), elementCount, arrayOfElements.Pointer())); err != nil {
		return nil, err
	}

	retVal = &Backend{
		internal:        internal,
		attributeName:   attributeName,
		attributeType:   attributeType,
		elementCount:    elementCount,
		arrayOfElements: arrayOfElements,
	}
	runtime.SetFinalizer(retVal, destroyBackend)
	return retVal, nil
}

// Descriptor returns the internal descriptor.
func (b *Backend) Descriptor() *Backend { return b.descriptor }

// AttributeName returns the internal attributeName.
func (b *Backend) AttributeName() BackendAttributeName { return b.attributeName }

// AttributeType returns the internal attributeType.
func (b *Backend) AttributeType() BackendAttributeType { return b.attributeType }

//TODO: "cudnnBackendSetAttribute": Parameter 3 Skipped "elementCount" of long - unmapped type

// ArrayOfElements returns the internal arrayOfElements.
func (b *Backend) ArrayOfElements() Memory { return b.arrayOfElements }

func destroyBackend(obj *Backend) { C.cudnnBackendDestroyDescriptor(obj.internal) }
