// Code generated by go generate. DO NOT EDIT.

// gotext is a tool for managing text in Go source code.
//
// Usage:
//
//	gotext [global options] command [arguments]
//
// The global options are:
//
//	-srclang=<code>            source code language used (default: en-US)
//	-declare-var=<name>        declare variable instead of overwriting the
//	                           message.DefaultCatalog package global
//	-go-build=<constraint>     include a //go:build line with the specified
//	                           constraint line conditions
//
// The commands are:
//
//	update      merge translations and generate catalog
//	extract     extracts strings to be translated from code
//	rewrite     rewrites fmt functions to use a message Printer
//	generate    generates code to insert translated messages
//
// Use "gotext help [command]" for more information about a command.
//
// Additional help topics:
//
// Use "gotext help [topic]" for more information about that topic.
//
// IMPORTANT:
//
// This version of gotext is a fork for the Go-Enjin project. Unless you're
// building enjin things, this is probably not the version of gotext you want
// to use.
//
// # Merge translations and generate catalog
//
// Usage:
//
//	gotext update <package>* [-out <gofile>]
//
// # Extracts strings to be translated from code
//
// Usage:
//
//	gotext extract <package>*
//
// # Rewrites fmt functions to use a message Printer
//
// Usage:
//
//	gotext rewrite [-w] <package>
//
// rewrite is typically done once for a project. It rewrites all usages of
// fmt to use x/text's message package whenever a message.Printer is in scope.
// It rewrites Print and Println calls with constant strings to the equivalent
// using Printf to allow translators to reorder arguments.
//
// The -w flag specifies to write files in place.
//
// # Generates code to insert translated messages
//
// Usage:
//
//	gotext generate <package> [-out <gofile>]
package main
