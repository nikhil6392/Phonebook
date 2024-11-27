# The  Phone Book
## Objective 
To Build an application where user can experience a fast an efficient phone book.
## Language -Go
### Methods 
- Insert
- Delete
- List
- Search

## Description
This application allows user to insert contact Information of people. And save the information for later uses. It 
allows users to search information based on Number, Which is working here as primary key. Users can also delete the 
Info too. And It also list all the contact information user have.

## Working 
- ### readCSVFile():
    - In this method First we check the csvfile exists or not.
    - If exists then we check it if it is empty or not.
    - If It has some values then we read them
- ### saveCSVFile():
    - First we check file exists or not.
    - If it exists then we save it into given path
- ### createIndex():
Which Creates Index based on Phone number

- ### insert():
  Insert methods enable user to insert new contact info
- ### delete():
delete() method gives user ability to delete the contact info

- ### list():
List method lists all the contact information which user have

- ### search() :
Search method gives user ability to search the contact information of users.

- ### Global Variables
- Structure Record : Which keeps the record of all the contacts info  
- CSVFILE  - It is the path of csvfile
- And mapping between index and info

- ### main():
In main function I have done few things
- I want to get a command to execute
- I want to check that I am working with valid CSV file
- I have given a long switch block where it has all the commands
   - insert
   - delete
   - search
   - list
   - default
 - I want to run the right command .

# What I have used in this 
- mapping
- CSV file operation
- os.Args
- os.Stat()
- .Mode()
- .IsRegular()
- switch()
- Regular expression
- UNIX time

# Packages I have worked with

- encoding/csc
- fmt
- os
- strings
- time
- regexp
- strconv


