// Copyright (c) 2017 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0
//

package annotations

const (
	kataAnnotationsPrefix     = "io.katacontainers."
	kataConfAnnotationsPrefix = kataAnnotationsPrefix + "config."
	kataAnnotHypervisorPrefix = kataConfAnnotationsPrefix + "hypervisor."
	kataAnnotContainerPrefix  = kataAnnotationsPrefix + "container."

	//
	// OCI
	//

	// BundlePathKey is the annotation key to fetch the OCI configuration file path.
	BundlePathKey = kataAnnotationsPrefix + "pkg.oci.bundle_path"

	// ContainerTypeKey is the annotation key to fetch container type.
	ContainerTypeKey = kataAnnotationsPrefix + "pkg.oci.container_type"

	SandboxConfigPathKey = kataAnnotationsPrefix + "config_path"
)

// Annotations related to Hypervisor configuration
const (
	//
	// Assets
	//
	KataAnnotationHypervisorPrefix = kataAnnotHypervisorPrefix

	// KernelPath is a sandbox annotation for passing a per container path pointing at the kernel needed to boot the container VM.
	KernelPath = kataAnnotHypervisorPrefix + "kernel"

	// ImagePath is a sandbox annotation for passing a per container path pointing at the guest image that will run in the container VM.
	ImagePath = kataAnnotHypervisorPrefix + "image"

	// InitrdPath is a sandbox annotation for passing a per container path pointing at the guest initrd image that will run in the container VM.
	InitrdPath = kataAnnotHypervisorPrefix + "initrd"

	// HypervisorPath is a sandbox annotation for passing a per container path pointing at the hypervisor that will run the container VM.
	HypervisorPath = kataAnnotHypervisorPrefix + "path"

	// JailerPath is a sandbox annotation for passing a per container path pointing at the jailer that will constrain the container VM.
	JailerPath = kataAnnotHypervisorPrefix + "jailer_path"

	// FirmwarePath is a sandbox annotation for passing a per container path pointing at the guest firmware that will run the container VM.
	FirmwarePath = kataAnnotHypervisorPrefix + "firmware"

	// FirmwareVolumePath is a sandbox annotation for passing a per container path pointing at the guest firmware volume
	// that will be passed to the container VM.
	FirmwareVolumePath = kataAnnotHypervisorPrefix + "firmware_volume"

	// KernelHash is a sandbox annotation for passing a container kernel image SHA-512 hash value.
	KernelHash = kataAnnotHypervisorPrefix + "kernel_hash"

	// ImageHash is an sandbox annotation for passing a container guest image SHA-512 hash value.
	ImageHash = kataAnnotHypervisorPrefix + "image_hash"

	// InitrdHash is an sandbox annotation for passing a container guest initrd SHA-512 hash value.
	InitrdHash = kataAnnotHypervisorPrefix + "initrd_hash"

	// HypervisorHash is an sandbox annotation for passing a container hypervisor binary SHA-512 hash value.
	HypervisorHash = kataAnnotHypervisorPrefix + "hypervisor_hash"

	// JailerHash is an sandbox annotation for passing a jailer binary SHA-512 hash value.
	JailerHash = kataAnnotHypervisorPrefix + "jailer_hash"

	// FirmwareHash is an sandbox annotation for passing a container guest firmware SHA-512 hash value.
	FirmwareHash = kataAnnotHypervisorPrefix + "firmware_hash"

	// FirmwareVolumeHash is an sandbox annotation for passing a container guest firmware volume SHA-512 hash value.
	FirmwareVolumeHash = kataAnnotHypervisorPrefix + "firmware_volume_hash"

	// AssetHashType is the hash type used for assets verification
	AssetHashType = kataAnnotationsPrefix + "asset_hash_type"

	//
	// Generic annotations
	//

	// KernelParams is a sandbox annotation for passing additional guest kernel parameters.
	KernelParams = kataAnnotHypervisorPrefix + "kernel_params"

	// MachineType is a sandbox annotation to specify the type of machine being emulated by the hypervisor.
	MachineType = kataAnnotHypervisorPrefix + "machine_type"

	// MachineAccelerators is a sandbox annotation to specify machine specific accelerators for the hypervisor.
	MachineAccelerators = kataAnnotHypervisorPrefix + "machine_accelerators"

	// CPUFeatures is a sandbox annotation to specify cpu specific features.
	CPUFeatures = kataAnnotHypervisorPrefix + "cpu_features"

	// DisableVhostNet is a sandbox annotation to specify if vhost-net is not available on the host.
	DisableVhostNet = kataAnnotHypervisorPrefix + "disable_vhost_net"

	// EnableVhostUserStore is a sandbox annotation to specify if vhost-user-blk/scsi is abailable on the host
	EnableVhostUserStore = kataAnnotHypervisorPrefix + "enable_vhost_user_store"

	// VhostUserStorePath is a sandbox annotation to specify the directory path where vhost-user devices
	// related folders, sockets and device nodes should be.
	VhostUserStorePath = kataAnnotHypervisorPrefix + "vhost_user_store_path"

	// VhostUserDeviceReconnect is a sandbox annotation to specify the timeout for reconnecting on
	// non-server sockets when the remote end goes away.
	VhostUserDeviceReconnect = kataAnnotHypervisorPrefix + "vhost_user_reconnect_timeout_sec"

	// GuestHookPath is a sandbox annotation to specify the path within the VM that will be used for 'drop-in' hooks.
	GuestHookPath = kataAnnotHypervisorPrefix + "guest_hook_path"

	// DisableImageNvdimm is a sandbox annotation to specify use of nvdimm device for guest rootfs image.
	DisableImageNvdimm = kataAnnotHypervisorPrefix + "disable_image_nvdimm"

	// ColdPlugVFIO is a sandbox annotation used to indicate if devices need to be coldplugged.
	ColdPlugVFIO = kataAnnotHypervisorPrefix + "cold_plug_vfio"

	// HotPlugVFIO is a sandbox annotation used to indicate if devices need to be hotplugged.
	HotPlugVFIO = kataAnnotHypervisorPrefix + "hot_plug_vfio"

	// PCIeRootPort is the number of PCIe root ports to create for the VM.
	PCIeRootPort = kataAnnotHypervisorPrefix + "pcie_root_port"

	// PCIeSwitchPort is the number of PCIe switch ports to create for the VM.
	PCIeSwitchPort = kataAnnotHypervisorPrefix + "pcie_switch_port"

	// EntropySource is a sandbox annotation to specify the path to a host source of
	// entropy (/dev/random, /dev/urandom or real hardware RNG device)
	EntropySource = kataAnnotHypervisorPrefix + "entropy_source"

	// UseLegacySerial sets legacy serial device for guest console if available and implemented for architecture
	UseLegacySerial = kataAnnotHypervisorPrefix + "use_legacy_serial"

	// GPU specific annotations used by remote hypervisor for instance selection
	// Number of GPUs required in the Kata VM
	DefaultGPUs = kataAnnotHypervisorPrefix + "default_gpus"
	// GPU model - tesla, h100, radeon etc..
	DefaultGPUModel = kataAnnotHypervisorPrefix + "default_gpu_model"

	//
	// CPU Annotations
	//

	// DefaultVCPUs is a sandbox annotation for passing the default vcpus assigned for a VM by the hypervisor.
	DefaultVCPUs = kataAnnotHypervisorPrefix + "default_vcpus"

	// DefaultVCPUs is a sandbox annotation that specifies the maximum number of vCPUs allocated for the VM by the hypervisor.
	DefaultMaxVCPUs = kataAnnotHypervisorPrefix + "default_max_vcpus"

	//
	// Memory related annotations
	//

	// DefaultMemory is a sandbox annotation for the memory assigned for a VM by the hypervisor.
	DefaultMemory = kataAnnotHypervisorPrefix + "default_memory"

	// MemSlots is a sandbox annotation to specify the memory slots assigned to the VM by the hypervisor.
	MemSlots = kataAnnotHypervisorPrefix + "memory_slots"

	// MemOffset is a sandbox annotation that specifies the memory space used for nvdimm device by the hypervisor.
	MemOffset = kataAnnotHypervisorPrefix + "memory_offset"

	// VirtioMem is a sandbox annotation that is used to enable/disable virtio-mem.
	VirtioMem = kataAnnotHypervisorPrefix + "enable_virtio_mem"

	// MemPrealloc is a sandbox annotation that specifies the memory space used for nvdimm device by the hypervisor.
	MemPrealloc = kataAnnotHypervisorPrefix + "enable_mem_prealloc"

	// ReclaimGuestFreedMemory is a sandbox annotation that specifies whether the memory freed by the guest will be reclaimed by the hypervisor or not.
	ReclaimGuestFreedMemory = kataAnnotHypervisorPrefix + "reclaim_guest_freed_memory"

	// HugePages is a sandbox annotation to specify if the memory should be pre-allocated from huge pages
	HugePages = kataAnnotHypervisorPrefix + "enable_hugepages"

	// Iommu is a sandbox annotation to specify if the VM should have a vIOMMU device
	IOMMU = kataAnnotHypervisorPrefix + "enable_iommu"

	// Enable Hypervisor Devices IOMMU_PLATFORM
	IOMMUPlatform = kataAnnotHypervisorPrefix + "enable_iommu_platform"

	// FileBackedMemRootDir is a sandbox annotation to soecify file based memory backend root directory
	FileBackedMemRootDir = kataAnnotHypervisorPrefix + "file_mem_backend"

	//
	// Shared File System related annotations
	//

	// Msize9p is a sandbox annotation to specify as the msize for 9p shares
	Msize9p = kataAnnotHypervisorPrefix + "msize_9p"

	// SharedFs is a sandbox annotation to specify the shared file system type, either virtio-9p or virtio-fs.
	SharedFS = kataAnnotHypervisorPrefix + "shared_fs"

	// VirtioFSDaemon is a sandbox annotations to specify virtio-fs vhost-user daemon path
	VirtioFSDaemon = kataAnnotHypervisorPrefix + "virtio_fs_daemon"

	// VirtioFSCache is a sandbox annotation to specify the cache mode for fs version cache
	VirtioFSCache = kataAnnotHypervisorPrefix + "virtio_fs_cache"

	// VirtioFSCacheSize is a sandbox annotation to specify the DAX cache size in MiB
	VirtioFSCacheSize = kataAnnotHypervisorPrefix + "virtio_fs_cache_size"

	// VirtioFSExtraArgs is a sandbox annotation to pass options to virtiofsd daemon
	VirtioFSExtraArgs = kataAnnotHypervisorPrefix + "virtio_fs_extra_args"

	//
	// Block Device related annotations
	//

	// BlockDeviceDriver specifies the driver to be used for block device either VirtioSCSI or VirtioBlock
	BlockDeviceDriver = kataAnnotHypervisorPrefix + "block_device_driver"

	// BlockDeviceAIO specifies I/O mechanism to be used with VirtioBlock for qemu
	BlockDeviceAIO = kataAnnotHypervisorPrefix + "block_device_aio"

	// DisableBlockDeviceUse  is a sandbox annotation that disallows a block device from being used.
	DisableBlockDeviceUse = kataAnnotHypervisorPrefix + "disable_block_device_use"

	// EnableIOThreads is a sandbox annotation to enable IO to be processed in a separate thread.
	// Supported currently for virtio-scsi driver.
	EnableIOThreads = kataAnnotHypervisorPrefix + "enable_iothreads"

	// BlockDeviceCacheSet is a sandbox annotation that specifies cache-related options will be set to block devices or not.
	BlockDeviceCacheSet = kataAnnotHypervisorPrefix + "block_device_cache_set"

	// BlockDeviceCacheDirect is a sandbox annotation that specifies cache-related options for block devices.
	// Denotes whether use of O_DIRECT (bypass the host page cache) is enabled.
	BlockDeviceCacheDirect = kataAnnotHypervisorPrefix + "block_device_cache_direct"

	// BlockDeviceCacheNoflush is a sandbox annotation that specifies cache-related options for block devices.
	// Denotes whether flush requests for the device are ignored.
	BlockDeviceCacheNoflush = kataAnnotHypervisorPrefix + "block_device_cache_noflush"

	// RxRateLimiterMaxRate is a sandbox annotation that specifies max rate on network I/O inbound bandwidth.
	RxRateLimiterMaxRate = kataAnnotHypervisorPrefix + "rx_rate_limiter_max_rate"

	// TxRateLimiter is a sandbox annotation that specifies max rate on network I/O outbound bandwidth
	TxRateLimiterMaxRate = kataAnnotHypervisorPrefix + "tx_rate_limiter_max_rate"

	// EnableGuestSwap is a sandbox annotation to enable swap in the guest.
	EnableGuestSwap = kataAnnotHypervisorPrefix + "enable_guest_swap"

	// EnableRootlessHypervisor is a sandbox annotation to enable rootless hypervisor (only supported in QEMU currently).
	EnableRootlessHypervisor = kataAnnotHypervisorPrefix + "rootless"

	// Initdata is the initdata passed in when CreateVM
	Initdata = kataConfAnnotationsPrefix + "runtime.cc_init_data"
)

// Runtime related annotations
const (
	kataAnnotRuntimePrefix = kataConfAnnotationsPrefix + "runtime."

	// DisableGuestSeccomp is a sandbox annotation that determines if seccomp should be applied inside guest.
	DisableGuestSeccomp = kataAnnotRuntimePrefix + "disable_guest_seccomp"

	// GuestSeLinuxLabel is a SELinux security policy that is applied to a container process inside guest.
	GuestSeLinuxLabel = kataAnnotRuntimePrefix + "guest_selinux_label"

	// SandboxCgroupOnly is a sandbox annotation that determines if kata processes are managed only in sandbox cgroup.
	SandboxCgroupOnly = kataAnnotRuntimePrefix + "sandbox_cgroup_only"

	// EnableVCPUsPinning is a sandbox annotation that controls bundling between vCPU threads and CPUs
	EnableVCPUsPinning = kataAnnotationsPrefix + "enable_vcpus_pinning"

	// EnablePprof is a sandbox annotation that determines if pprof enabled.
	EnablePprof = kataAnnotRuntimePrefix + "enable_pprof"

	// Experimental is a sandbox annotation that determines if experimental features enabled.
	Experimental = kataAnnotRuntimePrefix + "experimental"

	// InterNetworkModel is a sandbox annotaion that determines how the VM should be connected to the
	//the container network interface.
	InterNetworkModel = kataAnnotRuntimePrefix + "internetworking_model"

	// DisableNewNetNs is a sandbox annotation that determines if create a netns for hypervisor process.
	DisableNewNetNs = kataAnnotRuntimePrefix + "disable_new_netns"

	// VfioMode is a sandbox annotation to specify how attached VFIO devices should be treated
	// Overrides the runtime.vfio_mode parameter in the global configuration.toml
	VfioMode = kataAnnotRuntimePrefix + "vfio_mode"

	// CreateContainerTimeout is a sandbox annotaion that sets the create container timeout.
	CreateContainerTimeout = kataAnnotRuntimePrefix + "create_container_timeout"

	// ForceGuestPull is a sandbox annotation that sets experimental_force_guest_pull.
	ForceGuestPull = kataAnnotRuntimePrefix + "experimental_force_guest_pull"
)

// Agent related annotations
const (
	kataAnnotAgentPrefix = kataConfAnnotationsPrefix + "agent."

	// KernelModules is the annotation key for passing the list of kernel
	// modules and their parameters that will be loaded in the guest kernel.
	// Semicolon separated list of kernel modules and their parameters.
	// These modules will be loaded in the guest kernel using modprobe(8).
	// The following example can be used to load two kernel modules with parameters
	///
	//   annotations:
	//     io.katacontainers.config.agent.kernel_modules: "e1000e InterruptThrottleRate=3000,3000,3000 EEE=1; i915 enable_ppgtt=0"
	//
	// The first word is considered as the module name and the rest as its parameters.
	//
	KernelModules = kataAnnotAgentPrefix + "kernel_modules"

	// AgentTrace is a sandbox annotation to enable tracing for the agent.
	AgentTrace = kataAnnotAgentPrefix + "enable_tracing"

	// AgentContainerPipeSize is an annotation to specify the size of the pipes created for containers
	AgentContainerPipeSize       = kataAnnotAgentPrefix + ContainerPipeSizeOption
	ContainerPipeSizeOption      = "container_pipe_size"
	ContainerPipeSizeKernelParam = "agent." + ContainerPipeSizeOption
	CdhApiTimeoutOption          = "cdh_api_timeout"
	CdhApiTimeoutKernelParam     = "agent." + CdhApiTimeoutOption

	// Policy is an annotation containing the contents of an agent policy file, base64 encoded.
	Policy = kataAnnotAgentPrefix + "policy"
)

// Container resource related annotations
const (
	kataAnnotContainerResourcePrefix = kataAnnotContainerPrefix + "resource."

	// ContainerResourcesSwappiness is a container annotation to specify the Resources.Memory.Swappiness
	ContainerResourcesSwappiness = kataAnnotContainerResourcePrefix + "swappiness"

	// ContainerResourcesSwapInBytes is a container annotation to specify the Resources.Memory.Swap
	ContainerResourcesSwapInBytes = kataAnnotContainerResourcePrefix + "swap_in_bytes"
)

// Annotations related to file system options.
const (
	kataAnnotFsOptPrefix = kataAnnotationsPrefix + "fs-opt."

	// FileSystemLayer describes a layer of an overlay filesystem.
	FileSystemLayer = kataAnnotFsOptPrefix + "layer="

	// IsFileSystemLayer indicates that the annotated filesystem is a layer of an overlay fs.
	IsFileSystemLayer = kataAnnotFsOptPrefix + "is-layer"

	// IsFileBlockDevice indicates that the annotated filesystem is mounted on a block device
	// backed by a host file.
	IsFileBlockDevice = kataAnnotFsOptPrefix + "block_device=file"
)

const (
	// SHA512 is the SHA-512 (64) hash algorithm
	SHA512 string = "sha512"
)

// Third-party annotations - annotations defined by other projects or k8s plugins
// but that can change Kata Containers behaviour.

const (
	// This annotation enables SGX. Hardware-based isolation and memory encryption.
	// Supported suffixes are: Ki | Mi | Gi | Ti | Pi | Ei . For example: 4Mi
	// For more information about supported suffixes see https://physics.nist.gov/cuu/Units/binary.html
	SGXEPC = "sgx.intel.com/epc"
)
