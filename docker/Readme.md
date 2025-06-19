# Docker Basics

## 1. Pulling Images

`docker pull` is used to download a Docker image from a remote registry (like Docker Hub) to your local machine. This allows you to use that image to create containers without building the image yourself.

```sh
docker pull node:18
```

## 2. What is a Docker Image?

A Docker image is a lightweight, standalone, and executable package that includes everything needed to run a piece of software: the code, runtime, libraries, environment variables, and configuration files.

## 3. Listing Local Images

To check which Docker images you have on your local machine, use:

```sh
docker image ls
```
or
```sh
docker image
```

## 4. TAG

TAG means version.

## 5. What is a Docker Container?

A Docker container is a running instance of a Docker image. It is an isolated environment that contains everything needed to run your application, including the code, runtime, libraries, and dependencies.
Like in OOP: (Class → Object)

To create a Docker container:

```sh
docker run <name of the repository>
```

## 6. Interactive Mode

Example: `docker run -it ubuntu`

- `-i` (or `--interactive`): Keeps STDIN open so you can interact with the container.
- `-t` (or `--tty`): Allocates a pseudo-TTY (terminal), so you get a command-line interface.

Together, `-it` allows you to run the container in interactive mode, so you can type commands directly into the Ubuntu shell inside the container. This is commonly used for debugging or exploring the container environment.

## 7. Listing Containers

- `docker ps -a` lists all Docker containers on your system, including both running and stopped containers.
- `docker ps` by itself shows only running containers.
- The `-a` (or `--all`) flag shows all containers, regardless of their state (running, stopped, or exited).

## 8. Starting a Container

```sh
docker start <CONTAINER ID or name>
```

## 9. Stopping a Container

```sh
docker stop <CONTAINER ID or name>
```

## 10. Removing Images and Containers

- To remove an image:

  ```sh
  docker rmi <Image_Name>
  ```

- To remove a container:

  ```sh
  docker rm <Container_Name or ID>
  ```

If you want to remove an image, check if you used it as a container. If you did, first remove the container, then the image.

## 11. Pulling a Specific Version

To pull a specific version, use:

```sh
docker pull mysql:8
```

Otherwise, it will pull the latest version.

## 12. Detached Mode

`docker run -d <Image_Name>`
`-d` (or `--detach`): Runs the container in the background and prints the container ID.

## 13. Running MySQL with Environment Variables

```sh
docker run -d -e MYSQL_ROOT_PASSWORD=secret mysql
```

- `-d`: Runs the container in detached (background) mode.
- `-e MYSQL_ROOT_PASSWORD=secret`: Sets the MySQL root password to `secret` inside the container.
- `mysql`: The image to use (latest version by default).
This command starts a MySQL server container in the background with the root password set.

## 14. Running MySQL 8.0 with a Custom Container Name

You can run a specific version of MySQL (like 8.0) in a detached container with a custom name and environment variable:

```sh
docker run -d -e MYSQL_ROOT_PASSWORD=secret --name mysql-older mysql:8.0
```

-`-d`: Runs the container in detached (background) mode.

-`-e`: MYSQL_ROOT_PASSWORD=secret: Sets the root password to secret inside the MySQL container.

-`--name mysql-older`: Assigns the name mysql-older to the container for easier reference.

-`mysql:8.0`: Specifies the image and version of MySQL to use.

## 15. Port Binding

Port binding allows you to map a port on your local machine to a port inside the Docker container. This is useful when you want to access the service running inside the container (like a web server or database) from your host machine.

```sh
docker run -d -p 8080:80 nginx
```

-`-d`: Runs the container in detached mode (in the background).

-`-p 8080:80`: Maps port 80 inside the container (used by Nginx) to port 8080 on your local machine.

-`nginx`: The Docker image used (latest version by default).

This command starts an Nginx container in the background and makes it accessible at http://localhost:8080 on your browser.
like 8080--host port and 80 is container self port.. mapping porcess is called port binding
