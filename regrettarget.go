package CloudForest

import ()

/*
RegretTarget wraps a categorical feature for use in regret driven classification.
The ith entry in costs should contain the cost of misclassifying a case that actually
has the ith category.
*/
type RegretTarget struct {
	*Feature
	Costs []float64
}

//NewRegretTarget creates a RefretTarget and initializes RegretTarget.Costs to the proper length.
func NewRegretTarget(f *Feature) *RegretTarget {
	return &RegretTarget{f, make([]float64, f.NCats())}
}

/*RegretTarget.SetCosts puts costs in a map[string]float64 by feature name into the proper
entries in RegretTarget.Costs.*/
func (target *RegretTarget) SetCosts(costmap map[string]float64) {
	for i, c := range target.Back {
		target.Costs[i] = costmap[c]
	}
}

/*
RegretTarget.SplitImpurity is a version of Split Impurity that calls RegretTarget.Impurity
*/
func (target *RegretTarget) SplitImpurity(l []int, r []int, counter *[]int) (impurityDecrease float64) {
	nl := float64(len(l))
	nr := float64(len(r))

	impurityDecrease = nl * target.Impurity(&l, counter)
	impurityDecrease += nr * target.Impurity(&r, counter)

	impurityDecrease /= nl + nr
	return
}

//RegretTarget.Impurity implements a simple regret function that finds the average cost of
//a set using the misclassification costs in RegretTarget.Costs.
func (target *RegretTarget) Impurity(cases *[]int, counter *[]int) (e float64) {
	m := target.Modei(cases)
	t := 0
	for _, c := range *cases {
		if target.Missing[c] == false {
			t += 1
			cat := target.CatData[c]
			if cat != m {
				e += target.Costs[cat]
			}
		}

	}
	e /= float64(t)

	return
}
