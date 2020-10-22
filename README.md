# kubectl-creyaml

Generate Custom Resource Example Yaml (CREYaml) from Custom Resource Definition

## Intsall

From the root of this project, run

`make`

This will create a binary under $GOPATH/bin

## Usage 

To output all available properties:

`kubectl creyaml example.domain.com`

To output only required properties:

`kubectl creyaml example.domain.com --required`
