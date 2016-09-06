#!/bin/bash
fleetctl submit consul-agent@1.service
fleetctl submit consul-agent@2.service
fleetctl submit consul-agent@3.service
fleetctl submit consul-discovery@1.service
fleetctl submit consul-discovery@2.service
fleetctl submit consul-discovery@3.service
fleetctl submit python-service@1.service
fleetctl submit python-service@2.service
fleetctl submit python-service@3.service
fleetctl submit registrator@1.service
fleetctl submit registrator@2.service
fleetctl submit registrator@3.service
