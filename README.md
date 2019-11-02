# kottans-backend
## Git and GitHub
 - [x] Version Control with Git
 
    Course is pretty basic, but I did learn a bit about proper usage of tags and reverts of the commits. Mainly because I have always struggled with counting commits back from the HEAD
 - [x] try.github.io
 
    Have already passed it a while ago, though Learngitbranching was quite a tool for me at the time.
   
 -  Extras 

    - [ ] [Git за 30 хвилин](https://codeguida.com/post/453)

    - [ ] [Git tips](http://sixrevisions.com/web-development/git-tips/) — consolidate your knowledge of Git

    - [x] [Learn git branching](http://learngitbranching.js.org) — improve your understanding of branching

    - [ ] [Fork a repo](https://help.github.com/en/articles/fork-a-repo#step-2-create-a-local-clone-of-your-fork) - freely   experiment with changes without affecting the original project; [Sync a repo](https://help.github.com/en/articles/syncing-a-fork) - keep a fork of the repository up-to-date with the upstream repository

    - [ ] [Communicating using Markdown](https://lab.github.com/githubtraining/communicating-using-markdown)

    - [ ] [TypingClub](https://www.typingclub.com/) — improve your typing speed

    - [x] [Как учиться и справляться с негативными мыслями](https://guides.hexlet.io/learning/)
    
## Unix Shell
- [x] Linux Survival\
    Modules 1-3 were quite easy, since I use Ubuntu at university, but 4th module was helpful and full of new information, that takes too much time to find and read but saves a lot more than that. 
       
    
 - [x] linuxcommand.org\
    Learned a lot about writing scripts. Though I do not get across them too often, I think it nice to know at least a necessary minimum.
 *  Extra Materials
    - [ ] [Linux Bash Shell Cheat Sheet](https://annawilliford.github.io/2016-04-02-UTA/workshop/Linux/bash_cheat_sheet.pdf)
    - [ ] [Advanced Bash-Scripting Guide. An in-depth exploration of the art of shell scripting](http://www.tldp.org/LDP/abs/html/index.html)
    - [ ] [A Guide to 100 (ish) Useful Unix Commands ](http://oliverelliott.org/article/computing/ref_unix/)

## Git Collaboration
I found latest section on rebasing, cherry-picks and other useful as never really used those commands, mainly because did not need to
summary
![task_git_collaboration](task_git_collaboration/summary.jpg)
-  Extra materials
- [ ] [An Introduction to Git and GitHub by Brian Yu (CS50 course), video, ~40 min.](https://youtu.be/MJUJ4wbFm_A)
- [x] [Oh shit, git!](http://ohshitgit.com/)
- [ ] [Flight rules for git](https://github.com/k88hudson/git-flight-rules)
- [ ] [GitHub Learning Lab](https://lab.github.com/courses)
## Go basics 1

#### Theory
1. Passed "Tour of Go", Effective go, Golangbootcamp
2. Reported that EXL skills did not work. 
3. Finished articles
#### Practice
Please note, roman digits task uses numbers.go file, since I did not want to clutter the main.go

[roman digits](roman-digits/main.go)

- Extra materials

- [ ] [FreeCodeCamp Golang Course](https://www.youtube.com/playlist?list=PLJbE2Yu2zumCe9cO3SIyragJ8pLmVv0z9)
- [x] [Go By Example](https://gobyexample.com/)
- [x] [An Introduction to Programming in Go](https://www.golang-book.com/books/intro);
- [ ] [50 оттенков Go](https://habr.com/ru/company/mailru/blog/314804/)

## Memory Management
#### Answers
- What's going to happen if program reaches maximum limit of stack ?
    
    we have a stack overflow and the program receives a Segmentation Fault.

- What's going to happen if program requests a big (more then 128KB) memory allocation on heap ?

    the heap is enlarged via the brk() system call (implementation) to make room for the requested block
- What's the difference between Text and Data memory segments ?
    Text - read/execute - has code that was run
    Data - read/write - has initialized variables
    
    An extract from the output
```
7f67a4850000-7f67a4854000 r--p 00027000 08:11 134921                     /usr/lib64/libtinfo.so.6.0                                                                            
7f67a4854000-7f67a4855000 rw-p 0002b000 08:11 134921                     /usr/lib64/libtinfo.so.6.0                                                                            
7f67a4855000-7f67a487c000 r-xp 00000000 08:11 134641                     /usr/lib64/ld-2.25.so                                                                                 
7f67a4a27000-7f67a4a65000 r-xp 00000000 08:11 134826                     /usr/lib64/libnss_systemd.so.2                                                                        
7f67a4a65000-7f67a4a68000 r--p 0003d000 08:11 134826                     /usr/lib64/libnss_systemd.so.2                                                                        
7f67a4a68000-7f67a4a69000 rw-p 00040000 08:11 134826                     /usr/lib64/libnss_systemd.so.2                                                                        
7f67a4a69000-7f67a4a6c000 rw-p 00000000 00:00 0                                                                                                                                
7f67a4a79000-7f67a4a7b000 rw-p 00000000 00:00 0                                                                                                                                
7f67a4a7b000-7f67a4a7c000 r--p 00026000 08:11 134641                     /usr/lib64/ld-2.25.so                                                                                 
7f67a4a7c000-7f67a4a7e000 rw-p 00027000 08:11 134641                     /usr/lib64/ld-2.25.so                                                                                 
7f67a4a7e000-7f67a4b83000 r-xp 00000000 08:11 131601                     /usr/bin/bash                                                                                         
7f67a4d82000-7f67a4d86000 r--p 00104000 08:11 131601                     /usr/bin/bash                                                                                         
7f67a4d86000-7f67a4d8f000 rw-p 00108000 08:11 131601                     /usr/bin/bash
7f67a4054000-7f67a421b000 r-xp 00000000 08:11 134666                     /usr/lib64/libc-2.25.so                                                                                           
7f67a4d8f000-7f67a4d99000 rw-p 00000000 00:00 0                                                                                                                                
7f67a58a6000-7f67a58e8000 rw-p 00000000 00:00 0                          [heap]                                                                                                
7ffe34c43000-7ffe34c64000 rw-p 00000000 00:00 0                          [stack]                                                                                               
7ffe34cbb000-7ffe34cbd000 r-xp 00000000 00:00 0                          [vdso]                                                                                                
ffffffffff600000-ffffffffff601000 r-xp 00000000 00:00 0                  [vsyscall]
```
    Heap 7f67a58a6000-7f67a58e8000                                                                                                 
    Stack 7ffe34c43000-7ffe34c64000 
    MMS 7f67a4054000-7f67a421b000
    
Materials were not easy but some background from university helped a lot.
Besides, commands did not really work on my mac, but since, hey, Ubuntu is running on it as well, I found my way around these obstacles.
- Extra materials   
- [ ] [Debugging: Simple Memory Leaks in Go](https://medium.com/dm03514-tech-blog/sre-debugging-simple-memory-leaks-in-go-e0a9e6d63d4d)
- [ ] [Avoiding Memory Leak in Golang API](https://hackernoon.com/avoiding-memory-leak-in-golang-api-1843ef45fca8)
- [ ] [Как устроен garbage collector в Go 1.9](https://www.youtube.com/watch?v=CX4GSErFenI)

## TCP. UDP. Network
1. [Internet 101](https://www.khanacademy.org/computing/computer-science/internet-intro)

    Easy course with just little basics
1. [Networking for Web Developers](https://www.udacity.com/course/networking-for-web-developers--ud256)

    Had some interesting info in parts 4 and 5.
    Liked reading how sending packets and etc work(handshakes, flags, teardowns and so on)
    
    Also, the way traceroute works - sending suicide packets to get the route seems like a genius idea for the problem 
1. [How DNS Works](https://howdns.works/)

    Nice comic, enjoyed myself and sent to a friend
    ![Internet101](task_networks/Internet101.png)
    
    ![Udacity networks](task_networks/UdacityNetworks.png)

## Extra materials

- [ ] [Introduction to TCP/IP. Course](https://www.coursera.org/learn/tcpip)
- [ ] [What happens when you type google.com into your browser and press enter?](https://github.com/alex/what-happens-when)
- [ ] [Big list of TCP/UDP ports](https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers)
- [x] [Package 'net' Golang docs](https://golang.org/pkg/net/)
