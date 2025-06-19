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

- `-d`: Runs the container in detached (background) mode.

- `-e`: MYSQL_ROOT_PASSWORD=secret: Sets the root password to secret inside the MySQL container.

- `--name mysql-older`: Assigns the name mysql-older to the container for easier reference.

- `mysql:8.0`: Specifies the image and version of MySQL to use.

## 15. Port Binding

Port binding allows you to map a port on your local machine to a port inside the Docker container. This is useful when you want to access the service running inside the container (like a web server or database) from your host machine.

```sh
docker run -d -p 8080:80 nginx
```

- `-d`: Runs the container in detached mode (in the background).

- `-p 8080:80`: Maps port 80 inside the container (used by Nginx) to port 8080 on your local machine.

- `nginx`: The Docker image used (latest version by default).

This command starts an Nginx container in the background and makes it accessible at http://localhost:8080 on your browser.
like 8080--host port and 80 is container self port.. mapping porcess is called port binding


## container Id
To find out which port a Docker image (like MySQL) exposes by default, you can:
```sh
docker image inspect mysql
```
Output:
```sh
 "ExposedPorts": {
    "3306/tcp": {}
}
```

## 16. TroubleShoot CMD
This command shows the logs (output) from a running or stopped container.
```sh
docker logs <CONTAINER_ID or name>
```
```sh
docker exec -it <container_id or name> /bin/bash or /bin/sh
```
docker exec: Runs a command inside a running container.
-i: Keeps STDIN open (interactive).
-t: Allocates a pseudo-TTY (terminal).
<container_id or name>: The ID or name of the running container.
/bin/bash: The command to run (starts a Bash shell).
*** you can type and run any commands supported by the container’s operating system, just like you would in a regular terminal.


## 17. Docker Networks

Docker networks allow containers to communicate with each other and with the outside world. By default, Docker creates a bridge network, but you can create custom networks for better isolation and control.

- To list all networks:

  ```sh
  docker network ls
  ```

- To create a new network:

  ```sh
  docker network create my-network
  ```

- To run a container attached to a specific network:

  ```sh
  docker run -d --network my-network --name my-container nginx
  ```

- To connect an existing container to a network:

  ```sh
  docker network connect my-network my-container
  ```

Custom networks are useful when you want multiple containers to communicate securely and directly with each other.

## Difference Between Docker and Virtual Machines (VMs)

| Feature                | Docker (Containers)                                   | Virtual Machines (VMs)                        |
|------------------------|------------------------------------------------------|-----------------------------------------------|
| Virtualization Level   | OS-level virtualization                              | Hardware-level virtualization                 |
| Resource Usage         | Lightweight, shares host OS kernel                   | Heavyweight, each VM runs its own OS          |
| Startup Time           | Seconds                                              | Minutes                                       |
| Isolation              | Process-level isolation                              | Full OS isolation                             |
| Performance            | Near-native                                          | Slightly lower due to hypervisor overhead     |
| Portability            | Highly portable (runs anywhere Docker is supported)  | Less portable (depends on hypervisor/OS)      |
| Use Case               | Microservices, CI/CD, rapid deployment               | Running different OSes, legacy applications   |

**Summary:**
Docker containers are more lightweight and faster to start than VMs, as they share the host OS kernel. VMs provide stronger isolation by virtualizing hardware and running separate OS instances, but at the cost of higher resource usage and slower performance.
