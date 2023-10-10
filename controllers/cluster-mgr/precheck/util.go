package precheck

import (
	rainbondv1alpha1 "github.com/OPengXJ/rainbond-operator/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
)

func failConditoin(condition rainbondv1alpha1.RainbondClusterCondition, reason, msg string) rainbondv1alpha1.RainbondClusterCondition {
	condition.Status = corev1.ConditionFalse
	condition.Reason = reason
	condition.Message = msg
	return condition
}
