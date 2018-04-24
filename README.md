# Anventory
Experimental Ansible inventory management project

This is very WIP, nothing functional at the moment.

## Design goals

The goal is to create a REST API and hopefully eventually a web ui talking to this API to manage the Ansible inventory in a more flexible and easy way.

Some of it's goals are:

* Ability to fully render `host_vars` into the `_meta.host_vars` structure without letting Ansible resolve the variable precedence.
* Easier introspection of where a certain value of a `host_var` comes from (which group, ...)
* Better nested groups:
  * Add all hosts of child groups also to the hosts list of a parent group
  * Add all child groups of child groups also to che children of a parent group
* Ability to use from both a standalone Ansible setup and an AWX/Tower based setup

## Stuff to figure out

* Authentication: API keys/user management/rbac: what are the options?
* Use framework for API? (github.com/gorilla/mux ?)
* UI design - not my strongest point

