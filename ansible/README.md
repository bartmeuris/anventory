# Ansible inventory JSON structs

This part of the repository contains low-level Go structs intented to be marshaled into the JSON format Ansible understands for it's custom/dynamic inventory.

More info about how this is structured can be found in the official Ansible documentation: http://docs.ansible.com/ansible/latest/dev_guide/developing_inventory.html

## Remarks

* At the moment of writing, he ansible documentation mentions `"_meta": { "hostvars": ...`. This is wrong, this should be `"vars"` instead.
* This implementation assumes all host variables are stored in this `_meta.vars.<hostname>`, none in the hosts under groups (which is possible in theory, but ugly imho)
* Tests are not very good at the moment, they could break in future Go/JSON marshalling versions since it expects a certain order of the struct members in the outputted JSON.
* Not sure how to properly test this as a real ansible inventory. Currently this is done by:
  * to generate an `inventory.json` file, run: `ANVENTORY_WRITE=1 GOCACHE=off go test -v -run TestAnsible ./ansible/`
  * Running ansible ping on the `test` group: `ansible -i inventory.json -m ping test`. This assumes the current user can SSH to localhost.

## Example of generated config

```
{
  "_meta": {
    "vars": {
      "localhost": {
        "ansible_host": "localhost",
        "ansible_user": "bart",
        "has_something": false,
        "some_number": 50
      }
    }
  },
  "all": {
    "hosts": {
      "localhost": {}
    },
    "children": {
      "ungrouped": {}
    }
  },
  "test": {
    "hosts": {
      "localhost": {}
    }
  },
  "ungrouped": {}
}
```