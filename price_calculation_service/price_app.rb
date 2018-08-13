this_dir = File.expand_path(File.dirname(__FILE__))
lib_dir = File.join(this_dir, 'proto')
$LOAD_PATH.unshift(lib_dir) unless $LOAD_PATH.include?(lib_dir)

require 'rubygems'
require 'bundler/setup'

require 'sinatra/base'
require 'sinatra/json'
require 'grpc'
require './price'

require 'trippb_services_pb'

class App < Sinatra::Base
  get '/' do
    'Price Service'
  end

  get '/trip/price' do
    content_type :json

    request.body.rewind
    req_body = JSON.parse request.body.read

    if req_body['start_point'] && req_body['end_point']
      price = grpc_request(req_body['start_point'], req_body['end_point'])
      respond_with!(price: price)
    else
      respond_with!('Bad request', 400)
    end
  end

  def grpc_request(start_point, end_point)
    stub = Trippb::TripInfo::Stub.new(ENV['DISTANCE_SERVICE'], :this_channel_is_insecure)

    start_trip = Trippb::Point.new(lat: start_point['lat'], long: start_point['long'])
    end_trip = Trippb::Point.new(lat: end_point['lat'], long: end_point['long'])

    trip = Trippb::Trip.new(startPoint: start_trip, endPoint: end_trip)

    resp = stub.get_trip_info(trip)

    Price.new(resp.distance, resp.duration).calc
  end

  def respond_with!(body, status_code = 200)
    halt status_code, json(body)
  end
end

App.run!
