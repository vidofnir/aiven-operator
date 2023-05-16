// Code generated by user config generator. DO NOT EDIT.
// +kubebuilder:object:generate=true

package kafkamirrormakeruserconfig

// Kafka MirrorMaker configuration values
type KafkaMirrormaker struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=5242880
	// The minimum amount of data the server should return for a fetch request
	ConsumerFetchMinBytes *int `groups:"create,update" json:"consumer_fetch_min_bytes,omitempty"`

	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=5242880
	// The batch size in bytes producer will attempt to collect before publishing to broker.
	ProducerBatchSize *int `groups:"create,update" json:"producer_batch_size,omitempty"`

	// +kubebuilder:validation:Minimum=5242880
	// +kubebuilder:validation:Maximum=134217728
	// The amount of bytes producer can use for buffering data before publishing to broker.
	ProducerBufferMemory *int `groups:"create,update" json:"producer_buffer_memory,omitempty"`

	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=5000
	// The linger time (ms) for waiting new data to arrive for publishing.
	ProducerLingerMs *int `groups:"create,update" json:"producer_linger_ms,omitempty"`

	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=67108864
	// The maximum request size in bytes.
	ProducerMaxRequestSize *int `groups:"create,update" json:"producer_max_request_size,omitempty"`
}

// Integration user config
type KafkaMirrormakerUserConfig struct {
	// +kubebuilder:validation:MaxLength=128
	// +kubebuilder:validation:Pattern=`^[a-zA-Z0-9_.-]+$`
	// The alias under which the Kafka cluster is known to MirrorMaker. Can contain the following symbols: ASCII alphanumerics, '.', '_', and '-'.
	ClusterAlias *string `groups:"create,update" json:"cluster_alias,omitempty"`

	// Kafka MirrorMaker configuration values
	KafkaMirrormaker *KafkaMirrormaker `groups:"create,update" json:"kafka_mirrormaker,omitempty"`
}