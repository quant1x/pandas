package stat

import (
	"gitee.com/quant1x/gox/num"
	"math"
)

// PolyFit
//
//	Least squares polynomial fit.
//
//	.. note::
//		This forms part of the old polynomial API. Since version 1.4, the
//		new polynomial API defined in `numpy.polynomial` is preferred.
//		A summary of the differences can be found in the
//		:doc:`transition guide </reference/routines.polynomials>`.
//
//	Fit a polynomial ``p(x) = p[0] * x**deg + ... + p[deg]`` of degree `deg`
//	to points `(x, y)`. Returns a vector of coefficients `p` that minimises
//	the squared error in the order `deg`, `deg-1`, ... `0`.
//
//	The `Polynomial.fit <numpy.polynomial.polynomial.Polynomial.fit>` class
//	method is recommended for new code as it is more stable numerically. See
//	the documentation of the method for more information.
//
//	Parameters
//	----------
//	x : array_like, shape (M,)
//		x-coordinates of the M sample points ``(x[i], y[i])``.
//	y : array_like, shape (M,) or (M, K)
//		y-coordinates of the sample points. Several data sets of sample
//		points sharing the same x-coordinates can be fitted at once by
//		passing in a 2D-array that contains one dataset per column.
//	deg : int
//		Degree of the fitting polynomial
//
//	Returns
//	-------
//	p : ndarray, shape (deg + 1,) or (deg + 1, K)
//		Polynomial coefficients, highest power first.  If `y` was 2-D, the
//		coefficients for `k`-th data set are in ``p[:,k]``.
//
//	residuals, rank, singular_values, rcond
//		These values are only returned if ``full == True``
//
//		- residuals -- sum of squared residuals of the least squares fit
//		- rank -- the effective rank of the scaled Vandermonde
//			coefficient matrix
//		- singular_values -- singular values of the scaled Vandermonde
//			coefficient matrix
//		- rcond -- value of `rcond`.
//
//		For more details, see `numpy.linalg.lstsq`.
//
//	Warns
//	-----
//	RankWarning
//		The rank of the coefficient matrix in the least-squares fit is
//		deficient. The warning is only raised if ``full == False``.
//
//		The warnings can be turned off by
//
//		>>> import warnings
//		>>> warnings.simplefilter('ignore', np.RankWarning)
//
//	See Also
//	--------
//	polyval : Compute polynomial values.
//	linalg.lstsq : Computes a least-squares fit.
//	scipy.interpolate.UnivariateSpline : Computes spline fits.
//
//	Notes
//	-----
//	The solution minimizes the squared error
//
//	.. math::
//		E = \\sum_{j=0}^k |p(x_j) - y_j|^2
//
//	in the equations::
//
//		x[0]**n * p[0] + ... + x[0] * p[n-1] + p[n] = y[0]
//		x[1]**n * p[0] + ... + x[1] * p[n-1] + p[n] = y[1]
//		...
//		x[k]**n * p[0] + ... + x[k] * p[n-1] + p[n] = y[k]
//
//	The coefficient matrix of the coefficients `p` is a Vandermonde matrix.
//
//	`polyfit` issues a `RankWarning` when the least-squares fit is badly
//	conditioned. This implies that the best fit is not well-defined due
//	to numerical error. The results may be improved by lowering the polynomial
//	degree or by replacing `x` by `x` - `x`.mean(). The `rcond` parameter
//	can also be set to a value smaller than its default, but the resulting
//	fit may be spurious: including contributions from the small singular
//	values can add numerical noise to the result.
//
//	Note that fitting polynomial coefficients is inherently badly conditioned
//	when the degree of the polynomial is large or the interval of sample points
//	is badly centered. The quality of the fit should always be checked in these
//	cases. When polynomial fits are not satisfactory, splines may be a good
//	alternative.
//
//	References
//	----------
//	.. [1] Wikipedia, "Curve fitting",
//			https://en.wikipedia.org/wiki/Curve_fitting
//	.. [2] Wikipedia, "Polynomial interpolation",
//			https://en.wikipedia.org/wiki/Polynomial_interpolation
//	.. [3] numpy.polyfit(x, y, deg, rcond=None, full=False, w=None, cov=False)
//			https://numpy.org/doc/stable/reference/generated/numpy.polyfit.html
//
//	Examples
//	--------
//	>>> import warnings
//	>>> x = np.array([0.0, 1.0, 2.0, 3.0,  4.0,  5.0])
//	>>> y = np.array([0.0, 0.8, 0.9, 0.1, -0.8, -1.0])
//	>>> z = np.polyfit(x, y, 3)
//	>>> z
//	array([ 0.08703704, -0.81349206,  1.69312169, -0.03968254]) # may vary
func PolyFit(x, y []DType, deg int, args ...any) []DType {
	// 默认从右向左
	var __increasing = false
	if len(args) > 0 {
		// 第一个参数为是否copy
		if arg0, ok := args[0].(bool); ok {
			__increasing = arg0
		}
	}
	// Initialize matrix to store powers of x
	var X = make([][]float64, len(x))
	for i := range X {
		X[i] = make([]float64, deg+1)
	}

	// Fill matrix with powers of x
	for i := 0; i < len(x); i++ {
		for j := 0; j <= deg; j++ {
			k := j
			if !__increasing {
				k = deg - k
			}
			X[i][j] = math.Pow(x[i], float64(k))
		}
	}

	// Calculate transpose of X
	var XT = __transpose(X)

	// Multiply XT with X
	var XTX = __matmul(XT, X)

	// Invert XTX
	var XTXinv = __inv(XTX)

	// Multiply XTXinv with XT
	var XTXinvXT = __matmul(XTXinv, XT)

	// Multiply result with y
	var coef = __matvec(XTXinvXT, y)

	return coef
}

// PolyVal
//
//	Evaluate a polynomial at specific values.
//
//	.. note::
//	This forms part of the old polynomial API. Since version 1.4, the
//	new polynomial API defined in `numpy.polynomial` is preferred.
//	A summary of the differences can be found in the
//	:doc:`transition guide </reference/routines.polynomials>`.
//
//	If `p` is of length N, this function returns the value:
//
//	``p[0]*x**(N-1) + p[1]*x**(N-2) + ... + p[N-2]*x + p[N-1]``
//
//	If `x` is a sequence, then ``p(x)`` is returned for each element of ``x``.
//	If `x` is another polynomial then the composite polynomial ``p(x(t))``
//	is returned.
//
//	Parameters
//
// ----------
//
//	p : array_like or poly1d object
//		1D array of polynomial coefficients (including coefficients equal
//		to zero) from highest degree to the constant term, or an
//		instance of poly1d.
//	x : array_like or poly1d object
//		A number, an array of numbers, or an instance of poly1d, at
//		which to evaluate `p`.
//
//	Returns
//
// -------
//
//	values : ndarray or poly1d
//		If `x` is a poly1d instance, the result is the composition of the two
//		polynomials, i.e., `x` is "substituted" in `p` and the simplified
//		result is returned. In addition, the type of `x` - array_like or
//		poly1d - governs the type of the output: `x` array_like => `values`
//		array_like, `x` a poly1d object => `values` is also.
//
//	See Also
//	--------
//	poly1d: A polynomial class.
//
//	Notes
//	-----
//	Horner's scheme [1]_ is used to evaluate the polynomial. Even so,
//	for polynomials of high degree the values may be inaccurate due to
//	rounding errors. Use carefully.
//
//	If `x` is a subtype of `ndarray` the return value will be of the same type.
//
//	References
//	----------
//	.. [1] I. N. Bronshtein, K. A. Semendyayev, and K. A. Hirsch (Eng.
//	trans. Ed.), *Handbook of Mathematics*, New York, Van Nostrand
//	Reinhold Co., 1985, pg. 720.
//
//	Examples
//	--------
//	>>> np.polyval([3,0,1], 5)  # 3 * 5**2 + 0 * 5**1 + 1
//	76
//	>>> np.polyval([3,0,1], np.poly1d(5))
//	poly1d([76])
//	>>> np.polyval(np.poly1d([3,0,1]), 5)
//	76
//	>>> np.polyval(np.poly1d([3,0,1]), np.poly1d(5))
//	poly1d([76])
func PolyVal(p, x []DType) []DType {
	//p = NX.asarray(p)
	//if isinstance(x, poly1d):
	//	y = 0
	//else:
	//	x = NX.asanyarray(x)
	//	y = NX.zeros_like(x)
	//for pv in p:
	//y = y * x + pv
	y := Repeat(DType(0), len(x))
	for _, v := range p {
		num.Mul_Inplace(y, x)
		num.AddNumber_Inplace(y, v)
	}
	return y
}

// Function to calculate matrix transpose
func __transpose(a [][]float64) [][]float64 {
	var transposed = make([][]float64, len(a[0]))
	for i := range transposed {
		transposed[i] = make([]float64, len(a))
	}
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			transposed[j][i] = a[i][j]
		}
	}
	return transposed
}

// Function to multiply two matrices
func __matmul(a, b [][]float64) [][]float64 {
	var result = make([][]float64, len(a))
	for i := range result {
		result[i] = make([]float64, len(b[0]))
	}
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b[0]); j++ {
			for k := 0; k < len(b); k++ {
				result[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return result
}

// Function to multiply a matrix with a vector
func __matvec(a [][]float64, b []float64) []float64 {
	var result = make([]float64, len(a))
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			result[i] += a[i][j] * b[j]
		}
	}
	return result
}

func __inv(a [][]float64) [][]float64 {
	var n = len(a)

	// Create augmented matrix
	var augmented = make([][]float64, n)
	for i := range augmented {
		augmented[i] = make([]float64, 2*n)
		for j := 0; j < n; j++ {
			augmented[i][j] = a[i][j]
		}
	}
	for i := 0; i < n; i++ {
		augmented[i][i+n] = 1
	}

	// Perform Gaussian elimination
	for i := 0; i < n; i++ {
		var pivot = augmented[i][i]
		for j := i + 1; j < n; j++ {
			var factor = augmented[j][i] / pivot
			for k := i; k < 2*n; k++ {
				augmented[j][k] -= factor * augmented[i][k]
			}
		}
	}

	// Perform back-substitution
	for i := n - 1; i >= 0; i-- {
		var pivot = augmented[i][i]
		for j := i - 1; j >= 0; j-- {
			var factor = augmented[j][i] / pivot
			for k := i; k < 2*n; k++ {
				augmented[j][k] -= factor * augmented[i][k]
			}
		}
	}

	// Normalize rows
	for i := 0; i < n; i++ {
		var pivot = augmented[i][i]
		for j := 0; j < 2*n; j++ {
			augmented[i][j] /= pivot
		}
	}

	// Extract inverse from augmented matrix
	var inverse = make([][]float64, n)
	for i := range inverse {
		inverse[i] = make([]float64, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			inverse[i][j] = augmented[i][j+n]
		}
	}

	return inverse
}
