## Summary 

Given a vibe version VIBES.VIBES.VIBES, increment the:

1. VIBES version with an appropriate emoji to describe the vibes of what you are about to release.  

## Introduction

For a very long time we have used vibe-less versioning for our software, releasing it with no way to vibe check dependencies. With many consumers of dependencies not 
having a way to understand the vibes at a glance of the software they are about to consume. Dependency vibe hell is when versions don't offer vibes which prevents you from performing
vibe checks and lead to unexpected shifts in the current vibes of your project.

As a solution to this problem, we propose a simple set of rules and requirements that dictate how version vibes are assigned and incremented. These rules are new and not based on any 
pre-existing widespread common practices in both closed and open source software. For this system to work you will need to declare a public API. Once you identify your public API, 
you communicate vibe changes to it with specific increments to your vibe versioning.

We call this system "Vibe Versioning" or "VibVer" for short. Under this scheme the vibe versions and the way they change convey meaning about the underlying code and what vibes have changed
from one version to another.

## Vibe Versioning Specification (VibVer)

The key words "MUST", "MUST NOT", "REQUIRED", "SHALL", "SHALL NOT", "SHOULD",
"SHOULD NOT", "RECOMMENDED", "MAY", and "OPTIONAL" in this document are to be
interpreted as described in [RFC 2119](https://tools.ietf.org/html/rfc2119).

1. Software using vibe versioning MUST declare a public API. 
2. Software using vibe versioning MUST use emojis to declare the vibe version of their software.


### Examples: 
- 👍.🔥.🥲 
- 😈.🥳.🚀

You can also append a `vv` prefix to the version which stands for vibe version.
- vv👍.🔥.🥲
- vv🫡.😶‍🌫️.👾

## Backus–Naur Form Grammar for Valid VibVer Versions
```
<vibver>         ::= [<prefix>] <vibe-version>
<prefix>         ::= "vv"
<vibe-version>   ::= <emoji> "." <emoji> "." <emoji>
<emoji>          ::= UNICODE_EMOJI /* Valid Unicode emoji character */
```

## Why Use Vibe Versioning?

Spread the vibes with your software!

Reviews from vibe experts: 
>Finally, a versioning spec that speaks the universal language of emotion - VibVer lets your releases express their true energy, 
>making version communication both delightful and instantly relatable! 🥳.👍.🔥
>\- Vibe Expert, GitHub Copilot

>VibVer brilliantly transforms version numbers into emotional storytelling, making software releases not just updates, but shared experiences that developers can instantly vibe with! 🌟
>\- Vibe Expert, Claude Sonnet 3.5

>VibVer is a revolutionary approach to versioning that brings joy and clarity to software releases, making it a must-adopt for any project looking to connect with its users on a deeper, more emotional level! 🚀🥳👍 
>\- Vibe Expert, Gemini 2.0 Flash

## About

The Vibe Versioning specification was originally authored by [Arman Najafian](https://armannajafian.com), a vibes maximalist.

If you'd like to leave feedback, please [open an issue on
GitHub](https://github.com/vibver/vibver/issues)
