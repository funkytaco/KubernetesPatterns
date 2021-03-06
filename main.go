package main

import (
	// Foundational Patterns
	healthprobe "github.com/sharadbhat/KubernetesPatterns/FoundationalPatterns/HealthProbe"
	managedlifecycle "github.com/sharadbhat/KubernetesPatterns/FoundationalPatterns/ManagedLifecycle"

	// Behavioral Patterns
	batchjob "github.com/sharadbhat/KubernetesPatterns/BehavioralPatterns/BatchJob"
	daemonservice "github.com/sharadbhat/KubernetesPatterns/BehavioralPatterns/DaemonService"
	periodicjob "github.com/sharadbhat/KubernetesPatterns/BehavioralPatterns/PeriodicJob"

	// Configuration Patterns
	envvarconfiguration "github.com/sharadbhat/KubernetesPatterns/ConfigurationPatterns/EnvVarConfiguration"

	// Structural Patterns
	initcontainer "github.com/sharadbhat/KubernetesPatterns/StructuralPatterns/InitContainer"
	sidecar "github.com/sharadbhat/KubernetesPatterns/StructuralPatterns/Sidecar"
)

func main() {
	// Foundational Patterns
	managedlifecycle.GenerateManagedLifecycleSpec()
	healthprobe.GenerateLivenessProbeSpec()
	healthprobe.GenerateReadinessProbeSpec()

	// Behavioral Patterns
	batchjob.GenerateBatchJobSpec()
	periodicjob.GeneratePeriodicJobSpec()
	daemonservice.GenerateDaemonServiceSpec()

	// Configuration Patterns
	envvarconfiguration.GenerateEnvVarConfigurationSpec()

	// Structural Patterns
	initcontainer.GenerateInitContainerSpec()
	sidecar.GenerateSidecarSpec()
}
