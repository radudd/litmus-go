package types

import (
	corev1 "k8s.io/api/core/v1"
	clientTypes "k8s.io/apimachinery/pkg/types"
)

// ExperimentDetails is for collecting all the experiment-related details
type ExperimentDetails struct {
	ExperimentName      string
	EngineName          string
	ChaosDuration       int
	RampTime            int
	ChaosLib            string
	AppNS               string
	AppLabel            string
	AppKind             string
	ChaosUID            clientTypes.UID
	InstanceID          string
	ChaosNamespace      string
	ChaosPodName        string
	TargetContainer     string
	AuxiliaryAppInfo    string
	FillPercentage      int
	ContainerPath       string
	RunID               string
	Timeout             int
	Delay               int
	LIBImage            string
	LIBImagePullPolicy  string
	TargetPods          string
	PodsAffectedPerc    int
	Annotations         map[string]string
	Sequence            string
	Resources           corev1.ResourceRequirements
	ChaosServiceAccount string
	ImagePullSecrets    []corev1.LocalObjectReference
}
