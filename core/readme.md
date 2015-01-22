# core - event-driven infrastructure and specs

This folder contains core infrastructure for prototyping event-driven
back-ends. You can import it in your go and move from there.

* `root` - binary-sortable UUID and a definition of an event
* `api` - logic for hosting a simple JSON API (with some helpers)
* `bus` - event bus and an in-memory implementation
* `log` - helpers to setup logging
* `env` - environment for defining modules and specs (contracts)
* `specs` - express, verify and print event-driven specifications
* `hosting` - wire and run modules in a process



