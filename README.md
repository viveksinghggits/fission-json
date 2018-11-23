
# fission-json
fission wf functions where one function is processing json and sending the value to another function for process

# Pre-Req
You should have fission server and fission CLI installed on the K8S cluster to go through this example.

# Fision Workflow
This workflow (getandprocess.wg.yaml) consists of two functions, one (getapires) just makes an HTTP GET request to a dummy API and get a JSON response; this 
function then parses that JSON response to get a particular field and passes that field to another function in using Header.
The other function (proecessapires) get the Header value and returns its length as final output of the workflow.

# Creating the workflow function
You will have to create a standard fission function to create a workflow, just pass "workflow" as --env flag while creating this function. 
So the command may somewhat look like this

```fission fn create --name getandprocess --env workflow --src getandprocess.wg.yaml```

if we see the config file to create the workflow

```
apiVersion: 1
output: proecessapirestask
tasks:
  getapirestask:
    run: getapires

  proecessapirestask:
    run: proecessapires
    inputs:
      headers:
        previousres: "{$.Tasks.getapirestask.Output}"
        Content-Type: application/json
    requires:
    - getapirestask
```

we can clearly say that this workflow will try to output whatever is being returned from the `proecessapirestask` task, but since
this task requires another task i.e. `getapirestask` to run first, `getapirestask` will be called first and once this function is executed completely the task `proecessapirestask` will be executed.
The output of the first task i.e. `getapirestask` will be passed as input to the second task `proecessapirestask` in request headers, 
key of the header will be `previousres` and valye of the header will be the OP of the task `getapirestask`.


# Creating the first function
To create the function we can use standard fission CLI and the comamand will look like this 

```fission fn create --name <fun-name> --env <env-name> --src <srcfile-name> --entrypoint <func-entrypoint>```

in our example the valye fun-name, env-name, srcfile-name and func-entrypoint is `getapires`, `go`, `getapires.go` and  `Handler` respectively.

# Creating the second function 
We can create the second function like we create the first function by below command

```fission fn create --name <fun-name> --env <env-name> --src <srcfile-name> --entrypoint <func-entrypoint>```

the values for our particular scenario, for fun-name, env-name, srcfile-name and func-entrypoint will be proecessapires, go, proecessapires.go and Handler respectively.

I have written another [demo](https://github.com/viveksinghggits/fission-wf-dag) where I created a DAG (Directed Acyclic Graph) of fission function and associated them together using fission workflows.

