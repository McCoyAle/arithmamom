# Debian based imaged
FROM ubuntu

# Creating the working directory in the container
WORKDIR /app

# Copying local app code into the new directory
COPY . .

# Install Git
RUN apt-get update && apt-get install -y git

# Build the application
RUN apt-get install -y golang
RUN go build -o arithmamom .

# Exposing port 8080 on the container
EXPOSE 8080

# Copying the entrypoint script. This is due to the app being interactive
COPY arithmamom-entrypoint.sh /app

# Configuring entrypoint script as an executable
RUN chmod +x /app/arithmamom-entrypoint.sh

# Running entrypoint script at the start
CMD ["./arithmamom-entrypoint.sh"]

