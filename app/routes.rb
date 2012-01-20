get '/' do
  erb :main
end

get '/vinyls' do
  content_type :json
  Vinyl.all.to_json
end

post '/vinyls/?' do
  content_type :json
  @vinyl = Vinyl.create(title: params[:title])
  # @vinyl.to_json
end

get '/vinyls/:id' do
  content_type :json
  Vinyl.where("id = ?", params[:id]).first.to_json
end

put '/vinyls/:id' do
  @vinyl = Vinyl.where("id = ?", params[:id]).first
  @vinyl.update_attributes(params)
  @vinyl.save
  @vinyl.to_json
end
