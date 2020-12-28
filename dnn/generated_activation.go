package cudnn

/* Generated by gencudnn. DO NOT EDIT */

// #include <cudnn.h>
import "C"
import "runtime"

// Activation is a representation of cudnnActivationDescriptor_t.
type Activation struct {
	internal C.cudnnActivationDescriptor_t

	mode       ActivationMode
	reluNanOpt NanPropagation
	coef       float64
}

// NewActivation creates a new Activation.
func NewActivation(mode ActivationMode, reluNanOpt NanPropagation, coef float64) (retVal *Activation, err error) {
	var internal C.cudnnActivationDescriptor_t
	if err := result(C.cudnnCreateActivationDescriptor(&internal)); err != nil {
		return nil, err
	}

	if err := result(C.cudnnSetActivationDescriptor(internal, mode.C(), reluNanOpt.C(), C.double(coef))); err != nil {
		return nil, err
	}

	retVal = &Activation{
		internal:   internal,
		mode:       mode,
		reluNanOpt: reluNanOpt,
		coef:       coef,
	}
	runtime.SetFinalizer(retVal, destroyActivation)
	return retVal, nil
}

// ActivationDesc returns the internal activationDesc.
func (a *Activation) ActivationDesc() *Activation { return a.activationDesc }

// Mode returns the internal mode.
func (a *Activation) Mode() ActivationMode { return a.mode }

// ReluNanOpt returns the internal reluNanOpt.
func (a *Activation) ReluNanOpt() NanPropagation { return a.reluNanOpt }

// Coef returns the internal coef.
func (a *Activation) Coef() float64 { return a.coef }

func destroyActivation(obj *Activation) { C.cudnnDestroyActivationDescriptor(obj.internal) }
