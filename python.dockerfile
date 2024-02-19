FROM python:3.11.1

# Set the working directory in the container
WORKDIR /python-app
# Copy the current directory contents into the container at /app
COPY ./ai .
COPY ./idl/proto/internal.proto .

# Install any needed packages specified in requirements.txt
RUN pip install --no-cache-dir -r requirements.txt

RUN apt-get update && apt-get install -y --no-install-recommends apt-utils
RUN apt-get -y install curl
RUN apt-get -y install ffmpeg

RUN python -m grpc_tools.protoc -I. --python_out=. --pyi_out=. --grpc_python_out=. internal.proto

# Run server.py when the container launches
ENTRYPOINT ["python"]
CMD ["/python-app/server.py"]
