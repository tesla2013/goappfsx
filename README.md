# goappfsx
Ergonomic default file system interface for go applications.

Provides functions for referring to the current user's default configuration
directories and the location of the executable.  These functions are largely for
convenience as they are largely wrappers of the functions provided in the go
standard library.  Provided functions are useful for easily getting the
directory of, writing data to, reading data from, and removing files and
directories in the AppData directories and the location of the executable.
These are tasks I find myself commonly doing, and so this module gets me that
functionality more easily.
