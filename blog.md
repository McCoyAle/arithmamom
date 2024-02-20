# Building and Dockerizing Go Applications 
The goal of this article is to outline the process for building a golang application, containerizing it with Docker and then deploying it into a Kubernetes managed cluster. As many discussions as I have had with others regarding the process, I decided it best to outline what that looks like. In addition, I hope to shed some light on some of the different open sourced and cloud native solutions that integrate with it. 

In addition to the above, if you are reading this and are not exactly interested in building something. You can [checkout my blog](https://medium.com/@mccoyale/getting-started-with-open-source-ae0ea0e9bee6) on medium which discusses how you can leverage open sourced communities to contribute and improve your coding capabilities. 

## Introduction

### Who?
Kubernetes, Docker, Golang. 
Kubernetes, Docker, & Golang. 
Kubernetes, Docker, and Golang. 

### What?
The overarching summary is developing a golang application that is containerized utilizing Docker solution and deployed onto infrastructure with Kubernetes as the container orchestration solution. These are three names you will encounter frequently when working in a cloud-native realm. I am not advocating them as the best option or the one you should learn with. My motto has and always will be the right software and tooling should be used to solve the problem for a specific use case. 

### Why?
Container orchestrators provide a reliable way to manage applications without as much human intervention. If your use case requires you to build an application that packages the application dependencies to deploy to multiple operating system (OS) platforms then this is for you. 

Containerization also supports use cases where the architecture of the application requires certain software features or components to deploy separately. This is where monolithic vs microservices starts. That is not to say one is better than the other. This supports the same theory, the right architecture for your application is the one that best implements the features needed to solve your pain point. 

### How?
At a high level the overall process would look something like this:

1. Think about the problem you are trying to solve or application that you would like to create. If you are focused on merely coding, something simple will suffice.
2. Think, in a task oriented manner, about the functionality that is required. This includes:
    - Users logging into a system
    - A home page being populated
    - The list of workflows to be completed by each user.
    - How will your services communicate.
    - Does your application require a database to store specific information. 
*Note:* In this particular instance I did not go through a design and specification process. This application is  moreso for myself to get reaquainted with certain tooling and the respective workflows, including simply working on my coding capabilities.
3. Based on Step 2, use that information to determine the things that you need to code. 
4. Pick a language to code those things. (Again, I started this with learning golang in mind)
5. Once your code is done, how are you going to package it. 
6. If you do not have a cluster already you will need to get one deployed. If you are not sure about that you can get started with somehting like [Multipass](https://github.com/McCoyAle/kubernetes-materials/blob/main/multipass.md) and create a small cluster on your workstation. Permitted the size of your application and resources available for consumption. 
7. If you are expecting to maintain your code, you can them move into some CI/CD processes that automate and simplify how you manage your code base. If all goes well then maybe you invite some other developers and begin to create a community. 

Will need to create a diagram that depicts the process for:

- Golang Development (option) (not sure how to depict this)
- Containerization process specific to Docker
- Deployment process through Kubernetes deployment, depicting how the components interact with one another.

# Briefly introduce the importance of building and dockerizing Go applications.
# Highlight benefits like portability, reproducibility, and ease of deployment.

## Setting Up a Basic Go Project:

- Walk through the process of setting up a basic Go project.
- Discuss the structure of a typical Go project.

## Building Go Applications:

- Explain the basics of building Go applications using the go build command.
- Discuss the importance of managing dependencies with tools like go mod.

## Introduction to Docker:

- Provide an overview of Docker and containerization.
- Explain the advantages of using Docker for Go applications.

## Creating a Dockerfile:

- Guide readers through the steps of creating a Dockerfile for a Go application.
- Discuss considerations such as base images, environment variables, and dependencies.

## Building Docker Images:

- Explain how to use the docker build command to create Docker images.
- Highlight best practices for optimizing Docker builds.

## Running Docker Containers:

- Demonstrate how to use the docker run command to run containers based on your images.
- Discuss port mapping, volume mounting, and other container-related concepts.

## Multi-Stage Builds (Optional):

- Introduce the concept of multi-stage builds for optimizing Docker images.

Multi-Stage builds was introducted in Docker 17.05, which adds optimization capabilities to make your docker file more consumable or readable. You can read more about the functionality [here](https://docs.docker.com/build/building/multi-stage/). Multi-stage builds are enabled by allowing developers to incorporate two "FROM" statements in the Dockerfile. 

What this process does is allow the developer to build an image and complete tasks with required artifacts or tooling, and then build a new image with only the things required from the base image. Therefore, if there are any tools that you needed to complete one task, you do not need to consume resources or space by deploying them with the final image. This is especially beneficial when debugging or when implementing test to test certain features of your code. See the following examples below. 

- Discuss scenarios where multi-stage builds are beneficial.

## Integration with Databases (Optional):

- If applicable, discuss how to integrate Go applications with databases in a Dockerized environment.

## Troubleshooting and Best Practices:

### Overall Advice

- Become familiar with CLIs available for communicating with each software solution.
- Embrace the community and its respective members for each. There are so many involved that are simply passionate about what they do and would love to converse with you. 
- Get aquainted with the underlying foundations. This comment is specific to operating systems, networking and storage concepts. No matter what tool you work with these will be central to helping you fix an issue when the time comes. 

## Conclusion

If you are just starting your coding or containerization journey, then this process should at least help you get to a point where you have started and have something tangible to showcase. Any software solutions mentioned in this article are based off of my own experience, interest, and where my attention has been focused the last few years. I am not advocating for or against the associated programming language(s), container runtime(s), or database(s). 

To reiterate, my motto is the right software and tooling is the one that solves the problem for your use case. If there is no use case then just learn the ones you are interested in. Somone is using it somewhere and will need your skillset. You can also learn with what you have access to or what is being utilized within your current team or organization. Remember, software and technology is supposed to be fun and enjoyable, no matter how many bugs you run into. Learn what interests you and continue to "multi-stage" build upon it.






