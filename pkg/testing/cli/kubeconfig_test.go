// Copyright Jetstack Ltd. See LICENSE for details.
package cli

import (
	"fmt"
	"testing"
)

func TestKubeconfig_Success(t *testing.T) {

	dir, err := TmpDir()
	if err != nil {
		t.Errorf("unexpected error getting tmp dir: %v", err)
		return
	}

	path := fmt.Sprintf("%s/test", dir)
	config := fmt.Sprintf("%s/config", dir)

	var args [][]string
	for _, role := range []string{"master", "worker", "etcd", "all"} {
		for _, sign := range []string{"admin", "kubelet", "kube-apiserver", "kube-controller-manager", "kube-proxy", "kube-scheduler"} {
			args = append(args, []string{"kubeconfig", "test/pki/k8s/sign/" + sign, sign, path, config, "--init-role=test-" + role})
		}
	}

	for _, arg := range args {
		RunTest(arg, true, dir, t)
	}
}

func TestKubeconfig_Fail(t *testing.T) {

	dir, err := TmpDir()
	if err != nil {
		t.Errorf("unexpected error getting tmp dir: %v", err)
		return
	}

	path := fmt.Sprintf("%s/test", dir)
	config := fmt.Sprintf("%s/config", dir)

	var args [][]string
	for _, role := range []string{"foo", "bar"} {
		for _, sign := range []string{"admin", "kubelet", "kube-apiserver", "kube-controller-manager", "kube-proxy", "kube-scheduler"} {
			args = append(args, []string{"kubeconfig", "test/pki/k8s/sign/" + sign, sign, path, config, "--init-role=test-" + role})
		}
	}
	for _, role := range []string{"master", "worker", "etcd", "all"} {
		for _, sign := range []string{"bar", "foo"} {
			args = append(args, []string{"kubeconfig", "test/pki/k8s/sign/" + sign, sign, path, config, "--init-role=test-" + role})
		}
	}

	for _, arg := range args {
		RunTest(arg, false, dir, t)
	}
}
