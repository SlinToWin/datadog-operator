// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package v1alpha1

const (
	// AgentDeploymentNameLabelKey label key use to link a Resource to a DatadogAgent
	AgentDeploymentNameLabelKey = "agent.datadoghq.com/name"
	// AgentDeploymentComponentLabelKey label key use to know with component is it
	AgentDeploymentComponentLabelKey = "agent.datadoghq.com/component"
	// MD5AgentDeploymentAnnotationKey annotation key used on a Resource in order to identify which AgentDeployment have been used to generate it.
	MD5AgentDeploymentAnnotationKey = "agent.datadoghq.com/agentspechash"

	// DefaultAgentResourceSuffix use as suffix for agent resource naming
	DefaultAgentResourceSuffix = "agent"
	// DefaultClusterAgentResourceSuffix use as suffix for cluster-agent resource naming
	DefaultClusterAgentResourceSuffix = "cluster-agent"
	// DefaultClusterChecksRunnerResourceSuffix use as suffix for cluster-checks-runner resource naming
	DefaultClusterChecksRunnerResourceSuffix = "cluster-checks-runner"
	// DefaultMetricsServerResourceSuffix use as suffix for cluster-agent metrics-server resource naming
	DefaultMetricsServerResourceSuffix = "cluster-agent-metrics-server"
	// DefaultAPPKeyKey default app-key key (use in secret for instance).
	DefaultAPPKeyKey = "app_key"
	// DefaultAPIKeyKey default api-key key (use in secret for instance).
	DefaultAPIKeyKey = "api_key"
	// DefaultTokenKey default token key (use in secret for instance).
	DefaultTokenKey = "token"
	// DefaultClusterAgentServicePort default cluster-agent service port
	DefaultClusterAgentServicePort = 5005
	// DefaultMetricsServerServicePort default metrics-server port
	DefaultMetricsServerServicePort = 443
	// DefaultMetricsServerTargetPort default metrics-server pod port
	DefaultMetricsServerTargetPort = int(defaultMetricsProviderPort)
	// DefaultAdmissionControllerServicePort default admission controller service port
	DefaultAdmissionControllerServicePort = 443
	// DefaultAdmissionControllerTargetPort default admission controller pod port
	DefaultAdmissionControllerTargetPort = 8000
	// DefaultDogstatsdPort default dogstatsd port
	DefaultDogstatsdPort = 8125
)

// Datadog env var names
const (
	DatadogHost                                  = "DATADOG_HOST"
	DDAPIKey                                     = "DD_API_KEY"
	DDClusterName                                = "DD_CLUSTER_NAME"
	DDSite                                       = "DD_SITE"
	DDddURL                                      = "DD_DD_URL"
	DDHealthPort                                 = "DD_HEALTH_PORT"
	DDLogLevel                                   = "DD_LOG_LEVEL"
	DDPodLabelsAsTags                            = "DD_KUBERNETES_POD_LABELS_AS_TAGS"
	DDPodAnnotationsAsTags                       = "DD_KUBERNETES_POD_ANNOTATIONS_AS_TAGS"
	DDTags                                       = "DD_TAGS"
	DDCollectKubeEvents                          = "DD_COLLECT_KUBERNETES_EVENTS"
	DDLeaderElection                             = "DD_LEADER_ELECTION"
	DDLogsEnabled                                = "DD_LOGS_ENABLED"
	DDLogsConfigContainerCollectAll              = "DD_LOGS_CONFIG_CONTAINER_COLLECT_ALL"
	DDLogsContainerCollectUsingFiles             = "DD_LOGS_CONFIG_K8S_CONTAINER_USE_FILE"
	DDLogsConfigOpenFilesLimit                   = "DD_LOGS_CONFIG_OPEN_FILES_LIMIT"
	DDDogstatsdEnabled                           = "DD_USE_DOGSTATSD"
	DDDogstatsdOriginDetection                   = "DD_DOGSTATSD_ORIGIN_DETECTION"
	DDDogstatsdPort                              = "DD_DOGSTATSD_PORT"
	DDDogstatsdSocket                            = "DD_DOGSTATSD_SOCKET"
	DDDogstatsdMapperProfiles                    = "DD_DOGSTATSD_MAPPER_PROFILES"
	DDClusterAgentEnabled                        = "DD_CLUSTER_AGENT_ENABLED"
	DDClusterAgentKubeServiceName                = "DD_CLUSTER_AGENT_KUBERNETES_SERVICE_NAME"
	DDClusterAgentAuthToken                      = "DD_CLUSTER_AGENT_AUTH_TOKEN"
	DDMetricsProviderEnabled                     = "DD_EXTERNAL_METRICS_PROVIDER_ENABLED"
	DDMetricsProviderPort                        = "DD_EXTERNAL_METRICS_PROVIDER_PORT"
	DDMetricsProviderUseDatadogMetric            = "DD_EXTERNAL_METRICS_PROVIDER_USE_DATADOGMETRIC_CRD"
	DDMetricsProviderWPAController               = "DD_EXTERNAL_METRICS_PROVIDER_WPA_CONTROLLER"
	DDAppKey                                     = "DD_APP_KEY"
	DDClusterChecksEnabled                       = "DD_CLUSTER_CHECKS_ENABLED"
	DDIgnoreAutoConf                             = "DD_IGNORE_AUTOCONF"
	DDKubeStateMetricsCoreEnabled                = "DD_KUBE_STATE_METRICS_CORE_ENABLED"
	DDKubeStateMetricsCoreConfigMap              = "DD_KUBE_STATE_METRICS_CORE_CONFIGMAP_NAME"
	DDClcRunnerEnabled                           = "DD_CLC_RUNNER_ENABLED"
	DDClcRunnerHost                              = "DD_CLC_RUNNER_HOST"
	DDClcRunnerID                                = "DD_CLC_RUNNER_ID"
	DDExtraConfigProviders                       = "DD_EXTRA_CONFIG_PROVIDERS"
	DDExtraListeners                             = "DD_EXTRA_LISTENERS"
	DDHostname                                   = "DD_HOSTNAME"
	DDAPMEnabled                                 = "DD_APM_ENABLED"
	DDPPMReceiverSocket                          = "DD_APM_RECEIVER_SOCKET"
	DDProcessAgentEnabled                        = "DD_PROCESS_AGENT_ENABLED"
	DDSystemProbeAgentEnabled                    = "DD_SYSTEM_PROBE_ENABLED"
	DDSystemProbeSocketPath                      = "DD_SYSPROBE_SOCKET"
	DDSystemProbeCollectDNSStatsEnabled          = "DD_COLLECT_DNS_STATS"
	DDSystemProbeNPMEnabled                      = "DD_SYSTEM_PROBE_NETWORK_ENABLED"
	DDSystemProbeEnvPrefix                       = "DD_SYSTEM_PROBE_CONFIG_"
	DDSystemProbeDebugPort                       = DDSystemProbeEnvPrefix + "DEBUG_PORT"
	DDSystemProbeConntrackEnabled                = DDSystemProbeEnvPrefix + "ENABLE_CONNTRACK"
	DDSystemProbeBPFDebugEnabled                 = DDSystemProbeEnvPrefix + "BPF_DEBUG"
	DDSystemProbeTCPQueueLengthEnabled           = DDSystemProbeEnvPrefix + "ENABLE_TCP_QUEUE_LENGTH"
	DDSystemProbeOOMKillEnabled                  = DDSystemProbeEnvPrefix + "ENABLE_OOM_KILL"
	DDEnableMetadataCollection                   = "DD_ENABLE_METADATA_COLLECTION"
	DDKubeletHost                                = "DD_KUBERNETES_KUBELET_HOST"
	DDKubeletTLSVerify                           = "DD_KUBELET_TLS_VERIFY"
	DDKubeletCAPath                              = "DD_KUBELET_CLIENT_CA"
	DDCriSocketPath                              = "DD_CRI_SOCKET_PATH"
	DockerHost                                   = "DOCKER_HOST"
	DDAdmissionControllerEnabled                 = "DD_ADMISSION_CONTROLLER_ENABLED"
	DDAdmissionControllerMutateUnlabelled        = "DD_ADMISSION_CONTROLLER_MUTATE_UNLABELLED"
	DDAdmissionControllerInjectConfig            = "DD_ADMISSION_CONTROLLER_INJECT_CONFIG_ENABLED"
	DDAdmissionControllerInjectTags              = "DD_ADMISSION_CONTROLLER_INJECT_TAGS_ENABLED"
	DDAdmissionControllerServiceName             = "DD_ADMISSION_CONTROLLER_SERVICE_NAME"
	DDComplianceConfigEnabled                    = "DD_COMPLIANCE_CONFIG_ENABLED"
	DDComplianceConfigCheckInterval              = "DD_COMPLIANCE_CONFIG_CHECK_INTERVAL"
	DDComplianceConfigDir                        = "DD_COMPLIANCE_CONFIG_DIR"
	DDRuntimeSecurityConfigEnabled               = "DD_RUNTIME_SECURITY_CONFIG_ENABLED"
	DDRuntimeSecurityConfigPoliciesDir           = "DD_RUNTIME_SECURITY_CONFIG_POLICIES_DIR"
	DDRuntimeSecurityConfigSocket                = "DD_RUNTIME_SECURITY_CONFIG_SOCKET"
	DDRuntimeSecurityConfigSyscallMonitorEnabled = "DD_RUNTIME_SECURITY_CONFIG_SYSCALL_MONITOR_ENABLED"
	DDRuntimeSecurityConfigRemoteTaggerEnabled   = "DD_RUNTIME_SECURITY_CONFIG_REMOTE_TAGGER"
	DDExternalMetricsProviderEndpoint            = "DD_EXTERNAL_METRICS_PROVIDER_ENDPOINT"
	DDPrometheusScrapeEnabled                    = "DD_PROMETHEUS_SCRAPE_ENABLED"
	DDPrometheusScrapeServiceEndpoints           = "DD_PROMETHEUS_SCRAPE_SERVICE_ENDPOINTS"
	DDPrometheusScrapeChecks                     = "DD_PROMETHEUS_SCRAPE_CHECKS"
	DDExternalMetricsProviderAPIKey              = "DD_EXTERNAL_METRICS_PROVIDER_API_KEY"
	DDExternalMetricsProviderAppKey              = "DD_EXTERNAL_METRICS_PROVIDER_APP_KEY"
	DDAuthTokenFilePath                          = "DD_AUTH_TOKEN_FILE_PATH"

	// KubernetesEnvvarName Env var used by the Datadog Agent container entrypoint
	// to add kubelet config provider and listener
	KubernetesEnvvarName = "KUBERNETES"

	// Datadog volume names and mount paths

	LogDatadogVolumeName                 = "logdatadog"
	LogDatadogVolumePath                 = "/var/log/datadog"
	APMSocketVolumeName                  = "apmsocket"
	APMSocketVolumePath                  = "/var/run/datadog/apm"
	InstallInfoVolumeName                = "installinfo"
	InstallInfoVolumeSubPath             = "install_info"
	InstallInfoVolumePath                = "/etc/datadog-agent/install_info"
	InstallInfoVolumeReadOnly            = true
	ConfdVolumeName                      = "confd"
	ConfdVolumePath                      = "/conf.d"
	ChecksdVolumeName                    = "checksd"
	ChecksdVolumePath                    = "/checks.d"
	ConfigVolumeName                     = "config"
	ConfigVolumePath                     = "/etc/datadog-agent"
	AuthVolumeName                       = "datadog-agent-auth"
	AuthVolumePath                       = "/etc/datadog-agent/auth"
	ProcVolumeName                       = "procdir"
	ProcVolumePath                       = "/host/proc"
	ProcVolumeReadOnly                   = true
	PasswdVolumeName                     = "passwd"
	PasswdVolumePath                     = "/etc/passwd"
	GroupVolumeName                      = "group"
	GroupVolumePath                      = "/etc/group"
	HostRootVolumeName                   = "hostroot"
	HostRootVolumePath                   = "/host/root"
	CgroupsVolumeName                    = "cgroups"
	CgroupsVolumePath                    = "/host/sys/fs/cgroup"
	CgroupsVolumeReadOnly                = true
	SystemProbeSocketVolumeName          = "sysprobe-socket-dir"
	SystemProbeSocketVolumePath          = "/var/run/sysprobe"
	CriSocketVolumeName                  = "runtimesocketdir"
	CriSocketVolumeReadOnly              = true
	DogstatsdSocketVolumeName            = "dsdsocket"
	DogstatsdSocketVolumePath            = "/var/run/datadog/statsd"
	KubeStateMetricCoreVolumeName        = "ksm-core-config"
	PointerVolumeName                    = "pointerdir"
	PointerVolumePath                    = "/opt/datadog-agent/run"
	LogPodVolumeName                     = "logpodpath"
	LogPodVolumePath                     = "/var/log/pods"
	LogPodVolumeReadOnly                 = true
	LogContainerVolumeName               = "logcontainerpath"
	LogContainerVolumeReadOnly           = true
	SymlinkContainerVolumeName           = "symlinkcontainerpath"
	SymlinkContainerVolumeReadOnly       = true
	SystemProbeDebugfsVolumeName         = "debugfs"
	SystemProbeDebugfsVolumePath         = "/sys/kernel/debug"
	SystemProbeConfigVolumeName          = "system-probe-config"
	SystemProbeConfigVolumePath          = "/etc/datadog-agent/system-probe.yaml"
	SystemProbeConfigVolumeSubPath       = "system-probe.yaml"
	SystemProbeAgentSecurityVolumeName   = "datadog-agent-security"
	SystemProbeAgentSecurityVolumePath   = "/etc/config"
	SystemProbeSecCompRootVolumeName     = "seccomp-root"
	SystemProbeSecCompRootVolumePath     = "/host/var/lib/kubelet/seccomp"
	SystemProbeLibModulesVolumeName      = "modules"
	SystemProbeLibModulesVolumePath      = "/lib/modules"
	SystemProbeUsrSrcVolumeName          = "src"
	SystemProbeUsrSrcVolumePath          = "/usr/src"
	AgentCustomConfigVolumeName          = "custom-datadog-yaml"
	AgentCustomConfigVolumePath          = "/etc/datadog-agent/datadog.yaml"
	AgentCustomConfigVolumeSubPath       = "datadog.yaml"
	KubeletCAVolumeName                  = "kubelet-ca"
	DefaultKubeletAgentCAPath            = "/var/run/host-kubelet-ca.crt"
	OrchestratorExplorerConfigVolumeName = "orchestrator-explorer-config"

	HostCriSocketPathPrefix = "/host"

	SecurityAgentRuntimeCustomPoliciesVolumeName = "customruntimepolicies"
	SecurityAgentRuntimePoliciesDirVolumeName    = "runtimepoliciesdir"
	SecurityAgentRuntimePoliciesDirVolumePath    = "/etc/datadog-agent/runtime-security.d"
	SecurityAgentComplianceConfigDirVolumeName   = "compliancedir"
	SecurityAgentComplianceConfigDirVolumePath   = "/etc/datadog-agent/compliance.d"

	ClusterAgentCustomConfigVolumeName    = "custom-datadog-yaml"
	ClusterAgentCustomConfigVolumePath    = "/etc/datadog-agent/datadog-cluster.yaml"
	ClusterAgentCustomConfigVolumeSubPath = "datadog-cluster.yaml"

	SysteProbeAppArmorAnnotationKey   = "container.apparmor.security.beta.kubernetes.io/system-probe"
	SysteProbeSeccompAnnotationKey    = "container.seccomp.security.alpha.kubernetes.io/system-probe"
	SystemProbeOSReleaseDirVolumeName = "host-osrelease"
	SystemProbeOSReleaseDirVolumePath = "/etc/os-release"
	SystemProbeOSReleaseDirMountPath  = "/host/etc/os-release"

	// Extra config provider names

	KubeServicesConfigProvider              = "kube_services"
	KubeEndpointsConfigProvider             = "kube_endpoints"
	KubeServicesAndEndpointsConfigProviders = "kube_services kube_endpoints"
	ClusterChecksConfigProvider             = "clusterchecks"
	EndpointsChecksConfigProvider           = "endpointschecks"
	ClusterAndEndpointsConfigPoviders       = "clusterchecks endpointschecks"

	// Extra listeners

	KubeServicesListener              = "kube_services"
	KubeEndpointsListener             = "kube_endpoints"
	KubeServicesAndEndpointsListeners = "kube_services kube_endpoints"

	// Consts used to setup Rbac config
	// API Groups

	CoreAPIGroup           = ""
	ExtensionsAPIGroup     = "extensions"
	OpenShiftQuotaAPIGroup = "quota.openshift.io"
	RbacAPIGroup           = "rbac.authorization.k8s.io"
	AutoscalingAPIGroup    = "autoscaling"
	CertificatesAPIGroup   = "certificates.k8s.io"
	StorageAPIGroup        = "storage.k8s.io"
	CoordinationAPIGroup   = "coordination.k8s.io"
	DatadogAPIGroup        = "datadoghq.com"
	AdmissionAPIGroup      = "admissionregistration.k8s.io"
	AppsAPIGroup           = "apps"
	BatchAPIGroup          = "batch"
	PolicyAPIGroup         = "policy"
	NetworkingAPIGroup     = "networking.k8s.io"

	// Resources

	ServicesResource                    = "services"
	EventsResource                      = "events"
	EndpointsResource                   = "endpoints"
	PodsResource                        = "pods"
	NodesResource                       = "nodes"
	ComponentStatusesResource           = "componentstatuses"
	CertificatesSigningRequestsResource = "certificatesigningrequests"
	ConfigMapsResource                  = "configmaps"
	ResourceQuotasResource              = "resourcequotas"
	ReplicationControllersResource      = "replicationcontrollers"
	LimitRangesResource                 = "limitranges"
	PersistentVolumeClaimsResource      = "persistentvolumeclaims"
	PersistentVolumesResource           = "persistentvolumes"
	LeasesResource                      = "leases"
	ClusterResourceQuotasResource       = "clusterresourcequotas"
	NodeMetricsResource                 = "nodes/metrics"
	NodeSpecResource                    = "nodes/spec"
	NodeProxyResource                   = "nodes/proxy"
	NodeStats                           = "nodes/stats"
	HorizontalPodAutoscalersRecource    = "horizontalpodautoscalers"
	DatadogMetricsResource              = "datadogmetrics"
	DatadogMetricsStatusResource        = "datadogmetrics/status"
	WpaResource                         = "watermarkpodautoscalers"
	MutatingConfigResource              = "mutatingwebhookconfigurations"
	ValidatingConfigResource            = "validatingwebhookconfigurations"
	SecretsResource                     = "secrets"
	PodDisruptionBudgetsResource        = "poddisruptionbudgets"
	ReplicasetsResource                 = "replicasets"
	DeploymentsResource                 = "deployments"
	StatefulsetsResource                = "statefulsets"
	DaemonsetsResource                  = "daemonsets"
	JobsResource                        = "jobs"
	CronjobsResource                    = "cronjobs"
	StorageClassesResource              = "storageclasses"
	VolumeAttachments                   = "volumeattachments"
	ExtendedDaemonSetReplicaSetResource = "extendeddaemonsetreplicasets"
	ServiceAccountResource              = "serviceaccounts"
	NamespaceResource                   = "namespaces"
	PodSecurityPolicyResource           = "podsecuritypolicies"
	ClusterRoleBindingResource          = "clusterrolebindings"
	RoleBindingResource                 = "rolebindings"
	NetworkPolicyResource               = "networkpolicies"
	IngressesResource                   = "ingresses"

	// Resource names

	DatadogTokenResourceName           = "datadogtoken"
	DatadogLeaderElectionResourceName  = "datadog-leader-election"
	DatadogCustomMetricsResourceName   = "datadog-custom-metrics"
	DatadogClusterIDResourceName       = "datadog-cluster-id"
	ExtensionAPIServerAuthResourceName = "extension-apiserver-authentication"
	KubeSystemResourceName             = "kube-system"

	// Non resource URLs

	VersionURL = "/version"
	HealthzURL = "/healthz"
	MetricsURL = "/metrics"

	// Verbs

	GetVerb    = "get"
	ListVerb   = "list"
	WatchVerb  = "watch"
	UpdateVerb = "update"
	CreateVerb = "create"
	DeleteVerb = "delete"

	// Rbac resource kinds

	ClusterRoleKind    = "ClusterRole"
	RoleKind           = "Role"
	ServiceAccountKind = "ServiceAccount"
)
