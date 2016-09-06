#!/bin.bash

fleetctl stop consul-agent@1.service
fleetctl stop consul-agent@2.service
fleetctl stop consul-agent@3.service
fleetctl stop consul-discovery@1.service
fleetctl stop consul-discovery@2.service
fleetctl stop consul-discovery@3.service
fleetctl stop python-service@1.service
fleetctl stop python-service@2.service
fleetctl stop python-service@3.service
fleetctl stop registrator@1.service
fleetctl stop registrator@2.service
fleetctl stop registrator@3.service
