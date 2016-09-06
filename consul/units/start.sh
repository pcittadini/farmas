#!/bin/bash
fleetctl start consul-agent@1.service
fleetctl start consul-agent@2.service
fleetctl start consul-agent@3.service
fleetctl start consul-discovery@1.service
fleetctl start consul-discovery@2.service
fleetctl start consul-discovery@3.service
fleetctl start python-service@1.service
fleetctl start python-service@2.service
fleetctl start python-service@3.service
fleetctl start registrator@1.service
fleetctl start registrator@2.service
fleetctl start registrator@3.service
