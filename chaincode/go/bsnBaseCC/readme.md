#  about

The chaincode is written in golang language and depends on fabric 1.4 environment.
It includes the operation of adding, deleting, modifying and querying data. The data types supported by this chain code include string, integer, floating point, set (map, list), etc.

## chaincode functions:
### 1.Add data (set)
Input parameter description
baseKey：the unique primary key ID to be saved
baseValue：saved data information
```base
For example: ：{"baseKey":"str","baseValue":"this is string"}
```
basekey is a string that cannot be empty and basevalue can be any type of data. If the basekey already exists, it will directly return to the existing one and cannot be added. If it does not exist, data will be added.
### 2.Modify data (update)
Input parameter description
baseKey：the unique primary key ID to be modified
baseValue：saved data information
```base
For example:{"baseKey":"str","baseValue":"this is string"} 
```
basekey is a string that cannot be empty and basevalue can be any type of data. If the basekey does not exist, it cannot be updated. If it already exists, modify the data.
### 3.Delete data（delete）
Input parameter description
baseKey：the value of the unique primary key ID to be deleted
```base
For example："str"
```
The value of basekey cannot be empty and must exist, otherwise it cannot be deleted.
### 4.Get data（get）
Input parameter description
baseKey：the value of the unique primary key ID to be obtained
```base
For example："str"
```
The value of basekey cannot be empty and must exist, otherwise corresponding information will not be obtained.

### 5.Get history（getHistory）
Input parameter description
baseKey：the value of the unique primary key ID to be obtained
```base
Example: {"STR"}
```
the value of basekey cannot be empty.，response data：txId、txTime、isDelete、dataInfo.


## Introduction to chaincode catalog

* chaincode information

``` bash
bsnchaincode/
```

* entity classes

``` bash
models/
```

* tool classes

``` bash
utils/
```

* test

``` bash
test/
```

* index path

``` bash
META-INF/statedb/couchdb/indexes
```