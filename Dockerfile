FROM ruby:2.6.3

WORKDIR /app
COPY . .
RUN bundle install --system --without development

CMD ["bin/puma", "-C", "puma.rb"]
