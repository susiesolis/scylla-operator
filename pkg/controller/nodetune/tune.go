// Copyright (C) 2021 ScyllaDB

package nodetune

import (
	"context"
	"fmt"
	"regexp"
	"sort"

	"github.com/scylladb/go-set/strset"
	"github.com/scylladb/scylla-operator/pkg/controllerhelpers"
	"github.com/scylladb/scylla-operator/pkg/cri"
	"github.com/scylladb/scylla-operator/pkg/helpers/slices"
	"github.com/scylladb/scylla-operator/pkg/kubelet"
	"github.com/scylladb/scylla-operator/pkg/naming"
	"github.com/scylladb/scylla-operator/pkg/util/cpuset"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/klog/v2"
	"k8s.io/kubelet/pkg/apis/podresources/v1"
)

func getIRQCPUs(ctx context.Context, kubeletPodResourcesClient kubelet.PodResourcesClient, scyllaPods []*corev1.Pod, hostFullCpuset cpuset.CPUSet) (cpuset.CPUSet, error) {
	scyllaCPUs := cpuset.CPUSet{}
	for _, scyllaPod := range scyllaPods {
		if !controllerhelpers.IsPodTunable(scyllaPod) {
			continue
		}

		klog.Info("Getting cpuset from kubelet PodResources API")
		containerCpuSet, err := cpusetFromKubelet(ctx, kubeletPodResourcesClient, scyllaPod, naming.ScyllaContainerName)
		if err != nil {
			return cpuset.CPUSet{}, fmt.Errorf("can't get cpuset from kubelet PodResources API: %w", err)
		}

		if containerCpuSet.IsEmpty() {
			return cpuset.CPUSet{}, fmt.Errorf("can't find Scylla container cpuset")
		}

		scyllaCPUs = scyllaCPUs.Union(containerCpuSet)
	}

	// Use all CPUs *not* assigned to Scylla container for IRQs.
	return hostFullCpuset.Difference(scyllaCPUs), nil
}

func scyllaDataDirMountHostPaths(ctx context.Context, criClient cri.Client, scyllaPods []*corev1.Pod) ([]string, error) {
	dataDirs := strset.New()

	for _, pod := range scyllaPods {
		cid, err := getScyllaContainerIDInCRIFormat(pod)
		if err != nil {
			return nil, fmt.Errorf("get Scylla container ID: %w", err)
		}

		cs, err := criClient.Inspect(ctx, cid)
		if err != nil {
			return nil, fmt.Errorf("can't inspect container %q: %w", cid, err)
		}

		if cs != nil {
			for _, mount := range cs.Status.GetMounts() {
				if mount.ContainerPath != naming.DataDir {
					continue
				}
				dataDirs.Add(mount.HostPath)
			}
		}
	}

	return dataDirs.List(), nil
}

func cpusetFromKubelet(ctx context.Context, podResourcesClient kubelet.PodResourcesClient, scyllaPod *corev1.Pod, containerName string) (cpuset.CPUSet, error) {
	prs, err := podResourcesClient.List(ctx)
	if err != nil {
		return cpuset.CPUSet{}, fmt.Errorf("can't list pod resources: %w", err)
	}

	for _, pr := range prs {
		if pr.Namespace != scyllaPod.Namespace || pr.Name != scyllaPod.Name {
			continue
		}

		cr, _, ok := slices.Find(pr.Containers, func(cr *v1.ContainerResources) bool {
			return cr.Name == containerName
		})
		if !ok {
			klog.V(4).InfoS("Not found Scylla container in Scylla Pod", "Pod", klog.KObj(scyllaPod), "ContainerName", containerName)
			continue
		}

		cpuIDs := slices.ConvertSlice(cr.CpuIds, func(cpuid int64) int {
			return int(cpuid)
		})
		sort.Ints(cpuIDs)

		klog.V(4).InfoS("Found cpuset of Scylla Pod in kubelet PodResources API", "Pod", klog.KObj(scyllaPod), "Container", containerName, "CPUs", cpuIDs)
		return cpuset.NewCPUSet(cpuIDs...), nil
	}

	klog.V(2).InfoS("Cpuset of Scylla Pod is not available in kubelet PodResources API", "Pod", klog.KObj(scyllaPod), "Container", containerName)
	return cpuset.CPUSet{}, nil
}

func getScyllaContainerIDInCRIFormat(pod *corev1.Pod) (string, error) {
	cidURI, err := controllerhelpers.GetScyllaContainerID(pod)
	if err != nil {
		return "", err
	}

	cid, err := stripContainerID(cidURI)
	if err != nil {
		return "", fmt.Errorf("can't strip container ID prefix from %q: %w", cidURI, err)
	}

	return cid, nil
}

var containerIDRe = regexp.MustCompile(`[a-z]+://([a-z0-9]+)`)

func stripContainerID(containerID string) (string, error) {
	m := containerIDRe.FindStringSubmatch(containerID)
	if len(m) != 2 {
		return "", fmt.Errorf("unsupported containerID format %q", containerID)
	}
	return m[1], nil
}
