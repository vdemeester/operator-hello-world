# Code Kubernetes While Using It

## Intro

Hi Everyone, in the next 5 minutes I am going to talk about how I can update, compile and rollout a Kubernetes custom resource (but that can be to any Kubernest component really) using Kubernetes and my browser, nothing else.

## Who am I

I am Mario Loriedo, Italian, live in Paris - France (sometimes someone tell me that I have a french accent when I speak english!).

I work for Red Hat and I have been contributing the the open source project Che for the last 3 years. I am currently the tech lead of Che.

## What is Che

Che is an container based web IDE that runs in your Kubernetes cluster.

A Che development workspace runs in a Pod and everything (the editor, the plugins) runs as services in containers.

That makes it easy to run and debug in a Che development workspace an application that has a Kubernetes manifest.

Its plugin system is compatible with VS Code extensions (we are using the same API).

It's an eclipse foundation project

## The demo

I will run everything on my local minikube cluster.

I have deployed a custom resource and its controller using an operator on my local cluster.

We are going to start a Che workspace where we are going to clone the operator source code

We will update the source code and rollout the new version on the same cluster where my Che workspace is running.

Are you ready? Go!

## Demo

Finger crossed

## Recap

I have:
- navigated to the operator project page
- created a dev workspace with a just one click
- Updated the source code without leaving my browser
- Buid and deployed the operator from the embedded terminal

## Resources on the web

- github.com/eclipse/che
- twitter.com/eclipse-che
- mattermost
- CodeReady Workspaces

---

# Code Kubernetes While Using It - slides only version

## Intro

Hi Everyone, for the next 5 minutes I am going to talk about levereging Kubernetes for software development. And in particular to develop a Kubernetes controller and deployed on the same where.

## Who am I

I am Mario and I work for Red Hat in the Developer Tools Group.

In the last 3 years I have been contributor, committer and now project lead of Eclipse Che.

## What is Che?

These last years have been the years of the containerizations. We have seen any kind of software running in a container.

Amongst those we have seen people running their IDEs in their containers. Like eclipse

## Runs in Kubernetes

With Che we have taken this seriously and we have built an IDE from the ground up to run on Kubernetes
- microservice running in containers
- support vscode extensions
- define components using Kubernetes manifests
- Workspace Custom Resource Definition

## How does Che looks like

<Animated Gif>

## Coding Kubernetes on Kubernetes

Che is an IDE so what about coding a Kubernetes controller?
Che is a controller so what about updating Che while using it?

## Credit

In his book "Test-Driven Development by Example" Kent Beck develops a unit testing framework using TDD: he uses is own framework to test the framework's code.

## Resources on the web

- github.com/eclipse/che
- twitter.com/eclipse-che
- mattermost
- CodeReady Workspaces
- Demo at Red Hat booth