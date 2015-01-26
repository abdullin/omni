# omni - Work in Progress

This is a prototype sample for implementing event-driven microservices in golang ([HappyPancake style](http://abdullin.com/happypancake/)). 

This is currently work-in-progress without any modules implemented (just an infrastructure for defining modules, use cases and verifying these use-cases). There are almost no modules implemented at the time.



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

## Learn More

This project will be described in more detail in [GTD](http://abdullin.com/gtd/) series. 

At the moment, content in the series is very limited. So you can instead check out the background:

* Design is inspired by event-driven architecture we came up in [HappyPancake project](http://abdullin.com/happypancake/). Multiple blog posts cover that.
* Domain was covered extensively in [BeingTheWorst Podcast](http://beingtheworst.com/about).
