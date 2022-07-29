# k

K is a commandline tool to manage multi k8s.

How to Use:

1. First you must have installed:
    * kubectl
    * [kubecm](https://kubecm.cloud/)
    * [k9s](https://github.com/derailed/k9s)

3. Install it

```
go install -v github.com/luaxlou/k@latest
```

3. Use kubecm to manage k8s configs 
   
See [kubecm](https://github.com/sunny0826/kubecm) usage

4. Use k 

Example clusterName is "test":

`k o test`

Then will open k9s gui for "test" cluster.
   

Have fun.
