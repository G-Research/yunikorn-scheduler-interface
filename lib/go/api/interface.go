// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by make build. DO NOT EDIT

package api

import "github.com/G-Research/yunikorn-scheduler-interface/lib/go/si"

type SchedulerAPI interface {
	// Register a new RM, if it is a reconnect from previous RM, cleanup
	// all in-memory data and resync with RM.
	RegisterResourceManager(request *si.RegisterResourceManagerRequest, callback ResourceManagerCallback) (*si.RegisterResourceManagerResponse, error)

	// Update allocation request
	UpdateAllocation(request *si.AllocationRequest) error

	// Update application request
	UpdateApplication(request *si.ApplicationRequest) error

	// Update node info
	UpdateNode(request *si.NodeRequest) error

	// Notify scheduler to reload configuration and hot-refresh in-memory state based on configuration changes
	UpdateConfiguration(request *si.UpdateConfigurationRequest) error

	// Stops the scheduler API service
	Stop()
}

// RM side needs to implement this API
type ResourceManagerCallback interface {

	//Receive Allocation Update Response
	UpdateAllocation(response *si.AllocationResponse) error

	//Receive Application Update Response
	UpdateApplication(response *si.ApplicationResponse) error

	//Receive Node Update Response
	UpdateNode(response *si.NodeResponse) error

	// Run a certain set of predicate functions to determine if a proposed allocation
	// can be allocated onto a node.
	Predicates(args *si.PredicatesArgs) error

	// Run predicate functions to determine if a proposed allocation can be allocated
	// onto a node after preemption. The request contains a list of allocations to
	// speculatively remove.
	PreemptionPredicates(args *si.PreemptionPredicatesArgs) *si.PreemptionPredicatesResponse

	// This plugin is responsible for transmitting events to the shim side.
	// Events can be further exposed from the shim.
	SendEvent(events []*si.EventRecord)

	// Scheduler core can update container scheduling state to the RM,
	// the shim side can determine what to do incorporate with the scheduling state

	// update container scheduling state to the shim side
	// this might be called even the container scheduling state is unchanged
	// the shim side cannot assume to only receive updates on state changes
	// the shim side implementation must be thread safe
	UpdateContainerSchedulingState(request *si.UpdateContainerSchedulingStateRequest)
}

// RM can additionally implement this API to provide information during state dumps
type StateDumpPlugin interface {

	// This plugin is responsible for returning a JSON representation of the state of the shim
	GetStateDump() (string, error)
}
