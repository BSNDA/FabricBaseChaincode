This pre-made chaincode is written with Golang programming language. it runs on the framework of Fabric 1.4.3. In this chaincode, there are four operation methods: data adding, deleting, modifying and querying. The supported data types include string, integer, floating point, set (map, list), etc.
Add data (set)
Input parameter description
Basekey: the unique primary key ID to be saved
Basevalue: saved data information
For example: {"basekey": "STR", "basevalue": "this is string"}
Where basekey is a string that cannot be empty and basevalue can be any type of data. If the basekey already exists, it will directly return to the existing one and cannot be added. If it does not exist, data will be added.
Modify data (update)
Input parameter description
Basekey: the unique primary key ID to be modified
Basevalue: saved data information
For example: {"basekey": "STR", "basevalue": "this is string"}
Where basekey is a string that cannot be empty and basevalue can be any type of data. If the basekey does not exist, it cannot be updated. If it already exists, modify the data.
Delete data
Input parameter description
Basekey: the value of the unique primary key ID to be deleted
Example: {"STR"}
The value of basekey cannot be empty and must exist, otherwise it cannot be deleted.
Get data
Input parameter description
Basekey: the value of the unique primary key ID to be obtained
Example: {"STR"}
The value of basekey cannot be empty and must exist, otherwise corresponding information will not be obtained.
Get history
Input parameter description
Basekey: the value of the unique primary key ID to be obtained
Example: {"STR"}
Where the value of basekey cannot be empty.