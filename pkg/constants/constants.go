package constants

import (
	"fmt"

	batchv1 "k8s.io/api/batch/v1"
	"k8s.io/apimachinery/pkg/util/sets"
)

const (

	// DefaultJobReplicas is the default value for the ReplicatedJob replicas.
	DefaultJobReplicas = 1

	// JobSetKind is the Kind name for the JobSet.
	JobSetKind string = "JobSet"

	// LabelTrainJobAncestor is the label to identify relationship between
	// TrainJob and Pod template in the Runtime. The following labels are supported:
	// trainer.kubeflow.org/trainjob-ancestor-step: dataset-initializer  - trainJob.spec.initializer.dataset
	// trainer.kubeflow.org/trainjob-ancestor-step: model-initializer    - trainJob.spec.initializer.model
	// trainer.kubeflow.org/trainjob-ancestor-step: trainer              - trainJob.spec.trainer
	LabelTrainJobAncestor string = "trainer.kubeflow.org/trainjob-ancestor-step"

	// DatasetInitializer is the name of the Job, volume mount, container, and label value for the dataset initializer.
	DatasetInitializer string = "dataset-initializer"

	// DatasetMountPath is the volumeMount path for dataset.
	DatasetMountPath string = "/workspace/dataset"

	// ModelInitializer is the name of the Job, volume mount, container, and label value for the model initializer.
	ModelInitializer string = "model-initializer"

	// ModelMountPath is the volumeMount path for model.
	ModelMountPath string = "/workspace/model"

	// AncestorTrainer is the ancestor name for Trainer, which is mostly used for the value of
	// 'trainer.kubeflow.org/trainjob-ancestor-step'.
	AncestorTrainer string = "trainer"

	// Node is the name of the Job and container for the trainer node
	Node string = "node"

	// ContainerTrainerPort is the default port for the trainer nodes communication.
	ContainerTrainerPort int32 = 29500

	// PodGroupKind is the Kind name for the PodGroup.
	PodGroupKind string = "PodGroup"

	// TrainJobJobsCreationSucceededMessage is status condition message for the
	// {"type": "Created", "status": "True", "reason": "JobsCreationSucceeded"} condition.
	TrainJobJobsCreationSucceededMessage = "Succeeded to create Jobs"

	// TrainJobJobsBuildFailedMessage is status condition message for the
	// {"type": "Created", "status": "True", "reason": "JobsBuildFailed"} condition.
	TrainJobJobsBuildFailedMessage = "Failed to build Jobs"

	// TrainJobJobsCreationFailedMessage is status condition message for the
	// {"type": "Created", "status": "True", "reason": "JobsCreationFailed"} condition.
	TrainJobJobsCreationFailedMessage = "Failed to create Jobs"

	// TrainJobSuspendedMessage is status condition message for the
	// {"type": "Suspended", "status": "True", "reason": "Suspended"} condition.
	TrainJobSuspendedMessage = "TrainJob is suspended"

	// TrainJobResumedMessage is status condition message for the
	// {"type": "Suspended", "status": "True", "reason": "Resumed"} condition.
	TrainJobResumedMessage = "TrainJob is resumed"

	// Node is the name of the Job and container for the MPI launcher.
	// When RunLauncherAsNode: true, for the launcher Job the container name is node.
	Launcher string = "launcher"

	// MPISSHAuthSecretSuffix is the name suffix for Secret with MPI SSH keys.
	MPISSHAuthSecretSuffix string = "-mpi-ssh-auth"

	// MPISSHAuthVolumeName is the volume name for Secret with MPI SSH keys.
	MPISSHAuthVolumeName string = "mpi-ssh-auth"

	// MPISSHPrivateKeyFile is the file name for the private key.
	MPISSHPrivateKeyFile string = "id_rsa"

	// MPISSHPublicKey is the value in Secret data for the public key.
	MPISSHPublicKey string = "ssh-publickey"

	// MPISSHPublicKeyFile is the file name for the public key.
	MPISSHPublicKeyFile string = MPISSHPrivateKeyFile + ".pub"

	// MPISSHAuthorizedKeys is the file name for authorized keys.
	MPISSHAuthorizedKeys string = "authorized_keys"

	// MPIHostfilePath is the directory for the MPI hostfile.
	MPIHostfileDir string = "/etc/mpi"

	// MPIHostfileName is the file name for the MPI hostfile.
	MPIHostfileName string = "hostfile"

	// MPIHostfileConfigMapSuffix is the name suffix for ConfigMap with MPI hostfile.
	MPIHostfileConfigMapSuffix string = "-mpi-hostfile"

	// MPIHostfileVolumeName is the volume name for ConfigMap with MPI hostfile.
	MPIHostfileVolumeName string = "mpi-hostfile"

	// Distributed envs for mpirun.
	// Values for OpenMPI implementation.

	// OpenMPIEnvHostFileLocation is the OpenMPI default hostfile env key.
	OpenMPIEnvHostFileLocation string = "OMPI_MCA_orte_default_hostfile"

	// OpenMPIEnvKeyRSHArgs is the env key for the OpenMPI rsh arguments.
	OpenMPIEnvKeyRSHArgs string = "OMPI_MCA_plm_rsh_args"

	// OpenMPIEnvDefaultValueRSHArgs is the default env valur for the OpenMPI rsh arguments.
	OpenMPIEnvDefaultValueRSHArgs string = "-o ConnectionAttempts=10"

	// OpenMPIEnvKeepFQDNHostNames is the env key for OpenMPI if FQDN should be kept.
	OpenMPIEnvKeepFQDNHostNames string = "OMPI_MCA_orte_keep_fqdn_hostnames"

	// OpenMPIEnvDefaultSlots is the OpenMPI default number of slots env key.
	OpenMPIEnvDefaultSlots string = "OMPI_MCA_orte_set_default_slots"
	// Distributed envs for torchrun.
	// Ref: https://github.com/pytorch/pytorch/blob/3a0d0885171376ed610c8175a19ba40411fc6f3f/torch/distributed/argparse_util.py#L45
	// TorchEnvNumNodes is the env name for the number of training nodes.
	TorchEnvNumNodes string = "PET_NNODES"

	// TorchEnvNumProcPerNode is the env name for the number of procs per node (e.g. number of GPUs per Pod).
	TorchEnvNumProcPerNode string = "PET_NPROC_PER_NODE"

	// TorchEnvNodeRank is the env name for the node RANK
	TorchEnvNodeRank string = "PET_NODE_RANK"

	// TorchEnvMasterAddr is the env name for the master node address.
	TorchEnvMasterAddr string = "PET_MASTER_ADDR"

	// TorchEnvMasterPort is the env name for the master node port.
	TorchEnvMasterPort string = "PET_MASTER_PORT"
)

var (
	// JobCompletionIndexFieldPath is the field path for the Job completion index annotation.
	JobCompletionIndexFieldPath string = fmt.Sprintf("metadata.annotations['%s']", batchv1.JobCompletionIndexAnnotation)

	// Torchrun reserved env names
	TorchRunReservedEnvNames = sets.New(TorchEnvNumNodes, TorchEnvNumProcPerNode, TorchEnvNodeRank, TorchEnvMasterAddr, TorchEnvMasterPort)
)
