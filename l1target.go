package CloudForest

import (
	"math"
)

/*
L1Target wraps a numerical feature as a target for us in l1 norm regression.
*/
type L1Target struct {
	*Feature
}

/*
L1Target.SplitImpurity is an L1 version of SplitImpurity.
*/
func (target *L1Target) SplitImpurity(l []int, r []int, counter *[]int) (impurityDecrease float64) {
	nl := float64(len(l))
	nr := float64(len(r))

	impurityDecrease = nl * target.Impurity(&l, counter)
	impurityDecrease += nr * target.Impurity(&r, counter)

	impurityDecrease /= nl + nr
	return
}

//L1Target.Impurity is an L1 version of impurity returning L1 instead of squared error.
func (target *L1Target) Impurity(cases *[]int, counter *[]int) (e float64) {
	m := target.Mean(cases)
	e = target.MeanL1Error(cases, m)
	return

}

//L1Target.MeanL1Error returns the  Mean L1 norm error of the cases specified vs the predicted
//value. Only non missing cases are considered.
func (target *L1Target) MeanL1Error(cases *[]int, predicted float64) (e float64) {
	e = 0.0
	n := 0
	for _, i := range *cases {
		if !target.Missing[i] {
			e += math.Abs(predicted - target.NumData[i])
			n += 1
		}

	}
	e = e / float64(n)
	return

}
