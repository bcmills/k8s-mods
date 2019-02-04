// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mods

import (
	_ "k8s.io/api/core/v1"
	_ "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	_ "k8s.io/apimachinery/pkg/apis/meta/v1"
	_ "k8s.io/apiserver/pkg/apis/audit/v1"
	_ "k8s.io/cli-runtime/pkg/genericclioptions/resource"
	_ "k8s.io/client-go/kubernetes/typed/apps/v1"
	_ "k8s.io/cloud-provider"
	_ "k8s.io/cluster-bootstrap/token/api"
	_ "k8s.io/code-generator/pkg/util"
	_ "k8s.io/component-base/config/v1alpha1"
	_ "k8s.io/csi-api/pkg/apis/csi/v1alpha1"
	_ "k8s.io/gengo/generator"
	_ "k8s.io/heapster/metrics/api/v1/types"
	_ "k8s.io/klog"
	_ "k8s.io/kube-aggregator/pkg/apis/apiregistration/v1"
	_ "k8s.io/kube-controller-manager/config/v1alpha1"
	_ "k8s.io/kube-openapi/pkg/util/proto"
	_ "k8s.io/kube-proxy/config/v1alpha1"
	_ "k8s.io/kube-scheduler/config/v1alpha1"
	_ "k8s.io/kubelet/config/v1beta1"
	_ "k8s.io/kubernetes/pkg/apis/core/v1"
	_ "k8s.io/kubernetes/pkg/controller/podautoscaler"
	_ "k8s.io/metrics/pkg/apis/metrics/v1beta1"
	_ "k8s.io/sample-apiserver/pkg/client/listers/wardle/v1beta1"
	_ "k8s.io/sample-cli-plugin/pkg/cmd"
	_ "k8s.io/sample-controller/pkg/apis/samplecontroller/v1alpha1"
	_ "k8s.io/utils/clock"
)
