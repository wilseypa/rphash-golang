# Application #
For each server in your cluster (other than one which will be deemed master) deploy an agent server with the master's IP.
For the server that was not deemed to be master, deploy a master server with a list of agent IPs to the master.

# Outline #
Client -> Master -> Agents -> Master -> Client
