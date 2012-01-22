get '/' do
  erb :main
end

get '/vinyls' do
  content_type :json
  Vinyl.all.to_json(:include => [:author, :label])
end

post '/vinyls/?' do
  content_type :json
  attributes = JSON.parse request.body.read

  # raise attributes.inspect

  @vinyl = Vinyl.new(attributes["vinyl"])
  @vinyl.author = Author.find_or_create_by_name(attributes["author"]["name"])
  @vinyl.label = Label.find_or_create_by_name(attributes["label"]["name"])
  if @vinyl.save
    @vinyl.to_json
  end
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

delete '/vinyls/:id' do |id|
  content_type :json
  @vinyl = Vinyl.find(params[:id])
  @vinyl.destroy
  {:head => :ok}.to_json
end