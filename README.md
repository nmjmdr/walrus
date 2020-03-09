# walrus

> "The time has come," the walrus said, "To talk of other things..."

![Walrus and carpenter](https://upload.wikimedia.org/wikipedia/commons/thumb/2/2a/Briny_Beach.jpg/640px-Briny_Beach.jpg)

A task execution engine in GO


> Version 2.0 of Walrus called Jobber https://github.com/nmjmdr/jobber is now out. Jobber simplifies the design and improves 
> the project structure.Jobber does not have built in scheduler like Walrus. A scheduler is used to schedule jobs for 
> execution. Jobber implements dispatcher, worker and recoverer.
> Jobber can be easily extended to perform the functionality of Walrus by adding the scheduler component to it.


With Walrus you can schedule tasks to be executed immediately or at a point in time. It uses Redis to implement the scheduler queues.

Walrus has the following components:

* Scheduler API
* Dispatcher - Dispatches jobs to worker queues
* Worker - A worker process that picks jobs from worker queues and executes jobs 
* Recoverer - A recovery process that can recover jobs from workers that have failed or taking too long to execute and requeue them onto worker queues

A Job has a "Type". The type of the job identifies the name of the worker queue where the job needs to be queued. Workers specific to that job queue compete to pick jobs from the queue.

A Workers needs to be injected with an implementation of Handler interface. This interface defines the method that process the job. It also needs to define a "VisiblityTimeOut". This timeout identifies the time for which a worker is allowed to keep processing the job before being attempted to recovered from the "Recoverer".


Notes:

* Reference later: https://github.com/aphyr/tea-time 
Further enhancements: 
* Allow for repetitive task execution: "every" . 
  * https://lwn.net/Articles/646950/ . 
  * https://blog.acolyer.org/2015/11/23/hashed-and-hierarchical-timing-wheels/ . 
