The {{cmd .Name}} command replaces any line beginning
with a comment string with the line without the comment string
(default "// ").

If no arguments are passed, assumes standard input

If single argument is passed, assumes file name

Comment string can be passed as an argument. Otherwise
it is set as a configuration option.
