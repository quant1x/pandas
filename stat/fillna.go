package stat

import "golang.org/x/exp/slices"

// Fill 填充
//
//	Fill NA/NaN values using the specified method.
//	Parameters
//	----------
//	value : scalar, dict, Series, or DataFrame
//	   Value to use to fill holes (e.g. 0), alternately a
//	   dict/Series/DataFrame of values specifying which value to use for
//	   each index (for a Series) or column (for a DataFrame).  Values not
//	   in the dict/Series/DataFrame will not be filled. This value cannot
//	   be a list.
//	method : {{'backfill', 'bfill', 'pad', 'ffill', None}}, default None
//	   Method to use for filling holes in reindexed Series
//	   pad / ffill: propagate last valid observation forward to next valid
//	   backfill / bfill: use next valid observation to fill gap.
//	axis : {axes_single_arg}
//	   Axis along which to fill missing values. For `Series`
//	   this parameter is unused and defaults to 0.
//	inplace : bool, default False [√]
//	   If True, fill in-place. Note: this will modify any
//	   other views on this object (e.g., a no-copy slice for a column in a
//	   DataFrame).
//	limit : int, default None
//	   If method is specified, this is the maximum number of consecutive
//	   NaN values to forward/backward fill. In other words, if there is
//	   a gap with more than this number of consecutive NaNs, it will only
//	   be partially filled. If method is not specified, this is the
//	   maximum number of entries along the entire axis where NaNs will be
//	   filled. Must be greater than 0 if not None.
//	downcast : dict, default is None
//	   A dict of item->dtype of what to downcast if possible,
//	   or the string 'infer' which will try to downcast to an appropriate
//	   equal type (e.g. float64 to int64 if possible).
//
//	Returns
//	-------
//	[]T or None
func Fill[T BaseType](v []T, d T, args ...any) (rows []T) {
	// 默认不替换
	var __optInplace = false
	if len(args) > 0 {
		// 第一个参数为是否copy
		if _cp, ok := args[0].(bool); ok {
			__optInplace = _cp
		}
	}
	var dest []T
	if __optInplace {
		dest = v
	} else {
		dest = slices.Clone(v)
	}
	var values any = dest
	switch rows := values.(type) {
	case []string:
		ds := AnyToString(d)
		for idx, iv := range rows {
			if StringIsNaN(iv) {
				rows[idx] = ds
			}
		}
	case []float32:
		df32 := AnyToFloat32(d)
		for idx, iv := range rows {
			if Float32IsNaN(iv) {
				rows[idx] = df32
			}
		}
	case []float64:
		df64 := AnyToFloat64(d)
		for idx, iv := range rows {
			if Float64IsNaN(iv) {
				rows[idx] = df64
			}
		}
	default:
		return dest
	}
	return dest
}

// FillNa NaN填充默认值
func FillNa[T BaseType](x []T, v any, args ...any) []T {
	// 默认不copy
	var __optInplace = false
	if len(args) > 0 {
		// 第一个参数为是否copy
		if _cp, ok := args[0].(bool); ok {
			__optInplace = _cp
		}
	}
	var dest []T
	if __optInplace {
		dest = x
	} else {
		dest = slices.Clone(x)
	}
	var values any = dest
	switch rows := values.(type) {
	case []string:
		for idx, iv := range rows {
			if StringIsNaN(iv) {
				rows[idx] = AnyToString(v)
			}
		}
	case []int64:
		for idx, iv := range rows {
			if Float64IsNaN(float64(iv)) {
				rows[idx] = AnyToInt64(v)
			}
		}
	case []float32:
		for idx, iv := range rows {
			if Float32IsNaN(iv) {
				rows[idx] = AnyToFloat32(v)
			}
		}
	case []float64:
		for idx, iv := range rows {
			if Float64IsNaN(iv) {
				rows[idx] = AnyToFloat64(v)
			}
		}
	default:
		return dest
	}
	return dest
}
