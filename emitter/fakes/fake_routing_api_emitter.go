// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"code.cloudfoundry.org/route-emitter/emitter"
	"code.cloudfoundry.org/route-emitter/routingtable/schema/event"
)

type FakeRoutingAPIEmitter struct {
	EmitStub        func(routingEvents event.RoutingEvents) error
	emitMutex       sync.RWMutex
	emitArgsForCall []struct {
		routingEvents event.RoutingEvents
	}
	emitReturns struct {
		result1 error
	}
	emitReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRoutingAPIEmitter) Emit(routingEvents event.RoutingEvents) error {
	fake.emitMutex.Lock()
	ret, specificReturn := fake.emitReturnsOnCall[len(fake.emitArgsForCall)]
	fake.emitArgsForCall = append(fake.emitArgsForCall, struct {
		routingEvents event.RoutingEvents
	}{routingEvents})
	fake.recordInvocation("Emit", []interface{}{routingEvents})
	fake.emitMutex.Unlock()
	if fake.EmitStub != nil {
		return fake.EmitStub(routingEvents)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.emitReturns.result1
}

func (fake *FakeRoutingAPIEmitter) EmitCallCount() int {
	fake.emitMutex.RLock()
	defer fake.emitMutex.RUnlock()
	return len(fake.emitArgsForCall)
}

func (fake *FakeRoutingAPIEmitter) EmitArgsForCall(i int) event.RoutingEvents {
	fake.emitMutex.RLock()
	defer fake.emitMutex.RUnlock()
	return fake.emitArgsForCall[i].routingEvents
}

func (fake *FakeRoutingAPIEmitter) EmitReturns(result1 error) {
	fake.EmitStub = nil
	fake.emitReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRoutingAPIEmitter) EmitReturnsOnCall(i int, result1 error) {
	fake.EmitStub = nil
	if fake.emitReturnsOnCall == nil {
		fake.emitReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.emitReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRoutingAPIEmitter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.emitMutex.RLock()
	defer fake.emitMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeRoutingAPIEmitter) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ emitter.RoutingAPIEmitter = new(FakeRoutingAPIEmitter)
