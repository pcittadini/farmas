#!/bin/bash

fleetctl destroy consul-agent@1.service
fleetctl destroy consul-agent@2.service
fleetctl destroy consul-agent@3.service
fleetctl destroy consul-discovery@1.service
fleetctl destroy consul-discovery@2.service
fleetctl destroy consul-discovery@3.service
fleetctl destroy python-service@1.service
fleetctl destroy python-service@2.service
fleetctl destroy python-service@3.service
fleetctl destroy registrator@1.service
fleetctl destroy registrator@2.service
fleetctl destroy registrator@3.service
