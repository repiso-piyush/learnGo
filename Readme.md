 # Reference types like slices , maps , interfaces  and channels  are pass by reference by default (even though the pointer is not directly visible in the code).

 # You can’t use _ anywhere in the code. The blank identifier _ can be used to discard values, 


 # The defer allows us to guarantee that certain clean-up tasks are performed before we return from a function, for example:

## Closing a file stream
 Unlocking a locked resource (a mutex)
 Printing a footer in a report
 Closing a database connection