# Golang-Dictionary

This is a dictionary project. It is using the Badger Key/Value DB.
The dictionary contains entries, which consist of a word, its decsription and the date when it was created.
Here are the actions that you can perform on the entries:
- List all the entries
- Add an entry
- Describe an entry
- Remove an entry

You can precise the action that you want to perfom as an argument of the program.
Here are the respectives actions (mentionned before) with their arguments :
- list
- add <word> <definition>
- describe <word>
- remove <word>

In order to use the project properly, you can process this way :

- go build -o dict
- ./dict -action <action_name> <argument>

Here are some examples :
    Add an entry in the dictionary :
        - ./dict -action add Golang "A great language"
    List all entries :
        - ./dict -action list
    Remove an entry :
        - ./dict -action delete Golang