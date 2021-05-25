package cudnn

/* Generated by gencudnn. DO NOT EDIT */

// #include <cudnn.h>
import "C"

var _ActivationModeNames = map[ActivationMode]string{
	Sigmoid:     "Sigmoid",
	ReLU:        "ReLU",
	Tanh:        "Tanh",
	ClippedReLU: "ClippedReLU",
	Elu:         "Elu",
	Identity:    "Identity",
}

func (e ActivationMode) String() string { return _ActivationModeNames[e] }

var _BackendAttributeNameNames = map[BackendAttributeName]string{
	BEAttrNamePointwiseMode:                         "BEAttrNamePointwiseMode",
	BEAttrNamePointwiseMathPrec:                     "BEAttrNamePointwiseMathPrec",
	BEAttrNamePointwiseNanPropagation:               "BEAttrNamePointwiseNanPropagation",
	BEAttrNamePointwiseReluLowerClip:                "BEAttrNamePointwiseReluLowerClip",
	BEAttrNamePointwiseReluUpperClip:                "BEAttrNamePointwiseReluUpperClip",
	BEAttrNameConvolutionCompType:                   "BEAttrNameConvolutionCompType",
	BEAttrNameConvolutionConvMode:                   "BEAttrNameConvolutionConvMode",
	BEAttrNameConvolutionDilations:                  "BEAttrNameConvolutionDilations",
	BEAttrNameConvolutionFilterStrides:              "BEAttrNameConvolutionFilterStrides",
	BEAttrNameConvolutionPostPaddings:               "BEAttrNameConvolutionPostPaddings",
	BEAttrNameConvolutionPrePaddings:                "BEAttrNameConvolutionPrePaddings",
	BEAttrNameConvolutionSpatialDims:                "BEAttrNameConvolutionSpatialDims",
	BEAttrNameEngineheurMode:                        "BEAttrNameEngineheurMode",
	BEAttrNameEngineheurOperationGraph:              "BEAttrNameEngineheurOperationGraph",
	BEAttrNameEngineheurResults:                     "BEAttrNameEngineheurResults",
	BEAttrNameEnginecfgEngine:                       "BEAttrNameEnginecfgEngine",
	BEAttrNameEnginecfgIntermediateInfo:             "BEAttrNameEnginecfgIntermediateInfo",
	BEAttrNameEnginecfgKnobChoices:                  "BEAttrNameEnginecfgKnobChoices",
	BEAttrNameExecutionPlanHandle:                   "BEAttrNameExecutionPlanHandle",
	BEAttrNameExecutionPlanEngineConfig:             "BEAttrNameExecutionPlanEngineConfig",
	BEAttrNameExecutionPlanWorkspaceSize:            "BEAttrNameExecutionPlanWorkspaceSize",
	BEAttrNameExecutionPlanComputedIntermediateUids: "BEAttrNameExecutionPlanComputedIntermediateUids",
	BEAttrNameExecutionPlanRunOnlyIntermediateUids:  "BEAttrNameExecutionPlanRunOnlyIntermediateUids",
	BEAttrNameIntermediateInfoUniqueId:              "BEAttrNameIntermediateInfoUniqueId",
	BEAttrNameIntermediateInfoSize:                  "BEAttrNameIntermediateInfoSize",
	BEAttrNameIntermediateInfoDependentDataUids:     "BEAttrNameIntermediateInfoDependentDataUids",
	BEAttrNameIntermediateInfoDependentAttributes:   "BEAttrNameIntermediateInfoDependentAttributes",
	BEAttrNameKnobChoiceKnobType:                    "BEAttrNameKnobChoiceKnobType",
	BEAttrNameKnobChoiceKnobValue:                   "BEAttrNameKnobChoiceKnobValue",
	BEAttrNameOperationConvolutionForwardAlpha:      "BEAttrNameOperationConvolutionForwardAlpha",
	BEAttrNameOperationConvolutionForwardBeta:       "BEAttrNameOperationConvolutionForwardBeta",
	BEAttrNameOperationConvolutionForwardConvDesc:   "BEAttrNameOperationConvolutionForwardConvDesc",
	BEAttrNameOperationConvolutionForwardW:          "BEAttrNameOperationConvolutionForwardW",
	BEAttrNameOperationConvolutionForwardX:          "BEAttrNameOperationConvolutionForwardX",
	BEAttrNameOperationConvolutionForwardY:          "BEAttrNameOperationConvolutionForwardY",
	BEAttrNameOperationConvolutionBwdDataAlpha:      "BEAttrNameOperationConvolutionBwdDataAlpha",
	BEAttrNameOperationConvolutionBwdDataBeta:       "BEAttrNameOperationConvolutionBwdDataBeta",
	BEAttrNameOperationConvolutionBwdDataConvDesc:   "BEAttrNameOperationConvolutionBwdDataConvDesc",
	BEAttrNameOperationConvolutionBwdDataW:          "BEAttrNameOperationConvolutionBwdDataW",
	BEAttrNameOperationConvolutionBwdDataDx:         "BEAttrNameOperationConvolutionBwdDataDx",
	BEAttrNameOperationConvolutionBwdDataDy:         "BEAttrNameOperationConvolutionBwdDataDy",
	BEAttrNameOperationConvolutionBwdFilterAlpha:    "BEAttrNameOperationConvolutionBwdFilterAlpha",
	BEAttrNameOperationConvolutionBwdFilterBeta:     "BEAttrNameOperationConvolutionBwdFilterBeta",
	BEAttrNameOperationConvolutionBwdFilterConvDesc: "BEAttrNameOperationConvolutionBwdFilterConvDesc",
	BEAttrNameOperationConvolutionBwdFilterDw:       "BEAttrNameOperationConvolutionBwdFilterDw",
	BEAttrNameOperationConvolutionBwdFilterX:        "BEAttrNameOperationConvolutionBwdFilterX",
	BEAttrNameOperationConvolutionBwdFilterDy:       "BEAttrNameOperationConvolutionBwdFilterDy",
	BEAttrNameOperationPointwisePwDescriptor:        "BEAttrNameOperationPointwisePwDescriptor",
	BEAttrNameOperationPointwiseXdesc:               "BEAttrNameOperationPointwiseXdesc",
	BEAttrNameOperationPointwiseBdesc:               "BEAttrNameOperationPointwiseBdesc",
	BEAttrNameOperationPointwiseYdesc:               "BEAttrNameOperationPointwiseYdesc",
	BEAttrNameOperationPointwiseAlpha1:              "BEAttrNameOperationPointwiseAlpha1",
	BEAttrNameOperationPointwiseAlpha2:              "BEAttrNameOperationPointwiseAlpha2",
	BEAttrNameOperationGenstatsMode:                 "BEAttrNameOperationGenstatsMode",
	BEAttrNameOperationGenstatsMathPrec:             "BEAttrNameOperationGenstatsMathPrec",
	BEAttrNameOperationGenstatsXdesc:                "BEAttrNameOperationGenstatsXdesc",
	BEAttrNameOperationGenstatsSumdesc:              "BEAttrNameOperationGenstatsSumdesc",
	BEAttrNameOperationGenstatsSqsumdesc:            "BEAttrNameOperationGenstatsSqsumdesc",
	BEAttrNameOperationgraphHandle:                  "BEAttrNameOperationgraphHandle",
	BEAttrNameOperationgraphOps:                     "BEAttrNameOperationgraphOps",
	BEAttrNameOperationgraphEngineGlobalCount:       "BEAttrNameOperationgraphEngineGlobalCount",
	BEAttrNameTensorByteAlignment:                   "BEAttrNameTensorByteAlignment",
	BEAttrNameTensorDataType:                        "BEAttrNameTensorDataType",
	BEAttrNameTensorDimensions:                      "BEAttrNameTensorDimensions",
	BEAttrNameTensorStrides:                         "BEAttrNameTensorStrides",
	BEAttrNameTensorVectorCount:                     "BEAttrNameTensorVectorCount",
	BEAttrNameTensorVectorizedDimension:             "BEAttrNameTensorVectorizedDimension",
	BEAttrNameTensorUniqueId:                        "BEAttrNameTensorUniqueId",
	BEAttrNameTensorIsVirtual:                       "BEAttrNameTensorIsVirtual",
	BEAttrNameVariantPackUniqueIds:                  "BEAttrNameVariantPackUniqueIds",
	BEAttrNameVariantPackDataPointers:               "BEAttrNameVariantPackDataPointers",
	BEAttrNameVariantPackIntermediates:              "BEAttrNameVariantPackIntermediates",
	BEAttrNameVariantPackWorkspace:                  "BEAttrNameVariantPackWorkspace",
	BEAttrNameLayoutInfoTensorUid:                   "BEAttrNameLayoutInfoTensorUid",
	BEAttrNameLayoutInfoTypes:                       "BEAttrNameLayoutInfoTypes",
	BEAttrNameKnobInfoType:                          "BEAttrNameKnobInfoType",
	BEAttrNameKnobInfoMaximumValue:                  "BEAttrNameKnobInfoMaximumValue",
	BEAttrNameKnobInfoMinimumValue:                  "BEAttrNameKnobInfoMinimumValue",
	BEAttrNameKnobInfoStride:                        "BEAttrNameKnobInfoStride",
	BEAttrNameEngineOperationGraph:                  "BEAttrNameEngineOperationGraph",
	BEAttrNameEngineGlobalIndex:                     "BEAttrNameEngineGlobalIndex",
	BEAttrNameEngineKnobInfo:                        "BEAttrNameEngineKnobInfo",
	BEAttrNameEngineNumericalNote:                   "BEAttrNameEngineNumericalNote",
	BEAttrNameEngineLayoutInfo:                      "BEAttrNameEngineLayoutInfo",
}

func (e BackendAttributeName) String() string { return _BackendAttributeNameNames[e] }

var _BackendAttributeTypeNames = map[BackendAttributeType]string{
	BEAttrHandle:            "BEAttrHandle",
	BEAttrDataType:          "BEAttrDataType",
	BEAttrBoolean:           "BEAttrBoolean",
	BEAttrInt64:             "BEAttrInt64",
	BEAttrFloat:             "BEAttrFloat",
	BEAttrDouble:            "BEAttrDouble",
	BEAttrVoidPtr:           "BEAttrVoidPtr",
	BEAttrConvolutionMode:   "BEAttrConvolutionMode",
	BEAttrHeurMode:          "BEAttrHeurMode",
	BEAttrKnobType:          "BEAttrKnobType",
	BEAttrNanPropogation:    "BEAttrNanPropogation",
	BEAttrNumericalNote:     "BEAttrNumericalNote",
	BEAttrLayoutType:        "BEAttrLayoutType",
	BEAttrAttribName:        "BEAttrAttribName",
	BEAttrPointwiseMode:     "BEAttrPointwiseMode",
	BEAttrBackendDescriptor: "BEAttrBackendDescriptor",
	BEAttrGenstatsMode:      "BEAttrGenstatsMode",
}

func (e BackendAttributeType) String() string { return _BackendAttributeTypeNames[e] }

var _BackendDescriptorTypeNames = map[BackendDescriptorType]string{
	BEDescriptorPointwiseDescriptor:                          "BEDescriptorPointwiseDescriptor",
	BEDescriptorConvolutionDescriptor:                        "BEDescriptorConvolutionDescriptor",
	BEDescriptorEngineDescriptor:                             "BEDescriptorEngineDescriptor",
	BEDescriptorEnginecfgDescriptor:                          "BEDescriptorEnginecfgDescriptor",
	BEDescriptorEngineheurDescriptor:                         "BEDescriptorEngineheurDescriptor",
	BEDescriptorExecutionPlanDescriptor:                      "BEDescriptorExecutionPlanDescriptor",
	BEDescriptorIntermediateInfoDescriptor:                   "BEDescriptorIntermediateInfoDescriptor",
	BEDescriptorKnobChoiceDescriptor:                         "BEDescriptorKnobChoiceDescriptor",
	BEDescriptorKnobInfoDescriptor:                           "BEDescriptorKnobInfoDescriptor",
	BEDescriptorLayoutInfoDescriptor:                         "BEDescriptorLayoutInfoDescriptor",
	BEDescriptorOperationConvolutionForwardDescriptor:        "BEDescriptorOperationConvolutionForwardDescriptor",
	BEDescriptorOperationConvolutionBackwardFilterDescriptor: "BEDescriptorOperationConvolutionBackwardFilterDescriptor",
	BEDescriptorOperationConvolutionBackwardDataDescriptor:   "BEDescriptorOperationConvolutionBackwardDataDescriptor",
	BEDescriptorOperationPointwiseDescriptor:                 "BEDescriptorOperationPointwiseDescriptor",
	BEDescriptorOperationGenStatsDescriptor:                  "BEDescriptorOperationGenStatsDescriptor",
	BEDescriptorOperationgraphDescriptor:                     "BEDescriptorOperationgraphDescriptor",
	BEDescriptorVariantPackDescriptor:                        "BEDescriptorVariantPackDescriptor",
	BEDescriptorTensorDescriptor:                             "BEDescriptorTensorDescriptor",
}

func (e BackendDescriptorType) String() string { return _BackendDescriptorTypeNames[e] }

var _BackendHeurModeNames = map[BackendHeurMode]string{
	Instant: "Instant",
	SCount:  "SCount",
}

func (e BackendHeurMode) String() string { return _BackendHeurModeNames[e] }

var _BackendKnobTypeNames = map[BackendKnobType]string{
	SplitK:        "SplitK",
	Swizzle:       "Swizzle",
	TileSize:      "TileSize",
	UseTex:        "UseTex",
	Edge:          "Edge",
	Kblock:        "Kblock",
	Ldga:          "Ldga",
	Ldgb:          "Ldgb",
	ChunkK:        "ChunkK",
	SplitH:        "SplitH",
	WinoTile:      "WinoTile",
	Multiply:      "Multiply",
	SplitKBuf:     "SplitKBuf",
	Tilek:         "Tilek",
	Stages:        "Stages",
	ReductionMode: "ReductionMode",
	CtaSplitKMode: "CtaSplitKMode",
	SplitKSlc:     "SplitKSlc",
	IdxMode:       "IdxMode",
	Sliced:        "Sliced",
	SplitRs:       "SplitRs",
	Singlebuffer:  "Singlebuffer",
	Ldgc:          "Ldgc",
	Specfilt:      "Specfilt",
	Counts:        "Counts",
}

func (e BackendKnobType) String() string { return _BackendKnobTypeNames[e] }

var _BackendLayoutTypeNames = map[BackendLayoutType]string{
	BELayoutTypePreferredNchw:   "BELayoutTypePreferredNchw",
	BELayoutTypePreferredNhwc:   "BELayoutTypePreferredNhwc",
	BELayoutTypePreferredPad4ck: "BELayoutTypePreferredPad4ck",
	BELayoutTypePreferredPad8ck: "BELayoutTypePreferredPad8ck",
	BELayoutTypeCount:           "BELayoutTypeCount",
}

func (e BackendLayoutType) String() string { return _BackendLayoutTypeNames[e] }

var _BackendNumericalNoteNames = map[BackendNumericalNote]string{
	TensorCore:                "TensorCore",
	DownConvertInputs:         "DownConvertInputs",
	ReducedPrecisionReduction: "ReducedPrecisionReduction",
	Fft:                       "Fft",
	Nondeterministic:          "Nondeterministic",
	Winograd:                  "Winograd",
	TypeCount:                 "TypeCount",
}

func (e BackendNumericalNote) String() string { return _BackendNumericalNoteNames[e] }

var _BatchNormModeNames = map[BatchNormMode]string{
	PerActivation:     "PerActivation",
	Spatial:           "Spatial",
	SpatialPersistent: "SpatialPersistent",
}

func (e BatchNormMode) String() string { return _BatchNormModeNames[e] }

var _BatchNormOpsNames = map[BatchNormOps]string{
	BatchNorm:                 "BatchNorm",
	BatchNormOpsActivation:    "BatchNormOpsActivation",
	BatchNormOpsAddActivation: "BatchNormOpsAddActivation",
}

func (e BatchNormOps) String() string { return _BatchNormOpsNames[e] }

var _CTCLossAlgoNames = map[CTCLossAlgo]string{
	DeterministicCTCLoss:    "DeterministicCTCLoss",
	NonDeterministicCTCLoss: "NonDeterministicCTCLoss",
}

func (e CTCLossAlgo) String() string { return _CTCLossAlgoNames[e] }

var _DataTypeNames = map[DataType]string{
	Float:   "Float",
	Double:  "Double",
	Half:    "Half",
	Int8:    "Int8",
	Int32:   "Int32",
	Int8x4:  "Int8x4",
	Uint8:   "Uint8",
	Uint8x4: "Uint8x4",
	Int8x32: "Int8x32",
}

func (e DataType) String() string { return _DataTypeNames[e] }

var _DeterminismNames = map[Determinism]string{
	NonDeterministic: "NonDeterministic",
	Deterministic:    "Deterministic",
}

func (e Determinism) String() string { return _DeterminismNames[e] }

var _DirectionModeNames = map[DirectionMode]string{
	Unidirectional: "Unidirectional",
	Bidirectional:  "Bidirectional",
}

func (e DirectionMode) String() string { return _DirectionModeNames[e] }

var _DivNormModeNames = map[DivNormMode]string{
	PrecomputedMeans: "PrecomputedMeans",
}

func (e DivNormMode) String() string { return _DivNormModeNames[e] }

var _ErrQueryModeNames = map[ErrQueryMode]string{
	Rawcode:     "Rawcode",
	Nonblocking: "Nonblocking",
	Blocking:    "Blocking",
}

func (e ErrQueryMode) String() string { return _ErrQueryModeNames[e] }

var _FoldingDirectionNames = map[FoldingDirection]string{
	Fold:   "Fold",
	Unfold: "Unfold",
}

func (e FoldingDirection) String() string { return _FoldingDirectionNames[e] }

var _ForwardModeNames = map[ForwardMode]string{
	Inference: "Inference",
	Training:  "Training",
}

func (e ForwardMode) String() string { return _ForwardModeNames[e] }

var _FusedOpsConstParamLabelNames = map[FusedOpsConstParamLabel]string{
	Xdesc:                        "Xdesc",
	XdataPlaceholder:             "XdataPlaceholder",
	BnMode:                       "BnMode",
	BnEqscalebiasDesc:            "BnEqscalebiasDesc",
	BnEqscalePlaceholder:         "BnEqscalePlaceholder",
	BnEqbiasPlaceholder:          "BnEqbiasPlaceholder",
	ActivationDesc:               "ActivationDesc",
	ConvDesc:                     "ConvDesc",
	Wdesc:                        "Wdesc",
	WdataPlaceholder:             "WdataPlaceholder",
	Dwdesc:                       "Dwdesc",
	DwdataPlaceholder:            "DwdataPlaceholder",
	Ydesc:                        "Ydesc",
	YdataPlaceholder:             "YdataPlaceholder",
	Dydesc:                       "Dydesc",
	DydataPlaceholder:            "DydataPlaceholder",
	YstatsDesc:                   "YstatsDesc",
	YsumPlaceholder:              "YsumPlaceholder",
	YsqsumPlaceholder:            "YsqsumPlaceholder",
	BnScalebiasMeanvarDesc:       "BnScalebiasMeanvarDesc",
	BnScalePlaceholder:           "BnScalePlaceholder",
	BnBiasPlaceholder:            "BnBiasPlaceholder",
	BnSavedMeanPlaceholder:       "BnSavedMeanPlaceholder",
	BnSavedInvstdPlaceholder:     "BnSavedInvstdPlaceholder",
	BnRunningMeanPlaceholder:     "BnRunningMeanPlaceholder",
	BnRunningVarPlaceholder:      "BnRunningVarPlaceholder",
	Zdesc:                        "Zdesc",
	ZdataPlaceholder:             "ZdataPlaceholder",
	BnZEqscalebiasDesc:           "BnZEqscalebiasDesc",
	BnZEqscalePlaceholder:        "BnZEqscalePlaceholder",
	BnZEqbiasPlaceholder:         "BnZEqbiasPlaceholder",
	ActivationBitmaskDesc:        "ActivationBitmaskDesc",
	ActivationBitmaskPlaceholder: "ActivationBitmaskPlaceholder",
	Dxdesc:                       "Dxdesc",
	DxdataPlaceholder:            "DxdataPlaceholder",
	Dzdesc:                       "Dzdesc",
	DzdataPlaceholder:            "DzdataPlaceholder",
	BnDscalePlaceholder:          "BnDscalePlaceholder",
	BnDbiasPlaceholder:           "BnDbiasPlaceholder",
}

func (e FusedOpsConstParamLabel) String() string { return _FusedOpsConstParamLabelNames[e] }

var _FusedOpsPointerPlaceHolderNames = map[FusedOpsPointerPlaceHolder]string{
	NullPtr:        "NullPtr",
	PtrElemAligned: "PtrElemAligned",
	Ptr16:          "Ptr16",
}

func (e FusedOpsPointerPlaceHolder) String() string { return _FusedOpsPointerPlaceHolderNames[e] }

var _FusedOpsVariantParamLabelNames = map[FusedOpsVariantParamLabel]string{
	PtrXdata:                        "PtrXdata",
	PtrBnEqscale:                    "PtrBnEqscale",
	PtrBnEqbias:                     "PtrBnEqbias",
	PtrWdata:                        "PtrWdata",
	PtrDwdata:                       "PtrDwdata",
	PtrYdata:                        "PtrYdata",
	PtrDydata:                       "PtrDydata",
	PtrYsum:                         "PtrYsum",
	PtrYsqsum:                       "PtrYsqsum",
	PtrWorkspace:                    "PtrWorkspace",
	PtrBnScale:                      "PtrBnScale",
	PtrBnBias:                       "PtrBnBias",
	PtrBnSavedMean:                  "PtrBnSavedMean",
	PtrBnSavedInvstd:                "PtrBnSavedInvstd",
	PtrBnRunningMean:                "PtrBnRunningMean",
	PtrBnRunningVar:                 "PtrBnRunningVar",
	PtrZdata:                        "PtrZdata",
	PtrBnZEqscale:                   "PtrBnZEqscale",
	PtrBnZEqbias:                    "PtrBnZEqbias",
	PtrActivationBitmask:            "PtrActivationBitmask",
	PtrDxdata:                       "PtrDxdata",
	PtrDzdata:                       "PtrDzdata",
	PtrBnDscale:                     "PtrBnDscale",
	PtrBnDbias:                      "PtrBnDbias",
	ScalarSizeTWorkspaceSizeInBytes: "ScalarSizeTWorkspaceSizeInBytes",
	ScalarInt64TBnAccumulationCount: "ScalarInt64TBnAccumulationCount",
	ScalarDoubleBnExpAvgFactor:      "ScalarDoubleBnExpAvgFactor",
	ScalarDoubleBnEpsilon:           "ScalarDoubleBnEpsilon",
}

func (e FusedOpsVariantParamLabel) String() string { return _FusedOpsVariantParamLabelNames[e] }

var _FusedOpsNames = map[FusedOps]string{
	ScaleBiasActivationConvBnstats:   "ScaleBiasActivationConvBnstats",
	ScaleBiasActivationWgrad:         "ScaleBiasActivationWgrad",
	BnFinalizeStatisticsTraining:     "BnFinalizeStatisticsTraining",
	BnFinalizeStatisticsInference:    "BnFinalizeStatisticsInference",
	ConvScaleBiasAddActivation:       "ConvScaleBiasAddActivation",
	ScaleBiasAddActivationGenBitmask: "ScaleBiasAddActivationGenBitmask",
	DactivationForkDbatchnorm:        "DactivationForkDbatchnorm",
}

func (e FusedOps) String() string { return _FusedOpsNames[e] }

var _GenStatsModeNames = map[GenStatsMode]string{
	SumSq: "SumSq",
}

func (e GenStatsMode) String() string { return _GenStatsModeNames[e] }

var _IndicesTypeNames = map[IndicesType]string{
	Indices32: "Indices32",
	Indices64: "Indices64",
	Indices16: "Indices16",
	Indices8:  "Indices8",
}

func (e IndicesType) String() string { return _IndicesTypeNames[e] }

var _LRNModeNames = map[LRNMode]string{
	CrossChannelDim1: "CrossChannelDim1",
}

func (e LRNMode) String() string { return _LRNModeNames[e] }

var _LossNormalizationModeNames = map[LossNormalizationMode]string{
	LossNormModeNone:    "LossNormModeNone",
	LossNormModeSoftmax: "LossNormModeSoftmax",
}

func (e LossNormalizationMode) String() string { return _LossNormalizationModeNames[e] }

var _MathTypeNames = map[MathType]string{
	DefaultMath:                 "DefaultMath",
	TensorOpMath:                "TensorOpMath",
	TensorOpMathAllowConversion: "TensorOpMathAllowConversion",
	FmaMath:                     "FmaMath",
}

func (e MathType) String() string { return _MathTypeNames[e] }

var _MultiHeadAttnWeightKindNames = map[MultiHeadAttnWeightKind]string{
	QWeights: "QWeights",
	KWeights: "KWeights",
	VWeights: "VWeights",
	OWeights: "OWeights",
	QBiases:  "QBiases",
	KBiases:  "KBiases",
	VBiases:  "VBiases",
	OBiases:  "OBiases",
}

func (e MultiHeadAttnWeightKind) String() string { return _MultiHeadAttnWeightKindNames[e] }

var _NanPropagationNames = map[NanPropagation]string{
	NotPropagateNan: "NotPropagateNan",
	PropagateNan:    "PropagateNan",
}

func (e NanPropagation) String() string { return _NanPropagationNames[e] }

var _NormAlgoNames = map[NormAlgo]string{
	NormAlgoStandard: "NormAlgoStandard",
	NormAlgoPersist:  "NormAlgoPersist",
}

func (e NormAlgo) String() string { return _NormAlgoNames[e] }

var _NormModeNames = map[NormMode]string{
	NormModeActivation: "NormModeActivation",
	NormModeChannel:    "NormModeChannel",
}

func (e NormMode) String() string { return _NormModeNames[e] }

var _NormOpsNames = map[NormOps]string{
	Norm:                 "Norm",
	NormOpsActivation:    "NormOpsActivation",
	NormOpsAddActivation: "NormOpsAddActivation",
}

func (e NormOps) String() string { return _NormOpsNames[e] }

var _OpTensorOpNames = map[OpTensorOp]string{
	OpTensorOpAdd:  "OpTensorOpAdd",
	OpTensorOpMul:  "OpTensorOpMul",
	OpTensorOpMin:  "OpTensorOpMin",
	OpTensorOpMax:  "OpTensorOpMax",
	OpTensorOpSqrt: "OpTensorOpSqrt",
	OpTensorOpNot:  "OpTensorOpNot",
}

func (e OpTensorOp) String() string { return _OpTensorOpNames[e] }

var _PointwiseModeNames = map[PointwiseMode]string{
	PointWiseModeAdd:        "PointWiseModeAdd",
	PointWiseModeMul:        "PointWiseModeMul",
	PointWiseModeMin:        "PointWiseModeMin",
	PointWiseModeMax:        "PointWiseModeMax",
	PointWiseModeSqrt:       "PointWiseModeSqrt",
	PointWiseModeReluFwd:    "PointWiseModeReluFwd",
	PointWiseModeTanhFwd:    "PointWiseModeTanhFwd",
	PointWiseModeSigmoidFwd: "PointWiseModeSigmoidFwd",
	PointWiseModeEluFwd:     "PointWiseModeEluFwd",
}

func (e PointwiseMode) String() string { return _PointwiseModeNames[e] }

var _PoolingModeNames = map[PoolingMode]string{
	MaxPooling:                 "MaxPooling",
	AverageCountIncludePadding: "AverageCountIncludePadding",
	AverageCountExcludePadding: "AverageCountExcludePadding",
	MaxDeterministic:           "MaxDeterministic",
}

func (e PoolingMode) String() string { return _PoolingModeNames[e] }

var _RNNAlgoNames = map[RNNAlgo]string{
	RNNAlgoStandard:       "RNNAlgoStandard",
	RNNAlgoPersistStatic:  "RNNAlgoPersistStatic",
	RNNAlgoPersistDynamic: "RNNAlgoPersistDynamic",
	RNNAlgoCount:          "RNNAlgoCount",
}

func (e RNNAlgo) String() string { return _RNNAlgoNames[e] }

var _RNNBiasModeNames = map[RNNBiasMode]string{
	RNNBiasModeNoBias:        "RNNBiasModeNoBias",
	RNNBiasModeSingleInpBias: "RNNBiasModeSingleInpBias",
	RNNBiasModeDoubleBias:    "RNNBiasModeDoubleBias",
	RNNBiasModeSingleRecBias: "RNNBiasModeSingleRecBias",
}

func (e RNNBiasMode) String() string { return _RNNBiasModeNames[e] }

var _RNNClipModeNames = map[RNNClipMode]string{
	RNNClipModeNone:   "RNNClipModeNone",
	RNNClipModeMinmax: "RNNClipModeMinmax",
}

func (e RNNClipMode) String() string { return _RNNClipModeNames[e] }

var _RNNDataLayoutNames = map[RNNDataLayout]string{
	SeqMajorUnpacked:   "SeqMajorUnpacked",
	SeqMajorPacked:     "SeqMajorPacked",
	BatchMajorUnpacked: "BatchMajorUnpacked",
}

func (e RNNDataLayout) String() string { return _RNNDataLayoutNames[e] }

var _RNNInputModeNames = map[RNNInputMode]string{
	LinearInput: "LinearInput",
	SkipInput:   "SkipInput",
}

func (e RNNInputMode) String() string { return _RNNInputModeNames[e] }

var _RNNModeNames = map[RNNMode]string{
	RNNReLU: "RNNReLU",
	RNNTanh: "RNNTanh",
	LSTM:    "LSTM",
	GRU:     "GRU",
}

func (e RNNMode) String() string { return _RNNModeNames[e] }

var _ReduceTensorIndicesNames = map[ReduceTensorIndices]string{
	ReduceNoIndices:        "ReduceNoIndices",
	ReduceFlattenedIndices: "ReduceFlattenedIndices",
}

func (e ReduceTensorIndices) String() string { return _ReduceTensorIndicesNames[e] }

var _ReduceTensorOpNames = map[ReduceTensorOp]string{
	ReduceAdd:        "ReduceAdd",
	ReduceMul:        "ReduceMul",
	ReduceMin:        "ReduceMin",
	ReduceMax:        "ReduceMax",
	ReduceAmax:       "ReduceAmax",
	ReduceAvg:        "ReduceAvg",
	ReduceNorm1:      "ReduceNorm1",
	ReduceNorm2:      "ReduceNorm2",
	ReduceMulNoZeros: "ReduceMulNoZeros",
}

func (e ReduceTensorOp) String() string { return _ReduceTensorOpNames[e] }

var _ReorderTypeNames = map[ReorderType]string{
	DefaultReorder: "DefaultReorder",
	NoReorder:      "NoReorder",
}

func (e ReorderType) String() string { return _ReorderTypeNames[e] }

var _SamplerTypeNames = map[SamplerType]string{
	Bilinear: "Bilinear",
}

func (e SamplerType) String() string { return _SamplerTypeNames[e] }

var _SeqDataAxisNames = map[SeqDataAxis]string{
	TimeDim:  "TimeDim",
	BatchDim: "BatchDim",
	BeamDim:  "BeamDim",
	VectDim:  "VectDim",
}

func (e SeqDataAxis) String() string { return _SeqDataAxisNames[e] }

var _SeverityNames = map[Severity]string{
	Fatal:   "Fatal",
	Error:   "Error",
	Warning: "Warning",
	Info:    "Info",
}

func (e Severity) String() string { return _SeverityNames[e] }

var _SoftmaxAlgorithmNames = map[SoftmaxAlgorithm]string{
	Fast:     "Fast",
	Accurate: "Accurate",
	Log:      "Log",
}

func (e SoftmaxAlgorithm) String() string { return _SoftmaxAlgorithmNames[e] }

var _SoftmaxModeNames = map[SoftmaxMode]string{
	Instance: "Instance",
	Channel:  "Channel",
}

func (e SoftmaxMode) String() string { return _SoftmaxModeNames[e] }

var _TensorFormatNames = map[TensorFormat]string{
	NCHW:      "NCHW",
	NHWC:      "NHWC",
	NCHWVectC: "NCHWVectC",
}

func (e TensorFormat) String() string { return _TensorFormatNames[e] }

var _WgradModeNames = map[WgradMode]string{
	Add: "Add",
	Set: "Set",
}

func (e WgradMode) String() string { return _WgradModeNames[e] }
