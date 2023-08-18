//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright (c) 2022 Aiven, Helsinki, Finland. https://aiven.io/

// Code generated by controller-gen. DO NOT EDIT.

package kafkauserconfig

import ()

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IpFilter) DeepCopyInto(out *IpFilter) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IpFilter.
func (in *IpFilter) DeepCopy() *IpFilter {
	if in == nil {
		return nil
	}
	out := new(IpFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Kafka) DeepCopyInto(out *Kafka) {
	*out = *in
	if in.AutoCreateTopicsEnable != nil {
		in, out := &in.AutoCreateTopicsEnable, &out.AutoCreateTopicsEnable
		*out = new(bool)
		**out = **in
	}
	if in.CompressionType != nil {
		in, out := &in.CompressionType, &out.CompressionType
		*out = new(string)
		**out = **in
	}
	if in.ConnectionsMaxIdleMs != nil {
		in, out := &in.ConnectionsMaxIdleMs, &out.ConnectionsMaxIdleMs
		*out = new(int)
		**out = **in
	}
	if in.DefaultReplicationFactor != nil {
		in, out := &in.DefaultReplicationFactor, &out.DefaultReplicationFactor
		*out = new(int)
		**out = **in
	}
	if in.GroupInitialRebalanceDelayMs != nil {
		in, out := &in.GroupInitialRebalanceDelayMs, &out.GroupInitialRebalanceDelayMs
		*out = new(int)
		**out = **in
	}
	if in.GroupMaxSessionTimeoutMs != nil {
		in, out := &in.GroupMaxSessionTimeoutMs, &out.GroupMaxSessionTimeoutMs
		*out = new(int)
		**out = **in
	}
	if in.GroupMinSessionTimeoutMs != nil {
		in, out := &in.GroupMinSessionTimeoutMs, &out.GroupMinSessionTimeoutMs
		*out = new(int)
		**out = **in
	}
	if in.LogCleanerDeleteRetentionMs != nil {
		in, out := &in.LogCleanerDeleteRetentionMs, &out.LogCleanerDeleteRetentionMs
		*out = new(int)
		**out = **in
	}
	if in.LogCleanerMaxCompactionLagMs != nil {
		in, out := &in.LogCleanerMaxCompactionLagMs, &out.LogCleanerMaxCompactionLagMs
		*out = new(int)
		**out = **in
	}
	if in.LogCleanerMinCleanableRatio != nil {
		in, out := &in.LogCleanerMinCleanableRatio, &out.LogCleanerMinCleanableRatio
		*out = new(float64)
		**out = **in
	}
	if in.LogCleanerMinCompactionLagMs != nil {
		in, out := &in.LogCleanerMinCompactionLagMs, &out.LogCleanerMinCompactionLagMs
		*out = new(int)
		**out = **in
	}
	if in.LogCleanupPolicy != nil {
		in, out := &in.LogCleanupPolicy, &out.LogCleanupPolicy
		*out = new(string)
		**out = **in
	}
	if in.LogFlushIntervalMessages != nil {
		in, out := &in.LogFlushIntervalMessages, &out.LogFlushIntervalMessages
		*out = new(int)
		**out = **in
	}
	if in.LogFlushIntervalMs != nil {
		in, out := &in.LogFlushIntervalMs, &out.LogFlushIntervalMs
		*out = new(int)
		**out = **in
	}
	if in.LogIndexIntervalBytes != nil {
		in, out := &in.LogIndexIntervalBytes, &out.LogIndexIntervalBytes
		*out = new(int)
		**out = **in
	}
	if in.LogIndexSizeMaxBytes != nil {
		in, out := &in.LogIndexSizeMaxBytes, &out.LogIndexSizeMaxBytes
		*out = new(int)
		**out = **in
	}
	if in.LogMessageDownconversionEnable != nil {
		in, out := &in.LogMessageDownconversionEnable, &out.LogMessageDownconversionEnable
		*out = new(bool)
		**out = **in
	}
	if in.LogMessageTimestampDifferenceMaxMs != nil {
		in, out := &in.LogMessageTimestampDifferenceMaxMs, &out.LogMessageTimestampDifferenceMaxMs
		*out = new(int)
		**out = **in
	}
	if in.LogMessageTimestampType != nil {
		in, out := &in.LogMessageTimestampType, &out.LogMessageTimestampType
		*out = new(string)
		**out = **in
	}
	if in.LogPreallocate != nil {
		in, out := &in.LogPreallocate, &out.LogPreallocate
		*out = new(bool)
		**out = **in
	}
	if in.LogRetentionBytes != nil {
		in, out := &in.LogRetentionBytes, &out.LogRetentionBytes
		*out = new(int)
		**out = **in
	}
	if in.LogRetentionHours != nil {
		in, out := &in.LogRetentionHours, &out.LogRetentionHours
		*out = new(int)
		**out = **in
	}
	if in.LogRetentionMs != nil {
		in, out := &in.LogRetentionMs, &out.LogRetentionMs
		*out = new(int)
		**out = **in
	}
	if in.LogRollJitterMs != nil {
		in, out := &in.LogRollJitterMs, &out.LogRollJitterMs
		*out = new(int)
		**out = **in
	}
	if in.LogRollMs != nil {
		in, out := &in.LogRollMs, &out.LogRollMs
		*out = new(int)
		**out = **in
	}
	if in.LogSegmentBytes != nil {
		in, out := &in.LogSegmentBytes, &out.LogSegmentBytes
		*out = new(int)
		**out = **in
	}
	if in.LogSegmentDeleteDelayMs != nil {
		in, out := &in.LogSegmentDeleteDelayMs, &out.LogSegmentDeleteDelayMs
		*out = new(int)
		**out = **in
	}
	if in.MaxConnectionsPerIp != nil {
		in, out := &in.MaxConnectionsPerIp, &out.MaxConnectionsPerIp
		*out = new(int)
		**out = **in
	}
	if in.MaxIncrementalFetchSessionCacheSlots != nil {
		in, out := &in.MaxIncrementalFetchSessionCacheSlots, &out.MaxIncrementalFetchSessionCacheSlots
		*out = new(int)
		**out = **in
	}
	if in.MessageMaxBytes != nil {
		in, out := &in.MessageMaxBytes, &out.MessageMaxBytes
		*out = new(int)
		**out = **in
	}
	if in.MinInsyncReplicas != nil {
		in, out := &in.MinInsyncReplicas, &out.MinInsyncReplicas
		*out = new(int)
		**out = **in
	}
	if in.NumPartitions != nil {
		in, out := &in.NumPartitions, &out.NumPartitions
		*out = new(int)
		**out = **in
	}
	if in.OffsetsRetentionMinutes != nil {
		in, out := &in.OffsetsRetentionMinutes, &out.OffsetsRetentionMinutes
		*out = new(int)
		**out = **in
	}
	if in.ProducerPurgatoryPurgeIntervalRequests != nil {
		in, out := &in.ProducerPurgatoryPurgeIntervalRequests, &out.ProducerPurgatoryPurgeIntervalRequests
		*out = new(int)
		**out = **in
	}
	if in.RemoteLogStorageSystemEnable != nil {
		in, out := &in.RemoteLogStorageSystemEnable, &out.RemoteLogStorageSystemEnable
		*out = new(bool)
		**out = **in
	}
	if in.ReplicaFetchMaxBytes != nil {
		in, out := &in.ReplicaFetchMaxBytes, &out.ReplicaFetchMaxBytes
		*out = new(int)
		**out = **in
	}
	if in.ReplicaFetchResponseMaxBytes != nil {
		in, out := &in.ReplicaFetchResponseMaxBytes, &out.ReplicaFetchResponseMaxBytes
		*out = new(int)
		**out = **in
	}
	if in.SocketRequestMaxBytes != nil {
		in, out := &in.SocketRequestMaxBytes, &out.SocketRequestMaxBytes
		*out = new(int)
		**out = **in
	}
	if in.TransactionRemoveExpiredTransactionCleanupIntervalMs != nil {
		in, out := &in.TransactionRemoveExpiredTransactionCleanupIntervalMs, &out.TransactionRemoveExpiredTransactionCleanupIntervalMs
		*out = new(int)
		**out = **in
	}
	if in.TransactionStateLogSegmentBytes != nil {
		in, out := &in.TransactionStateLogSegmentBytes, &out.TransactionStateLogSegmentBytes
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Kafka.
func (in *Kafka) DeepCopy() *Kafka {
	if in == nil {
		return nil
	}
	out := new(Kafka)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaAuthenticationMethods) DeepCopyInto(out *KafkaAuthenticationMethods) {
	*out = *in
	if in.Certificate != nil {
		in, out := &in.Certificate, &out.Certificate
		*out = new(bool)
		**out = **in
	}
	if in.Sasl != nil {
		in, out := &in.Sasl, &out.Sasl
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaAuthenticationMethods.
func (in *KafkaAuthenticationMethods) DeepCopy() *KafkaAuthenticationMethods {
	if in == nil {
		return nil
	}
	out := new(KafkaAuthenticationMethods)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaConnectConfig) DeepCopyInto(out *KafkaConnectConfig) {
	*out = *in
	if in.ConnectorClientConfigOverridePolicy != nil {
		in, out := &in.ConnectorClientConfigOverridePolicy, &out.ConnectorClientConfigOverridePolicy
		*out = new(string)
		**out = **in
	}
	if in.ConsumerAutoOffsetReset != nil {
		in, out := &in.ConsumerAutoOffsetReset, &out.ConsumerAutoOffsetReset
		*out = new(string)
		**out = **in
	}
	if in.ConsumerFetchMaxBytes != nil {
		in, out := &in.ConsumerFetchMaxBytes, &out.ConsumerFetchMaxBytes
		*out = new(int)
		**out = **in
	}
	if in.ConsumerIsolationLevel != nil {
		in, out := &in.ConsumerIsolationLevel, &out.ConsumerIsolationLevel
		*out = new(string)
		**out = **in
	}
	if in.ConsumerMaxPartitionFetchBytes != nil {
		in, out := &in.ConsumerMaxPartitionFetchBytes, &out.ConsumerMaxPartitionFetchBytes
		*out = new(int)
		**out = **in
	}
	if in.ConsumerMaxPollIntervalMs != nil {
		in, out := &in.ConsumerMaxPollIntervalMs, &out.ConsumerMaxPollIntervalMs
		*out = new(int)
		**out = **in
	}
	if in.ConsumerMaxPollRecords != nil {
		in, out := &in.ConsumerMaxPollRecords, &out.ConsumerMaxPollRecords
		*out = new(int)
		**out = **in
	}
	if in.OffsetFlushIntervalMs != nil {
		in, out := &in.OffsetFlushIntervalMs, &out.OffsetFlushIntervalMs
		*out = new(int)
		**out = **in
	}
	if in.OffsetFlushTimeoutMs != nil {
		in, out := &in.OffsetFlushTimeoutMs, &out.OffsetFlushTimeoutMs
		*out = new(int)
		**out = **in
	}
	if in.ProducerBatchSize != nil {
		in, out := &in.ProducerBatchSize, &out.ProducerBatchSize
		*out = new(int)
		**out = **in
	}
	if in.ProducerBufferMemory != nil {
		in, out := &in.ProducerBufferMemory, &out.ProducerBufferMemory
		*out = new(int)
		**out = **in
	}
	if in.ProducerCompressionType != nil {
		in, out := &in.ProducerCompressionType, &out.ProducerCompressionType
		*out = new(string)
		**out = **in
	}
	if in.ProducerLingerMs != nil {
		in, out := &in.ProducerLingerMs, &out.ProducerLingerMs
		*out = new(int)
		**out = **in
	}
	if in.ProducerMaxRequestSize != nil {
		in, out := &in.ProducerMaxRequestSize, &out.ProducerMaxRequestSize
		*out = new(int)
		**out = **in
	}
	if in.ScheduledRebalanceMaxDelayMs != nil {
		in, out := &in.ScheduledRebalanceMaxDelayMs, &out.ScheduledRebalanceMaxDelayMs
		*out = new(int)
		**out = **in
	}
	if in.SessionTimeoutMs != nil {
		in, out := &in.SessionTimeoutMs, &out.SessionTimeoutMs
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaConnectConfig.
func (in *KafkaConnectConfig) DeepCopy() *KafkaConnectConfig {
	if in == nil {
		return nil
	}
	out := new(KafkaConnectConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaRestConfig) DeepCopyInto(out *KafkaRestConfig) {
	*out = *in
	if in.ConsumerEnableAutoCommit != nil {
		in, out := &in.ConsumerEnableAutoCommit, &out.ConsumerEnableAutoCommit
		*out = new(bool)
		**out = **in
	}
	if in.ConsumerRequestMaxBytes != nil {
		in, out := &in.ConsumerRequestMaxBytes, &out.ConsumerRequestMaxBytes
		*out = new(int)
		**out = **in
	}
	if in.ConsumerRequestTimeoutMs != nil {
		in, out := &in.ConsumerRequestTimeoutMs, &out.ConsumerRequestTimeoutMs
		*out = new(int)
		**out = **in
	}
	if in.ProducerAcks != nil {
		in, out := &in.ProducerAcks, &out.ProducerAcks
		*out = new(string)
		**out = **in
	}
	if in.ProducerCompressionType != nil {
		in, out := &in.ProducerCompressionType, &out.ProducerCompressionType
		*out = new(string)
		**out = **in
	}
	if in.ProducerLingerMs != nil {
		in, out := &in.ProducerLingerMs, &out.ProducerLingerMs
		*out = new(int)
		**out = **in
	}
	if in.ProducerMaxRequestSize != nil {
		in, out := &in.ProducerMaxRequestSize, &out.ProducerMaxRequestSize
		*out = new(int)
		**out = **in
	}
	if in.SimpleconsumerPoolSizeMax != nil {
		in, out := &in.SimpleconsumerPoolSizeMax, &out.SimpleconsumerPoolSizeMax
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaRestConfig.
func (in *KafkaRestConfig) DeepCopy() *KafkaRestConfig {
	if in == nil {
		return nil
	}
	out := new(KafkaRestConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaUserConfig) DeepCopyInto(out *KafkaUserConfig) {
	*out = *in
	if in.AdditionalBackupRegions != nil {
		in, out := &in.AdditionalBackupRegions, &out.AdditionalBackupRegions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CustomDomain != nil {
		in, out := &in.CustomDomain, &out.CustomDomain
		*out = new(string)
		**out = **in
	}
	if in.IpFilter != nil {
		in, out := &in.IpFilter, &out.IpFilter
		*out = make([]*IpFilter, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(IpFilter)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Kafka != nil {
		in, out := &in.Kafka, &out.Kafka
		*out = new(Kafka)
		(*in).DeepCopyInto(*out)
	}
	if in.KafkaAuthenticationMethods != nil {
		in, out := &in.KafkaAuthenticationMethods, &out.KafkaAuthenticationMethods
		*out = new(KafkaAuthenticationMethods)
		(*in).DeepCopyInto(*out)
	}
	if in.KafkaConnect != nil {
		in, out := &in.KafkaConnect, &out.KafkaConnect
		*out = new(bool)
		**out = **in
	}
	if in.KafkaConnectConfig != nil {
		in, out := &in.KafkaConnectConfig, &out.KafkaConnectConfig
		*out = new(KafkaConnectConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.KafkaRest != nil {
		in, out := &in.KafkaRest, &out.KafkaRest
		*out = new(bool)
		**out = **in
	}
	if in.KafkaRestAuthorization != nil {
		in, out := &in.KafkaRestAuthorization, &out.KafkaRestAuthorization
		*out = new(bool)
		**out = **in
	}
	if in.KafkaRestConfig != nil {
		in, out := &in.KafkaRestConfig, &out.KafkaRestConfig
		*out = new(KafkaRestConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.KafkaVersion != nil {
		in, out := &in.KafkaVersion, &out.KafkaVersion
		*out = new(string)
		**out = **in
	}
	if in.PrivateAccess != nil {
		in, out := &in.PrivateAccess, &out.PrivateAccess
		*out = new(PrivateAccess)
		(*in).DeepCopyInto(*out)
	}
	if in.PrivatelinkAccess != nil {
		in, out := &in.PrivatelinkAccess, &out.PrivatelinkAccess
		*out = new(PrivatelinkAccess)
		(*in).DeepCopyInto(*out)
	}
	if in.PublicAccess != nil {
		in, out := &in.PublicAccess, &out.PublicAccess
		*out = new(PublicAccess)
		(*in).DeepCopyInto(*out)
	}
	if in.SchemaRegistry != nil {
		in, out := &in.SchemaRegistry, &out.SchemaRegistry
		*out = new(bool)
		**out = **in
	}
	if in.SchemaRegistryConfig != nil {
		in, out := &in.SchemaRegistryConfig, &out.SchemaRegistryConfig
		*out = new(SchemaRegistryConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.StaticIps != nil {
		in, out := &in.StaticIps, &out.StaticIps
		*out = new(bool)
		**out = **in
	}
	if in.TieredStorage != nil {
		in, out := &in.TieredStorage, &out.TieredStorage
		*out = new(TieredStorage)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaUserConfig.
func (in *KafkaUserConfig) DeepCopy() *KafkaUserConfig {
	if in == nil {
		return nil
	}
	out := new(KafkaUserConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalCache) DeepCopyInto(out *LocalCache) {
	*out = *in
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalCache.
func (in *LocalCache) DeepCopy() *LocalCache {
	if in == nil {
		return nil
	}
	out := new(LocalCache)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateAccess) DeepCopyInto(out *PrivateAccess) {
	*out = *in
	if in.Kafka != nil {
		in, out := &in.Kafka, &out.Kafka
		*out = new(bool)
		**out = **in
	}
	if in.KafkaConnect != nil {
		in, out := &in.KafkaConnect, &out.KafkaConnect
		*out = new(bool)
		**out = **in
	}
	if in.KafkaRest != nil {
		in, out := &in.KafkaRest, &out.KafkaRest
		*out = new(bool)
		**out = **in
	}
	if in.Prometheus != nil {
		in, out := &in.Prometheus, &out.Prometheus
		*out = new(bool)
		**out = **in
	}
	if in.SchemaRegistry != nil {
		in, out := &in.SchemaRegistry, &out.SchemaRegistry
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateAccess.
func (in *PrivateAccess) DeepCopy() *PrivateAccess {
	if in == nil {
		return nil
	}
	out := new(PrivateAccess)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivatelinkAccess) DeepCopyInto(out *PrivatelinkAccess) {
	*out = *in
	if in.Jolokia != nil {
		in, out := &in.Jolokia, &out.Jolokia
		*out = new(bool)
		**out = **in
	}
	if in.Kafka != nil {
		in, out := &in.Kafka, &out.Kafka
		*out = new(bool)
		**out = **in
	}
	if in.KafkaConnect != nil {
		in, out := &in.KafkaConnect, &out.KafkaConnect
		*out = new(bool)
		**out = **in
	}
	if in.KafkaRest != nil {
		in, out := &in.KafkaRest, &out.KafkaRest
		*out = new(bool)
		**out = **in
	}
	if in.Prometheus != nil {
		in, out := &in.Prometheus, &out.Prometheus
		*out = new(bool)
		**out = **in
	}
	if in.SchemaRegistry != nil {
		in, out := &in.SchemaRegistry, &out.SchemaRegistry
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivatelinkAccess.
func (in *PrivatelinkAccess) DeepCopy() *PrivatelinkAccess {
	if in == nil {
		return nil
	}
	out := new(PrivatelinkAccess)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicAccess) DeepCopyInto(out *PublicAccess) {
	*out = *in
	if in.Kafka != nil {
		in, out := &in.Kafka, &out.Kafka
		*out = new(bool)
		**out = **in
	}
	if in.KafkaConnect != nil {
		in, out := &in.KafkaConnect, &out.KafkaConnect
		*out = new(bool)
		**out = **in
	}
	if in.KafkaRest != nil {
		in, out := &in.KafkaRest, &out.KafkaRest
		*out = new(bool)
		**out = **in
	}
	if in.Prometheus != nil {
		in, out := &in.Prometheus, &out.Prometheus
		*out = new(bool)
		**out = **in
	}
	if in.SchemaRegistry != nil {
		in, out := &in.SchemaRegistry, &out.SchemaRegistry
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicAccess.
func (in *PublicAccess) DeepCopy() *PublicAccess {
	if in == nil {
		return nil
	}
	out := new(PublicAccess)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchemaRegistryConfig) DeepCopyInto(out *SchemaRegistryConfig) {
	*out = *in
	if in.LeaderEligibility != nil {
		in, out := &in.LeaderEligibility, &out.LeaderEligibility
		*out = new(bool)
		**out = **in
	}
	if in.TopicName != nil {
		in, out := &in.TopicName, &out.TopicName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchemaRegistryConfig.
func (in *SchemaRegistryConfig) DeepCopy() *SchemaRegistryConfig {
	if in == nil {
		return nil
	}
	out := new(SchemaRegistryConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TieredStorage) DeepCopyInto(out *TieredStorage) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.LocalCache != nil {
		in, out := &in.LocalCache, &out.LocalCache
		*out = new(LocalCache)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TieredStorage.
func (in *TieredStorage) DeepCopy() *TieredStorage {
	if in == nil {
		return nil
	}
	out := new(TieredStorage)
	in.DeepCopyInto(out)
	return out
}
