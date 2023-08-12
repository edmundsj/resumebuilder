# resumebuilder

## The Problem
Applying to jobs sucks. One reason that it sucks is resume tailoring. Each job is looking for slightly different skills and experiences, and re-writing your resume each time to submit it to a different employer is hard.

It doesn't have to be this way. You know, in great detail, all the different experiences you have had, and what types of jobs those are relevant to. All you should have to do is write those down in a single place, and have a resume generated for you.

## The Solution
Enter `resumebuilder`. This tool allows you to write a massive resume once, and then filter that resume

## Getting started
1. Write your base resme however you like (google docs, text, markdown)
1. Add additional information, and tag that information with a particular job role. For example:
```
[devops] Implemented CI/CD pipeline in github actions
```
1. Run `resumebuilder` with all your desired tags. For example, say you are applying to a `devops` job:
```
go run resumebuilder <resume_file> devops
```

Any lines that start with `[devops]` or contain the `[devops]` tag 

## Why go?
Why is this program written in go? Mostly because go is a fun language I wanted to learn more about. It's also blazingly fast and trivially easy to publish packages with.


## TODO
- Add markdown support 
- Add HTML support

