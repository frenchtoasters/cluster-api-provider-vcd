/*
Copyright 2018 Tyler French

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

package machine

import (
        "context"
        "fmt"
        "log"

        clusterv1 "sigs.k8s.io/cluster-api/pkg/apis/cluster/v1alpha1"
        client "sigs.k8s.io/cluster-api/pkg/client/clientset_generated/clientset/typed/cluster/v1alpha1"
)

const (
        ProviderName = "vcd"
)

// Actuator is responsible for performing machine reconciliation
type Actuator struct {
        machinesGetter client.MachinesGetter
}

// ActuatorParams holds parameter information for Actuator
type ActuatorParams struct {
        MachinesGetter client.MachinesGetter
}

// NewActuator creates a new Actuator
func NewActuator(params ActuatorParams) (*Actuator, error) {
        return &Actuator{
                machinesGetter: params.MachinesGetter,
        }, nil
}

// Create creates a machine and is invoked by the Machine Controller
func (a *Actuator) Create(ctx context.Context, cluster *clusterv1.Cluster, machine *clusterv1.Machine) error {
        log.Printf("Creating machine %v for cluster %v.", machine.Name, cluster.Name)
        return fmt.Errorf("TODO: Not yet implemented")
}

// Delete deletes a machine and is invoked by the Machine Controller
func (a *Actuator) Delete(ctx context.Context, cluster *clusterv1.Cluster, machine *clusterv1.Machine) error {
        log.Printf("Deleting machine %v for cluster %v.", machine.Name, cluster.Name)
        return fmt.Errorf("TODO: Not yet implemented")
}

// Update updates a machine and is invoked by the Machine Controller
func (a *Actuator) Update(ctx context.Context, cluster *clusterv1.Cluster, machine *clusterv1.Machine) error {
        log.Printf("Updating machine %v for cluster %v.", machine.Name, cluster.Name)
        return fmt.Errorf("TODO: Not yet implemented")
}

// Exists test for the existance of a machine and is invoked by the Machine Controller
func (a *Actuator) Exists(ctx context.Context, cluster *clusterv1.Cluster, machine *clusterv1.Machine) (bool, error) {
        log.Printf("Checking if machine %v for cluster %v exists.", machine.Name, cluster.Name)
        return false, fmt.Errorf("TODO: Not yet implemented")
}

// The Machine Actuator interface must implement GetIP and GetKubeConfig functions as a workaround for issues
// cluster-api#158 (https://github.com/kubernetes-sigs/cluster-api/issues/158) and cluster-api#160
// (https://github.com/kubernetes-sigs/cluster-api/issues/160).

// GetIP returns IP address of the machine in the cluster.
func (a *Actuator) GetIP(cluster *clusterv1.Cluster, machine *clusterv1.Machine) (string, error) {
        log.Printf("Getting IP of machine %v for cluster %v.", machine.Name, cluster.Name)
        return "", fmt.Errorf("TODO: Not yet implemented")
}

// GetKubeConfig gets a kubeconfig from the master.
func (a *Actuator) GetKubeConfig(cluster *clusterv1.Cluster, master *clusterv1.Machine) (string, error) {
        log.Printf("Getting IP of machine %v for cluster %v.", master.Name, cluster.Name)
        return "", fmt.Errorf("TODO: Not yet implemented")
}
