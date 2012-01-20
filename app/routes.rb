get '/' do
  erb :main
end

get '/vinyls' do
  content_type :json
  Vinyl.all.to_json
end

post '/vinyls/?' do
  content_type :json
  attributes = JSON.parse request.body.read

  @vinyl = Vinyl.create(attributes)
  @vinyl.to_json
end

get '/vinyls/:id' do
  content_type :json
  attributes = JSON.parse request.body.read
  Vinyl.where("id = ?", attributes[:id]).first.to_json
end

put '/vinyls/:id' do
  attributes = JSON.parse request.body.read
  @vinyl = Vinyl.where("id = ?", attributes[:id]).first
  @vinyl.update_attributes(attributes)
  @vinyl.save
  @vinyl.to_json
end
