#syntax=docker/dockerfile:1

FROM python:3.9

ENV PYTHONDONTWRITEBYTECODE=1
ENV PYTHONBUFFERD=1

WORKDIR /api
COPY requirements.txt /api/
RUN pip3 install -r requirements.txt

COPY . /api/

RUN python manage.py migrate