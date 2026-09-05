# periscope
its like intel Vtune but it works. hardware monitoring support for apple silicon, perhaps other architectures in the future (how hard can it be).

# Modules
- `periscope-go`: go code for interfacing through a ui. not sure yet if this will have more of a role in the proj
- `periscope-c`: main interfacing with kernel and system monitoring during specific applications. this can also be called directly from xcode proj apps so might be easier to skip go entirely
- `periscope` (Deprecated): old rust project probably wont use this because I dont want to learn rust

## todo (maybe)
not sure if i want to keep `periscope/` directory, this was originally a rust ui and code for this, however, I think I plan to use go and c for now

# Structure
Ideally how this app will work is to be flexible and lightweight enough to record some diagnostics without impacting system performance that much

The main idea here is that there is a `Process Controller` which runs some command from the terminal. All other operations are wrapped around this process.

The `Sampler` will then 'sample' different diagnostics about the system. Some diagnostics I'd like to record are:
- code hotspots
- cpu/gpu usage
- memory usage
- failures
- time to complete
- time spent on each method (future)


## Process Control
The Process Controller will be a simple platform that is initiated when the program is first started. This will be either from a cli call or a call in the ui. The Process Controller will handle a few things initially:
- extracting crucial logging operations like `stdin`, `stdout`, and `stderr` from the terminal or command that were running
- maybe we can hand in some arguments to the program call to indicate which diagnostics we want to record, whether thats cpu/gpu usage, memory, hotspots, etc.
- indicate when and where errors may occur
- indicate how long the callee took to run
- organize provide some information to the user based on the diagnostics they'd like to collect
- spawn, pause or kill targets?

## Sampler
the sampler functionality is currently stored in the module `data-collection`. In this directory i plan to add tools for getting cpu usage, hotspots, and other system diagnostics. many of these methods may be intended to be used as a goroutine so they can be logged at the same time something is being executed.

In addition to the sampling functionality I plan to add some sort of diagnostics tracking or sumamry at the end of the simution, showing the user more detailed information about which parts of the code can be optimized.

# ideas
there is a library in go called `runtime/metrics` where i can get some information directly from the go runtime about cpu details, not sure if its what I want but it might be something to explore