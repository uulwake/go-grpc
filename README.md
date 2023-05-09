# Go gRPC

Learn about
1. Protobuf
2. gRPC client and server
3. Interceptors

How to run:
1. Server
    ```bash
        make run-grpc-server
    ```

2. Client
    ```bash
        make run-grpc-client
    ```

You can also run the server in watch mode if you have Node.js installed.
1. Install nodemon globally
    ```bash
        npm i -g nodemon
    ```
2. Run server
    ```bash
        make watch-grpc-server
    ```

Now, every time you change `*.go` file in this project, the server will be recompiled and run automatically.