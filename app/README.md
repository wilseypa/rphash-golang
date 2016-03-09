# Application #
For each server in your cluster (other than one which will be deemed master) deploy an agent server with the master's IP.

For the server that was not deemed to be master, deploy a master server with a list of agent IPs to the master.

## CREATE ##
`SendStreamCreateToAgents` will create a new stream on each of the nodes.

## SEND ##
`SendStreamDataToAgents` will send the data recieved at the master to all agents.

## CLUSTER ##
`SendStreamClusterToAgents` will issue a cluster command on all the nodes and they will report back to the master will all the responses.
