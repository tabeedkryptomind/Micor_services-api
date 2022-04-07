FROM python:3.9
# Update package manager (apt-get) 
# Python and Pip

WORKDIR /api


# Copy our script into the container
COPY . .


# Install our Python dependencies
RUN pip install -r requirements.txt

EXPOSE 5000
CMD [ "python3", "main.py" ]