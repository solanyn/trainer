// Copyright 2024 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/kubeflow/trainer/v2/pkg/client/clientset/versioned/typed/trainer/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeTrainerV1alpha1 struct {
	*testing.Fake
}

func (c *FakeTrainerV1alpha1) ClusterTrainingRuntimes() v1alpha1.ClusterTrainingRuntimeInterface {
	return newFakeClusterTrainingRuntimes(c)
}

func (c *FakeTrainerV1alpha1) TrainJobs(namespace string) v1alpha1.TrainJobInterface {
	return newFakeTrainJobs(c, namespace)
}

func (c *FakeTrainerV1alpha1) TrainingRuntimes(namespace string) v1alpha1.TrainingRuntimeInterface {
	return newFakeTrainingRuntimes(c, namespace)
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeTrainerV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
