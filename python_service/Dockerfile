FROM python:3.11

RUN mkdir /python-service

WORKDIR /python-service

COPY requirements.txt .

COPY .env .

RUN pip install -r requirements.txt

COPY . .

CMD gunicorn src.app:app --workers 1 --worker-class uvicorn.workers.UvicornWorker --bind=0.0.0.0:8000