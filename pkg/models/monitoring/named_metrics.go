/*
Copyright 2020 KubeSphere Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package monitoring

const (
	KubeSphereWorkspaceCount = "kubesphere_workspace_count"
	KubeSphereUserCount      = "kubesphere_user_count"
	KubeSphereClusterCount   = "kubesphere_cluser_count"
	KubeSphereAppTmplCount   = "kubesphere_app_template_count"

	WorkspaceNamespaceCount = "workspace_namespace_count"
	WorkspaceDevopsCount    = "workspace_devops_project_count"
	WorkspaceMemberCount    = "workspace_member_count"
	WorkspaceRoleCount      = "workspace_role_count"

	MetricMeterPrefix = "meter_"
)

var ClusterMetrics = []string{
	"cluster_cpu_utilisation",
	"cluster_cpu_usage",
	"cluster_cpu_total",
	"cluster_memory_utilisation",
	"cluster_memory_available",
	"cluster_memory_total",
	"cluster_memory_usage_wo_cache",
	"cluster_net_utilisation",
	"cluster_net_bytes_transmitted",
	"cluster_net_bytes_received",
	"cluster_net_packets_transmitted",
	"cluster_net_packets_received",
	"cluster_disk_read_iops",
	"cluster_disk_write_iops",
	"cluster_disk_read_throughput",
	"cluster_disk_write_throughput",
	"cluster_disk_size_usage",
	"cluster_disk_size_utilisation",
	"cluster_disk_size_capacity",
	"cluster_disk_size_available",
	"cluster_disk_inode_total",
	"cluster_disk_inode_usage",
	"cluster_disk_inode_utilisation",
	"cluster_namespace_count",
	"cluster_pod_count",
	"cluster_pod_quota",
	"cluster_pod_utilisation",
	"cluster_pod_running_count",
	"cluster_pod_succeeded_count",
	"cluster_pod_abnormal_count",
	"cluster_node_online",
	"cluster_node_offline",
	"cluster_node_total",
	"cluster_cronjob_count",
	"cluster_pvc_count",
	"cluster_daemonset_count",
	"cluster_deployment_count",
	"cluster_endpoint_count",
	"cluster_hpa_count",
	"cluster_job_count",
	"cluster_statefulset_count",
	"cluster_replicaset_count",
	"cluster_service_count",
	"cluster_secret_count",
	"cluster_pv_count",
	"cluster_ingresses_extensions_count",
	"cluster_load1",
	"cluster_load5",
	"cluster_load15",
	"cluster_pod_abnormal_ratio",
	"cluster_node_offline_ratio",

	// meter
	"meter_cluster_cpu_usage",
	"meter_cluster_memory_usage",
	"meter_cluster_net_bytes_transmitted",
	"meter_cluster_net_bytes_received",
	"meter_cluster_pvc_bytes_total",
}

var NodeMetrics = []string{
	"node_cpu_utilisation",
	"node_cpu_total",
	"node_cpu_usage",
	"node_memory_utilisation",
	"node_memory_usage_wo_cache",
	"node_memory_available",
	"node_memory_total",
	"node_net_utilisation",
	"node_net_bytes_transmitted",
	"node_net_bytes_received",
	"node_net_packets_transmitted",
	"node_net_packets_received",
	"node_disk_read_iops",
	"node_disk_write_iops",
	"node_disk_read_throughput",
	"node_disk_write_throughput",
	"node_disk_size_capacity",
	"node_disk_size_available",
	"node_disk_size_usage",
	"node_disk_size_utilisation",
	"node_disk_inode_total",
	"node_disk_inode_usage",
	"node_disk_inode_utilisation",
	"node_pod_count",
	"node_pod_quota",
	"node_pod_utilisation",
	"node_pod_running_count",
	"node_pod_succeeded_count",
	"node_pod_abnormal_count",
	"node_load1",
	"node_load5",
	"node_load15",
	"node_pod_abnormal_ratio",
	"node_pleg_quantile",

	"node_device_size_usage",
	"node_device_size_utilisation",

	// meter
	"meter_node_cpu_usage",
	"meter_node_memory_usage_wo_cache",
	"meter_node_net_bytes_transmitted",
	"meter_node_net_bytes_received",
	"meter_node_pvc_bytes_total",
}

var WorkspaceMetrics = []string{
	"workspace_cpu_usage",
	"workspace_memory_usage",
	"workspace_memory_usage_wo_cache",
	"workspace_net_bytes_transmitted",
	"workspace_net_bytes_received",
	"workspace_net_packets_transmitted",
	"workspace_net_packets_received",
	"workspace_pod_count",
	"workspace_pod_running_count",
	"workspace_pod_succeeded_count",
	"workspace_pod_abnormal_count",
	"workspace_ingresses_extensions_count",
	"workspace_cronjob_count",
	"workspace_pvc_count",
	"workspace_daemonset_count",
	"workspace_deployment_count",
	"workspace_endpoint_count",
	"workspace_hpa_count",
	"workspace_job_count",
	"workspace_statefulset_count",
	"workspace_replicaset_count",
	"workspace_service_count",
	"workspace_secret_count",
	"workspace_pod_abnormal_ratio",

	// meter
	"meter_workspace_cpu_usage",
	"meter_workspace_memory_usage",
	"meter_workspace_net_bytes_transmitted",
	"meter_workspace_net_bytes_received",
	"meter_workspace_pvc_bytes_total",
}

var NamespaceMetrics = []string{
	"namespace_cpu_usage",
	"namespace_memory_usage",
	"namespace_memory_usage_wo_cache",
	"namespace_net_bytes_transmitted",
	"namespace_net_bytes_received",
	"namespace_net_packets_transmitted",
	"namespace_net_packets_received",
	"namespace_pod_count",
	"namespace_pod_running_count",
	"namespace_pod_succeeded_count",
	"namespace_pod_abnormal_count",
	"namespace_pod_abnormal_ratio",
	"namespace_memory_limit_hard",
	"namespace_cpu_limit_hard",
	"namespace_pod_count_hard",
	"namespace_cronjob_count",
	"namespace_pvc_count",
	"namespace_daemonset_count",
	"namespace_deployment_count",
	"namespace_endpoint_count",
	"namespace_hpa_count",
	"namespace_job_count",
	"namespace_statefulset_count",
	"namespace_replicaset_count",
	"namespace_service_count",
	"namespace_secret_count",
	"namespace_configmap_count",
	"namespace_ingresses_extensions_count",
	"namespace_s2ibuilder_count",

	// meter
	"meter_namespace_cpu_usage",
	"meter_namespace_memory_usage_wo_cache",
	"meter_namespace_net_bytes_transmitted",
	"meter_namespace_net_bytes_received",
	"meter_namespace_pvc_bytes_total",
}

var ApplicationMetrics = []string{

	// meter
	"meter_application_cpu_usage",
	"meter_application_memory_usage_wo_cache",
	"meter_application_net_bytes_transmitted",
	"meter_application_net_bytes_received",
	"meter_application_pvc_bytes_total",
}

var WorkloadMetrics = []string{
	"workload_cpu_usage",
	"workload_memory_usage",
	"workload_memory_usage_wo_cache",
	"workload_net_bytes_transmitted",
	"workload_net_bytes_received",
	"workload_net_packets_transmitted",
	"workload_net_packets_received",
	"workload_deployment_replica",
	"workload_deployment_replica_available",
	"workload_statefulset_replica",
	"workload_statefulset_replica_available",
	"workload_daemonset_replica",
	"workload_daemonset_replica_available",
	"workload_deployment_unavailable_replicas_ratio",
	"workload_daemonset_unavailable_replicas_ratio",
	"workload_statefulset_unavailable_replicas_ratio",

	// meter
	"meter_workload_cpu_usage",
	"meter_workload_memory_usage_wo_cache",
	"meter_workload_net_bytes_transmitted",
	"meter_workload_net_bytes_received",
	"meter_workload_pvc_bytes_total",
}

var ServiceMetrics = []string{
	// meter
	"meter_service_cpu_usage",
	"meter_service_memory_usage_wo_cache",
	"meter_service_net_bytes_transmitted",
	"meter_service_net_bytes_received",
}

var VirtualmachineMetrics = []string{
	"virtualmachine_cpu_usage",
	"virtualmachine_memory_usage",
	"virtualmachine_disk_read_iops",
	"virtualmachine_disk_write_iops",
	"virtualmachine_disk_read_traffic_bytes",
	"virtualmachine_disk_write_traffic_bytes",
	"virtualmachine_disk_read_times_ms",
	"virtualmachine_disk_write_times_ms",
	"virtualmachine_network_bytes_transmitted",
	"virtualmachine_network_bytes_received",

	// meter
	"meter_virtualmachine_cpu_usage",
	"meter_virtualmachine_memory_usage_wo_cache",
	"meter_virtualmachine_net_bytes_transmitted",
	"meter_virtualmachine_net_bytes_received",
	"meter_virtualmachine_pvc_bytes_total",
}

var PodMetrics = []string{
	"pod_cpu_usage",
	"pod_memory_usage",
	"pod_memory_usage_wo_cache",
	"pod_net_bytes_transmitted",
	"pod_net_bytes_received",
	"pod_net_packets_transmitted",
	"pod_net_packets_received",

	// meter
	"meter_pod_cpu_usage",
	"meter_pod_memory_usage_wo_cache",
	"meter_pod_net_bytes_transmitted",
	"meter_pod_net_bytes_received",
	"meter_pod_pvc_bytes_total",
}

var ContainerMetrics = []string{
	"container_cpu_usage",
	"container_memory_usage",
	"container_memory_usage_wo_cache",
	"container_processes_usage",
	"container_threads_usage",
}

var PVCMetrics = []string{
	"pvc_inodes_available",
	"pvc_inodes_used",
	"pvc_inodes_total",
	"pvc_inodes_utilisation",
	"pvc_bytes_available",
	"pvc_bytes_used",
	"pvc_bytes_total",
	"pvc_bytes_utilisation",
}

var IngressMetrics = []string{
	"ingress_request_count",
	"ingress_request_5xx_count",
	"ingress_request_4xx_count",
	"ingress_active_connections",
	"ingress_success_rate",
	"ingress_request_duration_average",
	"ingress_request_duration_50percentage",
	"ingress_request_duration_95percentage",
	"ingress_request_duration_99percentage",
	"ingress_request_volume",
	"ingress_request_volume_by_ingress",
	"ingress_request_network_sent",
	"ingress_request_network_received",
	"ingress_request_memory_bytes",
	"ingress_request_cpu_usage",
}

var EtcdMetrics = []string{
	"etcd_server_list",
	"etcd_server_total",
	"etcd_server_up_total",
	"etcd_server_has_leader",
	"etcd_server_is_leader",
	"etcd_server_leader_changes",
	"etcd_server_proposals_failed_rate",
	"etcd_server_proposals_applied_rate",
	"etcd_server_proposals_committed_rate",
	"etcd_server_proposals_pending_count",
	"etcd_mvcc_db_size",
	"etcd_network_client_grpc_received_bytes",
	"etcd_network_client_grpc_sent_bytes",
	"etcd_grpc_call_rate",
	"etcd_grpc_call_failed_rate",
	"etcd_grpc_server_msg_received_rate",
	"etcd_grpc_server_msg_sent_rate",
	"etcd_disk_wal_fsync_duration",
	"etcd_disk_wal_fsync_duration_quantile",
	"etcd_disk_backend_commit_duration",
	"etcd_disk_backend_commit_duration_quantile",
}

var APIServerMetrics = []string{
	"apiserver_up_sum",
	"apiserver_request_rate",
	"apiserver_request_by_verb_rate",
	"apiserver_request_latencies",
	"apiserver_request_by_verb_latencies",
}

var SchedulerMetrics = []string{
	"scheduler_up_sum",
	"scheduler_schedule_attempts",
	"scheduler_schedule_attempt_rate",
	"scheduler_e2e_scheduling_latency",
	"scheduler_e2e_scheduling_latency_quantile",
}
