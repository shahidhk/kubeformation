// Kubeformation is a tool to generate cloud provider specific declarative
// templates that can be used for creating and managing their managed Kubernetes
// service. The templates are generated by reading a simple spec file, which
// describes a Kubernetes cluster in the simplest form.
//
// Cluster Spec
//
// The cluster spec defines a Kubernetes cluster in a minimalistic way. It takes
// some parameters for the master, like name and k8s version, and a list of node
// pools along with their properties. It can also take volumes for which it
// generates persistent disk in the template and Kubernetes Persistent Volume/Claim
// object as YAML file.
//
// A sample cluster spec is as follows:
//   version: v1
//   name: cluster-name
//   provider: gke
//   k8sVersion: "1.9"
//   nodePools:
//   - name: db-pool
//     type: n1-standard-1
//     labels:
//       app: postgres
//   - name: backend-pool
//     type: n1-standard-2
//     size: 2
//     labels:
//       app: backend
//   volumes:
//   - name: postgres
//     size: 10
// A detailed definition can found at https://godoc.org/github.com/hasura/kubeformation/pkg/spec/v1/#ClusterSpec.
//
// The following managed Kubernetes providers are supported:
//
// Google Kubernetes Engine - GKE
//
// A cluster spec is converted into Google Deployment Manager Template which
// defines are Kubernetes cluster. This template can be used by `gcloud` command
// to create and manage the cluster. GDM templates consists of a jinja template
// file along with a yaml file that defines the parameters.
//
// Azure Kubernetes Service - AKS
//
// For AKS, the cluster spec is converted into a Azure Deployment Manager
// Template. ARM templates consists of two JSON files, one file defines the
// deployment while other defines the parameters that can be used with the deployment.
//
package kubeformation
