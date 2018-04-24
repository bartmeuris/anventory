# Anventory
Experimental Ansible inventory management project

This is very WIP, nothing functional at the moment.

## Introduciton

The idea behind this project is to create an Ansible inventory management system being a bit more flexible and easier to manage at medium and possibly larger scale.

Current approach is to create a permanently running service which offers a REST API and hopefully eventually a web ui talking to this API.

## Design goals

Some of it's goals are:

* Ability to fully render `host_vars` into the `_meta.host_vars` structure without letting Ansible resolve the variable precedence.
* Easier introspection of where a certain value of a `host_var` comes from (which group, ...)
* Better nested groups:
  * Add all hosts of child groups also to the hosts list of a parent group
  * Add all child groups of child groups also to che children of a parent group
* Ability to use from both a standalone Ansible setup and an AWX/Tower based setup
* All changes should be track-able. This is a serious downside of AWX/Tower at this moment.

## Stuff to figure out/brainstorming

* Authentication: API keys/user management/rbac: what are the options?
* Use framework for API? (github.com/gorilla/mux ?)
* UI design - not my strongest point
* Ansible-Vault-like functionality? Options/possibilities to evaluate:
  * Full secret management in Anventory - most likely not desirable
  * Hashicorp Vault integration (how?)
  * Keep secrets in classic `group_vars` and `host_vars` files in a repository, and just reference them from Anventory. This creates a split brain scenario.
* How to track changes:
  * Keep history in database
    * Min: tricky to visualize
    * Min: pretty much creating own versioning system
  * Store all data in text-files on disk that can be committed in SC
    * -> Evaluate https://github.com/src-d/go-git as library
    * Min: tricky to keep stuff consistent, a bit creating our own database
    * Min: Possible performance impacts?
    * Min: hard for concurrent access
    * Plus: all changes can be tracked
    * Plus: versioning is done for us
    * Toughts:
      * define human editable on-disk text-format? YAML/JSON/TOML/...?
      * possible to work with post-commit hooks?
