package cudnn

/* Generated by gencudnn. DO NOT EDIT */

// #include <cudnn_v7.h>
import "C"
import "runtime"

// Reduction is a representation of cudnnReduceTensorDescriptor_t.
type Reduction struct {
	internal C.cudnnReduceTensorDescriptor_t

	reduceTensorOp          ReduceTensorOp
	reduceTensorCompType    DataType
	reduceTensorNanOpt      NanPropagation
	reduceTensorIndices     ReduceTensorIndices
	reduceTensorIndicesType IndicesType
}

// NewReduction creates a new Reduction.
func NewReduction(reduceTensorOp ReduceTensorOp, reduceTensorCompType DataType, reduceTensorNanOpt NanPropagation, reduceTensorIndices ReduceTensorIndices, reduceTensorIndicesType IndicesType) (retVal *Reduction, err error) {
	var internal C.cudnnReduceTensorDescriptor_t
	if err := result(C.cudnnCreateReduceTensorDescriptor(&internal)); err != nil {
		return nil, err
	}

	if err := result(C.cudnnSetReduceTensorDescriptor(internal, reduceTensorOp.C(), reduceTensorCompType.C(), reduceTensorNanOpt.C(), reduceTensorIndices.C(), reduceTensorIndicesType.C())); err != nil {
		return nil, err
	}

	retVal = &Reduction{
		internal:                internal,
		reduceTensorOp:          reduceTensorOp,
		reduceTensorCompType:    reduceTensorCompType,
		reduceTensorNanOpt:      reduceTensorNanOpt,
		reduceTensorIndices:     reduceTensorIndices,
		reduceTensorIndicesType: reduceTensorIndicesType,
	}
	runtime.SetFinalizer(retVal, destroyReduction)
	return retVal, nil
}

// ReduceTensorOp returns the internal reduceTensorOp.
func (r *Reduction) ReduceTensorOp() ReduceTensorOp { return r.reduceTensorOp }

// ReduceTensorCompType returns the internal reduceTensorCompType.
func (r *Reduction) ReduceTensorCompType() DataType { return r.reduceTensorCompType }

// ReduceTensorNanOpt returns the internal reduceTensorNanOpt.
func (r *Reduction) ReduceTensorNanOpt() NanPropagation { return r.reduceTensorNanOpt }

// ReduceTensorIndices returns the internal reduceTensorIndices.
func (r *Reduction) ReduceTensorIndices() ReduceTensorIndices { return r.reduceTensorIndices }

// ReduceTensorIndicesType returns the internal reduceTensorIndicesType.
func (r *Reduction) ReduceTensorIndicesType() IndicesType { return r.reduceTensorIndicesType }

func destroyReduction(obj *Reduction) { C.cudnnDestroyReduceTensorDescriptor(obj.internal) }
