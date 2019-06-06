FROM python:3.7
COPY . /app
WORKDIR /app

RUN pip install pipenv
RUN pipenv install --system --deploy --ignore-pipfile
EXPOSE 7004

ENTRYPOINT [ "python", "app.py" ]
