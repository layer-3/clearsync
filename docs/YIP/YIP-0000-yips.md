# 0000 - Yellow Improvements Proposals

## Status

Accepted

## Context

The Yellow Improvement Proposal (YIP) system is being introduced as an essential framework to document and guide enhancements in our open-source projects. This initiative is born out of the need to preserve the rationale behind key decisions and maintain a clear, historical context for our evolving codebase. Designed to be both accessible and detailed, YIPs will serve as a transparent, structured process for proposing, discussing, and adopting innovations, ensuring that all contributions align with our project's vision and long-term goals. This process will empower our developer community by providing a consistent and democratic platform for collaborative improvement.

## Decision

We will record decisions about architecture and other important decisions in clearsync repository.

YIPs are numbered by the order in which they were **committed**, not by the order in which they were decided. An YIP with a greater number overrides an YIP with a lesser number.

**Meta-(YIPs) are encouraged**. A meta-YIP records a decision about the YIP process itself -- such as a decision about the format, length or style of YIPs. This YIP is a meta YIP.

The format of an YIP shall follow the [template by Michael Nygard](https://github.com/joelparkerhenderson/architecture-decision-record/blob/main/templates/decision-record-template-by-michael-nygard/index.md) following the [suggestion](https://github.com/joelparkerhenderson/architecture-decision-record#suggestions-for-writing-good-adrs) of Joel Parker Henderson:

> ### Decision record template by Michael Nygard
>
> This is the template in [Documenting architecture decisions - Michael Nygard](http://thinkrelevance.com/blog/2011/11/15/documenting-architecture-decisions).
> You can use [adr-tools](https://github.com/npryce/adr-tools) for managing the YIP files.
>
> In each YIP file, write these sections:
>
> ### Title
>
> #### Status
>
> What is the status, such as proposed, accepted, rejected, deprecated, superseded, etc.?
>
> #### Context
>
> What is the issue that we're seeing that is motivating this decision or change?
>
> #### Decision
>
> What is the change that we're proposing and/or doing?
>
> #### Consequences
>
> What becomes easier or more difficult to do because of this change?

### Standard language e.g. shall, should, may and can

Mandatory requirements set within an Yellow Improvement Proposals are clearly distinguished by using specific standard verbs — specifically, shall, should, may, and can.

_Shall_, _should_, _may_ and _can_ are defined in [6.4.7](https://standards.ieee.org/about/policies/opman/sect6.html#6.4.7) of the _IEEE SA Standards Board Operations Manual_.

The word _shall_ indicates mandatory requirements strictly to be followed in order to conform to the standard and from which no deviation is permitted (_shall_ equals _is required to_).

Note that the word _must_ is deprecated and shall not be used when stating mandatory requirements; _must_ is used only to describe unavoidable situations. The word _will_ is deprecated and shall not be used when stating mandatory requirements; _will_ is only used in statements of fact.

The word _should_ indicates that among several possibilities, one is recommended as particularly suitablewithout mentioning or excluding others; or that a certain course of action is preferred but not necessarilyrequired (_should_ equals _is recommended that_).

The word _may_ is used to indicate a course of action permissible within the limits of the standard (_may_ equals _is permitted to_).

The word _can_ is used for statements of possibility and capability, whether material, physical, or causal (_can_ equals _is able to_).

### Consensus process

After opening pull request, author should ask for review from other team members.

Then, add the topic of specific YIP to the Architectural call agenda, discuss it and make a decision/vote.

## Consequences

The approach above provides an easily discoverable reference for developers in the future who question why the code is as it is. Being located in a single folder at the top level allows each YIP to "attach" to multiple files and folders.

There is no burden to make an YIP unless the maintainers of this repository deem it prudent.

It should reduce the chance of oscillating between multiple solutions to the same problem, or recommitting the same mistakes we have made in the past. Further motivation is provided in [this blog post](https://github.blog/2020-08-13-why-write-adrs/).
