# kubectl-creyaml

Generate Custom Resource Example Yaml (CREYaml) from Custom Resource Definition

kubectl-creyaml helps you to create an example Custom Reource(CR) definition for any Custom Resource Definition(CRD) installed in a Kubernetes cluster.

## Example

For a CRD defined as, 

```
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: crontabs.stable.example.com
spec:
  group: stable.example.com
  versions:
    - name: v1
      served: true
      storage: true
      schema:
        openAPIV3Schema:
          type: object
          properties:
            spec:
              type: object
              properties:
                cronSpec:
                  type: string
                image:
                  type: string
                replicas:
                  type: integer
  scope: Namespaced
  names:
    plural: crontabs
    singular: crontab
    kind: CronTab
    shortNames:
    - ct
```

```kubectl creyaml crontabs.stable.example.com``` will generate the below

```
apiVersion: stable.example.com/v1
kind: CronTab
metadata:
  name: ""
  annotations: {}
spec:
  cronSpec: string
  image: string
  replicas: integer
```

Notice how the mentioned type of each property in under spec.

## Install

#### Install with krew

creyaml is available as a [krew plugin](https://krew.sigs.k8s.io/)

Install creyaml by running:

`kubectl krew install creyaml`

#### From source

From the root of this project, run

`make`

This will create a binary under $GOPATH/bin

## Usage 

To output all available properties in a CRD:

```kubectl creyaml example.domain.com```

To output only required properties:

```kubectl creyaml example.domain.com --required```
