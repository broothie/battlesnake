#! env ruby
require 'sinatra'
require 'sinatra/json'
require 'sinatra/reloader' if development?

require_relative 'game/game'

also_reload 'game/*'

before do
  @body = JSON.parse(request.body.read) if request.body.size > 0
end

post '/start' do
  Game.new(@body).start
  json color: '#ff00e6', headType: 'silly', tailType: 'bolt'
end

post '/move' do
  json Game.new(@body).move
end

post '/end' do
  Game.new(@body).end
end

post '/ping' do
  200
end
