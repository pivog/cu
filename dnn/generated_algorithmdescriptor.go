package cudnn

/* Generated by gencudnn. DO NOT EDIT */

// #include <cudnn.h>
import "C"
import "runtime"

// AlgorithmDescriptor is a representation of cudnnAlgorithmDescriptor_t.
type AlgorithmDescriptor struct {
	internal C.cudnnAlgorithmDescriptor_t

	algorithm TODO
}

// NewAlgorithmDescriptor creates a new AlgorithmDescriptor.
func NewAlgorithmDescriptor(algorithm TODO) (retVal *AlgorithmDescriptor, err error) {
	var internal C.cudnnAlgorithmDescriptor_t
	if err := result(C.cudnnCreateAlgorithmDescriptor(&internal)); err != nil {
		return nil, err
	}

	if err := result(C.cudnnSetAlgorithmDescriptor(internal, algorithm)); err != nil {
		return nil, err
	}

	retVal = &AlgorithmDescriptor{
		internal:  internal,
		algorithm: algorithm,
	}
	runtime.SetFinalizer(retVal, destroyAlgorithmDescriptor)
	return retVal, nil
}

// AlgoDesc returns the internal algoDesc.
func (a *AlgorithmDescriptor) AlgoDesc() *AlgorithmDescriptor { return a.algoDesc }

//TODO: "cudnnSetAlgorithmDescriptor": Parameter 1 Skipped "algorithm" of struct{union Algorithm;} - unmapped type

func destroyAlgorithmDescriptor(obj *AlgorithmDescriptor) {
	C.cudnnDestroyAlgorithmDescriptor(obj.internal)
}
