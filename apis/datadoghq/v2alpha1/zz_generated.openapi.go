// +build !ignore_autogenerated

// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v2alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"./apis/datadoghq/v2alpha1.DatadogAgent":                 schema__apis_datadoghq_v2alpha1_DatadogAgent(ref),
		"./apis/datadoghq/v2alpha1.DatadogAgentGenericContainer": schema__apis_datadoghq_v2alpha1_DatadogAgentGenericContainer(ref),
		"./apis/datadoghq/v2alpha1.DatadogAgentStatus":           schema__apis_datadoghq_v2alpha1_DatadogAgentStatus(ref),
		"./apis/datadoghq/v2alpha1.DatadogFeatures":              schema__apis_datadoghq_v2alpha1_DatadogFeatures(ref),
		"./apis/datadoghq/v2alpha1.ImageConfig":                  schema__apis_datadoghq_v2alpha1_ImageConfig(ref),
	}
}

func schema__apis_datadoghq_v2alpha1_DatadogAgent(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DatadogAgent Deployment with Datadog Operator.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("./apis/datadoghq/v2alpha1.DatadogAgentSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("./apis/datadoghq/v2alpha1.DatadogAgentStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./apis/datadoghq/v2alpha1.DatadogAgentSpec", "./apis/datadoghq/v2alpha1.DatadogAgentStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema__apis_datadoghq_v2alpha1_DatadogAgentGenericContainer(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DatadogAgentGenericContainer is the generic structure describing any container's common configuration.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Description: "Name of the container that is overiden",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"env": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-map-keys": []interface{}{
									"name",
								},
								"x-kubernetes-list-type": "map",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "The different components (Datadog Agent, Cluster Agent, Cluster Check Runner) support many environment variables. See also: https://docs.datadoghq.com/agent/kubernetes/?tab=helm#environment-variables",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("k8s.io/api/core/v1.EnvVar"),
									},
								},
							},
						},
					},
					"volumeMounts": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-map-keys": []interface{}{
									"name",
									"mountPath",
								},
								"x-kubernetes-list-type": "map",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Specify additional volume mounts in the container.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("k8s.io/api/core/v1.VolumeMount"),
									},
								},
							},
						},
					},
					"resources": {
						SchemaProps: spec.SchemaProps{
							Description: "Make sure to keep requests and limits equal to keep the pods in the Guaranteed QoS class. See also: http://kubernetes.io/docs/user-guide/compute-resources/",
							Ref:         ref("k8s.io/api/core/v1.ResourceRequirements"),
						},
					},
					"command": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "atomic",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Command allows the specification of custom entrypoint for container",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"args": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "atomic",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Args allows the specification of extra args to `Command` parameter",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"healthPort": {
						SchemaProps: spec.SchemaProps{
							Description: "HealthPort of the agent container for internal liveness probe. Must be the same as the Liness/Readiness probes.",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"readinessProbe": {
						SchemaProps: spec.SchemaProps{
							Description: "Configure the Readiness Probe of the container",
							Ref:         ref("k8s.io/api/core/v1.Probe"),
						},
					},
					"livenessProbe": {
						SchemaProps: spec.SchemaProps{
							Description: "Configure the Liveness Probe of the container",
							Ref:         ref("k8s.io/api/core/v1.Probe"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/api/core/v1.EnvVar", "k8s.io/api/core/v1.Probe", "k8s.io/api/core/v1.ResourceRequirements", "k8s.io/api/core/v1.VolumeMount"},
	}
}

func schema__apis_datadoghq_v2alpha1_DatadogAgentStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DatadogAgentStatus defines the observed state of DatadogAgent",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"defaultOverride": {
						SchemaProps: spec.SchemaProps{
							Description: "DefaultOverride contains attributes that were not configured that the runtime defaulted.",
							Ref:         ref("./apis/datadoghq/v2alpha1.DatadogAgentSpec"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./apis/datadoghq/v2alpha1.DatadogAgentSpec"},
	}
}

func schema__apis_datadoghq_v2alpha1_DatadogFeatures(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DatadogFeatures are Features running on the Agent and Cluster Agent.",
				Type:        []string{"object"},
			},
		},
	}
}

func schema__apis_datadoghq_v2alpha1_ImageConfig(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ImageConfig Datadog agent container image config.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Description: "Define the image to use: Use \"gcr.io/datadoghq/agent:latest\" for Datadog Agent 7 Use \"datadog/dogstatsd:latest\" for Standalone Datadog Agent DogStatsD6 Use \"gcr.io/datadoghq/cluster-agent:latest\" for Datadog Cluster Agent Use \"agent\" with the registry and tag configurations for <registry>/agent:<tag> Use \"cluster-agent\" with the registry and tag configurations for <registry>/cluster-agent:<tag>",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"tag": {
						SchemaProps: spec.SchemaProps{
							Description: "Define the image version to use: To be used if the Name field does not correspond to a full image string.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"jmxEnabled": {
						SchemaProps: spec.SchemaProps{
							Description: "Define whether the Agent image should support JMX.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"pullPolicy": {
						SchemaProps: spec.SchemaProps{
							Description: "The Kubernetes pull policy: Use Always, Never or IfNotPresent.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"pullSecrets": {
						SchemaProps: spec.SchemaProps{
							Description: "It is possible to specify docker registry credentials. See https://kubernetes.io/docs/concepts/containers/images/#specifying-imagepullsecrets-on-a-pod",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("k8s.io/api/core/v1.LocalObjectReference"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/api/core/v1.LocalObjectReference"},
	}
}
