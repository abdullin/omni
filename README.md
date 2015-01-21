# omni - Work in Progress

This is a prototype sample for implementing event-driven microservices in golang ([HappyPancake style](http://abdullin.com/happypancake/)). 

This is currently work-in-progress without any modules implemented (just an infrastructure for defining modules, use cases and verifying these use-cases). There are no modules implemented at the time.



## Folder Structure

Solution is composed from the core infrastructure and actual domain implementation.

**Infrastructure**:

* `bin` - temporary folder with compiled binaries
* `core` - event-driven infrastructure and spec runners
* `etc` - misc files (no source code)
* `host` - entry point for the process


**Domain**:

* `lang` - domain language (value objects and events)
* `views` - module with basic views for the client
* _all the rest_ - event-driven modules with features
