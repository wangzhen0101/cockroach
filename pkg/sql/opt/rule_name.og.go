// Code generated by optgen; DO NOT EDIT.

package opt

const (
	startAutoRule RuleName = iota + NumManualRuleNames

	// ------------------------------------------------------------
	// Normalize Rule Names
	// ------------------------------------------------------------
	EliminateEmptyAnd
	EliminateEmptyOr
	EliminateSingletonAndOr
	SimplifyAnd
	SimplifyOr
	SimplifyFilters
	FoldNullAndOr
	NegateComparison
	EliminateNot
	NegateAnd
	NegateOr
	CommuteVarInequality
	CommuteConstInequality
	NormalizeCmpPlusConst
	NormalizeCmpMinusConst
	NormalizeCmpConstMinus
	NormalizeTupleEquality
	FoldNullComparisonLeft
	FoldNullComparisonRight
	EnsureJoinFiltersAnd
	EnsureJoinFilters
	PushDownJoinLeft
	PushDownJoinRight
	FoldPlusZero
	FoldZeroPlus
	FoldMinusZero
	FoldMultOne
	FoldOneMult
	FoldDivOne
	InvertMinus
	EliminateUnaryMinus
	EliminateProject
	FilterUnusedProjectCols
	FilterUnusedScanCols
	FilterUnusedSelectCols
	FilterUnusedLimitCols
	FilterUnusedOffsetCols
	FilterUnusedJoinLeftCols
	FilterUnusedJoinRightCols
	FilterUnusedAggCols
	FilterUnusedGroupByCols
	FilterUnusedValueCols
	CommuteVar
	CommuteConst
	EliminateCoalesce
	SimplifyCoalesce
	EliminateCast
	FoldNullCast
	FoldNullUnary
	FoldNullBinaryLeft
	FoldNullBinaryRight
	FoldNullInNonEmpty
	FoldNullInEmpty
	FoldNullNotInEmpty
	NormalizeInConst
	FoldInNull
	EnsureSelectFiltersAnd
	EnsureSelectFilters
	EliminateSelect
	MergeSelects
	PushDownSelectJoinLeft
	PushDownSelectJoinRight
	MergeSelectInnerJoin
	PushDownSelectGroupBy

	// ------------------------------------------------------------
	// Explore Rule Names
	// ------------------------------------------------------------
	GenerateIndexScans

	// NumRuleNames tracks the total count of rule names.
	NumRuleNames
)
