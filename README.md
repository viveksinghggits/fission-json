# fission-json
fission wf functions where one function is processing json and sending the value to another function for process


# Fision Workflow
This workflow (getandprocess.wg.yaml) consists of two functions, one (getapires) just makes an HTTP GET request to a dummy API and get a JSON response; this 
function then parses that JSON response to get a particular field and passes that field to another function in using Header.
The other function (proecessapires) get the Header value and returns its length as final output of the workflow.

# Creatinng the workflow function
You will have to create a standard fission function to create a workflow, just pass "workflow" as --env flag while creating this function. 
So the command may somewhat look like this

fission fn create --name getandprocess --env workflow --src getandprocess.wg.yaml

TODO
add details of the functions and describe the workflow script.
