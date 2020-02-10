FROM python:3.8

ARG USERNAME=streamlit
ARG USER_UID=1000
ARG USER_GID=$USER_UID

# Create a non-root user.
RUN groupadd --gid "$USER_GID" "$USERNAME" \
  && useradd -s /bin/bash --uid "$USER_UID" --gid "$USER_GID" -m "$USERNAME"

# Directory where the app
# will be deployed.
RUN mkdir -p /usr/src/streamlit
WORKDIR /usr/src/streamlit

# hadolint ignore=DL3013
RUN pip install streamlit

# Our requirements file could be empty
# but we will use it as a placeholder
# anyway.
COPY requirements.txt .
RUN pip install -r requirements.txt

# Copy sources.
COPY . /usr/src/streamlit

USER $USERNAME
ENTRYPOINT ["streamlit", "run", "/usr/src/streamlit/app.py"]
