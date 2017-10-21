### Implementation details

I have implemented a rest service with golang using postgres as the backend database with the following endpoints:

* GET /employees   
* DELETE /employees/{id}
* PUT /employees 
* GET /groups

The main endpoint is /groups endpoint which generates the random groups of employee for the Friday Lunch.

It does this by performing the following logic:
1. select all employees from postgres 
1. shuffle slice of employees
1. calculate number of groups for each size 
1. for each group size extract the number groups calculated on the previous step
1. marshall json output 
   

### Setup

* Install Dependencies 
  make deps

* Run tests
  make test

* Build Code
  make
