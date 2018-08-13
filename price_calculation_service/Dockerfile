FROM ruby:2.5.1

WORKDIR /app
ADD . /app

ENV RACK_ENV production

RUN bundle install --without development test
EXPOSE 4567
CMD ["thin", "--threaded", "start"]
