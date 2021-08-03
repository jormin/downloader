### About

Todo Manager

### Install

```
go install
```

### Usage

```shell
NAME:
   todo - A simple tool to manage your todo list

USAGE:
   todo [global options] command [command options] [arguments...]

VERSION:
   v1.0.0

DESCRIPTION:
   A simple tool to manage your todo list

COMMANDS:
   add      add todo
   discard  discard or reverse discard(with -r) todo
   do       do or undo(with -r) todo
   edit     edit todo
   ls       show config fund list
   rm       remove todo
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```

#### ls

```shell
todo ls [command options]
```

- -a: show all todo list (default: false)
- -d: show todo list of one date, it's invalid when there is the flag a (default: today)
- -l: the level of todo, optional values are 1-3, 1 is the minimum level and 3 is the maximum level (default: 1)
- -s: the status of todo, optional values are 0-2[uncompleted|completed|discarded] (default: 0)


###### show all todo list

```shell
➜  todo git:(master) ✗ todo ls -a
#     ID                       Date                Level     Status         Content
1     c3vs8ug6n88ksi3u9nr0     20210727(today)     Low       Uncompleted     test
2     c3vt4k86n88lrorie270     20210727(today)     Low       Uncompleted     aaa
3     c3vt4p86n88lrqd1j170     20210727(today)     Low       Uncompleted     bbb
4     c3vt4ug6n88lrsd7il8g     20210726            Low       Completed      ccc

```

###### show todo list of today

```shell
➜  todo git:(master) ✗ todo ls
#     ID                       Date         Level     Status         Content
1     c3vt4k86n88lrorie270     20210727     Low       Uncompleted     aaa
2     c3vt4p86n88lrqd1j170     20210727     Low       Uncompleted     bbb
3     c3vs8ug6n88ksi3u9nr0     20210727     Low       Uncompleted     test
```

###### show todo list of one date

```shell
➜  todo git:(master) ✗ todo ls -d 20210630
#     ID                       Date         Content     Level     Status         UpdateTime
1     c3emlm86n88invcs8s40     20210630     aaa         Low       Uncompleted     2021-07-01 15:12:25
```

#### add

```shell
todo add [command options] [content: todo content]
```

- -d: the date of todo (default: today)
- -l: the level of todo, optional values are 1-3, 1 is the minimum level and 3 is the maximum level (default: 1)
- -s: the status of todo, optional values are 0-1, 0 means uncompleted and 1 means completed (default: 0)

```shell
➜  todo git:(master) ✗ todo add -d 20210630 -l 3 -s 0 "test for add"
c3eog906n88kgk167ht0
```

#### edit

```shell
todo edit [command options] [id: unique todo id]
```

- -c: the content of todo
- -d: the date of todo (default: today)
- -l: the level of todo, optional values are 1-3, 1 is the minimum level and 3 is the maximum level (default: 1)
- -s: the status of todo, optional values are 0-1, 0 means uncompleted and 1 means completed (default: 0)

```shell
➜  todo git:(master) ✗ todo edit -c "test for edit" -l 3 -s 0 c3en56g6n88j8kosta40
edit todo c3en56g6n88j8kosta40 success
```

#### rm

- -a: remove all todo list (default: false)

```shell
todo rm [command options] [id1: unique todo id] [id2] ... [idn]
```

###### remove todo by id

```shell
➜  todo git:(master) ✗ todo rm c3emlko6n88innjbuf30
remove todo c3emlko6n88innjbuf30 success
```

###### remove all todos

```shell
➜  todo git:(master) ✗ todo rm -a
remove all todos success
```

#### do or undo

- -a: do or undo(with -r) all todo (default: false)
- -r: undo todo (default: false)

###### do or undo by id

```shell
➜  todo git:(master) ✗ todo do c4008v06n88n0ufiph60
do todo c4008v06n88n0ufiph60 success

➜  todo git:(master) ✗ todo do -r c4008v06n88n0ufiph60
undo todo c4008v06n88n0ufiph60 success
```

###### do or undo all todos

```shell
➜  todo git:(master) ✗ todo do -a
do todo c3vt9u86n88m1re06t1g success
do todo c3vtalg6n88m1v99t9j0 success
do todo c4008v06n88n0ufiph60 success

➜  todo git:(master) ✗ todo do -a -r
undo todo c3vt9u86n88m1re06t1g success
undo todo c3vtalg6n88m1v99t9j0 success
undo todo c4008v06n88n0ufiph60 success
```

#### discard or reverse discard

- -a: discard or reverse discard(with -r) all todo (default: false)
- -r: reverse discard todo (default: false)
- -s: the status to reverse, optional values are 0-1[uncompleted|completed] (default: 0)

###### discard or reverse discard by id

```shell
➜  todo git:(master) ✗ todo discard c40ck3o6n88nq44oli9g   
discard todo c40ck3o6n88nq44oli9g success

➜  todo git:(master) ✗ todo discard -r -s 1 c40ck3o6n88nq44oli9g
reverse discard todo c40ck3o6n88nq44oli9g success
```

###### discard or reverse discard all todos

```shell
➜  todo git:(master) ✗ todo discard -a 
discard todo c401cg06n88n6cr92q70 success
discard todo c401fng6n88n6qbf341g success
discard todo c40ck3o6n88nq44oli9g success
discard todo c3vt9u86n88m1re06t1g success
discard todo c3vtalg6n88m1v99t9j0 success

➜  todo git:(master) ✗ todo discard -a -r -s 1
reverse discard todo c3vtalg6n88m1v99t9j0 success
reverse discard todo c401cg06n88n6cr92q70 success
reverse discard todo c401fng6n88n6qbf341g success
reverse discard todo c40ck3o6n88nq44oli9g success
reverse discard todo c3vt9u86n88m1re06t1g success
```
