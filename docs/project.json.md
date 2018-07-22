# Project.JSON


## Commands


### Components

The `component[s]` command

#### List 
The _"Component List"_ command displays as list of components.

#####Simple List
A simple list will output those found in `project.json`:     
```bash
$ project components list 
```

#####Available List
An `--available` list w/o a value will display the components found to be available at 
the default host and group found in `project.json`:       
```
$ project components list --available 
```

#####Available List w/Host and Group
An `--available={host-group}` list with a value will display the components found to be available at 
the specified host and group:       
```
$ project components list --available=github.com/myorg 
```

#### Add
The _"Component Add"_ command adds the component reference to `project.json`:     
```bash
$ project component add {cref} 
```
The add command will validate the CRef for validity, but it will not modify 
the CRef; it will add it as-is to `project.json`

#### Remove
The _"Component Remove"_ command removes the component reference from `project.json`:     
```bash
$ project component remove [{ctype}|{cref}] 
```
The add command will validate the CRef for validity, but it will not modify 
the CRef; it will add it as-is to `project.json`



## Concepts

### CTypes
A _"**CType**"_ is a _"component type"_ which defines Host, Stack and Role and Version 
for the type of component.    

We use `{ctype}` in our command line examples to indicate a component reference.

#### CType Format

A CType takes the following form:

```
[[{host}/]{stack}/]role[:version]
```

Square brackets `[]` denote optional in the above format.

##### CType Versioning

CType versions should be a single integer with the default being `1` 
when omitted.


#### CType Examples

A CType can be one of:

- Fully Qualified
- Partially Qualified
- Non-Qualified

When one of major, minor or patch is omitted then the latest
available value will be used.   

##### Fully Qualified:
- **With Version**: `wplib.org/wordpress/dbserver:1`  
- **With No Version**: `wplib.org/wordpress/dbserver` 

##### Partially Qualified:
When not qualified we use the value for `{host}` from the `.defaults` 
section of the `project.json` file:
 
- **With Version**: `wordpress/dbserver:1`  
- **With No Version**: `wordpress/dbserver` 

##### Non-Qualified:
When not qualified we use the values for `{host}` and `{stack}` from 
the `.defaults` section of the `project.json` file:
 
- **With Version**: `dbserver:1`  
- **With No Version**: `dbserver` 


### CRefs

A _"**CRef**"_ is a _"component reference"_ which is an unambiguous method of 
referencing a needed component.  

We use `{cref}` in our command line examples to indicate a component reference.

#### CRef Format

A CRef takes the following form:

```
[[{host}/]{group}/]name[:latest]
[[{host}/]{group}/]name[:{major}[.{minor}[.{patch}]]]
```

Square brackets `[]` denote optional in the above format.

##### CRef Versioning

CRef versions may be `latest` or one of the following, according to a subset of [SemVer 2.0.0](https:/semver.org/): 

- `{major}.{minor}.{patch}`
- `{major}.{minor}`
- `{major}`


#### CRef Examples

A CRef can be one of:

- Fully Qualified
- Partially Qualified
- Non-Qualified

When one of major, minor or patch is omitted then the latest
available value will be used.   

##### Fully Qualified:
- **With Full Version**: `github.com/wplib/mysql:5.5.60`  
- **With Partial Version**: `github.com/wplib/mysql:5.5` 
- **With Latest Version**: `github.com/wplib/mysql:latest` 
- **With No Version**: `github.com/wplib/mysql` 

##### Partially Qualified:
When not qualified we use the value for `{host}` from the `.defaults` 
section of the `project.json` file:
 
- **With Full Version**: `wplib/mysql:5.5.60`  
- **With Partial Version**: `wplib/mysql:5.5` 
- **With Latest Version**: `wplib/mysql:latest` 
- **With No Version**: `wplib/mysql` 

##### Non-Qualified:
When not qualified we use the values for `{host}` and `{group}` from 
the `.defaults` section of the `project.json` file:
 
- **With Full Version**: `mysql:5.5.60`  
- **With Partial Version**: `mysql:5.5` 
- **With Latest Version**: `mysql:latest` 
- **With No Version**: `mysql` 




