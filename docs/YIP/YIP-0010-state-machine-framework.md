# 0010 - State Machine Framework

## Status

Proposed

## Context

Both clearing channel and settlement include a lot of operational states and logic. With the continuous growth of the code base and complexity of the protocol, it becomes harder to maintain and extend the code,
and to make sure that all the edge cases are covered and no security vulnerabilities are introduced.

For this reason, we need to develop simple, but yet extensible framework that allows to create, configure and operate state machines.

## Decision

The state machine framework should stem from the automata theory, given its robustness, proven correctness and wide adoption.

Moreover, given the reactive nature of the clearing and settlement protocols, meaning that the workflow starts with one event, which leads to other event generated and so on. The pushdown automaton is the best candidate for the described logic.

### Vocabulary

**Automata theory** - is the study of abstract machines and automata, as well as the computational problems that can be solved using them. It is a theory in theoretical computer science with close connections to mathematical logic.

**Automaton** - is an abstract self-propelled computing device which follows a predetermined sequence of operations automatically.

**State** represents the step of an automaton execution. In the context of this YIP we will only consider automata with finite number of states, which are called finite automaton (FA) or finite state machines (FSM).

**Symbol** - is an input that triggers the transition from one state to another.

**Alphabet** - is a finite set of symbols.

**Transition** - is a function that maps a state and an input symbol to the next state.

**Initial states** - is a set of states that the automaton can be in before any input.

**Final states** - is a set of states that the automaton can be in after the input is processed.

**Run** - is a sequence of states and input symbols that the automaton goes through during the processing of a string (a set of symbols).

The automaton **halts** when there are no more symbols in the input string.

**Complete automaton** - is an automaton that has a transition for every state and every symbol in the alphabet.

**Deterministic automaton** - is an automaton that has at most one transition for every state and every symbol in the alphabet, and has no epsilon transitions (transitions that do not consume any input).

### Pushdown automaton

**Pushdown automaton** (PDA) - is a finite automaton that employs a stack as a memory to store an arbitrary number of symbols from the alphabet of the PDA, but allows reading only the topmost symbol.

The term "pushdown" refers to the fact that the stack can be regarded as being "pushed down" like a tray dispenser at a cafeteria, since the operations never work on elements other than the top element. A stack automaton, by contrast, does allow access to and operations on deeper elements.

### State machine framework

The state machine framework shall be implemented as a library that can be used by the clearing and settlement protocols.

#### State machine framework vocabulary

To define the specifications of the state machine framework, we need to introduce the following terms:

**External event** - is an event pushed to the state machine by the protocol.

**Internal event** - is an event popped from the stack of the state machine during its run.

**Waiting states** - is a set of states that can accept external events after automaton has halted.

In other words, after the automaton has halted in one of the waiting states, after some time an external event can be pushed to the automaton, which will result in the automaton continuing the run.

#### State machine framework specifications

The framework shall implement a pushdown automaton with the following properties:

- The **alphabet** is a set of events that can be triggered by the protocol, e.g. Deposit, Challenge etc.
- The **states** are the operational states of the protocol, e.g. Created, Deposited, Challenged etc.
- The automaton is **incomplete**, because the flow of the clearing and settlement protocols is mostly linear, and there are no transitions for all the states and all the events.
- The automaton is **deterministic**, because there is at most one transition for every state and every event.
- The automaton starts the run with only 1 _external_ event.
- The stack of the automaton has a fixed size of 1. This means, that transition can result in either no or 1 event pushed to the stack.

The state machine shall:

- throw an error if the transition is not defined for the current state and the event.
- throw an error if it has halted in a non-waiting state and an external event is pushed to the automaton.
- throw an error if it has halted in a non-final or non-waiting state.

## Consequences

Below is the go-type preudo-code of the state machine framework:

```go
type Action func() (Event, error)

type StateConfig struct {
  // is executed during the transition to this state
  Action Action
  Transitions map[Event]State
}

type StateMachine struct {
  InitialState State
  WaitingStates []State
  FinalStates []State

  CurrentState State

  nextEvent Event

  StateConfigs map[State]StateConfig
}

type IStateMachine interface {
  // push external event to the state machine
  func ProcessEvent(event Event) error
  func IsWaitingState(state State) bool
  func IsFinalState(state State) bool
}

func (sm *StateMachine) nextState(event Event) (State, error) {
  if state, ok := sm.StateConfigs[sm.CurrentState]; ok {
    if state.Transitions != nil {
      if next, ok := state.Transitions[event]; ok {
        return next, nil
      }
    }
  }

  return NilState, ErrNoTransition
}

func (sm *StateMachine) run(event Event) error {
  var currEvent Event
  nextEvent := event

  for {
    currEvent = nextEvent

    // if no event
    if currEvent == NilEvent {
      // and current state is non-final and non-waiting, throw error
      if !sm.IsFinalState(sm.CurrentState) && !sm.IsWaitingState(sm.CurrentState) {
        return ErrIncorrectHalt
      }

      // if current state is waiting, do nothing and wait for the external event
      if isWaitingState(sm.CurrentState) {
        return nil
      }
    }

    nextState, err := sm.nextState(currEvent)
    if err != nil {
      return err
    }

    stateConfig, ok := sm.StateConfigs[nextState]
    if !ok {
      return ErrNoStateConfig
    }

    sm.nextEvent, err := stateConfig.Action()
    if err != nil {
      return err
    }

    sm.CurrentState = nextState
  }
}

func (sm *StateMachine) ProcessEvent(event Event) error {
  if !sm.IsWaitingState(sm.CurrentState) {
    return ErrNonWaitingState
  }

  return sm.run(event)
}
```
