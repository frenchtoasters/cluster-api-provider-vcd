/*
Copyright 2019 Tyler French.

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

package main

        logf.SetLogger(logf.ZapLogger(false))
        entryLog := log.WithName("entrypoint")

        // Setup a Manager
        mgr, err := manager.New(cfg, manager.Options{})
        if err != nil {
                entryLog.Error(err, "unable to set up overall controller manager")
                os.Exit(1)
        }

        cs, err := clientset.NewForConfig(cfg)
        if err != nil {
                panic(err)
        }

        clusterActuator, err := cluster.NewActuator(cluster.ActuatorParams{
                ClustersGetter: cs.ClusterV1alpha1(),
        })
        if err != nil {
                panic(err)
        }

        machineActuator, err := machine.NewActuator(machine.ActuatorParams{
                MachinesGetter: cs.ClusterV1alpha1(),
        })
        if err != nil {
                panic(err)
        }

        // Register our cluster deployer (the interface is in clusterctl and we define the Deployer interface on the actuator)
        common.RegisterClusterProvisioner("vcd", clusterActuator)

        if err := apis.AddToScheme(mgr.GetScheme()); err != nil {
                panic(err)
        }

        if err := clusterapis.AddToScheme(mgr.GetScheme()); err != nil {
                panic(err)
        }

        capimachine.AddWithActuator(mgr, machineActuator)

        capicluster.AddWithActuator(mgr, clusterActuator)

        if err := mgr.Start(signals.SetupSignalHandler()); err != nil {
                entryLog.Error(err, "unable to run manager")
                os.Exit(1)
        }
}
