# Go hacks

> Reference types like slices , maps , interfaces and channels are pass by reference by default (even though the pointer is not directly visible in the code).

> You can’t use _ anywhere in the code . The blank identifier _ can be used to discard values,

# Defer keyword

The defer allows us to guarantee that certain clean-up tasks are performed before we return from a function, for example:

> Closing a file stream
> Unlocking a locked resource (a mutex)
> Printing a footer in a report
> Closing a database connection

# Panic and Recover

> Panic is a built-in function that stops the ordinary flow of control and begins panicking. When the function F calls panic, execution of F stops, any deferred functions in F are executed normally, and then F returns to its caller. To the caller, F then behaves like a call to panic

