# How to create cnf extensions projects 
This repository proposes a method to extend the CNF test suite beyond the original test cases.
Steps to add private test cases:
- create a new repo for the extended test suite, this can be a private repo for instance https://github.com/myorg/example
- fork the https://github.com/test-network-function/cnf-certification-test.git project then go to:
```shell script
cd cnf-certification-test/cnf-certification-test
```
-  add a new import line in  suite_test.go: 
```shell script
import _ "github.com/myorg/example"
```
- build the test suite and run it as usual. The following line will run the existing networking suite and the new extended "mynewsuite" defined in the hypothetical https://github.com/myorg/example project
```shell script
./run-cnf-suites.sh -f networking mynewsuite
```